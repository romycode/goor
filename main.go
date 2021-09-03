package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	s := `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>title</title>
    <link rel="stylesheet" href="style.css">
    <script>
		window.onload = function() {
			console.info("!! <=> LOADED  <=> !!")
		};
	</script>
  </head>
  <body>
    <ul id="list" class="pretty-list">
		<li class="pretty-element-list">Element 1</li>
		<li class="pretty-element-list">Element 2</li>
		<li class="pretty-element-list">Element 3</li>
	</ul>
  </body>
</html>
`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	iterateNodes(doc)
}

func iterateNodes(n *html.Node) {
	for e := n.FirstChild; e != nil; e = e.NextSibling {
		if e.Type == html.ElementNode {
			fmt.Println("# <==============================> #")
			fmt.Printf("TAG: %v \n", e.Data)
			fmt.Printf("ATTRIBUTES: %#v \n", func() map[string]string {
				attr := map[string]string{}
				for _, attribute := range e.Attr {
					attr[attribute.Key] = attribute.Val
				}
				return attr
			}())
			fmt.Println("# <==============================> #")
		}

		iterateNodes(e)
	}

	//if n.Type == html.ElementNode {
	//buffer := new(bytes.Buffer)
	//_ = html.Render(buffer, n)

	//fmt.Println(string(buffer.Bytes()))
	//}
}
