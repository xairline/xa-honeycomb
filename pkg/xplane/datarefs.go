package xplane

import (
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/stretchr/testify/assert/yaml"
	"github.com/xairline/goplane/xplm/dataAccess"
	"github.com/xairline/xa-honeycomb/pkg"
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
	"os"
	"path"
	"reflect"
)

func (s *xplaneService) tryLoadProfile() {
	// Try to load the profile using the aircraft's UI name
	aircraftNameDrf, found := dataAccess.FindDataRef("sim/aircraft/view/acf_ui_name")
	if found {
		aircraftName := dataAccess.GetString(aircraftNameDrf)

		planeProfile, err := s.loadProfile(aircraftName)
		if err == nil {
			s.Logger.Infof("Loading BravoProfile for: %s", aircraftName)
			s.setupDataRefs(planeProfile)
			return
		} else {
			s.Logger.Errorf("Cannot loading BravoProfile for %s: %v", aircraftName, err)
		}
	}

	// Try to load the profile using the aircraft's ICAO
	aircraftIACODrf, found := dataAccess.FindDataRef("sim/aircraft/view/acf_ICAO")
	if found {
		aircraftIACO := dataAccess.GetString(aircraftIACODrf)

		planeProfile, err := s.loadProfile(aircraftIACO)
		if err == nil {
			s.Logger.Infof("Loading BravoProfile for: %s", aircraftIACO)
			s.setupDataRefs(planeProfile)
			return
		} else {
			s.Logger.Errorf("Cannot loading BravoProfile for %s: %v", aircraftIACO, err)
		}
	}

	s.Logger.Infof("Loading default BravoProfile")
	planeProfile, err := s.loadProfile("default")
	if err == nil {
		s.setupDataRefs(planeProfile)
	} else {
		s.Logger.Errorf("Error loading default BravoProfile: %v", err)
	}
}

