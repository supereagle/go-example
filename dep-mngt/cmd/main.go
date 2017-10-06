package main

import (
	"flag"
	"os"

	// "github.com/golang/glog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/supereagle/go-example/dep-mngt/cmd/options"
)

func main() {
	// Log to standard error instead of files.
	flag.Set("logtostderr", "true")

	// Flushes all pending log I/O.
	// defer glog.Flush()

	logrus.SetOutput(os.Stdout)

	// Initialize flags.
	f := &options.Flags{}
	f.AddFlags(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
}
