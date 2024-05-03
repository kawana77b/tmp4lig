package gitignoreio

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

const (
	API_URL = "https://gitignore.io/api/"
)

type GitignoreIO struct {
}

func NewGitignoreIO() *GitignoreIO {
	return &GitignoreIO{}
}

// GetGitignoreData returns the content of the .gitignore file for the given languages.
func (g GitignoreIO) GetGitignoreData(lang ...string) ([]byte, error) {
	//cf. https://gitignore.io/api/node,react
	u, err := url.JoinPath(API_URL, strings.Join(lang, ","))
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("cannot access to the API")
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf.Bytes(), nil
}

// GetList returns the list of all available languages.
func (g GitignoreIO) GetList() ([]string, error) {
	//cf. https://gitignore.io/api/list
	u, err := url.JoinPath(API_URL, "list")
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("cannot access to the API")
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	// split the response text by comma and newline
	split := func(text string) []string {
		res := []string{}
		for _, s := range strings.Split(text, "\n") {
			res = append(res, strings.Split(s, ",")...)
		}
		return res
	}

	s := buf.String()
	return split(s), nil
}
