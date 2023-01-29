package main

import (
	"bytes"
	"go/format"
	"os"
	"text/template"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

var (
	templateFile string
	valuesFile   string
	outputFile   string
	withTest     bool
)

func main() {
	pflag.StringVarP(&templateFile, "template", "t", "", "Code template file")
	pflag.StringVarP(&valuesFile, "value", "v", "", "Values file, yaml format")
	pflag.StringVarP(&outputFile, "output", "o", "", "Output file")
	pflag.Parse()

	rawTmpl, err := os.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}
	tmpl := template.Must(template.New("codegen").Funcs(template.FuncMap{
		"toSnake": toSnake,
		"toCamel": toCamel,
	}).Parse(string(rawTmpl)))

	buf := bytes.NewBuffer(nil)
	values := make(map[string]interface{})
	f, err := os.Open(valuesFile)
	if err != nil {
		panic(err)
	}
	if err := yaml.NewDecoder(f).Decode(values); err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, values); err != nil {
		panic(err)
	}
	raw, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(outputFile, raw, 0666); err != nil {
		panic(err)
	}
}

func toSnake(s string) string {
	out := []byte{}
	for i, b := range s {
		switch {
		case i == 0:
			out = append(out, toLower(byte(b)))
		case isUpper(byte(b)):
			out = append(out, '_', toLower(byte(b)))
		default:
			out = append(out, byte(b))
		}
	}
	return string(out)
}

func toCamel(s string) string {
	out := []byte{}
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch {
		case i == 0:
			out = append(out, toUpper(byte(b)))
		case b == '_':
			i++
			if i < len(s) {
				out = append(out, toUpper(byte(s[i])))
			}
		default:
			out = append(out, byte(b))
		}
	}
	return string(out)
}

func toUpper(b byte) byte {
	if isLower(b) {
		return 'A' + b - 'a'
	}
	return b
}
func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
func isLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}
func toLower(b byte) byte {
	if isUpper(b) {
		return 'a' + b - 'A'
	}
	return b
}
