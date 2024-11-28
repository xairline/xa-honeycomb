package pkg

import "github.com/expr-lang/expr/vm"

type BravoProfile struct {
	ProfileType string `yaml:"profile_type,omitempty"`
	Condition   string `yaml:"condition,omitempty"`
	Datarefs    []struct {
		Dataref_str string `yaml:"dataref_str,omitempty"`
		Dataref     interface{}
		Operator    string  `yaml:"operator,omitempty"`
		Threshold   float32 `yaml:"threshold,omitempty"`
		Index       int     `yaml:"index,omitempty"`
		Expr        *vm.Program
		Env         map[string]interface{}
	} `yaml:"datarefs,omitempty"`
	Commands []struct {
		Command_str string `yaml:"command_str,omitempty"`
		Command     interface{}
	}
	Data []struct {
		Dataref_str string `yaml:"dataref_str,omitempty"`
		Dataref     interface{}
	}
	On  func()
	Off func()
}

type Profile struct {
	// AP Knobs
	AP_HDG BravoProfile `yaml:"ap_hdg,omitempty"`
	AP_VS  BravoProfile `yaml:"ap_vs,omitempty"`
	AP_ALT BravoProfile `yaml:"ap_alt,omitempty"`
	AP_IAS BravoProfile `yaml:"ap_ias,omitempty"`
	AP_CRS BravoProfile `yaml:"ap_crs,omitempty"`
	// LEDs
	BUS_VOLTAGE        BravoProfile `yaml:"bus_voltage,omitempty"`
	HDG                BravoProfile `yaml:"hdg,omitempty"`
	NAV                BravoProfile `yaml:"nav,omitempty"`
	ALT                BravoProfile `yaml:"alt,omitempty"`
	APR                BravoProfile `yaml:"apr,omitempty"`
	VS                 BravoProfile `yaml:"vs,omitempty"`
	AP                 BravoProfile `yaml:"ap,omitempty"`
	IAS                BravoProfile `yaml:"ias,omitempty"`
	REV                BravoProfile `yaml:"rev,omitempty"`
	AP_STATE           BravoProfile `yaml:"ap_state,omitempty"`
	GEAR               BravoProfile `yaml:"gear,omitempty"`
	RETRACTABLE_GEAR   BravoProfile `yaml:"retractable_gear,omitempty"`
	MASTER_WARN        BravoProfile `yaml:"master_warn,omitempty"`
	MASTER_CAUTION     BravoProfile `yaml:"master_caution,omitempty"`
	FIRE               BravoProfile `yaml:"fire,omitempty"`
	OIL_LOW_PRESSURE   BravoProfile `yaml:"oil_low_pressure,omitempty"`
	FUEL_LOW_PRESSURE  BravoProfile `yaml:"fuel_low_pressure,omitempty"`
	ANTI_ICE           BravoProfile `yaml:"anti_ice,omitempty"`
	ENG_STARTER        BravoProfile `yaml:"eng_starter,omitempty"`
	APU                BravoProfile `yaml:"apu,omitempty"`
	VACUUM             BravoProfile `yaml:"vacuum,omitempty"`
	HYDRO_LOW_PRESSURE BravoProfile `yaml:"hydro_low_pressure,omitempty"`
	AUX_FUEL_PUMP      BravoProfile `yaml:"aux_fuel_pump,omitempty"`
	PARKING_BRAKE      BravoProfile `yaml:"parking_brake,omitempty"`
	VOLT_LOW           BravoProfile `yaml:"volt_low,omitempty"`
	DOORS              BravoProfile `yaml:"doors,omitempty"`
}
