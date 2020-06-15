package ch03

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"testing"
	"time"
)

func TestSimulateConcurrency(t *testing.T) {
	//ret := SimulateConcurrency(10)
	//fmt.Print(ret)

	h := fnv.New64a()
	h.Write([]byte("Ares"))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	fmt.Println(rg.Int63n(int64(time.Second)))
	fmt.Println(time.Duration(rg.Int63n(int64(time.Second))))
}

func TestSimulateCheckpoint(t *testing.T) {
	SimulateCheckpoint()
}

func TestSimulateDiningPhilosophers(t *testing.T) {
	SimulateDiningPhilosophers()
}
