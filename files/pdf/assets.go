package pdf

import (
	"bytes"
	"text/template"

	"github.com/satisfeet/go-pdf"
)

const path = "../../share/"

var styleXs = pdf.Style{
	Height:    3,
	Font:      "Helvetica",
	FontSize:  8,
	FontColor: [3]int{60, 60, 60},
}

var styleSm = pdf.Style{
	Height:    4,
	Font:      "Helvetica",
	FontSize:  9,
	FontColor: [3]int{60, 60, 60},
}

var styleMd = pdf.Style{
	Height:    5,
	Font:      "Helvetica",
	FontSize:  10,
	FontColor: [3]int{60, 60, 60},
}

var styleLg = pdf.Style{
	Height:    6,
	Font:      "Helvetica",
	FontSize:  12,
	FontColor: [3]int{60, 60, 60},
}

var styleLink = pdf.Style{
	Height:    6,
	FontColor: [3]int{70, 130, 180},
}

var styleTableHead = pdf.Style{
	Width:     32,
	Height:    6,
	Font:      "Helvetica",
	FontStyle: "B",
}

var styleTableBody = pdf.Style{
	Width:     32,
	Height:    4,
	Font:      "Helvetica",
	FontStyle: "",
}

var tmpl = template.Must(template.ParseGlob(path + "templates/*.tmpl"))

func render(name string, value interface{}) string {
	b := &bytes.Buffer{}

	if err := tmpl.ExecuteTemplate(b, name, value); err != nil {
		panic(err)
	}

	return b.String()
}
