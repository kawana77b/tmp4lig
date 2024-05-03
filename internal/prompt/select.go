package prompt

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
)

// SelectQuestion is a question that asks the user to select an option from a list.
type SelectQuestion struct {
	enableFilter bool
}

func NewSelectQuestion() *SelectQuestion {
	return &SelectQuestion{}
}

// EnableFilter enables or disables the filter feature of the select question.
func (q *SelectQuestion) EnableFilter(enable bool) {
	q.enableFilter = enable
}

// Ask asks a question with a list of options.
func (q *SelectQuestion) Ask(question string, options []string) (string, error) {
	if len(options) == 0 {
		return "", fmt.Errorf("options must not be empty")
	}

	answer := ""
	err := survey.AskOne(&survey.Select{
		Message: question,
		Options: options,
	}, &answer, survey.WithKeepFilter(q.enableFilter))

	if err != nil {
		return "", fmt.Errorf("failed to ask question: %w", err)
	}

	return answer, nil
}
