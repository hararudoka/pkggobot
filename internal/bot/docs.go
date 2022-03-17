package bot

import (
	"strings"

	"github.com/hararudoka/pkggobot/internal/service"
	tele "gopkg.in/telebot.v3"
)

// TODO: fix formatting
func (b Bot) onDoc(c tele.Context, a string) error {
	results := tele.Results{}

	title, text, url := "", "", ""

	args := strings.Split(a, ".")

	if len(args) > 0 { // TODO сделать
		pkgName := args[0]
		doc, err := service.NewDoc(pkgName)
		if err != nil {
			panic(err)
		}

		if len(args) == 2 {
			sym := doc.Find(args[1])
			text = sym.Description[0]
		} else if len(args) == 3 {
			doc.Find(args[1])
			// TODO get by 3d argument
		} else {
			text = doc.Overview
			title = doc.PkgName
			url = "https://pkg.go.dev/" + doc.PkgName
		}

		res := b.Result(c, "doc", map[string]interface{}{
			"Title": "found: " + title,
			"URL":   url,
		})
		res.SetContent(&tele.InputTextMessageContent{Text: text})

		results = append(results, res)
	} else {
		title, text = "doc <pkg>.[<methodOrType>[.<methodOrField>]]", ""

		res := b.Result(c, "nodoc", map[string]interface{}{
			"Title": title,
		})
		res.SetContent(&tele.InputTextMessageContent{Text: "not found: "}) // TODO: errors

		results = append(results, res)
	}

	return c.Answer(&tele.QueryResponse{
		Results:   results,
		CacheTime: -1,
	})
}
