package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/golang/glog"
	//"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Compare log packages")

	fmt.Println("==========Start Standard Log==========")
	stdLog()
	fmt.Println("===========End Standard Log===========")

	fmt.Println("==============Start Glog==============")
	gLog()
	fmt.Println("===============End Glog===============")
}

type person struct {
	name string
	age  int
}

var testPerson = person{
	name: "robin",
	age:  29,
}

func stdLog() {
	defer func() {
		// Need to sleep to ensure the log is after others. Why?
		time.Sleep(1 * time.Second)
		fmt.Println("Exit standard log.")
	}()

	log.Printf("%+v", testPerson)

	// Add prefix
	log.SetPrefix("Add prefix: ")
	log.Printf("%+v", testPerson)

	// Set flags
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Printf("%+v", testPerson)

	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Printf("%+v", testPerson)
}

func gLog() {
	defer func() {
		fmt.Println("Exit glog.")
	}()

	//Init the command-line flags.
	flag.Parse()

	// Flushes all pending log I/O.
	defer glog.Flush()

	glog.Infof("%+v", testPerson)
	glog.V(1).Infof("%+v", testPerson)
	glog.Errorf("%+v", testPerson)
}
