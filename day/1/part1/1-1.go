package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("1-1.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	var result int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		v, err := strconv.ParseInt(l, 10, 32)
		if err != nil {
			log.Panic(err)
		}
		result += v
	}
	log.Println(result)
}
