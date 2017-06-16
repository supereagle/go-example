package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
)

func main() {
	//Init the command-line flags.
	flag.Parse()

	fmt.Println("==============Start Glog==============")

	// Will be ignored as the program has exited in Fatal().
	defer func() {
		fmt.Println("Exit glog.")
	}()

	// Flushes all pending log I/O.
	defer glog.Flush()

	// The temp folder for log files when --log_dir is not set
	fmt.Printf("Temp folder for log files: %s\n", os.TempDir())

	glog.Info("Info")
	glog.V(1).Info("L1 info")
	glog.Error("Error")

	// Fatal() will call os.Exit() to exit the program,
	// so the following code and defer code will be ignored.
	glog.Fatal("Fatal")

	// Will be ignored as the program has exited in Fatal().
	fmt.Println("===============End Glog===============")
}
