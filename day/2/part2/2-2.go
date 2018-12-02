package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("2-2.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var boxes []string
	for scanner.Scan() {
		boxes = append(boxes, scanner.Text())
	}

	for i, box := range boxes {
		for _, candidate := range boxes[i+1:] {
			var d int
			for p := 0; p < len(box); p++ {
				if box[p] != candidate[p] {
					d++
				}
				/* if d > 1 {
					log.Printf("%v and %v are not a match (d:%v)", box, candidate, d)
					break
				} */
			}
			if d == 1 {
				log.Printf("Found a match! %v and %v differ by only one letter!", box, candidate)
				return
			}
			log.Printf("%v and %v are not a match (d:%v)", box, candidate, d)
		}
	}
}
