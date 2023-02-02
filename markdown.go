package main

import (
	"fmt"
	"strings"
)

func main() {
	d := NewDoc()
	d.AddTitle("test", 3).
		AddTitle("t*es#t2", 1)
	fmt.Println(d.content)
}

type markdownDoc struct {
	content *strings.Builder
}

func NewDoc() *markdownDoc {
	return &markdownDoc{
		content: &strings.Builder{},
	}
}

func (doc *markdownDoc) writeLine(content string) {
	doc.content.WriteString(content + "\n")
}

func (doc *markdownDoc) AddTitle(t string, lv int) *markdownDoc {
	if lv > 6 || lv < 1 {
		fmt.Sprintf("failed to add Title %s in level: %d", t, lv)
		return doc
	}
	t = replaceEscapeCharacter(t)
	mdSyntax := strings.Repeat("#", lv) + " " + t
	doc.writeLine(mdSyntax)
	return doc
}

func replaceEscapeCharacter(content string) string {
	escapeCharacters := []string{
		"\\", "`", "*",
		"_", "{", "}",
		"[", "]", "(",
		")", "#", "+",
		"-", ".", "!",
	}
	for _, c := range escapeCharacters {
		content = strings.Replace(content, c, "\\"+c, -1)
	}
	return content
}
