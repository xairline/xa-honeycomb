package xplane

import (
	"github.com/xairline/goplane/xplm/dataAccess"
	"github.com/xairline/xa-honeycomb/pkg"
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
	"reflect"
)

// flightloop, high freq code!
func (s *xplaneService) flightLoop(
	elapsedSinceLastCall,
	elapsedTimeSinceLastFlightLoop float32,
	counter int,
	ref interface{},
) float32 {

	if s.profile == nil {
		s.Logger.Info("Profile is nil, try to load it again")
		s.tryLoadProfile()
		s.lastCounter = 0
	}

	if counter-s.lastCounter > 200 {
		honeycomb.LED_STATE_CHANGED_LOCK.Lock()
		honeycomb.LED_STATE_CHANGED = true
		honeycomb.LED_STATE_CHANGED_LOCK.Unlock()
		s.lastCounter = counter
	}

	s.updateLeds()

	return 0.1
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
