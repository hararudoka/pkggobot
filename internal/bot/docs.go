package bot

import (
	"fmt"
	"os/exec"
	"strings"
	"unicode/utf8"

	"github.com/hararudoka/pkggobot/internal/service"
	tele "gopkg.in/telebot.v3"
)

// TODO: fix formatting
func (b Bot) onDoc(c tele.Context, a string) error {
	results := tele.Results{}

	title, text, url := "", "", ""

	args := strings.Split(a, ".")

	if len(args) == 0 {

	}

	if true { // TODO сделать
		if len(args) == 1 {
			_ = doc
		}

		res := b.Result(c, "doc", map[string]interface{}{
			"Title": "found: "+title,
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



func createDoc(str string) (title, text, url string) {
	out, err := exec.Command("/usr/local/go/bin/go", "doc", str).Output()
	if err != nil {
		fmt.Println(err)
	}

	title = str

	slice := strings.Split(str, ".")
	text = validate(string(out), len(slice))

	url = "https://pkg.go.dev/"
	if len(slice) >= 1 {
		url += slice[0]
		if len(slice) >= 2 {
			url += "#" + slice[1]
		}
	}
	return
}

func validate(text string, t int) string {
	if text == "" {
		return ""
	}

	if t == 1 {
		text = strings.Trim(text, "\n")
		sl := strings.Split(text, "\n")
		sl = sl[1:]
		text = strings.Join(sl, "\n")
		sl = strings.Split(text, ".")
		text = sl[0]
		text = strings.ReplaceAll(text, "\n", " ")+"."

	} else if t == 2 {

		//text = strings.Trim(text, "\n")
		sl := strings.Split(text, "\n")
		fmt.Println(len(sl))

		var sl2 []string
		for _, e := range sl {
			if !strings.Contains(e, "//") {
				sl2 = append(sl2, e)
			}
		}

		sl = sl2[2:]
		text = strings.Join(sl, "\n")

		text = strings.ReplaceAll(text, "\n\n", "\n")
	} else if t == 3 {

	}

	if utf8.RuneCount([]byte(text)) >= 400 {
		return string([]byte(text))
	}

	return "``` "+text+"```"
}