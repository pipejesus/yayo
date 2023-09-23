package Actions

import "errors"

type ActionInterface interface {
	Ask(previousInputs map[string]string) (string, error)
}

type ActionOption struct {
	value string
	label string
}

type ActionOptions []ActionOption

func (aopts ActionOptions) GetOptionValueByLabel(label string) (string, error) {
	for _, v := range aopts {
		if v.label != label {
			continue
		}
		return v.value, nil
	}
	return "", errors.New("Brak takiej opcji!")
}

func (aopts ActionOptions) GetOptionsLabels() []string {
	var flatOptions []string
	for _, v := range aopts {
		flatOptions = append(flatOptions, v.label)
	}

	return flatOptions
}

func (aopts ActionOptions) GetOptionsValues() []string {
	var flatOptions []string
	for _, v := range aopts {
		flatOptions = append(flatOptions, v.value)
	}

	return flatOptions
}
