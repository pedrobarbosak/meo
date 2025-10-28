package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/pedrobarbosak/meo/pkg/meo"
	"github.com/pedrobarbosak/meo/pkg/meo/requests"

	"github.com/pedrobarbosak/go-errors"
)

type Config struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Hostname string `json:"hostname" validate:"required"`
	Wifi     struct {
		PrivateNetwork struct {
			SSID     string `json:"ssid"`
			Password string `json:"password"`
		} `json:"network"`
		Band2_4GHz struct {
			Bandwidth     int `json:"bandwidth" validate:"omitempty,oneof=1 3"` // Bandwidth 20MHz = 1, 20MHz/40MHz = 3
			Channel       int `json:"channel"`
			TransmitPower int `json:"transmitPower"`
		} `json:"2.4ghz"`
		Band5GHz struct {
			Bandwidth     int `json:"bandwidth" validate:"omitempty,oneof=1 3 7"` // Bandwidth 20MHz = 1, 20MHz/40MHz = 3, 20MHz/40MHz/80MHz = 7
			Channel       int `json:"channel"`
			TransmitPower int `json:"transmitPower"`
		} `json:"5ghz"`
	}
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

	settings := requests.PutWifiSettings{
		PrivateNetwork: &requests.PrivateNetwork{
			SSID:     cfg.Wifi.PrivateNetwork.SSID,
			Password: cfg.Wifi.PrivateNetwork.Password,
		},
		Band2_4GHz: &requests.BandSettings{
			Bandwidth:     cfg.Wifi.Band2_4GHz.Bandwidth,
			Channel:       cfg.Wifi.Band2_4GHz.Channel,
			TransmitPower: cfg.Wifi.Band2_4GHz.TransmitPower,
		},
		Band5GHz: &requests.BandSettings{
			Bandwidth:     cfg.Wifi.Band5GHz.Bandwidth,
			Channel:       cfg.Wifi.Band5GHz.Channel,
			TransmitPower: cfg.Wifi.Band5GHz.TransmitPower,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	if err = meo.SetWifiSettings(ctx, settings); err != nil {
		log.Panicln(errors.New(err))
	}
	cancel()

	log.Println("Done")
}
