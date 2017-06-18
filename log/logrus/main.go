package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

//var log = logrus.New()

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	fmt.Println("=============Start Logrus=============")

	// Will be ignored as the program has exited in Fatal().
	defer func() {
		fmt.Println("Exit Logrus.")
	}()

	// Will be ignored as the log level is info serverity.
	log.Debug("Debug")

	// Log with fields
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Warn("Warn")

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Log without fields
	log.Error("Error")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "common field",
		"always": "always logged",
	})

	contextLogger.Info("Info one")
	contextLogger.Info("Info two")

	// Panic() will stop normal execution of the current goroutine,
	// so the following code will be ignored, but the defer code still executes.
	// log.Panic("Panic")

	// Fatal() will call os.Exit() to exit the program,
	// so the following code and defer code will be ignored.
	// log.Fatal("Fatal")

	// Will be ignored as the program has exited in Fatal().
	fmt.Println("==============End Logrus==============")
}
