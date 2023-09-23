package Lando

import (
	"fmt"
	"os"
	"yayo/Helpers"
)

func StartLando(projectName string) {
	err := os.Chdir(projectName)
	if err != nil {
		Helpers.CustomErrorPrefix().Printfln("Nie mogę znaleźć folderu z lando starterem, wychodzę!")
		os.Exit(0)
	}

	fmt.Println("")
	err = Helpers.OsCommand(
		fmt.Sprintf("lando"),
		[]string{"start"},
		true,
	)

	if err != nil {
		Helpers.CustomErrorPrefix().Printfln("Błąd przy odpalaniu projektu lando :(")
		os.Exit(0)
	}

}
