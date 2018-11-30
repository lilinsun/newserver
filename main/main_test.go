package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Test begin...")
	m.Run()
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print1to20()
	}
}
