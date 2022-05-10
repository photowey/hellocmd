package new

type Project struct {
	Path      string // project gen path
	Name      string // project name
	ModPrefix string // mod prefix
	ModName   string // mod name
}

var (
	project            Project
	DefaultProjectName = "helloapp"
)
