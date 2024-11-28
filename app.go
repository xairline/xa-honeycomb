package main

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert/yaml"
	"github.com/xairline/xa-honeycomb/pkg"
	"os"
	"path"
	"strings"
)

// App struct
type App struct {
	ctx      context.Context
	profiles []pkg.Profile
}

// NewApp creates a new App application struct
func NewApp() *App {
	exePath, _ := os.Getwd()
	profilesFolder := path.Join(exePath, "profiles") // list all yaml files under profiles folder
	entries, err := os.ReadDir(profilesFolder)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	var profiles []pkg.Profile
	// Loop through the entries
	for _, entry := range entries {
		if !entry.IsDir() && strings.Contains(entry.Name(), ".yaml") { // Skip directories, list only files
			fmt.Println(entry.Name()) // Prints file name only
			f, err := os.ReadFile(path.Join(profilesFolder, entry.Name()))
			if err != nil {
				fmt.Printf("Error opening file: %v", err)
				return nil
			}
			var res pkg.Profile
			err = yaml.Unmarshal(f, &res)
			if err != nil {
				fmt.Printf("Error reading file: %v", err)
				return nil
			}
			profiles = append(profiles, res)
		}
	}
	return &App{
		profiles: profiles,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	// get current dir of the app

	return fmt.Sprintf("Hello %v, It's show time!", a.profiles)
}
