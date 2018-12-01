package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("1-2.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var changes []int64
	for scanner.Scan() {
		l := scanner.Text()
		v, err := strconv.ParseInt(l, 10, 32)
		if err != nil {
			log.Panic(err)
		}
		changes = append(changes, v)
	}

	var freq int64
	seen := make(map[int64]bool)
	for {
		for _, v := range changes {
			freq += v
			log.Printf("%v: %v", v, freq)
			if seen[freq] {
				log.Printf("%v has been seen twice!", freq)
				return
			}
			seen[freq] = true
		}
	}
}
