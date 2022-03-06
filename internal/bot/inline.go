package bot

import (
	"log"

	tele "gopkg.in/telebot.v3"
)

func (b Bot) onQuery(c tele.Context) error {
	command, data := b.parseQuery(c.Data())

	log.Println("пришёл инлайн")
	switch command {
	case "doc", "d":
		return b.onDoc(c, data)
	default:
		return b.onHelp(c)
	}
}
