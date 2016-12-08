package basic

import (
	"testing"
	"fmt"
	"time"
	"encoding/json"
	"os/signal"
	"os"
	"syscall"
	"math/rand"
	"sync"
)

func TestTime0(t *testing.T) {
	a := 1
	//b := a * time.Second // invalid operation: a * time.Second (mismatched types int and time.Duration)
	b := time.Duration(a) * time.Second // ok
	fmt.Println(b)
}

func TestTime1(t *testing.T) {
	s, _ := json.Marshal(map[string]string{
		"time": time.Now().Format(time.RFC3339),
	})
	fmt.Println(string(s))

	var j map[string]interface{}
	json.Unmarshal(s, &j)
	ts := j["time"].(string)
	fmt.Println(time.Parse(time.RFC3339, ts))
}

type ticker struct {
	checkTimeDone chan struct{}
	checkTimeUpdateDone chan struct{}
	checkTimeUpdate chan struct{}
	tt *time.Ticker
}

func TestTime2(t *testing.T)  {
	tc := ticker{
		checkTimeDone: make(chan struct{}),
		checkTimeUpdateDone: make(chan struct{}),
		checkTimeUpdate: make(chan struct{}),
		tt: time.NewTicker(time.Second),
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		<-sigs
		fmt.Println("sig exit")
		signal.Stop(sigs)
		fmt.Println("tc check done")
		close(tc.checkTimeDone)
		fmt.Println("tc check update done")
		close(tc.checkTimeUpdateDone)
		wg.Done()
		fmt.Println("routine 0 done")
	}()

	wg.Add(1)
	go func() {
		outfor:
			for {
				select {
				case <- tc.tt.C:
					fmt.Println("tc ticker: ", time.Now(), &tc.tt.C)
				case <- tc.checkTimeUpdate:
					// 但是在routine 2内部进行同样的操作就不会产生该问题
					tc.tt.Stop()
					tc.tt = time.NewTicker(time.Second)
				case <- tc.checkTimeDone:
					tc.tt.Stop()
					break outfor
				}
			}
		wg.Done()
		fmt.Println("routine 2 done")
	}()

	wg.Add(1)
	go func() {
		time.Sleep(2 * time.Second)
		outfor:
			for i:=0; i < 10; i++ {
				r := rand.Intn(2000)
				tim := time.NewTimer(time.Duration(r) * time.Millisecond)
				select {
				case <- tim.C:
					fmt.Println("tc update, r=", r)
					tc.checkTimeUpdate <- struct {}{}
					// 在routine 1中对tc.tt进行new操作会导致routine 2中的tc.tt终止计时
					//tc.tt.Stop()
					//tc.tt = time.NewTicker(time.Second)
				case <- tc.checkTimeUpdateDone:
					break outfor
				}
			}
		fmt.Println("routine 1 done")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("routine main done")
}

func TestTime3(t *testing.T)  {
	timer := time.NewTimer(time.Second)
	go func() {
		<- timer.C
		println("Timer expired")
	}()
	stop := timer.Stop()
	println("Timer cancelled:", stop)
}

func TestTime4(t *testing.T)  {
	ticker := time.NewTicker(time.Second)
	done := make(chan struct{})
	ticker_done := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("ticker")
			case <-ticker_done:
				fmt.Println("ticker done")
				ticker.Stop()
				return
			}
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		done <- struct {}{}
	}()

	<-done
	ticker_done <- struct {}{}
}

func TestTime5(t *testing.T) {
	now := time.Now()
	last := now.Add(-time.Second)
	fmt.Println(now.Add(-time.Second).Before(last)) // false
	fmt.Println(now.Add(-2 * time.Second).Before(last)) // true
	fmt.Println(now.Sub(last).Seconds()) // 1
}
