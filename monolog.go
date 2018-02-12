package monolog

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Prompt struct {
	Msg string

	AcceptRegex  []string
	DeclineRegex []string
}

type Prompter struct {
	reader  io.Reader
	prompts []Prompt
}

func New(reader io.Reader) *Prompter {
	if reader == nil {
		reader = os.Stdin
	}

	return &Prompter{
		reader:  reader,
		prompts: make([]Prompt, 0),
	}
}

func (p *Prompter) Add(prompts ...Prompt) *Prompter {
	p.prompts = append(p.prompts, prompts...)
	return p
}

func (p *Prompter) Do() error {
	scanner := bufio.NewScanner(p.reader)

	for _, prompt := range p.prompts {
		for {
			fmt.Printf("%s: ", prompt.Msg)

			scanner.Scan()
			resp := scanner.Text()
			if resp != "y" {
				break
			}
		}
	}

	return nil
}
