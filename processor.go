//go:build tinygo

package main

import (
	"github.com/wasmvision/wasmvision-sdk-go/logging"
	"wasmcv.org/wasm/cv/mat"
)

//export process
func process(image mat.Mat) mat.Mat {
	loadConfig()

	logging.Info("Processor running")

	return image.Clone()
}
