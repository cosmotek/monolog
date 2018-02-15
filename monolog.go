package monolog

import (
	"bufio"
	"io"
	"os"
)

type Cmd uint

const (
	Continue  Cmd = 0
	Retry     Cmd = 1
	ExitChain Cmd = 2
)

type Prompt func(p *Prompter) Cmd

// Prompter is a  prompt chain object
type Prompter struct {
	scanner *bufio.Scanner
	errbuff error

	writer  io.Writer
	prompts []Prompt
}

func (p *Prompter) Read() string {
	p.scanner.Scan()
	return p.scanner.Text()
}

func (p *Prompter) Write(body string) {
	_, err := p.writer.Write([]byte(body))
	if err != nil {
		p.errbuff = err
	}
}

// New creates and returns a new prompter chain reading
// from either the provided reader parameter, or
// os.Stdin (if the reader parameter is nil).
// If the writer parameter is not nil, it will be used
// in place of os.Stdout.
func New(reader io.Reader, writer io.Writer) *Prompter {
	if reader == nil {
		reader = os.Stdin
	}

	scanner := bufio.NewScanner(reader)
	if writer == nil {
		writer = os.Stdout
	}

	return &Prompter{
		scanner: scanner,
		writer:  writer,
		prompts: make([]Prompt, 0),
	}
}

// Add appends any number of prompt objects to
// the chain and returns it. Keep in mind that the
// chain is first in, first out.
func (p *Prompter) Add(prompts ...Prompt) *Prompter {
	p.prompts = append(p.prompts, prompts...)
	return p
}

// Do executes a prompter chain and returns an error
// if it exists.
//
// This function will block but should
// be able to function in a go routine as long as the
// reader in use is async safe (may need a mutex-lock).
func (p *Prompter) Do() error {
	for _, prompt := range p.prompts {
		for {
			cmd := prompt(p)
			if p.errbuff != nil {
				return p.errbuff
			}

			if cmd == ExitChain {
				return nil
			} else if cmd == Continue {
				break
			}
		}
	}

	return nil
}
