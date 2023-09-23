package Actions

import "github.com/pterm/pterm"

type ActionDbVersion struct {
	Options ActionOptions
}

func (a ActionDbVersion) Ask(previousAnswers map[string]string) (string, error) {
	optionsLabels := a.Options.GetOptionsLabels()
	answer, err := pterm.DefaultInteractiveSelect.WithOptions(optionsLabels).WithDefaultText("Baza danych").Show()

	if err != nil {
		return "", err
	}

	answerValue, err := a.Options.GetOptionValueByLabel(answer)
	return answerValue, err
}

func NewActionDbVersion() ActionDbVersion {
	return ActionDbVersion{
		Options: DbOptions(),
	}
}
