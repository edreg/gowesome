package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	//What select lets you do is wait on multiple channels.
	//The first one to send a value "wins" and the code underneath the case is executed.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
		//time.After is a very handy function when using select.
		//Although it didn't happen in our case you can potentially write code that blocks forever
		//if the channels you're listening on never return a value. time.
		//After returns a chan (like ping) and will send a signal down it after the amount of time you define.
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

//func measureResponseTime(url string) time.Duration {
//	start := time.Now()
//	http.Get(url)
//	return time.Since(start)
//}
