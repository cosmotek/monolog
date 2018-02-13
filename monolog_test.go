package monolog_test

import "github.com/rucuriousyet/monolog"

func ExampleDo() {
	err := monolog.New(nil, nil).
		Add(monolog.Prompt{
			Msg: "Would you like to dance? (y/N)",
		}).
		Add(monolog.Prompt{
			Msg: "Would you like to sing? (y/N)",
		}).
		Add(monolog.Prompt{
			Msg: "Would you like to laugh? (y/N)",
		}).Do()

	if err != nil {
		panic(err)
	}
}
