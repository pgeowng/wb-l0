package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	NatsURL       string `envconfig:"NATS_URL"`
	NatsClusterId string `envconfig:"NATS_CLUSTER_ID"`
	NatsClientId  string `envconfig:"NATS_CLIENT_ID"`
	NatsSubject   string `envconfig:"NATS_SUBJECT"`

	PgDSN   string `envconfig:"PG_DSN"`
	PgReset bool   `envconfig:"PG_RESET"`

	HttpPort string `envconfig:"PORT"`
	LogFile  string `envconfig:"LOG_FILE"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Configuration: ", string(configBytes))
	})
	return &config
}
