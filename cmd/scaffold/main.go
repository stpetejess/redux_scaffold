package main

import (
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strings"

	args "github.com/alexflint/go-arg"
	strcase "github.com/iancoleman/strcase"
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

func (a *Args) transform() {
	a.Name = strings.ToLower(a.Name)
	a.SanitizedName = strcase.ToCamel(a.Name) //reg.ReplaceAllString(a.Name, "")
	a.UppercaseSanitizedName = strings.ToUpper(a.SanitizedName)
	if len(a.Location) < 1 {
		a.Location = "."
	}

	log("n: %s sn: %s usn: %s loc: %s", a.Name, a.SanitizedName, a.UppercaseSanitizedName, a.Location)
}

func main() {
	a := Args{}
	args.MustParse(&a)
	a.transform()

	files := []string{"constant", "action", "reducer", "style", "container", "component"}
	for _, fname := range files {
		// now do the template thing
		t, err := template.ParseFiles(fmt.Sprintf("../../templates/%s.tpl", fname))
		if err != nil {
			log("Error parsing template: %s", err)
			return
		}
		pre := fmt.Sprintf("%s/%s", a.Location, strcase.ToKebab(a.SanitizedName))
		if _, err := os.Stat(pre); os.IsNotExist(err) {
			log("Creating folder %s", pre)
			os.Mkdir(pre, os.ModePerm)
		}
		path := fmt.Sprintf("%s/%s", pre, fname)

		log("Creating %s.js", path)

		f, err := os.Create(fmt.Sprintf("%s.js", path))
		defer f.Close()
		if err != nil {
			log("create file: %s", err)
			return
		}
		if err = t.Execute(f, a); err != nil {
			log("Error executing template: %s", err)
			return
		}
		log(`Success!`)
	}
	return
}

func log(msg string, alltheargs ...interface{}) {
	msg = fmt.Sprintf("\n%s\n", msg)
	fmt.Printf(msg, alltheargs...)
}
