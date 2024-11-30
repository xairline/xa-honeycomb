package pkg

import "github.com/expr-lang/expr/vm"

type Command struct {
	CommandStr string `yaml:"command_str,omitempty" json:"command_str,omitempty"`
	Command    interface{}
}

type Dataref struct {
	DatarefStr string `yaml:"dataref_str,omitempty" json:"dataref_str,omitempty"`
	Dataref    interface{}
	Index      int `yaml:"index,omitempty" json:"index,omitempty"`
}

type DatarefCondition struct {
	DatarefStr string `yaml:"dataref_str,omitempty" json:"dataref_str,omitempty"`
	Dataref    interface{}
	Index      int     `yaml:"index,omitempty" json:"index,omitempty"`
	Operator   string  `yaml:"operator,omitempty" json:"operator,omitempty"`
	Threshold  float32 `yaml:"threshold,omitempty" json:"threshold,omitempty"`
	Expr       *vm.Program
	Env        map[string]interface{}
}

type Metadata struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	Enabled     bool   `yaml:"enabled,omitempty" json:"enabled,omitempty"`
}

type ConditionProfile struct {
	Datarefs  []DatarefCondition `yaml:"datarefs,omitempty" json:"datarefs,omitempty"`
	Condition string             `yaml:"condition,omitempty" json:"condition,omitempty"`
}

type DatarefProfile struct {
	Datarefs []Dataref `yaml:"datarefs,omitempty" json:"datarefs,omitempty"`
}

type LEDProfile struct {
	ConditionProfile `yaml:",inline"`
	On               func() `json:"-"`
	Off              func() `json:"-"`
}

type DataProfile struct {
	DatarefProfile `yaml:",inline"`
	Value          *float32 `yaml:"value,omitempty" json:"value,omitempty"`
}

type KnobProfile struct {
	DatarefProfile `yaml:",inline"`
	Commands       []Command `yaml:"commands,omitempty" json:"commands,omitempty"`
}

type Knobs struct {
	AP_HDG KnobProfile `yaml:"ap_hdg,omitempty" json:"ap_hdg,omitempty"`
	AP_VS  KnobProfile `yaml:"ap_vs,omitempty" json:"ap_vs,omitempty"`
	AP_ALT KnobProfile `yaml:"ap_alt,omitempty" json:"ap_alt,omitempty"`
	AP_IAS KnobProfile `yaml:"ap_ias,omitempty" json:"ap_ias,omitempty"`
	AP_CRS KnobProfile `yaml:"ap_crs,omitempty" json:"ap_crs,omitempty"`
}

type Leds struct {
	HDG                LEDProfile `yaml:"hdg,omitempty" json:"hdg,omitempty"`
	NAV                LEDProfile `yaml:"nav,omitempty" json:"nav,omitempty"`
	ALT                LEDProfile `yaml:"alt,omitempty" json:"alt,omitempty"`
	APR                LEDProfile `yaml:"apr,omitempty" json:"apr,omitempty"`
	VS                 LEDProfile `yaml:"vs,omitempty" json:"vs,omitempty"`
	AP                 LEDProfile `yaml:"ap,omitempty" json:"ap,omitempty"`
	IAS                LEDProfile `yaml:"ias,omitempty" json:"ias,omitempty"`
	REV                LEDProfile `yaml:"rev,omitempty" json:"rev,omitempty"`
	GEAR               LEDProfile `yaml:"gear,omitempty" json:"gear,omitempty"`
	MASTER_WARN        LEDProfile `yaml:"master_warn,omitempty" json:"master_warn,omitempty"`
	MASTER_CAUTION     LEDProfile `yaml:"master_caution,omitempty" json:"master_caution,omitempty"`
	FIRE               LEDProfile `yaml:"fire,omitempty" json:"fire,omitempty"`
	OIL_LOW_PRESSURE   LEDProfile `yaml:"oil_low_pressure,omitempty" json:"oil_low_pressure,omitempty"`
	FUEL_LOW_PRESSURE  LEDProfile `yaml:"fuel_low_pressure,omitempty" json:"fuel_low_pressure,omitempty"`
	ANTI_ICE           LEDProfile `yaml:"anti_ice,omitempty" json:"anti_ice,omitempty"`
	ENG_STARTER        LEDProfile `yaml:"eng_starter,omitempty" json:"eng_starter,omitempty"`
	APU                LEDProfile `yaml:"apu,omitempty" json:"apu,omitempty"`
	VACUUM             LEDProfile `yaml:"vacuum,omitempty" json:"vacuum,omitempty"`
	HYDRO_LOW_PRESSURE LEDProfile `yaml:"hydro_low_pressure,omitempty" json:"hydro_low_pressure,omitempty"`
	AUX_FUEL_PUMP      LEDProfile `yaml:"aux_fuel_pump,omitempty" json:"aux_fuel_pump,omitempty"`
	PARKING_BRAKE      LEDProfile `yaml:"parking_brake,omitempty" json:"parking_brake,omitempty"`
	VOLT_LOW           LEDProfile `yaml:"volt_low,omitempty" json:"volt_low,omitempty"`
	DOORS              LEDProfile `yaml:"doors,omitempty" json:"doors,omitempty"`
}

type Data struct {
	AP_STATE    DataProfile `yaml:"ap_state,omitempty" json:"ap_state,omitempty"`
	AP_ALT_STEP DataProfile `yaml:"ap_alt_step,omitempty" json:"ap_alt_step,omitempty"`
	AP_VS_STEP  DataProfile `yaml:"ap_vs_step,omitempty" json:"ap_vs_step,omitempty"`
	AP_IAS_STEP DataProfile `yaml:"ap_ias_step,omitempty" json:"ap_ias_step,omitempty"`
}

type Conditions struct {
	BUS_VOLTAGE      ConditionProfile `yaml:"bus_voltage,omitempty" json:"bus_voltage,omitempty"`
	RETRACTABLE_GEAR ConditionProfile `yaml:"retractable_gear,omitempty" json:"retractable_gear,omitempty"`
}

type Profile struct {
	Metadata   *Metadata   `yaml:"metadata" json:"metadata"`
	Knobs      *Knobs      `yaml:"knobs,omitempty" json:"knobs,omitempty"`
	Leds       *Leds       `yaml:"leds,omitempty" json:"leds,omitempty"`
	Data       *Data       `yaml:"data,omitempty" json:"data,omitempty"`
	Conditions *Conditions `yaml:"conditions,omitempty" json:"conditions,omitempty"`
}
