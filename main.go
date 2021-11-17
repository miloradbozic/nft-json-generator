package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
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
	var wg sync.WaitGroup

	creators := os.Args[2:]

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("error reading input .csv file: %v", err)
		os.Exit(1)
	}

	if err := ensureDir("./output"); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}

	list := string(data)
	lines := strings.Split(list, "\r\n")
	traitNames := getTraitNames(lines[0])
	for i, line := range lines[1:] {
		traitValues := strings.Split(line, ",")
		var traits []Trait
		for i, traitName := range traitNames {
			traits = append(traits, Trait{traitName, traitValues[i]})
		}
		nft := Nft{Name: fmt.Sprintf("Nft %d", i), Traits: traits, Creators: creators}

		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			generateFile(i, nft)
		}(i)

	}

	wg.Wait()
}

func getTraitNames(header string) []string {
	return strings.Split(header, ",")
}

func generateFile(i int, nft Nft) {
	fmt.Println("Generate ", i, "starting")
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
	fmt.Println("Generate ", i, "done")
}

func ensureDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModeDir)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}
