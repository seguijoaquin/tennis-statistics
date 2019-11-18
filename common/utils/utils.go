package utils

import (
	"log"
	"time"
)

// WaitForDependencies sleeps the execution until all dependencies are ready
func WaitForDependencies() {
	log.Println("[UTILS] This function is executed in package utils. Waiting 20 seconds for rabbit to come up...")
	time.Sleep(20 * time.Second)
	log.Println("[UTILS] Done")
}

// FailOnError logs an error msg and exits
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Check an error and panic accordingly
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
