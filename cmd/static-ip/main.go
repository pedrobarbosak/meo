package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/pedrobarbosak/meo/pkg/meo"

	"github.com/pedrobarbosak/go-errors"
)

type Config struct {
	Username string            `json:"username" validate:"required"`
	Password string            `json:"password" validate:"required"`
	Hostname string            `json:"hostname" validate:"required"`
	MACs     map[string]string `json:"macs" validate:"required,gte=1"`
}

func NewConfig() *Config {
	return &Config{
		Username: "meo",
		Password: "meo",
		Hostname: "192.168.1.254",
	}
}

func main() {
	cfg := NewConfig()

	configFile, err := os.ReadFile("cmd/static-ip/config.json")
	if err != nil {
		log.Panicln(errors.New(err))
	}

	if err = json.Unmarshal(configFile, cfg); err != nil {
		log.Panicln(errors.New(err))
	}

	meo, err := meo.New(cfg.Username, cfg.Password, cfg.Hostname)
	if err != nil {
		log.Panicln(errors.New(err))
	}

	for mac, ip := range cfg.MACs {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		if err = meo.AssignStaticIP(ctx, mac, ip); err != nil {
			log.Panicln(errors.New(err))
		}

		cancel()
	}

	log.Println("Done")
}
