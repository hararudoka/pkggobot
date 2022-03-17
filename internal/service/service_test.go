package service

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var url = "https://pkg.go.dev/strings"

func TestParseFunctions(t *testing.T) {
	page, err := getPage(url)
	if err != nil {
		panic(err)
	}

	fs := parseFunctions(page)

	log.Println(fs[0].Name)
}

func TestParseTypes(t *testing.T) {
	page, err := getPage(url)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name string
		want Type
	}{
		{
			name: "first test",
			want: Type{

				Symbol: Symbol{
					Signature:   "type Builder struct {\n\t// contains filtered or unexported fields\n}",
					Description: []string{"A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use. Do not copy a non-zero Builder."},
					Example:     "package main\n\nimport (\n\t\"fmt\"\n\t\"strings\"\n)\n\nfunc main() {\n\tvar b strings.Builder\n\tfor i := 3; i >= 1; i-- {\n\t\tfmt.Fprintf(&b, \"%d...\", i)\n\t}\n\tb.WriteString(\"ignition\")\n\tfmt.Println(b.String())\n\n}",
				},
				Methods: []Symbol{
					{
						Signature:   "func (b *Builder) Cap() int",
						Description: []string{"Cap returns the capacity of the builder's underlying byte slice. It is the total space allocated for the string being built and includes any bytes already written."},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := parseTypes(page)[0]
			fmt.Println(got)
			fmt.Println(test.want)
			if reflect.DeepEqual(got, test.want) { // this dont work, data actually not same
				t.Errorf("parseTypes()\nwant: %#v\ngot:  %#v", test.want, got)
			}
		})
	}
}
