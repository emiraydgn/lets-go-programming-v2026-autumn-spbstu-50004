package main

import (
	"flag"
	"fmt"

	"github.com/emiraydgn/task-3/internal/cbr"
	"github.com/emiraydgn/task-3/internal/config"
)

func main() {
	configPath := flag.String("config", "", "path to YAML configuration file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	currencies, err := cbr.Load(cfg.InputFile)
	if err != nil {
		panic(fmt.Errorf("load currencies: %w", err))
	}

	fmt.Println("Loaded currencies:", len(currencies))
}
