package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/ini.v1"
)

type Config struct {
	Server *Server
}

type Server struct {
	Port int
	Env  string
}

func (server *Server) IsLocal() bool {
	return server.Env == "local"
}

func New() *Config {
	return &Config{}
}

const iniFilePath = "internal/infrastructure/config/ini.config"

func (config *Config) Setup() error {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	iniFilePath := filepath.Join(rootPath, iniFilePath)
	cfg, err := ini.Load(iniFilePath)
	if err != nil {
		return err
	}
	if err := config.loadServerConfig(cfg); err != nil {
		return err
	}
	return nil
}

func (config *Config) loadServerConfig(cfg *ini.File) error {
	server := &Server{}
	portStr := cfg.Section("server").Key("port").String()
	if portStr == "" {
		return fmt.Errorf("port is required")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}
	server.Port = port

	envStr := cfg.Section("server").Key("env").String()
	if envStr == "" {
		return fmt.Errorf("env is required")
	}
	server.Env = envStr

	config.Server = server
	return nil
}
