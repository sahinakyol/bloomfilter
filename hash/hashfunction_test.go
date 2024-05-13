package hash

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMurmur3Hash(test *testing.T) {
	var testStrings = testStrings()
	var determinismFailCount = 0

	hashResults := make(map[string]uint32)

	for _, str := range testStrings {
		hash := Murmur3Hash(str, 123)
		if _, exists := hashResults[str]; !exists {
			hashResults[str] = hash
		} else {
			if hashResults[str] != hash {
				determinismFailCount++
			}
		}
	}

	seenHashes := make(map[uint32]bool)
	for str, hash := range hashResults {
		if _, exists := seenHashes[hash]; exists {
			fmt.Println("Collision detected for string:", str, "with hash:", hash)
		} else {
			seenHashes[hash] = true
		}
	}

	if determinismFailCount > 0 {
		fmt.Println("Determinism test failed times:", determinismFailCount)
	}
}

func TestFnv1aHash(t *testing.T) {
	var testStrings = testStrings()
	var determinismFailCount = 0

	hashResults := make(map[string]uint64)

	for _, str := range testStrings {
		hash := Fnv1aHash(str)
		if _, exists := hashResults[str]; !exists {
			hashResults[str] = hash
		} else {
			if hashResults[str] != hash {
				determinismFailCount++
			}
		}
	}

	seenHashes := make(map[uint64]bool)
	for str, hash := range hashResults {
		if _, exists := seenHashes[hash]; exists {
			fmt.Println("Collision detected for string:", str, "with hash:", hash)
		} else {
			seenHashes[hash] = true
		}
	}

	if determinismFailCount > 0 {
		fmt.Println("Determinism test failed times:", determinismFailCount)
	}
}

func TestStringSumHash(t *testing.T) {
	var testStrings = testStrings()
	var determinismFailCount = 0

	hashResults := make(map[string]uint64)

	for _, str := range testStrings {
		hash := StringSumHash(str)
		if _, exists := hashResults[str]; !exists {
			hashResults[str] = hash
		} else {
			if hashResults[str] != hash {
				determinismFailCount++
			}
		}
	}

	seenHashes := make(map[uint64]bool)
	for str, hash := range hashResults {
		if _, exists := seenHashes[hash]; exists {
			fmt.Println("Collision detected for string:", str, "with hash:", hash)
		} else {
			seenHashes[hash] = true
		}
	}

	if determinismFailCount > 0 {
		fmt.Println("Determinism test failed times:", determinismFailCount)
	}
}

func testStrings() []string {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var testStrings []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		testStrings = append(testStrings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during file scan: %s", err)
	}
	return testStrings
}

// Collision number is 0 for Fnv1aHash
