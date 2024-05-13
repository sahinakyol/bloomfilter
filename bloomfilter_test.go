package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestNewBloomFilter(t *testing.T) {
	n := uint64(1000)
	p := 0.01

	bf := NewBloomFilter(n, p)

	if bf.k != uint64(7) {
		t.Errorf("BloomFilter should calculate expected %v, result %v", bf.k, uint64(7))
	}

	if bf.size != uint64(9586) {
		t.Errorf("BloomFilter should calculate expected %v, result %v", bf.size, uint64(9586))
	}

	if len(bf.bitset) != 9586 {
		t.Errorf("BloomFilter should calculate expected %v, result %v", len(bf.bitset), uint64(9586))
	}
}

func TestBloomFilter_Add_And_Contains(t *testing.T) {
	n := uint64(1000)
	p := 0.01

	bf := NewBloomFilter(n, p)

	bf.Add("hello")

	if !bf.Contains("hello") {
		t.Error("BloomFilter should contain 'hello'")
	}
}

func TestBloomFilter(t *testing.T) {
	n := uint64(100000)
	p := 3000.0

	bf := NewBloomFilter(n, p)

	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var testStrings []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		bf.Add(scanner.Text())
		testStrings = append(testStrings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during file scan: %s", err)
	}

	for _, valueItem := range testStrings {
		if !bf.Contains(valueItem) {
			t.Errorf("BloomFilter should contain %s", valueItem)
		}
	}
}
