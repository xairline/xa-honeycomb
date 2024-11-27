//go:build !test

package main

import (
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
	"time"
)

func main() {
	Logger := honeycomb.NewConsoleLogger()
	Logger.Info("startup")

	bravoSvc := honeycomb.NewBravoService(Logger)

	honeycomb.OnLEDAP()
	time.Sleep(5 * time.Second)
	honeycomb.OffLEDAP()
	time.Sleep(5 * time.Second)

	bravoSvc.Exit()
}
