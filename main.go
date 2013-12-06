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
	simplifyGraph(g)
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
//	> Ben 45.67 Lea Tami
//	> Lea 60.75 Ben Tami
//	> Tami 33.20 Ben Lea
//	> Ben 20.99 Tami
//  > quit
//	I do not understand 'quit'; ignoring.
//	> CTRL-D
//
// And then it should return the map[string]map[string]float64 that it created
// in this process. It is not the job of buildGraphInteractively however, to
// simplify the "graph". Let that be anothers job.
func buildGraphInteractively() *graph {
	fmt.Println(`Please enter your finance graph here.

	<Person> <Sum> <Person> [Person...]

Lines have the above format, where those left of the sum
spend money on behalf of those to the right of the sum.
`)

	g := make(graph)

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
		obj := fields[2:]
		amount, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			fmt.Printf("error: cannot read number %q, ignoring line\n", fields[1])
			continue
		}
		amount /= float64(len(obj) + 1)

		// Put the information in the graph
		for _, v := range obj {
			if g[v] == nil {
				g[v] = make(map[string]float64)
			}
			g[v][subj] += amount
		}
	}

	fmt.Println()
	return &g
}

// simplifyGraph simplifies the "graph" g so that the number of transactions
// that need to take place is minimized.
//
// For example, if Ben owes Lea 12.00€, and Lea owes Ben 5.00€, then there
// should only be one transaction, namely that of Ben paying Lea 7.00€.
//
// Edges in the graph that have a weight of 0.0 should be deleted from the
// map, so that when we iterate over the map, we have the minimal number
// of transactions right there.
func simplifyGraph(g *graph) {
	// Implementation notes: we modify the graph g, that's why it is a pointer.
	// It doesn't really change the way you write this function though.
}

// printTransactions iterates through the simplified graph and prints all
// the edges contained in it as transactions from one person to another.
func printTransactions(g *graph) {
	fmt.Println("\nOUTSTANDING TRANSACTIONS:")
	for subj, others := range *g {
		fmt.Printf("%v must pay:\n", subj)
		for recv, sum := range others {
			fmt.Printf("    %.2f to %v\n", sum, recv)
		}
		fmt.Println()
	}
}
