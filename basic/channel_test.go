package basic

import (
	"fmt"
	"testing"
	"time"
)

// wrong usage, test will exit with status 2
func TestChannel0(t *testing.T)  {
	c := make(chan string)
	c <- "hello"
	fmt.Println(<-c)
}

// right usage, test print hello
func TestChannel1(t *testing.T)  {
	c := make(chan string)
	go func() {
		c <- "hello"
	}()
	fmt.Println(<-c)
}

// right usage, test print hello
func TestChannel2(t *testing.T)  {
	c := make(chan string, 1)
	c <- "hello"
	fmt.Println(<-c)
}

// wrong usage, fatal error: all goroutines are asleep - deadlock!
func TestChannel3(t *testing.T)  {
	c := make(chan string, 1)
	c <- "hello"
	c <- " world" // block goroutine
	fmt.Println(<-c)
}

// right usage, test print hello
func TestChannel4(t *testing.T)  {
	c := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		c <- "hello"
	}()
	fmt.Println(<-c) // block 3 seconds
}

func TestChannel5(t *testing.T)  {
	c := make(chan string, 2)
	c <- "hello"
	//fmt.Println(<-c)
	close(c)
	fmt.Printf("%T", <-c)
	fmt.Println(<-c)
}

func TestChannel6(t *testing.T) {
	messages := make(chan string)
	signals := make(chan bool)

	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

// 等待子进程结束用没有缓冲的channel

/*
wait
wait
done
child thread quit
main thread quit
*/
func TestChannel7(t *testing.T) {
	done := make(chan struct{})
	go func() {
		forDone:
			for {
				select {
				case <-done: // block
					fmt.Println("done")
					break forDone
				default:
					fmt.Println("wait")
					time.Sleep(500 * time.Millisecond)
				}
			}
		fmt.Println("child thread quit")
	}()
	time.Sleep(1 * time.Second)
	done <- struct {}{}
	fmt.Println("main thread quit")
}

/*
wait
wait
wait
main thread quit
*/
func TestChannel8(t *testing.T) {
	done := make(chan struct{}, 1)
	go func() {
		forDone:
			for {
				select {
				case <-done: // block
					fmt.Println("done")
					break forDone
				default:
					fmt.Println("wait")
					time.Sleep(500 * time.Millisecond)
				}
			}
		fmt.Println("child thread quit")
	}()
	time.Sleep(1 * time.Second)
	done <- struct {}{}
	fmt.Println("main thread quit")
}

/*
wait
wait
wait
done
child thread quit
main thread quit
*/
func TestChannel9(t *testing.T) {
	done := make(chan struct{}, 1)
	go func() {
		forDone:
			for {
				select {
				case <-done: // block
					fmt.Println("done")
					break forDone
				default:
					fmt.Println("wait")
					time.Sleep(500 * time.Millisecond)
				}
			}
		fmt.Println("child thread quit")
	}()
	time.Sleep(1 * time.Second)
	done <- struct {}{}
	time.Sleep(1 * time.Second)
	fmt.Println("main thread quit")
}