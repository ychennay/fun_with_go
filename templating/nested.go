package templating

import (
	"log"
	"os"
	"text/template"
)

var nestedTmp *template.Template

type nestedSage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom         []nestedSage
	Transportation []car
}

func init() {
	nestedTmp = template.Must(template.ParseFiles("templating/nested.html"))
}

func NestFill() {
	colonelSanders := nestedSage{Name: "Jeff", Motto: "Eat mo' chicken."}
	eminem := nestedSage{
		Name:  "Marshall Mathers",
		Motto: "Will the real slim shady please stand up?",
	}

	fordMustang := car{
		Manufacturer: "Ford",
		Model:        "Mustang",
		Doors:        4,
	}

	data := items{
		Wisdom:         []nestedSage{colonelSanders, eminem},
		Transportation: []car{fordMustang},
	}

	err := nestedTmp.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
