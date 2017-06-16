package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("==========Start Standard Log==========")

	// Will be ignored if call Fatal(), but still be executed if call Panic().
	defer func() {
		// Need to sleep to ensure the log is after others. Why?
		time.Sleep(1 * time.Second)
		fmt.Println("Exit standard log.")
	}()

	log.Print("Print")

	// Add prefix
	log.SetPrefix("Add prefix: ")
	log.Print("Print with prefix")

	// Set flags
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("Print with flags")

	// Update the flags
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Print("Print with flags")

	// Panic() will stop normal execution of the current goroutine,
	// so the following code will be ignored, but the defer code still executes.
	// log.Panic("Panic")

	// Fatal() will call os.Exit() to exit the program,
	// so the following code and defer code will be ignored.
	// log.Fatal("Fatal")

	// Will be ignored if call Panic() or Fatal() in front.
	fmt.Println("===========End Standard Log===========")
}
