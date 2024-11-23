//go:build !test

package main

import (
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
)

func main() {
	Logger := honeycomb.NewConsoleLogger()
	Logger.Info("startup")
}
