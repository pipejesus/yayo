package Actions

import (
	"github.com/pterm/pterm"
	"os"
	"yayo/Helpers"
)

type actionProjectName struct {
}

func (a actionProjectName) Ask(previousAnswers map[string]string) (string, error) {
	answer, err := pterm.DefaultInteractiveTextInput.WithDefaultText("Nazwa projektu").Show()

	if err != nil || len(answer) < 1 {
		Helpers.CustomErrorPrefix().Printfln("Brak nazwy projektu, wychodzę!")
		os.Exit(0)
	}

	return answer, err
}

func ActionProjectName() actionProjectName {
	return actionProjectName{}
}
