//go:build !test

package xplane

//go:generate mockgen -destination=./__mocks__/xplane.go -package=mocks -source=xplane.go

import (
	"github.com/xairline/goplane/extra"
	"github.com/xairline/goplane/xplm/dataAccess"
	"github.com/xairline/goplane/xplm/menus"
	"github.com/xairline/goplane/xplm/plugins"
	"github.com/xairline/goplane/xplm/utilities"
	"github.com/xairline/xa-honeycomb/pkg"
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
	"path/filepath"
	"sync"
)

var VERSION = "development"

type leds struct {
	rules string
	on    func()
	off   func()
}

type XplaneService interface {
	// init
	onPluginStateChanged(state extra.PluginState, plugin *extra.XPlanePlugin)
	onPluginStart()
	onPluginStop()
	// flight loop
	flightLoop(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop float32, counter int, ref interface{}) float32
	// cmd handlers
	Increase(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int
	Decrease(command utilities.CommandRef, phase utilities.CommandPhase, ref interface{}) int
	// menu handler
	menuHandler(menuRef interface{}, itemRef interface{})
	// datarefs
	setupDataRefs(airplaneICAO string)
}

type xplaneService struct {
	Plugin          *extra.XPlanePlugin
	BravoService    honeycomb.BravoService
	Logger          pkg.Logger
	debug           bool
	pluginPath      string
	myMenuId        menus.MenuID
	myMenuItemIndex int
	leds            map[string]leds
}

var xplaneSvcLock = &sync.Mutex{}
var xplaneSvc XplaneService

func NewXplaneService(
	logger pkg.Logger,
) XplaneService {
	if xplaneSvc != nil {
		logger.Info("Xplane SVC has been initialized already")
		return xplaneSvc
	} else {
		logger.Info("Xplane SVC: initializing")
		xplaneSvcLock.Lock()
		defer xplaneSvcLock.Unlock()

		systemPath := utilities.GetSystemPath()
		pluginPath := filepath.Join(systemPath, "Resources", "plugins", "xa-honeycomb")

		xplaneSvc := &xplaneService{
			Plugin:       extra.NewPlugin("xa honeycomb - "+VERSION, "com.github.xairline.xa-honeycomb", "honeycomb bridge"),
			BravoService: honeycomb.NewBravoService(logger),
			Logger:       logger,
			pluginPath:   pluginPath,
			leds:         nil,
		}
		xplaneSvc.Plugin.SetPluginStateCallback(xplaneSvc.onPluginStateChanged)
		xplaneSvc.Plugin.SetMessageHandler(xplaneSvc.messageHandler)
		return xplaneSvc
	}
}

func (s *xplaneService) messageHandler(message plugins.Message) {
	if message.MessageId == plugins.MSG_PLANE_LOADED {
		s.Logger.Info("Plane loaded")
		aircraftIACODrf, found := dataAccess.FindDataRef("sim/aircraft/view/acf_ICAO")
		if !found {
			s.Logger.Errorf("Failed to find ICAO")
		}
		aircraftIACO := dataAccess.GetString(aircraftIACODrf)
		s.Logger.Debugf("Plane ICAO: %s", aircraftIACO)
		s.setupDataRefs(aircraftIACO)
		honeycomb.AllOff()
	}
}
