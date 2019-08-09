package templating

import (
	"log"
	"os"
	"text/template"
)

var structTpl *template.Template

type sage struct {
	Name  string
	Motto string
	Age   int
}

func init() {
	structTpl = template.Must(template.ParseFiles("templating/structy.html"))
}

func StructFiller() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "Get rich or die trying",
		Age:   34,
	}

	gandhi := sage{
		Name:  "G-Unit",
		Motto: "Yo momma!",
		Age:   -1,
	}

	sages := []sage{gandhi, buddha}

	err := structTpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatal(err)
	}
}
