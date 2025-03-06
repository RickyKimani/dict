package scripts

import (
	"context"
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Terms struct {
	List []struct {
		Definition  string    `json:"definition"`
		Permalink   string    `json:"permalink"`
		ThumbsUp    int       `json:"thumbs_up"`
		Author      string    `json:"author"`
		Word        string    `json:"word"`
		Defid       int       `json:"defid"`
		CurrentVote string    `json:"current_vote"`
		WrittenOn   time.Time `json:"written_on"`
		Example     string    `json:"example"`
		ThumbsDown  int       `json:"thumbs_down"`
	} `json:"list"`
}

type TermsMsg struct {
	terms Terms
	err   error
}

func (t TermsMsg) Error() string {
	return t.err.Error()
}

func GetResults(t string) tea.Cmd {
	return func() tea.Msg {
		endpoint := fmt.Sprintf("https://api.urbandictionary.com/v0/define?term=%s", url.QueryEscape(t))
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
		if err != nil {
			return TermsMsg{err: err}
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return TermsMsg{err: err}
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
			}
		}(resp.Body)

		var terms Terms
		err = json.NewDecoder(resp.Body).Decode(&terms)
		if err != nil {
			return TermsMsg{err: err}
		}
		return TermsMsg{terms: terms}
	}
}
