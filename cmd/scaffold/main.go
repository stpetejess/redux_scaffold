package main

import (
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strings"

	args "github.com/alexflint/go-arg"
)

var (
	reg = regexp.MustCompile(`-`)
)

type Args struct {
	Name                   string `arg:"required" help:"name of folder to create"`
	SanitizedName          string `arg:"-"`
	UppercaseSanitizedName string `arg:"-"`
	Location               string `help:"where to create the js files...defaults to ."`
}

func (a *Args) Transform() {
	a.Name = strings.ToLower(a.Name)
	a.SanitizedName = reg.ReplaceAllString(a.Name, "")
	a.UppercaseSanitizedName = strings.ToUpper(a.SanitizedName)

	fmt.Printf("n: %s sn: %s usn: %s\n", a.Name, a.SanitizedName, a.UppercaseSanitizedName)
}

func main() {
	a := Args{}
	args.MustParse(&a)
	a.Transform()

	paths := []string{"constant", "action", "reducer", "style", "container", "component"}
	for _, path := range paths {
		// now do the template thing
		t, err := template.ParseFiles(fmt.Sprintf("../../templates/%s.tpl", path))
		if err != nil {
			fmt.Printf("Error parsing template: %s\n", err)
			return
		}
		fmt.Printf("Creating ./%s.js\n", path)

		f, err := os.Create(fmt.Sprintf("./%s.js", path))
		defer f.Close()
		if err != nil {
			fmt.Println("create file: ", err)
			return
		}
		if err = t.Execute(f, a); err != nil {
			fmt.Printf("Error executing template: %s\n", err)
			return
		}
		fmt.Printf(`Success!`)
	}
	return
}
