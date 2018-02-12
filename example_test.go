package monolog_test

import "github.com/rucuriousyet/monolog"

func ExamplePromptChain() {
	err := monolog.New(nil).
		Add(monolog.Prompt{
			Msg: "Would you like to dance? (y/N)",
		}).
		Add(monolog.Prompt{
			Msg: "Would you like to sing? (y/N)",
		}).
		Add(monolog.Prompt{
			Msg: "Would you like to kiss? (y/N)",
		}).Do()

	if err != nil {
		panic(err)
	}

	// Output: nil
}
