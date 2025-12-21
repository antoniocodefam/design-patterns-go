package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type myLogger struct{
  asterisk string
}

var (
   instance *myLogger
   once sync.Once
   mutex = &sync.Mutex{}
)

func getInstance() *myLogger {
	once.Do(func(){
		instance = &myLogger{}
	})

	return instance
}

func (l *myLogger) log(text string) {
	if l.asterisk == "" {
		mutex.Lock()
		defer mutex.Unlock()

		if l.asterisk == "" {
           fmt.Println("Creating random asterisk...")
		   time.Sleep(2 * time.Second)

		   asteriskCount := rand.Intn(5) +1
		   l.asterisk = strings.Repeat("*", asteriskCount)
		}
	}

	fmt.Printf("%s %s %s\n", l.asterisk, text, l.asterisk)
}