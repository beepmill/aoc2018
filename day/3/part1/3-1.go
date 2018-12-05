package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Claim TODO
type Claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Height int
}

// Fabric TODO
type Fabric [][]int

// NewFabric TODO
func NewFabric(width, height int) (fabric Fabric) {
	log.Printf("Creating fabric of dimensions %vx%v", width, height)
	fabric = make(Fabric, width)
	for i := 0; i < height; i++ {
		fabric[i] = make([]int, height)
	}
	return
}

// NewClaim creates a new Claim from a string formatted like this:
// #16 @ 570,515: 23x18
func NewClaim(c string) (claim Claim, err error) {
	log.Println(c)
	claimMatch := *regexp.MustCompile(`^#(?P<id>\d+)\s@\s(?P<x>\d+),(?P<y>\d+):\s(?P<w>\d+)x(?P<h>\d+)$`)
	captures := make(map[string]string)
	m := claimMatch.FindStringSubmatch(c)
	for i, name := range claimMatch.SubexpNames() {
		if i != 0 {
			captures[name] = m[i]
			log.Printf("%v: %v", name, m[i])
		}
	}
	claim.ID, err = strconv.Atoi(captures["id"])
	if err != nil {
		return Claim{}, err
	}
	claim.X, err = strconv.Atoi(captures["x"])
	if err != nil {
		return Claim{}, err
	}
	claim.Y, err = strconv.Atoi(captures["y"])
	if err != nil {
		return Claim{}, err
	}
	claim.Width, err = strconv.Atoi(captures["w"])
	if err != nil {
		return Claim{}, err
	}
	claim.Height, err = strconv.Atoi(captures["h"])
	if err != nil {
		return Claim{}, err
	}
	log.Printf("%#v", claim)
	return
}

// AddClaim TODO
func (f *Fabric) AddClaim(c Claim) (err error) {
	//log.Printf("Adding claim #%v to fabric", c.ID)
	for x := c.X; x < c.X+c.Width; x++ {
		for y := c.Y; y < c.Y+c.Height; y++ {
			//log.Printf("%v:%v", x, y)
			(*f)[x][y]++
		}
	}
	return
}

// ConflictedSquareInches TODO
func (f *Fabric) ConflictedSquareInches() (squares int, err error) {
	for _, row := range *f {
		for _, inch := range row {
			if inch > 1 {
				squares++
			}
		}
	}
	return
}

func main() {
	f, err := os.Open("3-1.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	var claims []Claim
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		c, err := NewClaim(l)
		if err != nil {
			log.Panic(err)
		}
		claims = append(claims, c)
	}
	// Determine fabric size
	var w, h int
	for _, claim := range claims {
		if w < claim.X+claim.Width {
			w = claim.X + claim.Width
		}
		if h < claim.Y+claim.Height {
			h = claim.Y + claim.Height
		}
	}
	fabric := NewFabric(w, h)
	for _, claim := range claims {
		fabric.AddClaim(claim)
	}
	conflicted, err := fabric.ConflictedSquareInches()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("%v square inches are in contention!", conflicted)
}
