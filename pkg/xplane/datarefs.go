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

func (s *xplaneService) setupDataRefs(airplaneICAO string) {
	s.Logger.Infof("Setup Datarefs for: %s", airplaneICAO)

	s.Logger.Infof("Loading BravoProfile for: %s", airplaneICAO)
	var planeProfile pkg.Profile
	planeProfile, err := s.loadProfile(airplaneICAO)
	if err != nil {
		s.Logger.Errorf("Error loading BravoProfile: %v", err)
		s.Logger.Infof("Loading defalt BravoProfile for: %s", airplaneICAO)
		planeProfile, err = s.loadProfile("default")
		if err != nil {
			s.Logger.Errorf("Error loading default BravoProfile: %v", err)
			return
		}
	}
	if planeProfile.Metadata == nil {
		planeProfile.Metadata = &pkg.Metadata{}
	}
	if planeProfile.Data == nil {
		planeProfile.Data = &pkg.Data{}
	}
	if planeProfile.Knobs == nil {
		planeProfile.Knobs = &pkg.Knobs{}
	}
	if planeProfile.Leds == nil {
		planeProfile.Leds = &pkg.Leds{}
	}
	err = s.CompileRules(&planeProfile)
	if err != nil {
		s.Logger.Errorf("Error compiling rules: %v", err)
		s.profile = nil
		return
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

func (s *xplaneService) loadProfile(airplaneICAO string) (pkg.Profile, error) {
	// load datarefs for the airplane from csv
	csvFilePath := path.Join(s.pluginPath, "profiles", fmt.Sprintf("%s.yaml", airplaneICAO))
	s.Logger.Infof("Loading datarefs from: %s", csvFilePath)
	f, err := os.ReadFile(csvFilePath)
	if err != nil {
		s.Logger.Errorf("Error opening file: %v", err)
		return pkg.Profile{}, err
	}
	var res pkg.Profile
	err = yaml.Unmarshal(f, &res)
	if err != nil {
		s.Logger.Errorf("Error reading file: %v", err)
		return pkg.Profile{}, err
	}
	return res, nil
}

func (s *xplaneService) CompileRules(p *pkg.Profile) error {
	err := s.compileRules(p.Leds, p.Data)
	if err != nil {
		s.Logger.Errorf("Error compiling rules for leds: %v", err)
		return err
	}
	s.Logger.Infof("Rules compiled successfully")
	return nil
}

func (s *xplaneService) updateLeds() {

	// special case for bus voltage
	dataref := s.profile.Data.BUS_VOLTAGE.Datarefs[0]
	output, err := expr.Run(dataref.Expr, dataref.Env)
	if err != nil {
		s.Logger.Errorf("BUS_VOLTAGE - Error running expression: %v", err)
	}
	if !output.(bool) {
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
		fieldValue, ok := fieldVal.Interface().(pkg.BravoProfile)
		if !ok {
			s.Logger.Errorf("Field %s is not of type BravoProfile", fieldName)
			continue
		}

		if fieldValue.Datarefs == nil && fieldValue.Commands == nil {
			s.Logger.Debugf("No datarefs found for: %s", fieldName)
			continue
		}

		if fieldName == "GEAR" {
			// special case for gear
			retractable_gear_dataref := s.profile.Data.RETRACTABLE_GEAR.Datarefs[0]
			retractable_gear_output, retractable_gear_err := expr.Run(retractable_gear_dataref.Expr, retractable_gear_dataref.Env)
			if retractable_gear_err != nil {
				s.Logger.Errorf("GEAR - Error running retractable_gear expression: %v", retractable_gear_err)
				continue
			}

			if retractable_gear_output.(bool) {
				dataref := s.profile.Leds.GEAR.Datarefs[0]
				output := dataAccess.GetFloatArrayData(dataref.Dataref.(dataAccess.DataRef))
				s.updateGearLEDs(output)
			} else {
				s.updateGearLEDs([]float32{0, 0, 0})
			}
			continue
		}

		var result bool
		if fieldValue.Condition == "any" {
			result = false
		} else {
			result = true
		}
		for _, dataref := range fieldValue.Datarefs {
			if dataref.Expr == nil {
				continue
			}
			output, err := expr.Run(dataref.Expr, dataref.Env)
			if err != nil {
				s.Logger.Errorf("Error running expression: %v", err)
				result = false
				break
			}
			if fieldValue.Condition == "any" {
				result = result || output.(bool)
			} else {
				// all or nothing (single value)
				result = result && output.(bool)
			}
		}
		if result {
			fieldValue.On()
		} else {
			fieldValue.Off()
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

func (s *xplaneService) compileRules(l *pkg.Leds, d *pkg.Data) error {
	var vals []reflect.Value
	vals = append(vals, reflect.ValueOf(l).Elem(), reflect.ValueOf(d).Elem())
	for index, _ := range vals {
		val := vals[index]
		typ := val.Type()
		s.Logger.Infof("Compiling rules for: %s", typ.Name())
		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i) // Get the field metadata
			fieldName := field.Name
			s.Logger.Infof("-- Compiling rules for: %s", fieldName)
			// Get the field value as a reflect.Value
			fieldVal := val.Field(i)
			// Perform type assertion to BravoProfile
			fieldValue, ok := fieldVal.Interface().(pkg.BravoProfile)
			if !ok {
				s.Logger.Errorf("Field %s is not of type BravoProfile", fieldName)
				return fmt.Errorf("Field %s is not of type BravoProfile", fieldName)
				continue
			}

			// Modify the fieldValue
			if fieldValue.Datarefs != nil {
				for j := range fieldValue.Datarefs {
					dataref := &fieldValue.Datarefs[j]
					s.Logger.Infof("---- Compiling dataref: %s", dataref.DatarefStr)
					// Get a pointer to the actual element
					myDataref, found := dataAccess.FindDataRef(dataref.DatarefStr)
					if !found {
						s.Logger.Errorf("Dataref not found: %s", dataref.DatarefStr)
						continue
					}
					dataref.Dataref = myDataref

					datarefType := dataAccess.GetDataRefTypes(myDataref)

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
						s.Logger.Errorf("Dataref type not supported: %v", datarefType)
					}

					s.Logger.Infof("---- Compiling expression: %s - %s: %s", code, fieldName, dataref.DatarefStr)
					env := map[string]interface{}{
						"GetFloatData":      dataAccess.GetFloatData,
						"GetIntData":        dataAccess.GetIntData,
						"GetFloatArrayData": dataAccess.GetFloatArrayData,
						"GetIntArrayData":   dataAccess.GetIntArrayData,
						"myDataref":         myDataref,
					}
					program, err := expr.Compile(code, expr.Env(env))
					if err != nil {
						s.Logger.Errorf("Error compiling expression: %v", err)
						return err
					}
					dataref.Expr = program
					dataref.Env = env
				}
				fieldValue.On, fieldValue.Off = s.assignOnAndOffFuncs(fieldName)
			}

			// Assign the modified value back to the struct field
			fieldVal.Set(reflect.ValueOf(fieldValue))
			s.Logger.Infof("-- Rules compiled successfully for: %s", fieldName)
		}
		s.Logger.Infof("Rules compiled successfully: %s", typ.Name())
	}

	return nil
}
