package main

import (
	"html/template"
	"os"
)

type Info struct {
	Name  string
	Age   int
	Loves []string
}

type Species struct {
	Human Info
	Dog   Info
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := Species{
		Human: Info{
			Name:  "Francesca Cantor",
			Age:   29,
			Loves: []string{"pizza", "cuddling", "napping", "video games"},
		},
		Dog: Info{
			Name:  "Lizzy",
			Age:   7,
			Loves: []string{"food", "cuddling", "borking", "chasing critters"},
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