func (s *xplaneService) setupDataRefs(planeProfile pkg.Profile) {
	// Fill in any missing sections of the profile
	if planeProfile.Metadata == nil {
		planeProfile.Metadata = &pkg.Metadata{}
	}
	if planeProfile.Data == nil {
		planeProfile.Data = &pkg.Data{}
	}
	if planeProfile.Conditions == nil {
		planeProfile.Conditions = &pkg.Conditions{}
	}
	if planeProfile.Knobs == nil {
		planeProfile.Knobs = &pkg.Knobs{}
	}
	if planeProfile.Leds == nil {
		planeProfile.Leds = &pkg.Leds{}
	}

	var err error
	hasErrors := false

	s.Logger.Infof("Loading LEDs")
	err = rangeStruct(planeProfile.Leds, s.loadProfileElement)
	if err != nil {
		s.Logger.Errorf("Error loading LEDs: %v", err)
		hasErrors = true
	}

	s.Logger.Infof("Loading Datas")
	err = rangeStruct(planeProfile.Data, s.loadProfileElement)
	if err != nil {
		s.Logger.Errorf("Error loading Datas: %v", err)
		hasErrors = true
	}

	s.Logger.Infof("Loading Knobs")
	err = rangeStruct(planeProfile.Knobs, s.loadProfileElement)
	if err != nil {
		s.Logger.Errorf("Error loading Knobs: %v", err)
		hasErrors = true
	}

	s.Logger.Infof("Loading Conditions")
	err = rangeStruct(planeProfile.Conditions, s.loadProfileElement)
	if err != nil {
		s.Logger.Errorf("Error loading Conditions: %v", err)
		hasErrors = true
	}

	if hasErrors {
		s.Logger.Infof("Loaded profile with errors")
	} else {
		s.Logger.Infof("Successfully loaded profile")
	}
	s.profile = &planeProfile
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
		return func() {
			return
		}, honeycomb.AllOff
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
	case "OIL_LOW_PRESSURE":
		return honeycomb.OnLEDLowOilPress, honeycomb.OffLEDLowOilPress
	case "FUEL_LOW_PRESSURE":
		return honeycomb.OnLEDLowFuelPress, honeycomb.OffLEDLowFuelPress
	case "ANTI_ICE":
		return honeycomb.OnLEDAntiIce, honeycomb.OffLEDAntiIce
	case "ENG_STARTER":
		return honeycomb.OnLEDStarter, honeycomb.OffLEDStarter
	case "APU":
		return honeycomb.OnLEDApu, honeycomb.OffLEDApu
	case "VACUUM":
		return honeycomb.OnLEDVacuum, honeycomb.OffLEDVacuum
	case "HYDRO_LOW_PRESSURE":
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

func (s *xplaneService) loadProfile(airplaneConfig string) (pkg.Profile, error) {
	// load datarefs for the airplane from YAML
	configFilePath := path.Join(s.pluginPath, "profiles", fmt.Sprintf("%s.yaml", airplaneConfig))
	f, err := os.ReadFile(configFilePath)
	if err != nil {
		return pkg.Profile{}, err
	}
	s.Logger.Infof("Loading datarefs from: %s", configFilePath)
	var res pkg.Profile
	err = yaml.Unmarshal(f, &res)
	if err != nil {
		return pkg.Profile{}, err
	}
	return res, nil
}

func (s *xplaneService) updateLeds() {
	if s.profile == nil {
		return
	}

	// special case for bus voltage
	busVoltage, busVoltageOK := s.evaluateCondition(&s.profile.Conditions.BUS_VOLTAGE)
	if busVoltageOK && !busVoltage {
		honeycomb.AllOff()
		return
	}

	val := reflect.ValueOf(s.profile.Leds).Elem() // Get the actual struct value
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i) // Get the field metadata
		fieldName := field.Name
		// Get the field value as a reflect.Value
		fieldVal := val.Field(i)
		// Perform type assertion to BravoProfile
		fieldValue, ok := fieldVal.Interface().(pkg.LEDProfile)
		if !ok {
			s.Logger.Errorf("Field %s is not of type LEDProfile", fieldName)
			continue
		}

		if fieldValue.Datarefs == nil {
			continue
		}

		if fieldName == "GEAR" {
			// special case for gear
			retractableGear, retractableGearOK := s.evaluateCondition(&s.profile.Conditions.RETRACTABLE_GEAR)
			if retractableGearOK && !retractableGear {
				s.updateGearLEDs([]float32{0, 0, 0})
				continue
			}

			dataref := s.profile.Leds.GEAR.Datarefs[0]
			if dataref.Dataref != nil {
				output := dataAccess.GetFloatArrayData(dataref.Dataref.(dataAccess.DataRef))
				s.updateGearLEDs(output)
			}
			continue
		}

		result, resultOK := s.evaluateCondition(&fieldValue.ConditionProfile)
		if resultOK {
			if result {
				fieldValue.On()
			} else {
				fieldValue.Off()
			}
		}
	}
}

func (s *xplaneService) updateGearLEDs(output []float32) {
	if output[0] >= 0.99 {
		honeycomb.OnLEDNoseGearGreen()
		honeycomb.OffLEDNoseGearRed()
	}
	if output[1] >= 0.99 {
		honeycomb.OnLEDLeftGearGreen()
		honeycomb.OffLEDLeftGearRed()
	}
	if output[2] >= 0.99 {
		honeycomb.OnLEDRightGearGreen()
		honeycomb.OffLEDRightGearRed()
	}

	if output[0] <= 0.01 {
		honeycomb.OffLEDNoseGearGreen()
		honeycomb.OffLEDNoseGearRed()
	}
	if output[1] <= 0.01 {
		honeycomb.OffLEDLeftGearGreen()
		honeycomb.OffLEDLeftGearRed()
	}
	if output[2] <= 0.01 {
		honeycomb.OffLEDRightGearGreen()
		honeycomb.OffLEDRightGearRed()
	}

	if output[0] > 0.01 && output[0] < 0.99 {
		honeycomb.OffLEDNoseGearGreen()
		honeycomb.OnLEDNoseGearRed()
	}
	if output[1] > 0.01 && output[1] < 0.99 {
		honeycomb.OffLEDLeftGearGreen()
		honeycomb.OnLEDLeftGearRed()
	}
	if output[2] > 0.01 && output[2] < 0.99 {
		honeycomb.OffLEDRightGearGreen()
		honeycomb.OnLEDRightGearRed()
	}
}

