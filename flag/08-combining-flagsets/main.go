package main

import (
	"flag"
	"fmt"
	"os"
)

type config1 struct {
	Animal string
	Food   string
}

func NewConfig1() (*config1, *flag.FlagSet) {
	cfg := config1{
		Animal: "dog",
		Food:   "bones",
	}

	fs := flag.NewFlagSet("config1", flag.ContinueOnError)
	fs.StringVar(&cfg.Animal, "animal", "dog", "kind of animal")
	fs.StringVar(&cfg.Food, "food", "bones", "food for the animal")

	return &cfg, fs

}

type config2 struct {
	Vehicle string
	Color   string
}

func NewConfig2() (*config2, *flag.FlagSet) {
	cfg := config2{
		Vehicle: "car",
		Color:   "red",
	}

	fs := flag.NewFlagSet("config2", flag.ContinueOnError)
	fs.StringVar(&cfg.Vehicle, "vehicle", "car", "kind of vehicle")
	fs.StringVar(&cfg.Color, "color", "red", "color of the vehicle")

	return &cfg, fs

}

func main() {

	cfg1, fs1 := NewConfig1()
	cfg2, fs2 := NewConfig2()

	fs := flag.NewFlagSet("config", flag.ContinueOnError)

	// Copy flags from fs1 to fs
	fs1.VisitAll(func(f *flag.Flag) {
		fs.Var(f.Value, f.Name, f.Usage)
	})

	// Copy flags from fs2 to fs
	fs2.VisitAll(func(f *flag.Flag) {
		fs.Var(f.Value, f.Name, f.Usage)
	})

	// Parse flags and handle help/usage flags.
	// Since -h or --help are not defined in any of the flagsets, it
	// will fail to parse. We can then look at the original args given
	// and check if -h or --help was given, and if so, print usage.
	err := fs.Parse(os.Args[1:])
	if err != nil {
		// Check for help or h flag by looking at the original args
		for _, arg := range os.Args[1:] {
			if arg == "--help" || arg == "-h" {
				fs.Usage()
				os.Exit(0)
			}
		}
		// If it's not help, exit with the error
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("cfg1: %+v\n", cfg1)
	fmt.Printf("cfg2: %+v\n", cfg2)

}
