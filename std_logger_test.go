package clog

import (
	"github.com/pkg/errors"
	stdLog "log"
	"os"
	"testing"
)

var testLogger Logger

func TestMain(m *testing.M) {
	testLogger = NewStdLogger(LevelDebug, os.Stdout, "", stdLog.LstdFlags)

	// Run Test
	exitCode := m.Run()

	// Exit
	os.Exit(exitCode)
}

func TestDebug(t *testing.T) {
	testLogger.Debug("Debug")
	testLogger.Debugf("%s", "Debugf")
}

func TestInfo(t *testing.T) {
	testLogger.Info("Info")
	testLogger.Infof("%s", "Infof")
}

func TestWarn(t *testing.T) {
	testLogger.Warn("Warn")
	testLogger.Warnf("%s", "Warnf")
}

func TestError(t *testing.T) {
	testLogger.Error("Error", errors.New("an error occurred"))
	testLogger.Errorf("%s", "Errorf")
}

func TestFatal(t *testing.T) {
	testLogger.Fatal("Fatal", errors.New("a fatal error occurred"))
	testLogger.Fatalf("%s", "Fatalf")
}
