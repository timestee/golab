package main

import (
	"math/rand"
	"time"
)
import "fmt"

type Work struct {
	index int
}

func (w *Work) Refuse() {
}

func (w *Work) Do() {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(1000)))
}
func (w *Work) Quit() {
}

func worker(i int, ch chan Work, quit chan struct{}) {
	for {
		select {
		case w := <-ch:
			if quit == nil {
				fmt.Println("worker", i, "refused", w)
				w.Refuse()
				break
			}
			w.Do()
			fmt.Println("worker", i, "processed", w)
		case <-quit:
			fmt.Println("worker", i, "quitting")
			quit = nil
		}
	}
}

func main() {
	nilTestChan := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		nilTestChan <- struct{}{}
	}()
	for {
		select {
		case <-nilTestChan:
			nilTestChan = nil
			fmt.Println("sellect nil")
		default:
			if nilTestChan == nil {
				goto selectNil
			}
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
			fmt.Println("doing")
		}
	}
selectNil:
	ch, quit := make(chan Work), make(chan struct{})
	go func() {
		for i := 0; i < 100; i++ {
			ch <- Work{index: i}
		}
	}()
	for i := 0; i < 4; i++ {
		go worker(i, ch, quit)
	}
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(1000)))
	close(quit)
	time.Sleep(2 * time.Second)
}
