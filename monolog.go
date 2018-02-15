package monolog

import (
	"bufio"
	"io"
	"os"
)

// Cmd is a simple type used for command constants returned to
// the prompter by prompt functions
type Cmd uint

const (
	// Continue tells the prompter to continue to the next prompt
	Continue Cmd = 0

	// Retry tells the prompter to re-run the currentd prompt
	Retry Cmd = 1

	// ExitChain tells the prompter to exit the chain now
	// (skipping all other prompts)
	ExitChain Cmd = 2
)

// Prompt is a simple function type that takes a prompter
// pointer and returns a command for the prompter to exec.
// This should be used to call the prompter.Read and
// prompter.Write methods, validate the received input and
// return the appropriate command.
type Prompt func(p *Prompter) Cmd

// Prompter is a  prompt chain object
type Prompter struct {
	scanner *bufio.Scanner
	errbuff error

	writer  io.Writer
	prompts []Prompt
}

// Read reads a string from the reader provided to the
// prompter. This function will block until input is
// received.
func (p *Prompter) Read() string {
	p.scanner.Scan()
	return p.scanner.Text()
}

// Write writes a string to the writer provided to the
// prompter itself. If an error is encountered, it should
// be returned by the Do function after the prompt finishes
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

// Add appends any number of prompt functions to
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
