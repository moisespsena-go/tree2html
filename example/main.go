package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/moisespsena-go/tree2html"
)

var New = tree2html.New

func main() {
	f, err := os.Create("example.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("<!DOCTYPE html>\n<html><head><meta charset=\"UTF-8\"></head><body><h1>Example</h1>")

	write(f, New("",
		New(""),
		New(""),
	).Build())

	write(f,
		New("",
			New("",
				New("",
					New(""),
					New("",
						New(""),
						New(""),
						New("",
							New("", New("", New(""))),
							New(""),
						),
						New(""),
					),
					New(""),
				),
				New("",
					New(""),
					New(""),
				),
			),
			New("",
				New(""),
				New("",
					New(""),
					New(""),
					New("",
						New("", New("", New(""))),
						New(""),
					),
					New(""),
				),
				New(""),
			),
		).Build())

	f.WriteString("</body></html>")
}

func write(f *os.File, tree *tree2html.Tree) {
	// make labels based on indexes
	tree.Walk(func(p []*tree2html.Tree, t *tree2html.Tree, i int) {
		t.Value = buildLabel(append(p[1:], t))
	})

	// Vertical
	f.WriteString("<p>Vertical</p>\n")
	f.WriteString("<table border=\"1\"><tbody>\n")
	tree.VTable().WriteTo(f)
	f.WriteString("</tbody></table>\n")

	// Horizontal
	f.WriteString("<p>Horizontal</p>\n")
	f.WriteString("<table border=\"1\"><tbody>\n")
	tree.HTable().WriteTo(f)
	f.WriteString("</tbody></table>\n")
	f.WriteString("<hr />\n")
}

func buildLabel(t []*tree2html.Tree) string {
	indexes := make([]string, len(t))
	for i, t := range t {
		indexes[i] = strconv.Itoa(t.Index() + 1)
	}
	return strings.Join(indexes, ".")
}
