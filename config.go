//go:build tinygo

package main

import (
	"github.com/wasmvision/wasmvision-sdk-go/config"
	"github.com/wasmvision/wasmvision-sdk-go/logging"
)

var (
	something string
)

const (
	defaultThing = "your value"
)

func loadConfig() {
	if something == "" {
		ok, _, isErr := config.GetConfig("processor-thing").Result()
		if isErr {
			something = defaultThing
		} else {
			something = ok
		}

		logging.Info("Using processor-thing " + something)
	}
}
