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
	s.myMenuItemIndex = menus.AppendMenuItem(s.myMenuId, "Enable Debug", 0, true)
	if s.debug {
		menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Checked)
	} else {
		menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Unchecked)
	}
	increaseCmd := utilities.CreateCommand("Honeycomb Bravo/increase", "Increase the value of the autopilot mode selected with the rotary encoder.")
	decreaseCmd := utilities.CreateCommand("Honeycomb Bravo/decrease", "Decrease the value of the autopilot mode selected with the rotary encoder.")
	// set up command handlers
	utilities.RegisterCommandHandler(increaseCmd, s.Increase, true, nil)
	utilities.RegisterCommandHandler(decreaseCmd, s.Decrease, true, nil)
}

func (s *xplaneService) onPluginStop() {
	s.Logger.Info("Plugin stopped")
}
