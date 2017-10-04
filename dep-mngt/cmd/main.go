package main

import (
	"flag"

	// "github.com/golang/glog"
	"github.com/spf13/pflag"

	"github.com/supereagle/go-example/dep-mngt/cmd/options"
)



func main() {
	// Log to standard error instead of files.
	flag.Set("logtostderr", "true")

	// Flushes all pending log I/O.
	// defer glog.Flush()

	// Initialize flags.
	f := &options.Flags{}
	f.AddFlags(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
}
