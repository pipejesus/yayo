package Actions

import (
	"github.com/pterm/pterm"
	"os"
)

type actionConfirmProject struct {
}

func (a actionConfirmProject) Ask(previousAnswers map[string]string) (string, error) {

	projectName := previousAnswers["projectName"]
	selectedPhp := previousAnswers["phpVersion"]
	selectedDb := previousAnswers["dbVersion"]

	pterm.Printfln(
		"Utworzyć projekt o nazwie %s na bazie PHP v%s i %s (https://%[1]s.lndo.site)?",
		pterm.Yellow(projectName),
		pterm.Yellow(selectedPhp),
		pterm.Yellow(selectedDb),
	)

	confirmed, err := pterm.DefaultInteractiveConfirm.WithConfirmText("t").WithRejectText("n").WithDefaultValue(true).Show("Zgoda?")

	if !confirmed {
		os.Exit(0)
	}

	var answer string
	if confirmed {
		answer = "tak"
	} else {
		answer = "nie"
	}

	return answer, err
}

func ActionConfirmProject() actionConfirmProject {
	return actionConfirmProject{}
}
