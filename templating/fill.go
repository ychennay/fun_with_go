package templating

import (
	"log"
	"os"
	"text/template"
)

var advTpl *template.Template

func init() {
	advTpl = template.Must(template.ParseFiles("templating/boing.html"))
}

func SageFill() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := advTpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
