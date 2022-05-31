package new

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr/v2"
	"github.com/urfave/cli"

	"github.com/photowey/hellocmd/cmd/hellocmd/helper"
)

// CreateProject create a template project for Hellocmd
func CreateProject(ctx *cli.Context) (err error) {
	newArgs := ctx.Args()
	if len(newArgs) <= 0 {
		fmt.Println("Command line new execution error, please use [hellocmd new -h] for details")
		return
	}
	name := newArgs[0]

	project.Name = DefaultProjectName
	if name != "" && len(strings.TrimSpace(name)) > 0 {
		project.Name = name
	}

	if project.Path != "" {
		if project.Path, err = filepath.Abs(project.Path); err != nil {
			return
		}
		project.Path = filepath.Join(project.Path, project.Name)
	} else {
		pwd, _ := os.Getwd()
		project.Path = filepath.Join(pwd, project.Name)
	}
	modPath := helper.DetermineModPath(project.Path)
	fmt.Println("new project modPrefix:", modPath)
	project.ModPrefix = modPath
	project.ModName = project.ModPrefix + name

	if err = doCreateProject(); err != nil {
		return
	}

	fmt.Println("---------------- hellocmd new report ----------------")
	fmt.Println("Project dir:", project.Path)
	fmt.Println("Project created successfully")
	fmt.Println("Run cmd:")
	fmt.Println("$ cd " + project.Path)
	fmt.Println("$ go mod tidy")
	fmt.Println("---------------- hellocmd new report ----------------")

	return
}

//go:generate packr2
func doCreateProject() (err error) {
	box := packr.New("all", "./templates")
	if err = os.MkdirAll(project.Path, 0755); err != nil {
		return
	}
	for _, name := range box.List() {
		if project.ModPrefix != "" && name == "go.mod.tmpl" {
			// continue // not skipping
		}
		tmpl, _ := box.FindString(name)
		i := strings.LastIndex(name, string(os.PathSeparator))
		if i > 0 {
			dir := name[:i]
			if err = os.MkdirAll(filepath.Join(project.Path, dir), 0755); err != nil {
				return
			}
		}
		name = strings.TrimSuffix(name, ".tmpl")
		if err = doWriteFile(filepath.Join(project.Path, name), tmpl); err != nil {
			return
		}
	}

	return
}

func doWriteFile(path, tmpl string) (err error) {
	data, err := parseTmpl(tmpl)
	if err != nil {
		return
	}
	fmt.Println("File generated----------------------->", path)

	return ioutil.WriteFile(path, data, 0755)
}

func parseTmpl(tmpl string) ([]byte, error) {
	tmp, err := template.New("").Parse(tmpl)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = tmp.Execute(&buf, project); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
