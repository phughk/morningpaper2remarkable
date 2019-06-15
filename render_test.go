package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
)

func TestRenderStoresToFile(t *testing.T) {
	// Given a filename
	filename := "testRenderStoresToFile.pdf"

	// And contents
	item := &gofeed.Item{}
	item.Content = "<h1>title</h1><p>paragraph contents</p>"

	// When parsed
	render(filename, item)

	// File is created
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		assert.Fail(t, fmt.Sprintf("File `%s` was not created", filename))
	}
}
