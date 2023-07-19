package common

import "go-micro.dev/v4/config"

// MySQL相关配置
type MySQLConfig struct {
	Host        string
	User        string
	Password    string
	Port        int
	TablePrefix string
	Name        string
}

func GetMysqlFromConsul(config config.Config, path ...string) *MySQLConfig {
	mysqlConfig := &MySQLConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}
