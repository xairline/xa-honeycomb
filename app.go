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
	ctx          context.Context
	profiles     []pkg.Profile
	profileFiles []string
}

// NewApp creates a new App application struct
func NewApp() *App {
	exePath, _ := os.Executable()
	fmt.Println("exePath:", exePath)
	profilesFolder := path.Join(exePath, "..", "..", "..", "..", "profiles") // list all yaml files under profiles folder
	entries, err := os.ReadDir(profilesFolder)
	if err != nil {
		fmt.Println("Error:", err)
		exePath, _ = os.Getwd()
		profilesFolder = path.Join(exePath, "profiles")
		entries, err = os.ReadDir(profilesFolder)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
	}
	var profiles []pkg.Profile
	var profileFiles []string
	// Loop through the entries
	for _, entry := range entries {
		if !entry.IsDir() && strings.Contains(entry.Name(), ".yaml") { // Skip directories, list only files
			fileName := path.Join(profilesFolder, entry.Name())
			f, err := os.ReadFile(fileName)
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
			profileFiles = append(profileFiles, fileName)
		}
	}
	return &App{
		profiles:     profiles,
		profileFiles: profileFiles,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) pkg.Profile {
	// get current dir of the app
	for _, profile := range a.profiles {
		if profile.Name == name {
			return profile
		}
	}

	return pkg.Profile{}
}

func (a *App) GetProfiles() []pkg.Profile {
	//return "default,A339"
	return a.profiles
}

func (a *App) GetProfileFiles() []string {
	//return "default,A339"
	return a.profileFiles
}
