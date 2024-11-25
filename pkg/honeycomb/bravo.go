package honeycomb

import (
	"context"
	"github.com/sstallion/go-hid"
	"github.com/xairline/xa-honeycomb/pkg"
	"sync"
	"time"
)

var Vendor uint16 = 0x294B
var Product uint16 = 0x1901

type BravoService interface {
	UpdateLeds()
	Exit()
}

type bravoService struct {
	Logger pkg.Logger
	Bravo  *hid.Device
	ctx    context.Context

	hidReportBuffer []byte
	cancelFunc      context.CancelFunc
}

func (b *bravoService) UpdateLeds() {

	go func() {
		for {
			select {
			case <-b.ctx.Done():
				b.Logger.Infof("UpdateLeds: Context canceled, exiting goroutine")
				return
			default:
				if LED_STATE_CHANGED {
					b.hidReportBuffer[0] = 0x0
					b.hidReportBuffer[1] = AUTO_PILOT_W
					b.hidReportBuffer[2] = LANDING_GEAR_W
					b.hidReportBuffer[3] = ANUNCIATOR_W1
					b.hidReportBuffer[4] = ANUNCIATOR_W2
				}

				if x, err := b.Bravo.SendFeatureReport(b.hidReportBuffer); err != nil {
					b.Logger.Errorf("failed to write to device: %v", err)
					b.Logger.Infof("bytes written: %d\n", x)
				}
				LED_STATE_CHANGED = false
				time.Sleep(100 * time.Millisecond) // Simulated delay
			}
		}
	}()
}

func (b bravoService) Exit() {
	b.cancelFunc()

	b.hidReportBuffer[0] = 0x0
	b.hidReportBuffer[1] = 0x0
	b.hidReportBuffer[2] = 0x0
	b.hidReportBuffer[3] = 0x0
	b.hidReportBuffer[4] = 0x0

	if x, err := b.Bravo.SendFeatureReport(b.hidReportBuffer); err != nil {
		b.Logger.Errorf("failed to write to device: %v", err)
		b.Logger.Infof("bytes written: %d\n", x)
	}

	if err := b.Bravo.Close(); err != nil {
		b.Logger.Errorf("failed to close device: %v", err)
	}
	if err := hid.Exit(); err != nil {
		b.Logger.Errorf("failed to exit hidapi: %v", err)
	}
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
		ctx, cancel := context.WithCancel(context.Background())
		bravoSvc = &bravoService{
			Logger:          logger,
			Bravo:           bravo,
			ctx:             ctx,
			hidReportBuffer: make([]byte, 65),
			cancelFunc:      cancel,
		}
		bravoSvc.UpdateLeds()
		return bravoSvc
	}
}
