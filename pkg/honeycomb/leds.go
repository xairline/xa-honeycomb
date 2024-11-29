package honeycomb

import "strings"

const (
	LED_FUEL_PUMP     = 1
	LED_PARKING_BRAKE = 2
	LED_LOW_VOLTS     = 4
	LED_DOOR          = 8

	LED_LOW_OIL_PRESS  = 1
	LED_LOW_FUEL_PRESS = 2
	LED_ANTI_ICE       = 4
	LED_STARTER        = 8
	LED_APU            = 16
	LED_MASTER_CAUTION = 32
	LED_VACUUM         = 64
	LED_LOW_HYD_PRESS  = 128

	LED_LEFT_GEAR_GREEN  = 1
	LED_LEFT_GEAR_RED    = 2
	LED_NOSE_GEAR_GREEN  = 4
	LED_NOSE_GEAR_RED    = 8
	LED_RIGHT_GEAR_GREEN = 16
	LED_RIGHT_GEAR_RED   = 32
	LED_MASTER_WARNING   = 64
	LED_ENGINE_FIRE      = 128

	LED_HEADING = 1
	LED_NAV     = 2
	LED_APR     = 4
	LED_REV     = 8
	LED_ALT     = 16
	LED_VS      = 32
	LED_IAS     = 64
	LED_AP      = 128
)

// Global variables for state
var (
	ANUNCIATOR_W2     byte = 0
	ANUNCIATOR_W1     byte = 0
	LANDING_GEAR_W    byte = 0
	AUTO_PILOT_W      byte = 0
	LED_STATE_CHANGED      = false
)

func setBit(val byte, bit byte) byte {
	return val | bit
}

func clearBit(val byte, bit byte) byte {
	return val &^ bit
}

func OnLEDFuelPump() {
	ANUNCIATOR_W2_BEFORE := ANUNCIATOR_W2
	ANUNCIATOR_W2 = setBit(ANUNCIATOR_W2, LED_FUEL_PUMP)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W2_BEFORE != ANUNCIATOR_W2
}

func OffLEDFuelPump() {
	ANUNCIATOR_W2_BEFORE := ANUNCIATOR_W2
	ANUNCIATOR_W2 = clearBit(ANUNCIATOR_W2, LED_FUEL_PUMP)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W2_BEFORE != ANUNCIATOR_W2
}

func OnLEDParkingBrake() {
	ANUNCIATOR_W2_BEFORE := ANUNCIATOR_W2
	ANUNCIATOR_W2 = setBit(ANUNCIATOR_W2, LED_PARKING_BRAKE)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W2_BEFORE != ANUNCIATOR_W2
}

func OffLEDParkingBrake() {
	ANUNCIATOR_W2_BEFORE := ANUNCIATOR_W2
	ANUNCIATOR_W2 = clearBit(ANUNCIATOR_W2, LED_PARKING_BRAKE)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W2_BEFORE != ANUNCIATOR_W2
}

