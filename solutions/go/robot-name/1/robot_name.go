package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var usedNames = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(usedNames) >= 676000 {
		return "", errors.New("no more unique names available")
	}

	for {
		c1 := rune('A' + rand.Intn(26))
		c2 := rune('A' + rand.Intn(26))
		num := rand.Intn(1000)
		candidate := fmt.Sprintf("%c%c%03d", c1, c2, num)

		if !usedNames[candidate] {
			usedNames[candidate] = true
			r.name = candidate
			return r.name, nil
		}
	}
}

func (r *Robot) Reset() {
	r.name = ""
}   