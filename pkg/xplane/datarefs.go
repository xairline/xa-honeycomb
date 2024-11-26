package xplane

import (
	"fmt"
	"github.com/stretchr/testify/assert/yaml"
	"github.com/xairline/goplane/xplm/dataAccess"
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
	"os"
	"path"
	"reflect"
	"strings"
)

func (s *xplaneService) setupDataRefs(airplaneICAO string) {
	s.Logger.Infof("Setup Datarefs for: %s", airplaneICAO)

	s.Logger.Infof("Loading defalt profile for: %s", airplaneICAO)
	defaultProfile := s.loadProfile("default")
	s.compileRules(&defaultProfile)

	//s.Logger.Infof("Loading profile for: %s", airplaneICAO)
	//records := s.loadProfile(airplaneICAO)
	//rules := s.compileRules(records)
	//
	//// merge default and airplane specific records
	//for name, led := range rules {
	//	defaultRules[name] = led
	//	s.Logger.Debugf("Replace record: %s", name)
	//}
	//s.leds = defaultRules
	//s.datarefs = make(map[string][]dataAccess.DataRef)
}

func (s *xplaneService) assignOnAndOffFuncs(name string) (func(), func()) {
	switch name {
	case "APR":
		return honeycomb.OnLEDAPR, honeycomb.OffLEDAPR
	case "ALT":
		return honeycomb.OnLEDAlt, honeycomb.OffLEDAlt
	case "VS":
		return honeycomb.OnLEDVS, honeycomb.OffLEDVS
	case "HDG":
		return honeycomb.OnLEDHeading, honeycomb.OffLEDHeading
	case "NAV":
		return honeycomb.OnLEDNav, honeycomb.OffLEDNav
	case "REV":
		return honeycomb.OnLEDREV, honeycomb.OffLEDREV
	case "IAS":
		return honeycomb.OnLEDIAS, honeycomb.OffLEDIAS
	case "AP":
		return honeycomb.OnLEDAP, honeycomb.OffLEDAP
	case "BUS_VOLTAGE":
		return honeycomb.OnLEDLowVolts, honeycomb.OffLEDLowVolts
	case "GEAR":
		return honeycomb.OnLedGearGreen, honeycomb.OnLedGearRed
	case "MASTER_WARN":
		return honeycomb.OnLEDMasterWarning, honeycomb.OffLEDMasterWarning
	case "MASTER_CAUTION":
		return honeycomb.OnLEDMasterCaution, honeycomb.OffLEDMasterCaution
	case "FIRE":
		return honeycomb.OnLEDEngineFire, honeycomb.OffLEDEngineFire
	case "VOLT_LOW":
		return honeycomb.OnLEDLowVolts, honeycomb.OffLEDLowVolts
	case "OIL_LOW_P":
		return honeycomb.OnLEDLowOilPress, honeycomb.OffLEDLowOilPress
	case "FUEL_LOW_P":
		return honeycomb.OnLEDLowFuelPress, honeycomb.OffLEDLowFuelPress
	case "ANTI_ICE":
		return honeycomb.OnLEDAntiIce, honeycomb.OffLEDAntiIce
	case "ENG_STARTER":
		return honeycomb.OnLEDStarter, honeycomb.OffLEDStarter
	case "APU":
		return honeycomb.OnLEDApu, honeycomb.OffLEDApu
	case "VACUUM":
		return honeycomb.OnLEDVacuum, honeycomb.OffLEDVacuum
	case "HYDRO_LOW_P":
		return honeycomb.OnLEDLowHydPress, honeycomb.OffLEDLowHydPress
	case "PARKING_BRAKE":
		return honeycomb.OnLEDParkingBrake, honeycomb.OffLEDParkingBrake
	case "DOORS":
		return honeycomb.OnLEDDoor, honeycomb.OffLEDDoor
	case "AUX_FUEL_PUMP":
		return honeycomb.OnLEDFuelPump, honeycomb.OffLEDFuelPump
	default:
		s.Logger.Warningf("No on/off functions found for: %s", name)
		return nil, nil
	}
}

func (s *xplaneService) loadProfile(airplaneICAO string) Profile {
	// load datarefs for the airplane from csv
	csvFilePath := path.Join(s.pluginPath, "profiles", fmt.Sprintf("%s.yaml", airplaneICAO))
	s.Logger.Debugf("Loading datarefs from: %s", csvFilePath)
	f, err := os.ReadFile(csvFilePath)
	if err != nil {
		s.Logger.Errorf("Error opening file: %v", err)
	}
	var res Profile
	err = yaml.Unmarshal(f, &res)
	if err != nil {
		s.Logger.Errorf("Error reading file: %v", err)
	}
	return res
}

func (s *xplaneService) compileRules(p *Profile) {
	val := reflect.ValueOf(p).Elem() // Get the actual struct value
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i) // Get the field metadata
		fieldName := field.Name

		// Get the field value as a reflect.Value
		fieldVal := val.Field(i)

		// Perform type assertion to profile
		fieldValue, ok := fieldVal.Interface().(profile)
		if !ok {
			s.Logger.Errorf("Field %s is not of type profile", fieldName)
			continue
		}

		// Modify the fieldValue
		switch fieldValue.ProfileType {
		case "dataref":
			for j := range fieldValue.Datarefs {
				dataref := &fieldValue.Datarefs[j] // Get a pointer to the actual element
				myDataref, found := dataAccess.FindDataRef(dataref.Dataref_str)
				if !found {
					s.Logger.Errorf("Dataref not found: %s", dataref.Dataref_str)
				}
				dataref.Dataref = myDataref
				// TODO: add expr eval here
			}
			fieldValue.on, fieldValue.off = s.assignOnAndOffFuncs(fieldName)
		case "data":
			for j := range fieldValue.Data {
				data := &fieldValue.Data[j] // Get a pointer to the actual element
				myDataref, found := dataAccess.FindDataRef(data.Dataref_str)
				if !found {
					s.Logger.Errorf("Dataref not found: %s", data.Dataref_str)
				}
				data.Dataref = myDataref
			}
		}

		// Assign the modified value back to the struct field
		fieldVal.Set(reflect.ValueOf(fieldValue))
	}
	s.Logger.Debugf("Compiled rules: %+v", p)
}

func (s *xplaneService) updateLeds() {
	for name, led := range s.leds {
		if s.evaluateRules(name, led.rules) {
			led.on()
		} else {
			led.off()
		}
	}
}

func (s *xplaneService) evaluateRules(name, rules string) bool {
	rules_parsed := strings.Split(rules, ",")
	var rules_expr []string
	var rules_operators string
	if len(rules_parsed) >= 3 {
		rules_expr = rules_parsed[0 : len(rules_parsed)-2]
		rules_operators = rules_parsed[len(rules_parsed)-1]
	} else {
		rules_expr = append(rules_expr, rules_parsed[0])
		rules_operators = rules_parsed[1]
	}

	if s.datarefs[name] == nil {
		s.datarefs[name] = make([]dataAccess.DataRef, len(rules_parsed)-1)
		for _, rule := range rules_expr {
			dataref_str := strings.Split(rule, ":")[0]
			dr, found := dataAccess.FindDataRef(dataref_str)
			if !found {
				s.Logger.Errorf("Dataref not found: %s", rules)
				return false
			}
			s.datarefs[name] = append(s.datarefs[name], dr)
		}

	}
	s.Logger.Debugf("Evaluating rules: %+v, opeartor: %s", s.datarefs, rules_operators)

	return true
}
