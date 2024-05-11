package conf

import "github.com/spf13/viper"

var Config *Global

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
}

type Spider struct {
}

type Mysql struct {
	Host     string
	Port     string
	UsrName  string
	Password string
}

type Strategy struct {
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
}
