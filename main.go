package main

import (
	"yayo/Actions"
	"yayo/Github"
	"yayo/Helpers"
	"yayo/InputsCollector"
	"yayo/Lando"
	//"github.com/wk8/go-ordered-map/v2"
)

func main() {
	Helpers.DisplayLogoNewProject()

	inputsCollector := InputsCollector.NewCollector(InputsCollector.UserInputs{
		InputsCollector.CreateUserInput("projectName", Actions.ActionProjectName()),
		InputsCollector.CreateUserInput("phpVersion", Actions.NewActionPhpVersion()),
		InputsCollector.CreateUserInput("dbVersion", Actions.NewActionDbVersion()),
		InputsCollector.CreateUserInput("confirmProject", Actions.ActionConfirmProject()),
	})

	inputsCollector.RunUserInputs()

	Github.DownloadStarter(inputsCollector)
	Lando.UpdateLandoFile(inputsCollector)

}
