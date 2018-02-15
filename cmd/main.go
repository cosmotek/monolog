package main

import (
	"strings"

	"github.com/rucuriousyet/monolog"
)

func main() {
	err := monolog.New(nil, nil).
		Add(func(p *monolog.Prompter) monolog.Cmd {
			p.Write("are you a human? (y/N): ")
			res := strings.ToLower(p.Read())

			if res == "y" || res == "yes" {
				return monolog.Continue

			} else if res == "n" || res == "no" {
				p.Write("goodbye.")
				return monolog.ExitChain
			}

			p.Write("invalid input, please retry.\n\n")
			return monolog.Retry
		}).
		Add(func(p *monolog.Prompter) monolog.Cmd {
			p.Write("are you a male? (y/N): ")
			res := strings.ToLower(p.Read())

			if res == "y" || res == "yes" {
				p.Write("hey dude.")
				return monolog.Continue

			} else if res == "n" || res == "no" {
				p.Write("flower power!")
				return monolog.Continue
			}

			p.Write("invalid input, please retry.\n\n")
			return monolog.Retry
		}).Do()

	if err != nil {
		panic(err)
	}
}
