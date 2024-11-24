//go:build !test

package main

import (
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
)

func main() {
	Logger := honeycomb.NewConsoleLogger()
	Logger.Info("startup")

	bravoSvc := honeycomb.NewBravoService(Logger)

	//Logger.Infof("BravoService is ready: %v", bravoSvc.IsReady())

	bravoSvc.Exit()
}
