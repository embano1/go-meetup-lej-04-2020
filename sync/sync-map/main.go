package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	k := "hello"
	v := "world"

	// store k/v in the sync-map
	sm.Store(k, v)

	// retrieve the value for k from the sync-map
	res, ok := sm.Load(k)
	if ok {
		fmt.Printf("found value %q for key %q\n", res.(string), k)
	}

	// delete and retrieve
	sm.Delete(k)
	_, ok = sm.Load(k)
	if !ok {
		fmt.Printf("key %q successfully deleted\n", k)
	}
}
