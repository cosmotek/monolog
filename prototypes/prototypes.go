package prototypes

import (
	"fmt"
	"strings"

	"github.com/rucuriousyet/monolog"
)

// YesNo returns a yes or no prompt with the provided msg
func YesNo(promptMsg string, gotYes *bool) monolog.Prompt {
	return func(p *monolog.Prompter) monolog.Cmd {
		p.Write(fmt.Sprintf("%s (y/N): ", promptMsg))
		res := strings.ToLower(p.Read())

		if res == "y" || res == "yes" {
			*gotYes = true
			return monolog.Continue

		} else if res == "n" || res == "no" {
			return monolog.Continue
		}

		p.Write("invalid input, please retry (yes or no).\n\n")
		return monolog.Retry
	}
}

// Str returns a string prompt using the provided pointer and msg
func Str(promptMsg string, got *string) monolog.Prompt {
	return func(p *monolog.Prompter) monolog.Cmd {
		p.Write(fmt.Sprintf("%s: ", promptMsg))

		*got = p.Read()
		return monolog.Continue
	}
}
