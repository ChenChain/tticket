package tuser

import (
	"context"
	"time"
	"tticket/pkg/dal/localcache"
	"tticket/pkg/model"
)

var CACHE_USER_KEY = "cache_user"

func GetUserMails(ctx context.Context) []string {
	res, _ := localcache.Cache().Get(CACHE_USER_KEY)
	return res.([]string)
}

func CacheUserMail(ctx context.Context) error {
	users, err := model.FindUsers(ctx)
	if err != nil {
		return err
	}
	mails := []string{}
	for _, u := range users {
		mails = append(mails, u.Mail)
	}
	localcache.Cache().Set(CACHE_USER_KEY, mails, 20*time.Minute)
	return nil
}

// TODO 每15分钟刷新users
