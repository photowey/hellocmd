package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/photowey/hellocmd/pkg/regexz"
)

func DetermineModPath(projectPath string) (modPath string) {
	dir := filepath.Dir(projectPath)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod.tmpl")); err == nil {
			content, _ := ioutil.ReadFile(filepath.Join(dir, "go.mod.tmpl"))
			mod := regexz.RegexpExtract(`module\s+(?P<alias>[\S]+)`, string(content), "$alias")
			name := strings.TrimPrefix(filepath.Dir(projectPath), dir)
			name = strings.TrimPrefix(name, string(os.PathSeparator))
			if name == "" {
				return fmt.Sprintf("%s/", mod)
			}
			return fmt.Sprintf("%s/%s/", mod, name)
		}
		parent := filepath.Dir(dir)
		if dir == parent {
			return ""
		}
		dir = parent
	}
}
