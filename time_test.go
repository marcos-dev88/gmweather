package gmweather

import (
	"log"
	"testing"
	"time"
)

func TestTimeClock(t *testing.T) {
	start := time.Now()

	for {
		var since = time.Since(start)
		// s := time.Now().Add(time.Second * 5)

		if since.Seconds() > 5 {
			log.Printf("break!")
			break
		}
	}

	// log.Printf("time took -> %v", since.Seconds())
}
