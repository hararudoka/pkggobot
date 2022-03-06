package bot

import (
	"fmt"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func (b Bot) parseQuery(text string) (command, data string) {
	parts := strings.Split(text, " ")
	if len(parts) >= 1 {
		command = strings.ToLower(parts[0])
		data = strings.Join(parts[1:], " ")
	}
	return
}

func (b Bot) messageString(m *tele.Message) string {
	return fmt.Sprintf("%d_%d", m.Chat.ID, m.ID)
}

func (b Bot) sendLoading(c tele.Context) (*tele.Message, error) {
	return b.Reply(
		c.Message(),
		b.Text(c, "loading"),
	)
}

func (b Bot) sendInlineError(c tele.Context, name string) error {
	return c.Answer(&tele.QueryResponse{
		SwitchPMText:      b.Text(c, name+"_inline"),
		SwitchPMParameter: name + "_inline",
		CacheTime:         -1,
		IsPersonal:        true,
	})
}

func (b Bot) sendCallbackError(c tele.Context, name string) error {
	return c.Respond(&tele.CallbackResponse{Text: b.Text(c, name+"_callback")})
}

func (b Bot) edit(msg tele.Editable, what interface{}, opts ...interface{}) error {
	_, err := b.Edit(msg, what, opts...)
	return err
}
