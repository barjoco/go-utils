package tabs

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Write is used to write tabbed output
// array 1st dimension is each line
// array 2nd dimension is items in the line
// eg [][]string{ []string{"ANIMAL", "COLOUR", "EATS"}, []string{"Dog", "Black", "Biscuits"}, []string{"Cat", "Ginger", "Fish"} }
// outputs:
// ANIMAL COLOUR EATS
// Dog    Black  Biscuits
// Cat    Ginger Fish
func Write(tabLines [][]string) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)
	for _, line := range tabLines {
		var linef string
		for _, lineCol := range line {
			linef += lineCol + "\t"
		}
		fmt.Fprintf(tw, "%s\n", linef)
	}
	tw.Flush()
}
