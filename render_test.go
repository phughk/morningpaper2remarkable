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

	// When parsed
	item.Content = "<h1>title</h1><br><p>paragraph contents</p>"
	render(filename, item)

	// Then file is created
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		assert.Fail(t, fmt.Sprintf("File `%s` was not created", filename))
	}
}