func rangeStruct(s interface{}, modify func(name string, value interface{}) (interface{}, error)) error {
	v := reflect.ValueOf(s)

	// Dereference to get the underlying struct
	v = v.Elem()
	t := v.Type()

	// Iterate over the fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		value := field.Interface()
		newValue, err := modify(fieldType.Name, value)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(newValue))
	}
	return nil
}

func (s *xplaneService) loadProfileElement(fieldName string, value interface{}) (interface{}, error) {
	dataProfileValue, ok := value.(pkg.DataProfile)
	if ok {
		s.Logger.Infof("-- Loading Data: %s", fieldName)
		err := s.loadDatarefProfile(fieldName, &dataProfileValue.DatarefProfile)
		return dataProfileValue, err
	}

	knobProfileValue, ok := value.(pkg.KnobProfile)
	if ok {
		s.Logger.Infof("-- Loading Knob: %s", fieldName)
		err := s.loadDatarefProfile(fieldName, &knobProfileValue.DatarefProfile)
		return knobProfileValue, err
	}

	conditionProfileValue, ok := value.(pkg.ConditionProfile)
	if ok {
		s.Logger.Infof("-- Loading Condition: %s", fieldName)
		err := s.loadConditionProfile(fieldName, &conditionProfileValue)
		return conditionProfileValue, err
	}

	ledProfileValue, ok := value.(pkg.LEDProfile)
	if ok {
		s.Logger.Infof("-- Loading LED: %s", fieldName)
		err := s.loadLedProfile(fieldName, &ledProfileValue)
		return ledProfileValue, err
	}

	return value, fmt.Errorf("Field %s is not of a known type", fieldName)
}

func (s *xplaneService) loadDatarefProfile(fieldName string, fieldValue *pkg.DatarefProfile) error {
	if fieldValue.Datarefs != nil {
		for j := range fieldValue.Datarefs {
			dataref := &fieldValue.Datarefs[j]
			dataref.Dataref = s.getDataref(dataref.DatarefStr)
		}
	} else {
		s.Logger.Infof("---- No datarefs specified for %s", fieldName)
	}
	return nil
}

func (s *xplaneService) loadConditionProfile(fieldName string, fieldValue *pkg.ConditionProfile) error {
	if fieldValue.Datarefs == nil {
		s.Logger.Infof("---- No datarefs specified")
		return nil
	}

	for j := range fieldValue.Datarefs {
		dataref := &fieldValue.Datarefs[j]
		myDataref := s.getDataref(dataref.DatarefStr)

		if myDataref == nil {
			continue
		}

		dataref.Dataref = myDataref
		datarefType := dataAccess.GetDataRefTypes(myDataref)

		if dataref.Operator != "" {
			if !isOperatorSupported(dataref.Operator) {
				return fmt.Errorf("Unsupported operator found: %s", dataref.Operator)
			}

			var code string
			switch datarefType {
			case dataAccess.TypeFloat:
				code = fmt.Sprintf("GetFloatData(myDataref) %s %f", dataref.Operator, dataref.Threshold)
			case dataAccess.TypeInt:
				code = fmt.Sprintf("GetIntData(myDataref) %s %d", dataref.Operator, int(dataref.Threshold))
			case dataAccess.TypeFloatArray:
				code = fmt.Sprintf("GetFloatArrayData(myDataref)[%d] %s %f", dataref.Index, dataref.Operator, dataref.Threshold)
			case dataAccess.TypeIntArray:
				code = fmt.Sprintf("GetIntArrayData(myDataref)[%d] %s %d", dataref.Index, dataref.Operator, int(dataref.Threshold))
			default:
				return fmt.Errorf("Dataref type not supported: %v", datarefType)
			}

			s.Logger.Infof("---- Compiling expression: %s - %s[%d]: %s", code, fieldName, j, dataref.DatarefStr)
			env := map[string]interface{}{
				"GetFloatData":      dataAccess.GetFloatData,
				"GetIntData":        dataAccess.GetIntData,
				"GetFloatArrayData": dataAccess.GetFloatArrayData,
				"GetIntArrayData":   dataAccess.GetIntArrayData,
				"myDataref":         myDataref,
			}
			program, err := expr.Compile(code, expr.Env(env))
			if err != nil {
				return fmt.Errorf("Error compiling expression: %v", err)
			}
			dataref.Expr = program
			dataref.Env = env
		} else {
			return fmt.Errorf("Condition missing operator: %s", fieldName)
		}
	}

	s.Logger.Infof("-- Rules compiled successfully for: %s", fieldName)
	return nil
}

