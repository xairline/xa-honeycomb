package main

import (
	"context"
	"fmt"
	"github.com/xairline/xa-honeycomb/pkg"
	"os"
)

// App struct
type App struct {
	ctx      context.Context
	profiles map[string]pkg.Profile
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	// get current dir of the app
	exePath, _ := os.Getwd()
	return fmt.Sprintf("Hello %s, It's show time!", exePath)
}
