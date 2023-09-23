package InputsCollector

import "yayo/Actions"

type UserInput struct {
	Question string
	Answer   string
	Action   Actions.ActionInterface
	IsError  bool
}

func (u *UserInput) Run(collector *Collector) {
	answer, err := u.Action.Ask(collector.CurrentQuestionAnswerMap())

	if err != nil {
		u.IsError = true
		return
	}

	u.Answer = answer
	collector.QuestionsAnswerMap = collector.CurrentQuestionAnswerMap()
}

func CreateUserInput(question string, action Actions.ActionInterface) *UserInput {
	userInput := &UserInput{}
	userInput.Question = question
	userInput.Action = action

	return userInput
}
