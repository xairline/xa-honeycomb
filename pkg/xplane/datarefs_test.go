package xplane

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock logger to capture logs for testing
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(msg string) {
	m.Called(msg)
}

func (m *MockLogger) Debug(msg string) {
	m.Called(msg)
}

func (m *MockLogger) Error(msg string) {
	m.Called(msg)
}

func (m *MockLogger) Warningf(format string, a ...interface{}) {
	m.Called(format, a)
}

func (m *MockLogger) Warning(msg string) {
	m.Called(msg)
}

func (m *MockLogger) Infof(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Debugf(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Errorf(format string, args ...interface{}) {
	m.Called(format, args)
}

func TestSetupDataRefs(t *testing.T) {
	// Create a temporary directory to simulate the plugin path
	// current path of the file
	_, b, _, _ := runtime.Caller(0)
	// path of the file
	basepath := path.Dir(b)
	pluginPath := path.Join(basepath, "..", "..")
	err := os.MkdirAll(pluginPath, 0755)
	assert.NoError(t, err)

	// Create a test CSV file
	airplaneICAO := "B738"

	// Create a mock logger
	mockLogger := new(MockLogger)
	mockLogger.On("Infof", mock.Anything, mock.Anything).Return()
	mockLogger.On("Debugf", mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything).Return()
	mockLogger.On("Warningf", mock.Anything, mock.Anything).Return()

	// Create the xplaneService
	xpService := &xplaneService{
		Logger:     mockLogger,
		pluginPath: pluginPath,
	}

	// Call the setupDataRefs method
	xpService.setupDataRefs(airplaneICAO)

}

func TestSetupDataRefsWithAdvancedExpression(t *testing.T) {
	// Create a temporary directory to simulate the plugin path
	// current path of the file
	_, b, _, _ := runtime.Caller(0)
	// path of the file
	basepath := path.Dir(b)
	pluginPath := path.Join(basepath, "..", "..")
	err := os.MkdirAll(pluginPath, 0755)
	assert.NoError(t, err)

	// Create a test CSV file
	airplaneICAO := "DH8D"

	// Create a mock logger
	mockLogger := new(MockLogger)
	mockLogger.On("Infof", mock.Anything, mock.Anything).Return()
	mockLogger.On("Debugf", mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything).Return()
	mockLogger.On("Warningf", mock.Anything, mock.Anything).Return()

	// Create the xplaneService
	xpService := &xplaneService{
		Logger:     mockLogger,
		pluginPath: pluginPath,
	}

	// Call the setupDataRefs method
	xpService.setupDataRefs(airplaneICAO)

}

func TestUpdateLeds(t *testing.T) {
	// Create a temporary directory to simulate the plugin path
	// current path of the file
	_, b, _, _ := runtime.Caller(0)
	// path of the file
	basepath := path.Dir(b)
	pluginPath := path.Join(basepath, "..", "..")
	err := os.MkdirAll(pluginPath, 0755)
	assert.NoError(t, err)

	// Create a test CSV file
	airplaneICAO := "A339"

	// Create a mock logger
	mockLogger := new(MockLogger)
	mockLogger.On("Infof", mock.Anything, mock.Anything).Return()
	mockLogger.On("Debugf", mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything).Return()
	mockLogger.On("Warningf", mock.Anything, mock.Anything).Return()

	// Create the xplaneService
	xpService := &xplaneService{
		Logger:     mockLogger,
		pluginPath: pluginPath,
	}

	// Call the setupDataRefs method
	xpService.setupDataRefs(airplaneICAO)
	xpService.updateLeds()
}
