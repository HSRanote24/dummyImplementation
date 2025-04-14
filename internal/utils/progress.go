package utils

import (
	"log"
	"time"
)

// TrackProgress logs the progress of a given operation.
func TrackProgress(operation string, totalSteps int) {
	for i := 1; i <= totalSteps; i++ {
		log.Printf("Progress: %s - Step %d of %d", operation, i, totalSteps)
		time.Sleep(1 * time.Second) // Simulate work being done
	}
	log.Printf("Operation %s completed.", operation)
}