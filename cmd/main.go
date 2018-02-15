package main

import (
	"fmt"
	"strings"

	"github.com/rucuriousyet/monolog"
	"github.com/rucuriousyet/monolog/prototypes"
)

func main() {
	likesCake := false
	name := "<nil>"

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
				p.Write("hey dude.\n")
				return monolog.Continue

			} else if res == "n" || res == "no" {
				p.Write("flower power!\n")
				return monolog.Continue
			}

			p.Write("invalid input, please retry.\n\n")
			return monolog.Retry
		}).
		Add(prototypes.YesNo("do you like cake?", &likesCake)).
		Add(prototypes.Str("whats your name?", &name)).Do()

	if err != nil {
		panic(err)
	}

	fmt.Println("name =", name, "likes cake =", likesCake)
}
