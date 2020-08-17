package flags

import (
	"os"
	"strings"

	"github.com/barjoco/utils/log"
	"github.com/fatih/color"
)

// flagMap stores data about a flag; how many values it expects, and the id of the function it associated with
// funcMap stores a map of functions, where the key is a number the flags can associate with in the flagMap
// funcCounter counts the number of functions that have been registered
var flagMap = map[string][]int{}
var funcMap = map[int]func([]string){}
var funcCounter = 0

// Set a flag's action
// `names` are the optional names for the flag
// `values` is the number of values the flag is expecting. Use -1 if expecting a variable number of values
// `action` is the executed when the flag is parsed. Action provides the number of values supplied to the flag as an argument
// Example: Set([]string{"i", "input"}, -1, func(vals []string){ fmt.Println("Input files:", vals) })
func Set(names []string, values int, action func([]string)) {
	for _, name := range names {
		flagMap[name] = []int{values, funcCounter}
	}
	if values < -1 {
		log.ErrorFatal("Invalid number for expected values (must be -1 <= x)")
	}
	funcMap[funcCounter] = action
	funcCounter++
}

// Parse a list of arguments
func Parse(args []string) {
	// Separate argument list into groups of flag + values
	argString := strings.Join(args, " ")
	flags := strings.Split(argString, "-")[1:]

	var flagGroup []string
	var flagName string
	var numValsExp int
	var numVals int

	var flag string
	for i := 0; i < len(flags); i++ {
		flag = strings.TrimSpace(flags[i])

		// Remove empty elements and re-prepend flags with `-` or `--`
		if flag == "" {
			flags[i+1] = "-" + flags[i+1]
			continue
		}
		flag = "-" + flag

		// Add switch collections as indivdual switches to the flags array
		if flag[:2] != "--" && len(flag) > 2 && len(strings.Fields(flag)) == 1 {
			flags = append(flags, strings.Split(flag[1:], "")...)
			continue
		}

		// Parse flag group
		flagGroup = strings.Split(flag, " ")
		flagName = flagGroup[0]

		if flagData, ok := flagMap[flagName]; ok {
			// Check correct number of values are supplied
			numValsExp = flagData[0]
			numVals = len(flagGroup[1:])
			if numValsExp != -1 && numVals != numValsExp {
				color.Red("Unexpected number of values supplied to `%s` (has %d, wants %d)", flagName, numVals, numValsExp)
				os.Exit(1)
			}
			// Execute function associated with this flag
			funcMap[flagData[1]](flagGroup[1:])
		} else {
			color.Red("Unrecognised flag: %s", flagName)
			os.Exit(1)
		}
	}
}
