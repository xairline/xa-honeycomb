//go:build !test

package xplane

//go:generate mockgen -destination=./__mocks__/xplane.go -package=mocks -source=xplane.go

import (
	"github.com/expr-lang/expr/vm"
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

type profile struct {
	ProfileType string `yaml:"profile_type,omitempty"`
	Condition   string `yaml:"condition,omitempty"`
	Datarefs    []struct {
		Dataref_str string `yaml:"dataref_str,omitempty"`
		Dataref     dataAccess.DataRef
		Operator    string  `yaml:"operator,omitempty"`
		Threshold   float32 `yaml:"threshold,omitempty"`
		expr        *vm.Program
		env         map[string]interface{}
	} `yaml:"datarefs,omitempty"`
	Commands []struct {
		Command_str string `yaml:"command_str,omitempty"`
		Command     utilities.CommandRef
	}
	Data []struct {
		Dataref_str string `yaml:"dataref_str,omitempty"`
		Dataref     dataAccess.DataRef
	}
	on  func()
	off func()
}

type Profile struct {
	BUS_VOLTAGE        profile `yaml:"bus_voltage,omitempty"`
	HDG                profile `yaml:"hdg,omitempty"`
	NAV                profile `yaml:"nav,omitempty"`
	ALT                profile `yaml:"alt,omitempty"`
	APR                profile `yaml:"apr,omitempty"`
	VS                 profile `yaml:"vs,omitempty"`
	AP                 profile `yaml:"ap,omitempty"`
	IAS                profile `yaml:"ias,omitempty"`
	REV                profile `yaml:"rev,omitempty"`
	AP_STATE           profile `yaml:"ap_state,omitempty"`
	GEAR               profile `yaml:"gear,omitempty"`
	RETRACTABLE_GEAR   profile `yaml:"retractable_gear,omitempty"`
	MASTER_WARN        profile `yaml:"master_warn,omitempty"`
	MASTER_CAUTION     profile `yaml:"master_caution,omitempty"`
	FIRE               profile `yaml:"fire,omitempty"`
	OIL_LOW_PRESSURE   profile `yaml:"oil_low_pressure,omitempty"`
	FUEL_LOW_PRESSURE  profile `yaml:"fuel_low_pressure,omitempty"`
	ANTI_ICE           profile `yaml:"anti_ice,omitempty"`
	ENG_STARTER        profile `yaml:"eng_starter,omitempty"`
	APU                profile `yaml:"apu,omitempty"`
	VACUUM             profile `yaml:"vacuum,omitempty"`
	HYDRO_LOW_PRESSURE profile `yaml:"hydro_low_pressure,omitempty"`
	AUX_FUEL_PUMP      profile `yaml:"aux_fuel_pump,omitempty"`
	PARKING_BRAKE      profile `yaml:"parking_brake,omitempty"`
	VOLT_LOW           profile `yaml:"volt_low,omitempty"`
	DOORS              profile `yaml:"doors,omitempty"`
}

type XplaneService interface {
	// init
	onPluginStateChanged(state extra.PluginState, plugin *extra.XPlanePlugin)
	onPluginStart()
	onPluginStop()
	// flight loop
	flightLoop(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop float32, counter int, ref interface{}) float32
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
	profile         *Profile
	apSelector      string
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
			profile:      nil,
			apSelector:   "",
		}
		xplaneSvc.Plugin.SetPluginStateCallback(xplaneSvc.onPluginStateChanged)
		xplaneSvc.Plugin.SetMessageHandler(xplaneSvc.messageHandler)
		return xplaneSvc
	}
}

func (s *xplaneService) messageHandler(message plugins.Message) {
	if message.MessageId == plugins.MSG_PLANE_LOADED {
		s.Logger.Info("Plane loaded")
		s.profile = nil
		honeycomb.AllOff()
	}
}
