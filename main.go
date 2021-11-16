package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Nft struct {
	Name     string
	Traits   []Trait
	Creators []string
}

type Trait struct {
	Name  string
	Value string
}

var fns = template.FuncMap{
	"plus1": func(x int) int {
		return x + 1
	},
}

func main() {
	creators := os.Args[1:]

	data, err := os.ReadFile("list.csv")
	if err != nil {
		fmt.Printf("error %v", err)
		os.Exit(1)
	}

	list := string(data)
	lines := strings.Split(list, "\r\n")
	traitNames := getTraitNames(lines[0])
	for i, line := range lines {
		traitValues := strings.Split(line, ",")
		var traits []Trait
		for i, traitName := range traitNames {
			traits = append(traits, Trait{traitName, traitValues[i]})
		}
		nft := Nft{Name: fmt.Sprintf("Nft %d", i), Traits: traits, Creators: creators}
		generateFile(i, nft)
	}

}

func getTraitNames(header string) []string {
	return strings.Split(header, ",")
}

func generateFile(i int, nft Nft) {
	paths := []string{
		"json.tmpl",
	}
	fmt.Printf("%v\n", nft)
	f, err := os.Create(fmt.Sprintf("./output/%d.json", i))
	if err != nil {
		fmt.Printf("Error creating file %d.json: %v\n", i, err)
	}
	t := template.Must(template.New("json.tmpl").Funcs(fns).ParseFiles(paths...))
	err = t.Execute(f, nft)
	//t.Execute(os.Stdout, nft)
	if err != nil {
		panic(err)
	}
}
