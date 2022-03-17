package bot

import (
	"log"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"gopkg.in/telebot.v3/middleware"
)

type Bot struct {
	*tele.Bot
	*layout.Layout
}

// New inits bot
func New(path string) (*Bot, error) {
	lt, err := layout.New(path)
	if err != nil {
		return nil, err
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		return nil, err
	}

	if cmds := lt.Commands(); cmds != nil {
		if err := b.SetCommands(cmds); err != nil {
			return nil, err
		}
	}

	return &Bot{
		Bot:    b,
		Layout: lt,
	}, nil
}

// common Inline handler
func (b Bot) onQuery(c tele.Context) error {
	command, data := b.parseQuery(c.Data())

	log.Println("got inline rn")
	switch command {
	case "doc", "d":
		return b.onDoc(c, data)
	default:
		return b.onHelp(c)
	}
}

func (b *Bot) Start() {
	// Middlewares
	b.Use(middleware.Logger())
	b.Use(middleware.AutoRespond())
	b.Use(b.Middleware("ru"))

	// Handlers
	b.Handle("/start", b.onStart)
	b.Handle(tele.OnQuery, b.onQuery)

	b.Bot.Start()
}
