package ch03

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"sync"
	"time"
)

var ph = []string{"P1", "P2", "P3", "P4", "P5"}

const hunger = 3               // number of times each philosopher eats
const think = time.Second * 60 // Mean think time
const eat = time.Second * 60   // Mean eat time

var dining sync.WaitGroup

func diningProblem(phName string, dominantHand, otherHand *sync.Mutex) {
	fmt.Println(phName, "Seated")
	rg := rand.New(rand.NewSource(genSeed(phName)))

	// 平均时间t，t/2加上0到t的随机数？？？
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}

	for h := hunger; h > 0; h-- {
		fmt.Println(phName, "Hungry")
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		fmt.Println(phName, "Eating")
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		fmt.Println(phName, "Thinking")
		rSleep(think)
	}

	fmt.Println(phName, "Satisfied")
	dining.Done()
	fmt.Println(phName, "Left the table")
}

func SimulateDiningPhilosophers() {
	fmt.Println("Table empty")

	dining.Add(5)
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(ph); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(ph[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	go diningProblem(ph[0], forkLeft, fork0)
	dining.Wait() // wait for philosophers to finish
	fmt.Println("Table empty")
}

func genSeed(phName string) int64 {
	h := fnv.New64a()
	h.Write([]byte(phName))
	return int64(h.Sum64())
}
