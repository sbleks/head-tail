package headtail_test

import (
	"fmt"
	"log"

	headtail "github.com/sbleks/head-tail"
)

func Example_head() {

	lines, err := headtail.Head(`testdata/somefile`, 3)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(lines)

	// Output:
	// [alpha beta gamma]
}

func Example_head_over_limit() {

	lines, err := headtail.Head(`testdata/somefile`, 100)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(lines)

	// Output:
	// [alpha beta gamma delta]
}

func Example_tail() {

	lines, err := headtail.Tail(`testdata/somefile`, 3)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(lines)

	// Output:
	// [beta gamma delta]
}

func Example_tail_over_limit() {

	lines, err := headtail.Tail(`testdata/somefile`, 100)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(lines)

	// Output:
	// [alpha beta gamma delta]
}
