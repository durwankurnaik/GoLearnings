package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // here, the empty struct takes 0 bytes of memory allocation so this is actually very efficient

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{} // don't know why this was written, the program also works fine even without this line
	// in this example, the logCh is being forcefully closed because main function call was finished and in this case
	// it does not matter that much, but in some cases, we may get memory leaks and that could cause our application to shut down forcefully
	// we could do a defer statement with a anonymous function that closes the channel, but we most of the time, an empty struct and select statement
	// by this way, we have successfully closed our logCh channel without causing any memory leak that could have happened if it was to be close forcefully
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		} // here, using this method is preferred over using defer statement
	}
}
