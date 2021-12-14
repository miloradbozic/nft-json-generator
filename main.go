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
	list, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("error reading input .csv file: %v", err)
		os.Exit(1)
	}

	creators := os.Args[2:]
	generateNftJsons(string(list), creators, true)
}

func generateNftJsons(list string, creators []string, paralelize bool) {
	var wg sync.WaitGroup

	if err := ensureDir("./output"); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}

	lines := strings.Split(list, "\r\n")
	traitNames := getTraitNames(lines[0])
	for i, line := range lines[1:] {
		traitValues := strings.Split(line, ",")
		var traits []Trait
		for i, traitName := range traitNames {
			traits = append(traits, Trait{traitName, traitValues[i]})
		}
		nft := Nft{Name: fmt.Sprintf("NFT %d", i+1), Traits: traits, Creators: creators}

		if paralelize {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				generateJson(i, nft)
			}(i)
		} else {
			generateJson(i, nft)
		}
	}

	wg.Wait()
}

func getTraitNames(header string) []string {
	return strings.Split(header, ",")
}

func generateJson(i int, nft Nft) {
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
