package pdf

import (
	"io"

	"github.com/satisfeet/go-pdf"
	"github.com/satisfeet/hoopoe/store"
)

func NewInvoice(m store.Order) io.WriterTo {
	d := pdf.NewDocument(pdf.DocumentInfo{
		Title:   "Invoice",
		Author:  "Bodo Kaiser",
		Header:  header,
		Footer:  footer,
		FontDir: path + "fonts",
	})

	d.Text(56, 42, styleMd, render("customer", m))
	d.Text(10, 90, styleLg, render("subject", m))
	d.Text(10, 105, styleSm, render("content", m))

	d.TableRow(10, 170, styleTableHead, []string{
		"Artikel", "Größe", "Farbe", "Anzahl", "Stückpreis", "Gesamtpreis",
	})

	for _, i := range m.Items {
		d.TableRow(10, 180, styleTableBody, []interface{}{
			i.Product.Name,
			i.Variation.Size,
			i.Variation.Color,
			i.Quantity,
			i.Product.Pricing.String(),
			i.Pricing.String(),
		})
	}

	return d
}

func header(d *pdf.Document) {
	d.Image(0, 0, pdf.Style{Width: 140}, path+"images/brand.png")

	e := "info@satisfeet.me"
	w := "www.satisfeet.me"
	p := "+49 162 2635327"

	d.Text(160, 40, styleSm, render("company", nil))
	d.Link(160, 64, styleLink, e, "mailto:"+e)
	d.Link(160, 68, styleLink, w, "http://"+w)
	d.Link(160, 72, styleLink, p, "tel:"+p)
}

func footer(d *pdf.Document) {
	d.Text(0, -18, styleXs, render("account", nil))
	d.Text(178, -12, styleXs, render("pages", d))
}
