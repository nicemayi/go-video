package main

import (
	"fmt"
	"testing"
)

// special name
func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	m.Run()
}

func TestPrint(t *testing.T) {
	res := Print(2)
	if res != 2 {
		t.Errorf("Error!")
	}
}

func TestPrint2(t *testing.T) {
	t.SkipNow()
	res := Print(3)
	if res != 3 {
		t.Errorf("Error!")
	}
}

func TestPrint3(t *testing.T) {
	t.Run("a2", func(t *testing.T) {
		fmt.Println("1")
	})
	t.Run("a3", func(t *testing.T) {
		fmt.Println("1")
	})
}

func BenchmarkAll1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print(n)
	}
}

func BenchmarkAll2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print(n)
	}
}
