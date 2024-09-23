package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/moisespsena-go/tree2html"
)

func Node(child ...*tree2html.Tree) *tree2html.Tree {
	return tree2html.Node(nil, child...)
}

func main() {
	f, err := os.Create("example/example.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("<!DOCTYPE html>\n<html><head><meta charset=\"UTF-8\"></head><body><h1>Example</h1>")

	write(f, tree2html.New(
		Node(),
		Node(),
	))

	write(f, tree2html.New(
		Node(
			Node(
				Node(
					Node(
						Node(
							Node(Node(Node(Node()))),
							Node(),
						),
					),
				),
			),
		),
	))

	write(f,
		tree2html.New(
			Node(
				Node(),
			),
			Node(
				Node(),
				Node(
					Node(
						Node(Node(Node())),
						Node(),
					),
					Node(),
				),
				Node(),
			),
		))

	f.WriteString("</body></html>")
}

func write(f *os.File, tree *tree2html.Tree) {
	// make labels based on indexes
	tree.Walk(func(p []*tree2html.Tree, t *tree2html.Tree, i int) {
		t.Value = buildLabel(append(p[1:], t))
	})

	// Vertical
	f.WriteString("<p>The Tree:</p>\n<pre>")
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	enc.Encode(tree.Children)
	f.WriteString("</pre>\n")

	// Vertical
	f.WriteString("<p>Vertical</p>\n")
	f.WriteString(`<table style="text-align:center" border="1"><tbody>` + "\n")
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
