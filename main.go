// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Fairsplit is a program to split expenses among a group of people.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Fairsplit – no one's cheating no one here.\n")
	g := buildGraphInteractively()
	printTransactions(g)
}

type graph map[string]map[string]float64

// buildGraphInteractively prints an introduction message and then proceeds to
// take input from the standard input until EOF is encountered (which can be
// inserted with Ctrl-D or Ctrl-Z, depending on OS.)
//
// If a line is erroneous, an error message should be printed, but processing
// should continue. Input from the user should be indicated by a greater than
// mark >. For example, the input might look like this:
//
//	Please enter your finance graph here.
//
//		<Person> <Sum> <Person> [Person...]
//
//	Lines have the above format, where those left of the sum
//	spend money on behalf of those to the right of the sum.
//
//	> Ben 45.67 Reed Sonia
//	> Reed 60.75 Ben Sonia
//	> Sonia 33.20 Ben Reed
//	> Ben 20.99 Sonia
//	> CTRL-D
//
func buildGraphInteractively() graph {
	fmt.Println(`Please enter your finance graph here.

	<Person> <Sum> <Person> [Person...]

Lines have the above format, where those left of the sum
spend money on behalf of those to the right of the sum.
`)

	amends := make(map[string]float64)

	// Implementation tip: use fmt.*Scan* functions.
	r := bufio.NewReader(os.Stdin)
	for {
		// Get a line from the standard input
		fmt.Print("> ")
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("error: %v\n", err)
			break
		}

		// Check that we have enough information
		fields := strings.Fields(line)
		if len(fields) < 3 {
			if len(fields) > 0 {
				fmt.Println("error: invalid format, ignoring line")
			} else if err == io.EOF {
				break
			}
			continue
		}

		// Assign the information to the correct entities
		subj := fields[0]
		others := fields[2:]
		amount, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			fmt.Printf("error: cannot read number %q, ignoring line\n", fields[1])
			continue
		}

		// Put the information in the graph
		part := amount / float64(len(others)+1)
		amends[subj] -= amount - part
		for _, o := range others {
			amends[o] += part
		}
	}
	fmt.Println()

	// Convert the amends data-structure to a graph
	type obj struct {
		k string  // key = name
		v float64 // value = sum
	}
	gets := make([]obj, 0, len(amends))
	puts := make([]obj, 0, len(amends))
	for k, v := range amends {
		if v < 0 {
			gets = append(gets, obj{k, -v})
		} else if v > 0 {
			puts = append(puts, obj{k, v})
		}
	}

	g := make(graph)
	for i, j := 0, 0; i < len(puts); i++ {
		put := &puts[i]
		for put.v >= 0.01 {
			if g[put.k] == nil {
				g[put.k] = make(map[string]float64)
			}

			get := &gets[j]
			if get.v > put.v {
				g[put.k][get.k] = put.v
				get.v -= put.v
				put.v = 0
			} else if get.v <= put.v {
				g[put.k][get.k] = get.v
				put.v -= get.v
				get.v = 0
				j++
			}
		}
	}

	return g
}

// printTransactions iterates through the simplified graph and prints all
// the edges contained in it as transactions from one person to another.
func printTransactions(g graph) {
	fmt.Println("\nOUTSTANDING TRANSACTIONS:")

	for subj, others := range g {
		if len(others) == 0 {
			continue
		}
		fmt.Printf("%v must pay:\n", subj)
		for recv, sum := range others {
			fmt.Printf("    %.2f to %v\n", sum, recv)
		}
		fmt.Println()
	}
}
