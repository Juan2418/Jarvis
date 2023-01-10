package utils

import (
	"os"
	"strconv"
)

func CreateFullBufferedChannel(capacity int) chan bool {
	sync := make(chan bool, capacity)

	for i := 0; i < capacity; i++ {
		sync <- true
	}
	return sync
}

func GetConcurrentRoutines() int {
	concurrentRoutines := 10
	CONCURRENT_ROUTINES := os.Getenv("CONCURRENT_ROUTINES")
	if CONCURRENT_ROUTINES != "" {
		concurrentRoutinesEnvValue, err := strconv.Atoi(CONCURRENT_ROUTINES)
		if err != nil {
			panic(err)
		}
		concurrentRoutines = concurrentRoutinesEnvValue
	}

	return concurrentRoutines
}

func GetSyncBuffer() chan bool {
	concurrentRoutines := GetConcurrentRoutines()
	sync := CreateFullBufferedChannel(concurrentRoutines)
	return sync
}
