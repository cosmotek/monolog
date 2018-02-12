package monolog_test

import "github.com/rucuriousyet/monolog"

func ExampleDo() {
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

	// WouldOutput: Would you like to dance? (y/N): Would you like to sing? (y/N): Would you like to kiss?
}
