package main

import (
	"github.com/mmcdole/gofeed"
)

func render(filename string, item *gofeed.Item) {
	pdf := &PdfCreatorState{}
	pdf.Init()
	pdf.Title("title")
	pdf.Paragraph("<p>paragraph</p>")
	pdf.Save(filename)
}
