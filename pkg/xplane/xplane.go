//go:build !test

package xplane

//go:generate mockgen -destination=./__mocks__/xplane.go -package=mocks -source=xplane.go

import (
	"github.com/xairline/goplane/extra"
	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/menus"
	"github.com/xairline/goplane/xplm/plugins"
	"github.com/xairline/xa-honeycomb/pkg"
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
	"runtime"
	"sync"
)

var VERSION = "development"

type XplaneService interface {
	// init
	onPluginStateChanged(state extra.PluginState, plugin *extra.XPlanePlugin)
	onPluginStart()
	onPluginStop()
	// flight loop
	flightLoop(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop float32, counter int, ref interface{}) float32
}

type xplaneService struct {
	Plugin          *extra.XPlanePlugin
	BravoService    honeycomb.BravoService
	Logger          pkg.Logger
	debug           bool
	myMenuId        menus.MenuID
	myMenuItemIndex int
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
		xplaneSvc := &xplaneService{
			Plugin:       extra.NewPlugin("xa honeycomb - "+VERSION, "com.github.xairline.xa-honeycomb", "honeycomb bridge"),
			BravoService: honeycomb.NewBravoService(logger, "TODO"),
			Logger:       logger,
		}
		xplaneSvc.Plugin.SetPluginStateCallback(xplaneSvc.onPluginStateChanged)
		xplaneSvc.Plugin.SetMessageHandler(xplaneSvc.messageHandler)
		return xplaneSvc
	}
}

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

	//systemPath := utilities.GetSystemPath()
	//err := godotenv.Load(s.configFilePath)
	//if err != nil {
	//	s.Logger.Errorf("Some error occured. Err: %s", err)
	//}
	//if os.Getenv("DEBUG") == "true" {
	//	s.debug = true
	//}
	//
	//// API drefs are available at plugin start
	//s.lat_dr, _ = dataAccess.FindDataRef("sim/flightmodel/position/latitude")
	//s.lon_dr, _ = dataAccess.FindDataRef("sim/flightmodel/position/longitude")
	//s.weatherMode_dr, _ = dataAccess.FindDataRef("sim/weather/region/weather_source")
	//s.sysTime_dr, _ = dataAccess.FindDataRef("sim/time/use_system_time")
	//s.simCurrentMonth_dr, _ = dataAccess.FindDataRef("sim/cockpit2/clock_timer/current_month")
	//s.simCurrentDay_dr, _ = dataAccess.FindDataRef("sim/cockpit2/clock_timer/current_day")
	//s.simLocalHours_dr, _ = dataAccess.FindDataRef("sim/cockpit2/clock_timer/local_time_hours")
	//
	//// start with delay to let the dust settle
	//processing.RegisterFlightLoopCallback(s.flightLoop, 5.0, nil)
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
	//
	//// set internal vars to known "no snow" state
	//s.snowNow, s.rwySnowCover, s.iceNow = s.p2x.SnowDepthToXplaneSnowNow(0)
}

func (s *xplaneService) onPluginStop() {
	s.Logger.Info("Plugin stopped")
}

// flightloop, high freq code!
func (s *xplaneService) flightLoop(
	elapsedSinceLastCall,
	elapsedTimeSinceLastFlightLoop float32,
	counter int,
	ref interface{},
) float32 {

	return -1
}

func (s *xplaneService) messageHandler(message plugins.Message) {
	if message.MessageId == plugins.MSG_PLANE_LOADED || message.MessageId == plugins.MSG_SCENERY_LOADED {
		s.Logger.Info("Plane/Scenery loaded")
	}
}

func (s *xplaneService) menuHandler(menuRef interface{}, itemRef interface{}) {
	s.debug = !s.debug

	if s.debug {
		logging.MinLevel = logging.Debug_Level
		menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Checked)
	} else {
		logging.MinLevel = logging.Info_Level
		menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Unchecked)
	}

	s.Logger.Infof("menu clicked: %v", itemRef)
}
