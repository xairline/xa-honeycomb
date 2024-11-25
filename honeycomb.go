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

	honeycomb.OnLEDAlt()
	honeycomb.OnLEDMasterWarning()
	time.Sleep(5 * time.Second)
	honeycomb.OffLEDAlt()
	honeycomb.OnLEDFuelPump()
	honeycomb.OnLEDLeftGearGreen()
	time.Sleep(5 * time.Second)

	bravoSvc.Exit()
}
