package main

import (
	//"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	//COMMAND-LINE ARGUMENTS
	
	/* argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg) */

	//COMMAND-LINE ARGUMENTS

	//COMMAND-LINE FLAGS

/* 	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args()) */
	
	//COMMAND-LINE FLAGS


	//COMMAND-LINE SUBCOMMANDS

/* 	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.String("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' sbcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(" enable:", *fooEnable)
		fmt.Println(" name:", *fooName)
		fmt.Println(" tail:". fooCmd.Args())
	
	case "bar":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println(" name:", *fbarLevel)
		fmt.Println(" tail:". barCmd.Args())
	
	case "foo":
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	} */
	//COMMAND-LINE SUBCOMMANDS


	//ENVIRONMENT VARIABLES

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

	//ENVIRONMENT VARIABLES
}