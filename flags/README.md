# flags

Flag parsing modelled after classic GNU utils.

## Features

### Aliases

A flag can have multiple aliases. For example, `-l` and `--long`.

### Expected values

A flag can expect a number of values to be supplied. For example, `--output result.pdf`

Use `-1` to uncap the number of values to be expected.

### Switches

Switches can be provided, indivually or as a collection. For example, `-l -r -s -a` or `-lrsa`

Switches within a switch collection won't recieve any values, and are always processed at the end of the arguments list (but can appear anywhere).

## Quick start

Use `flags.Set` to set flags. The function expects an array of flag aliases, the number of values it expects, and a function which is executed when the flag is met.

The function provides an array containing the values supplied to the flag.

Once all flags are set, call `flags.Parse`, giving it an array of arguments to parse.

```go
// Has one alias, --name, and expects one value
flags.Set([]string{"--name"}, 1, func(vals []string) {
    fmt.Println("Hello,", vals[0])
})

// Has two aliases, -l and --leave, and expects no values
flags.Set([]string{"-l", "--leave"}, 0, func(_ []string) {
    fmt.Println("Bye!")
})

// Parse flags
flags.Parse(os.Args[1:])
```

```sh
myapp --name John -l
```

```
Outputs:
Hello, John
Bye!
```

## Syntactic notes

Switch collections are syntactic sugar for single-letter flags, allowing groups of them to take up less space.

Flags are executed in the order they appear on the command line, except for switch collections, which are executed after all other flags, in the order they appear in the collection.