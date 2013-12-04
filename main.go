package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fairsplit – no one's cheating no one here.\n")
	g := buildGraphInteractively()
	simplifyGraph(g)
	printTransactions(g)
}

type graph map[string]map[string]float64

const intro = `Please enter your finance graph here.

	<Person> [Person...] <Sum> <Person> [Person...]

Lines have the above format, where those left of the sum
spend money on behalf of those to the right of the sum.
`

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
//		<Person> [Person...] <Sum> <Person> [Person...]
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
	fmt.Println(intro)
	g := make(graph)

	// Implementation tip: use fmt.*Scan* functions.

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

const outro = `== OUTSTANDING TRANSACTIONS ============================================`

// printTransactions iterates through the simplified graph and prints all
// the edges contained in it as transactions from one person to another.
func printTransactions(g *graph) {
	fmt.Println(outro)
	for subj, others := range g {
		fmt.Fprint("%v must pay:\n", subj)
		for recv, sum := range others {
			fmt.Fprint("    %v to %v\n", sum, recv)
		}
		fmt.Println()
	}
}
