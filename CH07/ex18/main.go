// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

//Node is node interface
type Node interface{}

//CharData is character data interface
type CharData string

//Element is element interface
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var temp Element

	root := Element{Type: xml.Name{Space: "root", Local: "root"}, Attr: []xml.Attr{}, Children: []Node{}}
	stack := []Element{root}
	i := 0
	for {
		i++
		if i > 15 {
			break
		}
		tok, err := dec.Token()
		if err == io.EOF {
			fmt.Println(stack[0].Type.Local)
			//root = stack[0].Children[0].(Element)
			root = stack[0]
			fmt.Println("Parse done!")
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			//fmt.Printf("len(stack): %d\n", len(stack))
			fmt.Println(tok.Name.Local)
			child := Element{Type: tok.Name, Attr: tok.Attr, Children: []Node{}}
			fmt.Printf("stack[stlen-1].children: %v\n", stack[len(stack)-1].Children)
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, child)
			stack = append(stack, child)
			fmt.Printf("sstack[].c[0]: %v\n", stack[0].Children[0])

			if len(stack) == 6 {
				temp = stack[1]
				fmt.Print("aaaa")
				showChild(temp)
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			//fmt.Println(string(tok.Copy()))
			children := stack[len(stack)-1].Children
			stack[len(stack)-1].Children = append(children, tok.Copy())
		}
	}
	showChild(temp)
	//	showChild(stack[0])
}

func showChild(node Node) {
	switch nn := node.(type) {
	case Element:
		fmt.Printf("<%s>\n", nn.Type.Local)
		for _, c := range nn.Children {
			fmt.Printf("<%s>\n", c)
			showChild(c)
		}
	case CharData:
		fmt.Println(nn)
	}
}
