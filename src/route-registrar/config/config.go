package config

import (
	"github.com/fraenkel/candiedyaml"
	"os"
)

type MessageBusServer struct {
	Host     string
	User     string
	Password string
}

type HealthCheckerConf struct {
	Name              string
	Interval          float64 `yaml:"interval_in_seconds"`
	HealthcheckScript string  `yaml:"healthcheck_script_path"`
}

type Config struct {
	MessageBusServers []MessageBusServer `yaml:"message_bus_servers"`
	ExternalHost      string             `yaml:"external_host"`
	ExternalIp        string             `yaml:"external_ip"`
	Port              int
	HealthChecker     *HealthCheckerConf `yaml:"health_checker"`
}

func InitConfigFromFile(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	decoder := candiedyaml.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
