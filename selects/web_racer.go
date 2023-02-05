package selects

import (
	"fmt"
	"net/http"
	"time"
)

type EmptyChan struct{}
type HttpFunc func(url string) (res *http.Response, err error)

const tenSecondsTimeOut = 10 * time.Second

func Racer(urlOne, urlTwo string) (winner string, err error) {
	return ConfigurableRacer(urlOne, urlTwo, tenSecondsTimeOut)

}

func ConfigurableRacer(urlOne, urlTwo string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(urlOne):
		return urlOne, nil
	case <-ping(urlTwo):
		return urlTwo, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Racer: Http request took more than %q", timeout.String())
	}

}

func ping(url string) chan EmptyChan {
	ch := make(chan EmptyChan)
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
