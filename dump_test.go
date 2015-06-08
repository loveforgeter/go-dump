package dump_test

import (
	"fmt"
	"os"
	"testing"

	. "."
)

func TestDump(t *testing.T) {
	b := true
	i := 123
	f := 321.123
	s := "string"
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	fmt.Fprint(os.Stdout, "Dump int: ")
	Dump(i, os.Stdout)
	fmt.Fprintln(os.Stdout)

	fmt.Fprint(os.Stdout, "Dump float: ")
	Dump(f, os.Stdout)
	fmt.Fprintln(os.Stdout)

	fmt.Fprint(os.Stdout, "Dump bool: ")
	Dump(b, os.Stdout)
	fmt.Fprintln(os.Stdout)

	fmt.Fprint(os.Stdout, "Dump string: ")
	Dump(s, os.Stdout)
	fmt.Fprintln(os.Stdout)

	fmt.Fprint(os.Stdout, "Dump map: ")
	Dump(m, os.Stdout)
	fmt.Fprintln(os.Stdout)
}
