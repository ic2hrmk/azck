package conf

import "flag"

type Flags struct {
	Kind    string
	EnvFile string
}

func LoadFlags() *Flags {
	flags := &Flags{}

	flag.StringVar(&flags.Kind, "kind", "", "Kind of micro service")
	flag.StringVar(&flags.EnvFile, "env", "", "Environment file")

	flag.Parse()

	return flags
}
