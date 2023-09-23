package InputsCollector

import (
	"errors"
)

type UserInputs []*UserInput

type Collector struct {
	Inputs             []*UserInput
	QuestionsAnswerMap map[string]string
}

func NewCollector(inputs UserInputs) *Collector {
	return &Collector{
		inputs,
		make(map[string]string),
	}
}

func (c *Collector) RunUserInputs() {
	for _, userInput := range c.Inputs {
		userInput.Run(c)
	}
}

func (c *Collector) GetAnswer(question string) (string, error) {
	if answer, ok := c.QuestionsAnswerMap[question]; ok {
		return answer, nil
	}

	return "", errors.New("Brakuje odpowiedzi na pytanie: " + question)
}

func (c *Collector) CurrentQuestionAnswerMap() map[string]string {
	qaMap := make(map[string]string)

	if len(c.Inputs) < 1 {
		return qaMap
	}

	for _, v := range c.Inputs {
		qaMap[v.Question] = v.Answer
	}

	return qaMap
}
