package hxx_test

import (
	"fmt"

	"github.com/aalpar/hxx"
)

func ExampleFormats() {

	// a simple string example
	input0 := `Lorem Ipsum.`
	fmt.Printf("%v\n", hxx.NewDump(input0))

	// structures and other complex objects can be dumped.  These objects must be supplied as pointers into `NewDump`.
	// Pointers within objects are printed as they are stored.  Pointers are not followed.
	input1 := struct {
		Ipsum rune
		Array [5]int
	}{
		' ',
		[5]int{21, 32, 43, 54, 65},
	}
	fmt.Printf("%v\n", hxx.NewDump(&input1))

	// Output:
	// 0: 4c 6f 72 65 6d 20 49 70 73 75 6d 2e              Lorem Ipsum.
	//  0: 20  0  0  0  0  0  0  0 15  0  0  0  0  0  0  0   ...............
	// 10: 20  0  0  0  0  0  0  0 2b  0  0  0  0  0  0  0   .......+.......
	// 20: 36  0  0  0  0  0  0  0 41  0  0  0  0  0  0  0  6.......A.......

}
