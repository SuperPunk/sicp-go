package ch03

import (
	"log"
	"testing"
)

func TestSimulateConcurrency(t *testing.T) {
	ret := SimulateConcurrency(1000000)
	log.Println(ret)
}

func TestSimulateCheckpoint(t *testing.T) {
	SimulateCheckpoint()
}

func TestSimulateDiningPhilosophers(t *testing.T) {
	SimulateDiningPhilosophers()
}