func (s *xplaneService) loadLedProfile(fieldName string, fieldValue *pkg.LEDProfile) error {
	err := s.loadConditionProfile(fieldName, &fieldValue.ConditionProfile)

	fieldValue.On, fieldValue.Off = s.assignOnAndOffFuncs(fieldName)
	return err
}

func (s *xplaneService) getDataref(datarefStr string) dataAccess.DataRef {
	s.Logger.Infof("---- Finding dataref: %s", datarefStr)
	// Get a pointer to the actual element
	myDataref, found := dataAccess.FindDataRef(datarefStr)
	if !found {
		s.Logger.Errorf("Dataref not found: %s", datarefStr)
		return nil
	}

	return myDataref
}

// Check whether the given expression operator is allowed in our boolean expressions
// This prevents arbitrary code execution from user input.
func isOperatorSupported(operator string) bool {
	return operator == "==" ||
		operator == ">" ||
		operator == "<" ||
		operator == ">=" ||
		operator == "<=" ||
		operator == "!="
}

// Evaluate a condition
// Returns:
// 1. The result of the condition evaluation
// 2. Whether the condition was valid (if false then it should be ignored)
func (s *xplaneService) evaluateCondition(condition *pkg.ConditionProfile) (bool, bool) {
	var valid = false
	var result bool
	if condition.Condition == "any" {
		result = false
	} else {
		result = true
	}
	for _, dataref := range condition.Datarefs {
		if dataref.Expr == nil {
			continue
		}
		output, err := expr.Run(dataref.Expr, dataref.Env)
		if err != nil {
			s.Logger.Errorf("Error running expression: %v", err)
			continue
		}
		if condition.Condition == "any" {
			result = result || output.(bool)
		} else {
			// all or nothing (single value)
			result = result && output.(bool)
		}
		valid = true
	}
	return result, valid
}

// Extract a value from the given data profile
// Returns:
// 1. The value if found, or 0.0
// 2. Whether a value was found or not
func (s *xplaneService) dataValue(bp *pkg.DataProfile) (float64, bool) {
	if len(bp.Datarefs) > 0 {
		// TODO support something like "condition" that can aggregate multiple datarefs or array datarefs
		// e.g. "max" or "min" or "sum" or "avg"
		myDataref := bp.Datarefs[0]
		if myDataref.Dataref == nil {
			return 0.0, false
		}
		datarefType := dataAccess.GetDataRefTypes(myDataref.Dataref.(dataAccess.DataRef))
		switch datarefType {
		case dataAccess.TypeFloat:
			return float64(dataAccess.GetFloatData(myDataref.Dataref.(dataAccess.DataRef))), true
		case dataAccess.TypeInt:
			return float64(dataAccess.GetIntData(myDataref.Dataref.(dataAccess.DataRef))), true
		case dataAccess.TypeFloatArray:
			return float64(dataAccess.GetFloatArrayData(myDataref.Dataref.(dataAccess.DataRef))[0]), true
		case dataAccess.TypeIntArray:
			return float64(dataAccess.GetIntArrayData(myDataref.Dataref.(dataAccess.DataRef))[0]), true
		default:
			s.Logger.Errorf("Dataref type not supported: %v", datarefType)
			return 0.0, false
		}
	} else if bp.Value != nil {
		return float64(*bp.Value), true
	} else {
		return 0.0, false
	}
}
