package Github

import (
	"github.com/pterm/pterm"
	"os"
	"yayo/Helpers"
	"yayo/InputsCollector"
)

const landoStarterUrl string = "https://github.com/pipejesus/luntek-starter.git"

func DownloadStarter(inputsCollector *InputsCollector.Collector) {
	folderName, err := inputsCollector.GetAnswer("projectName")

	spinnerInfo, _ := pterm.DefaultSpinner.Start("Ściągam starter na Twoją dyskietkę.")

	if err != nil {
		spinnerInfo.Fail("Brak nazwy projektu!")
		os.Exit(0)
	}

	err = Helpers.OsCommand("git", []string{"clone", landoStarterUrl, folderName}, false)

	if err != nil {
		spinnerInfo.Fail("Nieudane ściąganie startera z GitHub!")
		os.Exit(0)
	}

	err = os.Chdir(folderName)

	if err != nil {
		spinnerInfo.Fail("Nie mogę przejść do folderu ze starterem")
		os.Exit(0)
	}

	err = Helpers.OsCommand("rm", []string{"-rf", ".git"}, false)

	if err != nil {
		spinnerInfo.Fail("Nie mogę wyczyścić folderu ze starterem")
		os.Exit(0)
	}

	err = os.Chdir("..")

	if err != nil {
		spinnerInfo.Fail("Nie mogę wrócić do folderu nadrzędnego!")
		os.Exit(0)
	}

	spinnerInfo.Info("Ok, ściągnięte!")
}
