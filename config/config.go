package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type AthenaConfig struct {
	DataSource DataSourceConfig `json:"dataSource"`

	Kafka      KafkaConfig      `json:"kafka"`
	PushServer PushServerConfig `json:"pushServer"`
	RestServer RestServerConfig `json:"restServer"`
	JwtSecret  string           `json:"jwtSecret"`
}

type DataSourceConfig struct {
	DriverName        string `json:"driverName"`
	Host              string `json:"host"`
	Port              int64  `json:"port"`
	Database          string `json:"database"`
	User              string `json:"user"`
	Password          string `json:"password"`
	EnableAutoMigrate bool   `json:"enableAutoMigrate"`
}

type KafkaConfig struct {
	Brokers []string `json:"brokers"`
}

type PushServerConfig struct {
	Addr string `json:"addr"`
	Path string `json:"path"`
}

type RestServerConfig struct {
	Addr string `json:"addr"`
}

var config AthenaConfig
var configOnce sync.Once

func GetConfig() *AthenaConfig {
	configOnce.Do(func() {
		bytes, err := ioutil.ReadFile("config.json")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}
