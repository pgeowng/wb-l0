package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	NatsClusterId string `envconfig:"NATS_CLUSTER_ID"`
	NatsClientId  string `envconfig:"NATS_CLIENT_ID"`
	NatsSubject   string `envconfig:"NATS_SUBJECT"`
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