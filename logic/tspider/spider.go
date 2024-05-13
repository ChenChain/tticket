package tspider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"tticket/pkg/log"
	"tticket/pkg/model"
	"tticket/pkg/thttp"
)

var url = "https://jc.zhcw.com/port/client_json.php?transactionType=10001001&lotteryId=1&issueCount=50&startIssue=&endIssue=&startDate=&endDate=&type=0&pageNum=1&pageSize=30"

type LotteryBody struct {
	PageNum  int `json:"string"`
	PageSize int `json:"string"`
	Total    int `json:"string"`
	Pages    int `json:"string"`
	Data     []*LotteryData
	Message  string
}

type LotteryData struct {
	OpenTime         string
	Issue            string // 期号
	FrontWinningNum  string
	FrontWinningNums []string
	BackWinningNum   string
	RedColorNums     []int64
	BlueColorNum     int64
}

func (m *LotteryData) parse() error {
	res := strings.Split(m.FrontWinningNum, " ")
	if len(res) != 6 {
		return errors.New(fmt.Sprintf("error lottery data:%v", m))
	}

	m.RedColorNums = make([]int64, 6)
	for i := 0; i < 6; i++ {
		m.RedColorNums[i], _ = strconv.ParseInt(res[i], 10, 64)
	}
	m.BlueColorNum, _ = strconv.ParseInt(m.BackWinningNum, 10, 64)
	return nil
}

func getLotteryData(ctx context.Context) (string, error) {
	resp, err := thttp.GetClient().Do(ctx, thttp.GET_METHOD, url, map[string]string{"Referer": "https://www.zhcw.com/"})
	log.Info(ctx, "lottery data", zap.String("data", resp))
	if err != nil {
		return "", err
	}
	return resp, nil
}

func parseLottery(ctx context.Context, lotteryContent string) ([]*model.Ball, error) {
	body := &LotteryBody{}
	if err := json.Unmarshal([]byte(lotteryContent), body); err != nil {
		return nil, err
	}

	res := make([]*model.Ball, 0)
	for _, v := range body.Data {
		if err := v.parse(); err != nil {
			return nil, err
		}
		ball := &model.Ball{
			LotteryDrawingTime: v.OpenTime,
			Num1:               v.RedColorNums[0],
			Num2:               v.RedColorNums[1],
			Num3:               v.RedColorNums[2],
			Num4:               v.RedColorNums[3],
			Num5:               v.RedColorNums[4],
			Num6:               v.RedColorNums[5],
			Num7:               v.BlueColorNum,
		}
		res = append(res, ball)
	}
	return res, nil
}

func Logic(ctx context.Context) error {
	data, err := getLotteryData(ctx)
	if err != nil {
		log.Error(ctx, "failed to get lottery data", zap.Error(err))
		return err
	}
	balls, err := parseLottery(ctx, data)
	if err != nil {
		log.Error(ctx, "failed to parse lottery data", zap.Error(err))
		return err
	}
	if err := model.InsertBalls(ctx, balls); err != nil {
		log.Error(ctx, "failed to insert ball", zap.Error(err))
		return err
	}
	return nil
}
