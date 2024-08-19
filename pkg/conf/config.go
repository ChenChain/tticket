package conf

import "github.com/spf13/viper"

var Config = &Global{}

type Global struct {
	Mail
	Spider
	Strategy

	Mysql
}

type Mail struct {
	Address  string
	Port     string
	UserName string
	Password string
	Host     string
	FromUser string `mapstructure:"from_user"`
}

type Spider struct {
}

type Mysql struct {
	DbName      string `mapstructure:"db_name"`
	Host        string
	Port        string
	UserName    string
	Password    string
	LogLocation string
}

type Strategy struct {
	UserStrategyNum int
}

func Init() {
	viper.SetConfigName("tticket")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/tticket")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(Config); err != nil {
		panic(err)
	}

	if Config.Strategy.UserStrategyNum == 0 {
		Config.Strategy.UserStrategyNum = 100
	}
}
