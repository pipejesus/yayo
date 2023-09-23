package Actions

import (
	"github.com/pterm/pterm"
)

type ActionPhpVersion struct {
	Options ActionOptions
}

func (a ActionPhpVersion) Ask(previousAnswers map[string]string) (string, error) {
	optionsLabels := a.Options.GetOptionsLabels()
	answer, err := pterm.DefaultInteractiveSelect.WithOptions(optionsLabels).WithDefaultText("Wersja PHP").WithFilter(false).Show()

	if err != nil {
		return "", err
	}

	answerValue, err := a.Options.GetOptionValueByLabel(answer)

	return answerValue, err
}

func NewActionPhpVersion() ActionPhpVersion {
	return ActionPhpVersion{
		Options: PhpOptions(),
	}
}
