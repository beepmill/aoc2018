package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("2-1.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	var doubles, triples int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		log.Print(l)
		a := make(map[rune]int)
		for _, c := range l {
			a[c]++
		}
		var d, t bool
		for k, v := range a {
			switch v {
			case 2:
				d = true
				log.Printf("%c is a double!", k)
			case 3:
				t = true
				log.Printf("%c is a triple!", k)
			default:
				continue
			}
		}
		if d {
			doubles++
		}
		if t {
			triples++
		}
	}
	log.Printf("doubles:%v, triples:%v, checksum:%v", doubles, triples, doubles*triples)
}