func OnLEDLowVolts() {
	ANUNCIATOR_W2_BEFORE := ANUNCIATOR_W2
	ANUNCIATOR_W2 = setBit(ANUNCIATOR_W2, LED_LOW_VOLTS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W2_BEFORE != ANUNCIATOR_W2
}

func OffLEDLowVolts() {
	ANUNCIATOR_W2_BEFORE := ANUNCIATOR_W2
	ANUNCIATOR_W2 = clearBit(ANUNCIATOR_W2, LED_LOW_VOLTS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W2_BEFORE != ANUNCIATOR_W2
}

func OnLEDDoor() {
	ANUNCIATOR_W2_BEFORE := ANUNCIATOR_W2
	ANUNCIATOR_W2 = setBit(ANUNCIATOR_W2, LED_DOOR)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W2_BEFORE != ANUNCIATOR_W2
}

func OffLEDDoor() {
	ANUNCIATOR_W2_BEFORE := ANUNCIATOR_W2
	ANUNCIATOR_W2 = clearBit(ANUNCIATOR_W2, LED_DOOR)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W2_BEFORE != ANUNCIATOR_W2
}

// ANUNCIATOR_W1 LEDs
func OnLEDLowOilPress() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = setBit(ANUNCIATOR_W1, LED_LOW_OIL_PRESS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}
func OffLEDLowOilPress() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = clearBit(ANUNCIATOR_W1, LED_LOW_OIL_PRESS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OnLEDLowFuelPress() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = setBit(ANUNCIATOR_W1, LED_LOW_FUEL_PRESS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OffLEDLowFuelPress() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = clearBit(ANUNCIATOR_W1, LED_LOW_FUEL_PRESS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OnLEDAntiIce() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = setBit(ANUNCIATOR_W1, LED_ANTI_ICE)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OffLEDAntiIce() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = clearBit(ANUNCIATOR_W1, LED_ANTI_ICE)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OnLEDStarter() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = setBit(ANUNCIATOR_W1, LED_STARTER)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OffLEDStarter() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = clearBit(ANUNCIATOR_W1, LED_STARTER)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OnLEDApu() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = setBit(ANUNCIATOR_W1, LED_APU)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OffLEDApu() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = clearBit(ANUNCIATOR_W1, LED_APU)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OnLEDMasterCaution() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = setBit(ANUNCIATOR_W1, LED_MASTER_CAUTION)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OffLEDMasterCaution() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = clearBit(ANUNCIATOR_W1, LED_MASTER_CAUTION)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OnLEDVacuum() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = setBit(ANUNCIATOR_W1, LED_VACUUM)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OffLEDVacuum() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = clearBit(ANUNCIATOR_W1, LED_VACUUM)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OnLEDLowHydPress() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = setBit(ANUNCIATOR_W1, LED_LOW_HYD_PRESS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

func OffLEDLowHydPress() {
	ANUNCIATOR_W1_BEFORE := ANUNCIATOR_W1
	ANUNCIATOR_W1 = clearBit(ANUNCIATOR_W1, LED_LOW_HYD_PRESS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || ANUNCIATOR_W1_BEFORE != ANUNCIATOR_W1
}

// LANDING_GEAR_W LEDs
func OnLEDLeftGearGreen() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = setBit(LANDING_GEAR_W, LED_LEFT_GEAR_GREEN)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OffLEDLeftGearGreen() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = clearBit(LANDING_GEAR_W, LED_LEFT_GEAR_GREEN)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OnLEDLeftGearRed() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = setBit(LANDING_GEAR_W, LED_LEFT_GEAR_RED)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OffLEDLeftGearRed() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = clearBit(LANDING_GEAR_W, LED_LEFT_GEAR_RED)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OnLEDNoseGearGreen() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = setBit(LANDING_GEAR_W, LED_NOSE_GEAR_GREEN)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OffLEDNoseGearGreen() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = clearBit(LANDING_GEAR_W, LED_NOSE_GEAR_GREEN)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OnLEDNoseGearRed() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = setBit(LANDING_GEAR_W, LED_NOSE_GEAR_RED)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OffLEDNoseGearRed() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = clearBit(LANDING_GEAR_W, LED_NOSE_GEAR_RED)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OnLEDRightGearGreen() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = setBit(LANDING_GEAR_W, LED_RIGHT_GEAR_GREEN)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OffLEDRightGearGreen() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = clearBit(LANDING_GEAR_W, LED_RIGHT_GEAR_GREEN)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OnLEDRightGearRed() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = setBit(LANDING_GEAR_W, LED_RIGHT_GEAR_RED)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OffLEDRightGearRed() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = clearBit(LANDING_GEAR_W, LED_RIGHT_GEAR_RED)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OnLedGearGreen() {
	OnLEDLeftGearGreen()
	OffLEDLeftGearRed()
	OnLEDNoseGearGreen()
	OffLEDNoseGearRed()
	OnLEDRightGearGreen()
	OffLEDRightGearRed()
}

func OnLedGearRed() {
	OnLEDLeftGearRed()
	OffLEDLeftGearGreen()
	OnLEDNoseGearRed()
	OffLEDNoseGearGreen()
	OnLEDRightGearRed()
	OffLEDRightGearGreen()
}

func OnLEDMasterWarning() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = setBit(LANDING_GEAR_W, LED_MASTER_WARNING)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OffLEDMasterWarning() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = clearBit(LANDING_GEAR_W, LED_MASTER_WARNING)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OnLEDEngineFire() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = setBit(LANDING_GEAR_W, LED_ENGINE_FIRE)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

func OffLEDEngineFire() {
	LANDING_GEAR_W_BEFORE := LANDING_GEAR_W
	LANDING_GEAR_W = clearBit(LANDING_GEAR_W, LED_ENGINE_FIRE)
	LED_STATE_CHANGED = LED_STATE_CHANGED || LANDING_GEAR_W_BEFORE != LANDING_GEAR_W
}

// AUTO_PILOT_W LEDs
func OnLEDHeading() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = setBit(AUTO_PILOT_W, LED_HEADING)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OffLEDHeading() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = clearBit(AUTO_PILOT_W, LED_HEADING)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OnLEDNav() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = setBit(AUTO_PILOT_W, LED_NAV)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OffLEDNav() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = clearBit(AUTO_PILOT_W, LED_NAV)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OnLEDAPR() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = setBit(AUTO_PILOT_W, LED_APR)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OffLEDAPR() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = clearBit(AUTO_PILOT_W, LED_APR)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OnLEDREV() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = setBit(AUTO_PILOT_W, LED_REV)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OffLEDREV() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = clearBit(AUTO_PILOT_W, LED_REV)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OnLEDAlt() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = setBit(AUTO_PILOT_W, LED_ALT)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OffLEDAlt() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = clearBit(AUTO_PILOT_W, LED_ALT)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OnLEDVS() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = setBit(AUTO_PILOT_W, LED_VS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OffLEDVS() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = clearBit(AUTO_PILOT_W, LED_VS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OnLEDIAS() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = setBit(AUTO_PILOT_W, LED_IAS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OffLEDIAS() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = clearBit(AUTO_PILOT_W, LED_IAS)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OnLEDAP() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = setBit(AUTO_PILOT_W, LED_AP)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

func OffLEDAP() {
	AUTO_PILOT_W_BEFORE := AUTO_PILOT_W
	AUTO_PILOT_W = clearBit(AUTO_PILOT_W, LED_AP)
	LED_STATE_CHANGED = LED_STATE_CHANGED || AUTO_PILOT_W_BEFORE != AUTO_PILOT_W
}

// DebugPrintLEDStates prints the current state of all LEDs
func (b *bravoService) DebugPrintLEDStates() {
	var sb strings.Builder

	sb.WriteString("\nANUNCIATOR_W2 LEDs:\n")
	if ANUNCIATOR_W2&LED_FUEL_PUMP != 0 {
		sb.WriteString("- Fuel Pump is ON\n")
	}
	if ANUNCIATOR_W2&LED_PARKING_BRAKE != 0 {
		sb.WriteString("- Parking Brake is ON\n")
	}
	if ANUNCIATOR_W2&LED_LOW_VOLTS != 0 {
		sb.WriteString("- Low Volts is ON\n")
	}
	if ANUNCIATOR_W2&LED_DOOR != 0 {
		sb.WriteString("- Door is ON\n")
	}

	sb.WriteString("ANUNCIATOR_W1 LEDs:\n")
	if ANUNCIATOR_W1&LED_LOW_OIL_PRESS != 0 {
		sb.WriteString("- Low Oil Pressure is ON\n")
	}
	if ANUNCIATOR_W1&LED_LOW_FUEL_PRESS != 0 {
		sb.WriteString("- Low Fuel Pressure is ON\n")
	}
	if ANUNCIATOR_W1&LED_ANTI_ICE != 0 {
		sb.WriteString("- Anti-Ice is ON\n")
	}
	if ANUNCIATOR_W1&LED_STARTER != 0 {
		sb.WriteString("- Starter is ON\n")
	}
	if ANUNCIATOR_W1&LED_APU != 0 {
		sb.WriteString("- APU is ON\n")
	}
	if ANUNCIATOR_W1&LED_MASTER_CAUTION != 0 {
		sb.WriteString("- Master Caution is ON\n")
	}
	if ANUNCIATOR_W1&LED_VACUUM != 0 {
		sb.WriteString("- Vacuum is ON\n")
	}
	if ANUNCIATOR_W1&LED_LOW_HYD_PRESS != 0 {
		sb.WriteString("- Low Hydraulic Pressure is ON\n")
	}

	sb.WriteString("LANDING_GEAR_W LEDs:\n")
	if LANDING_GEAR_W&LED_LEFT_GEAR_GREEN != 0 {
		sb.WriteString("- Left Gear Green is ON\n")
	}
	if LANDING_GEAR_W&LED_LEFT_GEAR_RED != 0 {
		sb.WriteString("- Left Gear Red is ON\n")
	}
	if LANDING_GEAR_W&LED_NOSE_GEAR_GREEN != 0 {
		sb.WriteString("- Nose Gear Green is ON\n")
	}
	if LANDING_GEAR_W&LED_NOSE_GEAR_RED != 0 {
		sb.WriteString("- Nose Gear Red is ON\n")
	}
	if LANDING_GEAR_W&LED_RIGHT_GEAR_GREEN != 0 {
		sb.WriteString("- Right Gear Green is ON\n")
	}
	if LANDING_GEAR_W&LED_RIGHT_GEAR_RED != 0 {
		sb.WriteString("- Right Gear Red is ON\n")
	}
	if LANDING_GEAR_W&LED_MASTER_WARNING != 0 {
		sb.WriteString("- Master Warning is ON\n")
	}
	if LANDING_GEAR_W&LED_ENGINE_FIRE != 0 {
		sb.WriteString("- Engine Fire is ON\n")
	}

	sb.WriteString("AUTO_PILOT_W LEDs:\n")
	if AUTO_PILOT_W&LED_HEADING != 0 {
		sb.WriteString("- Heading is ON\n")
	}
	if AUTO_PILOT_W&LED_NAV != 0 {
		sb.WriteString("- Navigation is ON\n")
	}
	if AUTO_PILOT_W&LED_APR != 0 {
		sb.WriteString("- Approach is ON\n")
	}
	if AUTO_PILOT_W&LED_REV != 0 {
		sb.WriteString("- Reverse is ON\n")
	}
	if AUTO_PILOT_W&LED_ALT != 0 {
		sb.WriteString("- Altitude is ON\n")
	}
	if AUTO_PILOT_W&LED_VS != 0 {
		sb.WriteString("- Vertical Speed is ON\n")
	}
	if AUTO_PILOT_W&LED_IAS != 0 {
		sb.WriteString("- Indicated Airspeed is ON\n")
	}
	if AUTO_PILOT_W&LED_AP != 0 {
		sb.WriteString("- Autopilot is ON\n")
	}

	// Log the complete message at once
	b.Logger.Debug(sb.String())
}

func AllOff() {
	ANUNCIATOR_W2 = 0
	ANUNCIATOR_W1 = 0
	LANDING_GEAR_W = 0
	AUTO_PILOT_W = 0
	LED_STATE_CHANGED = true
}
