package main

import "github.com/jung-kurt/gofpdf"

// PdfCreator is used to neatly interact with the creation of the PDF documents
type PdfCreator interface {
	Init()
	Save()
	Title(string)
	Paragraph(string)
}

// PdfCreatorState holds basic rendering state that must be initialise
type PdfCreatorState struct {
	pdf        gofpdf.Fpdf
	html       gofpdf.HTMLBasicType
	lineHeight float64
}

// Init this should probably be a constructor
func (state *PdfCreatorState) Init() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	state.pdf = *pdf
	state.lineHeight = 12
	state.pdf.SetFont("Helvetica", "", 20)
	state.pdf.SetFontSize(state.lineHeight)
	state.pdf.SetTextColor(20, 20, 20)
	state.pdf.AddPage()
	state.html = state.pdf.HTMLBasicNew()
}

// Title insert a title into the document
func (state *PdfCreatorState) Title(title string) {
	scaledLineHeight := float64(uint(state.lineHeight * 1.2))
	state.html.Write(scaledLineHeight, title)
}

// Paragraph insert a paragraph into the document
func (state *PdfCreatorState) Paragraph(content string) {
	state.html.Write(state.lineHeight, content)
}

// Save this will store the render to a file and close
func (state *PdfCreatorState) Save(filename string) {
	state.pdf.OutputFileAndClose(filename)
}
