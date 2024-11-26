package xplane

import (
	"encoding/csv"
	"fmt"
	"github.com/xairline/goplane/xplm/dataAccess"
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
	"os"
	"path"
	"strings"
)

// TODO: mess
func (s *xplaneService) setupDataRefs(airplaneICAO string) {
	s.Logger.Infof("Setup Datarefs for: %s", airplaneICAO)

	s.Logger.Infof("Loading defalt profile for: %s", airplaneICAO)
	defaultRecords := s.loadProfile("sample")
	defaultRules := s.compileRules(defaultRecords)

	s.Logger.Infof("Loading profile for: %s", airplaneICAO)
	records := s.loadProfile(airplaneICAO)
	rules := s.compileRules(records)

	// merge default and airplane specific records
	for name, led := range rules {
		defaultRules[name] = led
		s.Logger.Debugf("Replace record: %s", name)
	}
	s.leds = defaultRules
	s.datarefs = make(map[string][]dataAccess.DataRef)
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

func (s *xplaneService) loadProfile(airplaneICAO string) [][]string {
	// load datarefs for the airplane from csv
	csvFilePath := path.Join(s.pluginPath, "profiles", fmt.Sprintf("%s.csv", airplaneICAO))
	s.Logger.Debugf("Loading datarefs from: %s", csvFilePath)
	f, err := os.Open(csvFilePath)
	if err != nil {
		s.Logger.Errorf("Error opening file: %v", err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		s.Logger.Errorf("Error reading csv: %v", err)
	}
	return records
}

func (s *xplaneService) compileRules(records [][]string) map[string]leds {
	res := make(map[string]leds)
	for i, record := range records {
		if i == 0 {
			continue
		}
		name := record[0]
		onFunc, offFunc := s.assignOnAndOffFuncs(name)
		if onFunc == nil || offFunc == nil {
			s.Logger.Debugf("No on/off functions found for: %s", name)
			continue
		}
		dataref_strs := strings.Split(record[1], ";")
		rules_str := fmt.Sprintf("%s:%s:%s", dataref_strs[0], record[2], record[3])
		rules_str = strings.ReplaceAll(rules_str, " or ", " || ")
		rules_str = strings.ReplaceAll(rules_str, " and ", " && ")
		rules_str = strings.ReplaceAll(rules_str, " x", fmt.Sprintf(" %s", dataref_strs[0]))
		if len(dataref_strs) > 1 {
			for i, dataref_str := range dataref_strs {
				if i == 0 {
					continue
				}

				rules_str += "," + fmt.Sprintf("%s:%s:%s", dataref_str, record[2], record[3])
			}

		}
		my_operator := "&&"
		if record[4] == "any" {
			my_operator = "||"
		}
		rules_str = strings.ReplaceAll(rules_str, " || ", ",")
		rules_str = strings.ReplaceAll(rules_str, " && ", ",")
		rules_str += "," + my_operator
		res[name] = leds{
			rules: rules_str,
			on:    onFunc,
			off:   offFunc,
		}
	}
	return res
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
