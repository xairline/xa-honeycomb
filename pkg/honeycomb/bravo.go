package honeycomb

import (
	"github.com/xairline/xa-honeycomb/pkg"
	"os"
	"sync"
)

type BravoService interface {
	IsReady() bool // ready to retrieve values
}

type bravoService struct {
	Logger          pkg.Logger
	bravoFilePath   string
	bravoFileFolder string
}

func (g bravoService) IsReady() bool {
	//TODO implement me
	panic("implement me")
}

var bravoSvcLock = &sync.Mutex{}
var bravoSvc BravoService

func NewBravoService(logger pkg.Logger, dir string) BravoService {
	if bravoSvc != nil {
		logger.Info("Bravo SVC has been initialized already")
		return bravoSvc
	} else {
		logger.Info("Bravo SVC: initializing")
		bravoSvcLock.Lock()
		defer bravoSvcLock.Unlock()
		logger.Infof("Bravo SVC: initializing with folder %s", dir)
		bravoSvc = &bravoService{
			Logger:          logger,
			bravoFileFolder: dir,
			bravoFilePath:   "",
		}
		// make sure bravo file folder exists
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
		return bravoSvc
	}
}
