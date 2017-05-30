//timers are for when you want to do sth once in the future
//tickers are for when you want to sth repeatedly at regular intervals
//attention:
//Tickers use a similar mechanism to timers: a channel that is sent values
//ticker is useful for some heart beat operations

package BasicFunctions

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	//ticker could also be stopped just like the timer
	//once the ticker is stopped it will not
	//receive any values on its channel
	//if the main thread is finished the ticker goroutines
	//are also expired automatically

	//ticker.Stop()
	//time.Sleep(time.Millisecond * 1500)

	fmt.Println("Ticker stopped")

}
