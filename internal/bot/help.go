package bot

import (
	"github.com/google/uuid"
	tele "gopkg.in/telebot.v3"
)

func (b Bot) onHelp(c tele.Context) error {
	results := tele.Results{}

	cmds, err := b.Bot.Commands()
	if err != nil {
		return err
	}

	for _, e := range cmds {
		result := &tele.ArticleResult{
			Title:       e.Text,
			Description: e.Description,
			Text:        e.Description,
		}
		result.SetResultID(uuid.New().String())
		results = append(results, result)
	}

	return c.Answer(&tele.QueryResponse{
		Results:   results,
		CacheTime: -1,
	})
}
