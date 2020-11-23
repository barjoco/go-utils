package flags

import (
	"os"
	"strings"

	"github.com/barjoco/utils/log"
	"github.com/fatih/color"
)

// flagGroups stores an array of flag groups, that is the flag name + the values supplied to it
// flagMap stores data about a flag; how many values it expects and the id of the function it is associated with
// funcMap stores a map of functions, where the key is a number the flagGroups can associate with in the flagMap
// funcCounter counts the number of functions that have been registered
var flagGroups []string
var flagMap = map[string][]int{}
var funcMap = map[int]func([]string){0: func([]string) {}}
var funcCounter = 0

// Set a flag's action
// `names` are the optional names for the flag
// `values` is the number of values the flag is expecting. Use -1 if expecting a variable number of values
// `action` is the executed when the flag is parsed. Action provides the values supplied to the flag as an argument
// Example: Set([]string{"i", "input"}, -1, func(vals []string){ fmt.Println("Input files:", vals) })
func Set(names []string, values int, action func([]string)) {
	for _, name := range names {
		if _, ok := flagMap[name]; ok {
			log.ErrorFatal("Flag `%s` has already been set", name)
		}
		flagMap[name] = []int{values, funcCounter}
	}
	if values < -1 {
		log.ErrorFatal("Invalid number for expected values (must be -1 <= x < ∞)")
	}
	funcMap[funcCounter] = action
	funcCounter++
}

// Parse a list of arguments
func Parse(args []string) {
	// Separate argument list into groups of flag + values
	argString := strings.Join(args, " ")
	flagGroups = strings.Split(argString, "-")[1:]

	var flagGroup []string
	var flagName string
	var numValsExp int
	var numVals int

	var flag string
	for i := 0; i < len(flagGroups); i++ {
		flag = strings.TrimSpace(flagGroups[i])

		// Ignore empty elements and re-prepend flags with `-` or `--`
		if flag == "" {
			flagGroups[i+1] = "-" + flagGroups[i+1]
			continue
		}
		flag = "-" + flag

		// Add switch collections as indivdual switches to the flags array
		if flag[:2] != "--" && len(strings.Fields(flag)[0]) > 2 {
			flagGroups = append(flagGroups, strings.Split(flag[1:], "")...)
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
			color.Red("Unrecognised flag: `%s`", flagName)
			os.Exit(1)
		}
	}
}
