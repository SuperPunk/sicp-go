package ch03

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	partList    = []string{"A", "B", "C", "D"}
	nAssemblies = 3
	wg          sync.WaitGroup
)

func worker(part string) {
	log.Println(part, "worker begins part")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println(part, "worker completes part")
	wg.Done()
}

func SimulateCheckpoint() {
	rand.Seed(time.Now().UnixNano())
	for c := 1; c <= nAssemblies; c++ {
		log.Println("begin assembly cycle", c)
		wg.Add(len(partList))
		for _, v := range partList {
			go worker(v)
		}
		wg.Wait()
		log.Println("assemble.  cycle", c, "complete")
	}
}
