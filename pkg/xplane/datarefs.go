package xplane

import (
	"encoding/csv"
	"fmt"
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
	"os"
	"path"
	"strings"
)

type profile struct {
	dataref_strs []string
	operator     string
	value        string
	conditions   string
	on           func()
	off          func()
}

func (s *xplaneService) setupDataRefs(airplaneICAO string) {
	s.Logger.Infof("Setup Datarefs for: %s", airplaneICAO)
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
	res := make(map[string]profile)
	for i, record := range records {
		if i == 0 {
			continue
		}
		name := record[0]
		onFunc, offFunc := s.assignOnAndOffFuncs(name)
		res[name] = profile{
			dataref_strs: strings.Split(record[1], ";"),
			operator:     record[2],
			value:        record[3],
			conditions:   record[4],
			on:           onFunc,
			off:          offFunc,
		}
	}
	s.Logger.Debugf("res: %v", res)
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
	//[xa honeycomb - development] WARNING: No on/off functions found for: VOLT_LOW
	//[xa honeycomb - development] WARNING: No on/off functions found for: OIL_LOW_P
	//[xa honeycomb - development] WARNING: No on/off functions found for: FUEL_LOW_P
	//[xa honeycomb - development] WARNING: No on/off functions found for: ANTI_ICE
	//[xa honeycomb - development] WARNING: No on/off functions found for: ENG_STARTER
	//[xa honeycomb - development] WARNING: No on/off functions found for: APU
	//[xa honeycomb - development] WARNING: No on/off functions found for: VACUUM
	//[xa honeycomb - development] WARNING: No on/off functions found for: HYDRO_LOW_P
	//[xa honeycomb - development] WARNING: No on/off functions found for: PARKING_BRAKE
	//[xa honeycomb - development] WARNING: No on/off functions found for: DOORS
	default:
		s.Logger.Warningf("No on/off functions found for: %s", name)
		return nil, nil
	}
}
