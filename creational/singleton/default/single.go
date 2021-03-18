package single

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

type single struct{}

var singleInstance *single

func newInstance() *single {
	if singleInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleInstance == nil {
			fmt.Println("creating single instance now")
			singleInstance = &single{}
		} else {
			fmt.Println("single instance already created")
		}
	}else {
		fmt.Println("single instance already created")
	}
	return singleInstance
}
