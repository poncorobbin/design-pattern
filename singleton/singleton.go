package main

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}
var once sync.Once

type db struct{}

var instance *db

func GetInstanceWithLock() *db {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance = &db{}
		fmt.Println("creating single instance now")
	} else {
		fmt.Println("single instance already created")
	}
	return instance
}

func GetInstanceWithOnce() *db {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("creating single instance now")
				instance = &db{}
			})
	} else {
		fmt.Println("single instance already created")
	}

	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		go GetInstanceWithOnce()
	}

	fmt.Scanln()
}
