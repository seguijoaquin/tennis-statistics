package common

import (
	"log"
	"time"
)

// WaitForDependencies sleeps the execution until all dependencies are ready
func WaitForDependencies() {
	log.Println("[COMMON] This function is executed in package common. Waiting 20 seconds for rabbit to come up...")
	time.Sleep(20 * time.Second)
	log.Println("[COMMON] Done")
}
