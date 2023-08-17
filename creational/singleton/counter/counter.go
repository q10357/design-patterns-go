package counter

//https://medium.com/@ramseyjiang_22278/singleton-design-pattern-in-golang-with-unit-tests-324952bdc77f

import (
	"fmt"
	"sync"
)

/*
Once is an object that will perform exactly one action.
A Once must not be copied after first use.
In the terminology of the Go memory model, the return from f “synchronizes before” the return from any call of once.Do(f).
https://pkg.go.dev/sync#Once
*/
var once sync.Once

type Singleton interface {
	IncrementByOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func (s *singleton) IncrementByOne() int {
	s.count++
	return s.count
}

// If instance exists, no replicas
func GetInstance() *singleton {
	if instance == nil {
		//init out singleton
		once.Do(
			func() {
				fmt.Println("Initializing singleton instance...")
				instance = new(singleton)
			},
		)
	} else {
		//already initialized
		fmt.Println("Singleton instance already initialized.")
	}

	return instance
}
