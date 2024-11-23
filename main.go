//go:build !test

package main

import (
	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/plugins"
	"github.com/xairline/goplane/xplm/utilities"
	"github.com/xairline/xa-honeycomb/pkg/xplane"
	"path/filepath"
)

func main() {
}

func init() {
	xplaneLogger := xplane.NewXplaneLogger()
	plugins.EnableFeature("XPLM_USE_NATIVE_PATHS", true)
	logging.MinLevel = logging.Info_Level
	logging.PluginName = "xa honeycomb - " + xplane.VERSION
	// get plugin path
	systemPath := utilities.GetSystemPath()
	pluginPath := filepath.Join(systemPath, "Resources", "plugins", "xa-honeycomb")
	xplaneLogger.Infof("Plugin path: %s", pluginPath)

	// entrypoint
	xplane.NewXplaneService(
		xplaneLogger,
	)
}
