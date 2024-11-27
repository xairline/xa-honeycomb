package xplane

import (
	"github.com/xairline/goplane/extra"
	"github.com/xairline/goplane/xplm/menus"
	"github.com/xairline/goplane/xplm/processing"
	"github.com/xairline/goplane/xplm/utilities"
	"runtime"
)

func (s *xplaneService) onPluginStateChanged(state extra.PluginState, plugin *extra.XPlanePlugin) {
	switch state {
	case extra.PluginStart:
		s.onPluginStart()
	case extra.PluginStop:
		s.onPluginStop()
	case extra.PluginEnable:
		s.Logger.Infof("Plugin: %s enabled", plugin.GetName())
	case extra.PluginDisable:
		s.Logger.Infof("Plugin: %s disabled", plugin.GetName())
	}
}

func (s *xplaneService) onPluginStart() {
	s.Logger.Info("Plugin started")
	runtime.GOMAXPROCS(runtime.NumCPU())

	processing.RegisterFlightLoopCallback(s.flightLoop, 5.0, nil)
	//
	// setup menu
	menuId := menus.FindPluginsMenu()
	menuContainerId := menus.AppendMenuItem(menuId, "XA Honeycomb", 0, false)
	s.myMenuId = menus.CreateMenu("XA Honeycomb", menuId, menuContainerId, s.menuHandler, nil)
	menus.AppendMenuItem(s.myMenuId, "Reload Profile", 0, true)
	menus.AppendMenuSeparator(s.myMenuId)
	s.myMenuItemIndex = menus.AppendMenuItem(s.myMenuId, "Enable Debug", 1, true)

	if s.debug {
		menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Checked)
	} else {
		menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Unchecked)
	}
	increaseCmd := utilities.CreateCommand("Honeycomb Bravo/increase", "Increase the value of the autopilot mode selected with the rotary encoder.")
	decreaseCmd := utilities.CreateCommand("Honeycomb Bravo/decrease", "Decrease the value of the autopilot mode selected with the rotary encoder.")

	mode_ias := utilities.CreateCommand("Honeycomb Bravo/mode_ias", "Set the autopilot mode to IAS.")
	mode_alt := utilities.CreateCommand("Honeycomb Bravo/mode_alt", "Set the autopilot mode to ALT.")
	mode_vs := utilities.CreateCommand("Honeycomb Bravo/mode_vs", "Set the autopilot mode to VS.")
	mode_hdg := utilities.CreateCommand("Honeycomb Bravo/mode_hdg", "Set the autopilot mode to HDG.")
	mode_crs := utilities.CreateCommand("Honeycomb Bravo/mode_crs", "Set the autopilot mode to CRS.")

	// set up command handlers
	utilities.RegisterCommandHandler(increaseCmd, s.changeApValue, true, "up")
	utilities.RegisterCommandHandler(decreaseCmd, s.changeApValue, true, "down")
	utilities.RegisterCommandHandler(mode_ias, s.changeAPMode, true, "ias")
	utilities.RegisterCommandHandler(mode_alt, s.changeAPMode, true, "alt")
	utilities.RegisterCommandHandler(mode_vs, s.changeAPMode, true, "vs")
	utilities.RegisterCommandHandler(mode_hdg, s.changeAPMode, true, "hdg")
	utilities.RegisterCommandHandler(mode_crs, s.changeAPMode, true, "crs")

}

func (s *xplaneService) onPluginStop() {
	s.BravoService.Exit()
	s.Logger.Info("Plugin stopped")
}
