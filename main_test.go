package main

import (
	"sync"
	"testing"
	"time"
)

func TestF1(t *testing.T) {
	ch := make(chan int)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
	}()
	select {
	case <-ch:
	case <-time.After(300 * time.Millisecond):
	}
	<-ch
	<-ch
	<-ch

	select {
	case <-time.After(300 * time.Millisecond):
	case <-time.After(600 * time.Millisecond):
	case <-time.After(900 * time.Millisecond):
		// Test other primitive
		mu := sync.Mutex{}
		mu.Lock()
		mu.Unlock()
		mu.Lock()

		go func() {
			mu.Lock()
			mu.Unlock()
			mu.Lock()
		}()
	}

	select {
	case <-time.After(300 * time.Millisecond):
	case <-time.After(600 * time.Millisecond):
	case <-time.After(900 * time.Millisecond):
		// Test channel created in SDK
		ch2 := time.After(300 * time.Millisecond)
		<-ch2
	}
}
