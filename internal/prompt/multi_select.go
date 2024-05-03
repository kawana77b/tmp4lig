package prompt

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
)

type MultiSelectQuestion struct {
	enableFilter bool
}

func NewMultiSelectQuestion() *MultiSelectQuestion {
	return &MultiSelectQuestion{}
}

func (q *MultiSelectQuestion) EnableFilter(enable bool) {
	q.enableFilter = enable
}

func (q *MultiSelectQuestion) Ask(question string, options []string) ([]string, error) {
	if len(options) == 0 {
		return []string{}, fmt.Errorf("options must not be empty")
	}

	answers := []string{}
	err := survey.AskOne(&survey.MultiSelect{
		Message: question,
		Options: options,
	}, &answers, survey.WithKeepFilter(q.enableFilter))

	if err != nil {
		return answers, fmt.Errorf("failed to ask question: %w", err)
	}

	return answers, nil
}
