package email

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"path/filepath"

	"github.com/satisfeet/go-pdf"
	"github.com/satisfeet/hoopoe/model"
)

const (
	ImageDir    = "images"
	TemplateDir = "templates"
)

var holder = struct {
	Name      string
	Street    string
	Place     string
	Email     string
	Web       string
	Phone     string
	TaxNumber string
	IBAN      string
	BIC       string
}{
	"Bodo Kaiser",
	"Geiserichstr. 3",
	"12105 Berlin",
	"info@satisfeet.me",
	"www.satisfeet.me",
	"+49 162 2635326",
	"DE291325845",
	"DE67100900002451723009",
	"BEVODEBB",
}

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

var tmpl = template.Must(template.ParseGlob(
	filepath.Join(TemplateDir, "*.tmpl"),
))

func NewInvoice(m model.Order) io.WriterTo {
	d := pdf.NewDocument(pdf.DocumentInfo{
		Title:  "Invoice",
		Author: "Bodo Kaiser",
		Header: header,
		Footer: footer,
	})

	d.Text(56, 42, styleMd, render("customer", m))
	d.Text(10, 90, styleLg, render("subject", m))
	d.Text(10, 105, styleSm, render("content", m))

	d.TableRow(10, 170, styleTableHead, []string{
		"Artikel", "Größe", "Farbe", "Anzahl", "Stückpreis", "Gesamtpreis",
	})

	for _, oi := range m.Items {
		d.TableRow(10, 180, styleTableBody, []interface{}{
			oi.Product.Name,
			oi.Variation.Size,
			oi.Variation.Color,
			oi.Quantity,
			oi.Product.Pricing.String(),
			oi.Pricing.String(),
		})
	}

	return d
}

func header(d *pdf.Document) {
	d.Image(0, 0, pdf.Style{Width: 140}, "images/brand.png")

	d.Text(160, 40, styleSm, holder.Name)
	d.Text(160, 46, styleSm, holder.Street)
	d.Text(160, 50, styleSm, holder.Place)
	d.Link(160, 58, styleLink, holder.Email, "mailto:"+holder.Email)
	d.Link(160, 62, styleLink, holder.Web, "http://"+holder.Web)
	d.Link(160, 70, styleLink, holder.Phone, "tel:"+holder.Phone)
}

func footer(d *pdf.Document) {
	d.Text(0, -18, styleXs, fmt.Sprintf("%s\n\nIBAN: %s - BIC: %s",
		holder.Name,
		holder.IBAN,
		holder.BIC,
	))
	d.Text(178, -12, styleXs, fmt.Sprintf("Seite %d von {nb}",
		d.Page(),
	))
}

func render(name string, value interface{}) string {
	b := &bytes.Buffer{}

	tmpl.ExecuteTemplate(b, name, value)

	return b.String()
}
