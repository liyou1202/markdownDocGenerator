package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var exportPath string
	for i, arg := range os.Args[1:] {

		if i >= 2 {
			fmt.Println("accept only one argument file Path, e.g. /path/to/export/doc.md")
			return
		}
		exportPath = arg
	}

	d := NewDoc()
	d.AddTitle("test", 3).
		AddBlankLines(2).
		AddTitle("t*es#t2", 1).
		AddCodeBlock("var text = \"123\"", "go").
		AddLink("this is a link", "https://www.google.com")

	err := d.Export(exportPath)
	if err != nil {
		log.Fatal(err)
	}
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

func (doc *markdownDoc) AddBlankLines(lv int) *markdownDoc {
	if lv > 0 {
		for i := 1; i <= lv; i++ {
			doc.writeLine("")
		}
	}

	return doc
}

func (doc *markdownDoc) AddCodeBlock(code, language string) *markdownDoc {
	code = replaceEscapeCharacter(code)
	mdSyntax := fmt.Sprintf("``` %s\n%s\n```\n", language, code)
	doc.writeLine(mdSyntax)

	return doc
}

func (doc *markdownDoc) AddLink(text, path string) *markdownDoc {
	text = replaceEscapeCharacter(text)
	mdSyntax := fmt.Sprintf("[%s](%s)", text, path)
	doc.writeLine(mdSyntax)

	return doc
}

func (doc *markdownDoc) Export(filename string) error {
	return ioutil.WriteFile(filename, []byte(doc.content.String()), os.ModePerm)
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
