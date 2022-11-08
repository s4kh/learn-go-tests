package main

import (
	"os"
	"time"

	"github.com/s4kh/learn-go-tests/mocking"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

type DefaultSleeper struct{}

func (ds *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	// Power of using interface
	ds := &DefaultSleeper{}
	mocking.Countdown(os.Stdout, ds)
	// fmt.Println(Hello("test"))
}
