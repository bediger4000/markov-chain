package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type TableEntry struct {
	state       rune
	nextState   rune
	probability float64
}

func main() {
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))
	startState := rune(os.Args[1][0])
	steps, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Chain runs from %c for %d steps\n", startState, steps)
	fileName := os.Args[3]
	fmt.Printf("Table in %q\n", fileName)

	buffer, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	table := make(map[rune][]*TableEntry)
	var states []rune

	for lineNo, line := range bytes.Split(buffer, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		var state, nextState rune
		var probability float64
		n, err := fmt.Sscanf(string(line), "%c %c %f", &state, &nextState, &probability)
		if err != nil {
			log.Fatalf("line %d: %v\n", lineNo, err)
		}
		if n != 3 {
			log.Fatalf("read %d items from %s:%d, wanted to read 3\n",
				n, fileName, lineNo)
		}

		if _, ok := table[state]; !ok {
			states = append(states, state)
		}

		if len(table[state]) > 0 {
			probability += table[state][len(table[state])-1].probability
		}

		table[state] = append(
			table[state],
			&TableEntry{
				state:       state,
				nextState:   nextState,
				probability: probability,
			},
		)
	}

	sort.Slice(states, func(i, j int) bool { return states[i] < states[j] })

	for _, ary := range table {
		for _, entry := range ary {
			fmt.Printf("%c  %c  %f\n", entry.state, entry.nextState, entry.probability)
		}
	}

	visited := make(map[rune]int)

	state := startState

	for steps > 0 {
		prob := rand.Float64()
		ary := table[state]
		for i := range ary {
			if prob <= ary[i].probability {
				visited[state]++
				state = ary[i].nextState
				break
			}
		}
		steps--
	}

	sum := 0
	for _, state := range states {
		fmt.Printf("%c: %d\n", state, visited[state])
		sum += visited[state]
	}
	fmt.Printf("total state visits: %d\n", sum)
}
