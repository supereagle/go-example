package options

import (
	"github.com/spf13/pflag"
)

// Flags represents the flags needed to start the application.
type Flags struct {
	Address string
	Port    int32
}

// AddFlags adds the application flags to the flagset of pflag.
func (f *Flags) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&f.Address, "host", f.Address, "The IP address to serve on.")
	fs.Int32Var(&f.Port, "port", f.Port, "The port to serve on.")
}
