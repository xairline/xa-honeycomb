package honeycomb

import (
	"fmt"
	"github.com/sstallion/go-hid"
	"github.com/xairline/xa-honeycomb/pkg"
	"log"
	"sync"
)

var Vendor uint16 = 0x294B
var Product uint16 = 0x1901

type BravoService interface {
	IsReady() bool // ready to retrieve values
	Exit()
}

type bravoService struct {
	Logger pkg.Logger
	Bravo  *hid.Device
}

func (b bravoService) Exit() {
	if err := hid.Exit(); err != nil {
		b.Logger.Errorf("failed to exit hidapi: %v", err)
	}
}

func (b bravoService) IsReady() bool {
	//TODO implement me
	panic("implement me")
}

var bravoSvcLock = &sync.Mutex{}
var bravoSvc BravoService

func NewBravoService(logger pkg.Logger) BravoService {
	if bravoSvc != nil {
		logger.Info("Bravo SVC has been initialized already")
		return bravoSvc
	} else {
		logger.Info("Bravo SVC: initializing")
		bravoSvcLock.Lock()
		defer bravoSvcLock.Unlock()

		if err := hid.Init(); err != nil {
			logger.Errorf("failed to initialize hidapi: %v", err)
			return nil
		}
		bravo, err := hid.OpenFirst(Vendor, Product)
		if err != nil {
			logger.Errorf("failed to open device: %v", err)
			return nil
		}

		s, err := bravo.GetMfrStr()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Manufacturer String: %s\n", s)

		bravoSvc = &bravoService{
			Logger: logger,
			Bravo:  bravo,
		}
		return bravoSvc
	}
}
