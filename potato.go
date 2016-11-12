package main

import (
	"fmt"
	log "logrus"
	helper "helper"
)

/*
func Init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
*/

func main() {
	fmt.Println("Hello World")
	log.Info("Hello Potato")
	helper.HelperInit()
}
