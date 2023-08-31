package main

import (
	_ "fmt"

	"github.com/spf13/pflag"
)

func flags() {
	// Define flags
	flag := pflag.NewFlagSet("default", pflag.ContinueOnError)
	flag.String("name", "User", "Your name")
	flag.Int("age", 0, "Your age")

	// Access and print flag values
	//fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
	//fmt.Println(flag.Args())
	// args := flag.Args()
	// for i := 0; i < len(args); i += 2 {
	// 	flagName := args[i]
	// 	flagValue := args[i+1]
	// 	fmt.Printf("Flag: %s, Value: %s\n", flagName, flagValue)
	// }
	// flag.VisitAll(func(f *flag.Flag) {
	// 	fmt.Printf("%s: %s\n", f.Name, f.Value)
	// })
}
