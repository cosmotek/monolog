package monolog

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// Prompt is a simple prompt object containing the
// various situational messages, rules and handlers
// of a command line prompt. It can use regex, function
// handlers or boolean evaluation to respond, quit the
// chain etc.
type Prompt struct {
	Msg string

	AcceptRegex  []string
	DeclineRegex []string
}

// Eval takes an input string, evaluates it against
// it's ruleset, and returns a "canContinue" boolean
// an error. If the error is not nil, but canContinue
// is true, the prompt will be repeated. If canContinue
// is false and the error is not nil, the chain will
// exit and print out the error
func (p Prompt) Eval(input string) (bool, error) {
	switch input {
	case "y":
		return true, nil
	case "n":
		return false, errors.New("goodbye.")
	default:
		return true, errors.New("unrecognized input, please try again.\n\n")
	}

	return true, nil
}

// Prompter is a  prompt chain object
type Prompter struct {
	reader  io.Reader
	writer  io.Writer
	prompts []Prompt
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

	if writer == nil {
		writer = os.Stdout
	}

	return &Prompter{
		reader:  reader,
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
	scanner := bufio.NewScanner(p.reader)

	for _, prompt := range p.prompts {
		for {
			_, err := p.writer.Write([]byte(prompt.Msg))
			if err != nil {
				return err
			}

			scanner.Scan()
			canContinue, resperr := prompt.Eval(scanner.Text())

			if resperr != nil {
				_, err = p.writer.Write([]byte(resperr.Error()))
				if err != nil {
					return err
				}
			}

			if !canContinue {
				return nil
			}

			if resperr == nil {
				break
			}
		}
	}

	return nil
}
