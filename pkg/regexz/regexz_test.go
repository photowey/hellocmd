package regexz

import (
	"testing"
)

func TestRegexpReplace(t *testing.T) {
	type args struct {
		regex string
		src   string
		temp  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test regex replace",
			args: args{
				regex: `module\s+(?P<name>[\S]+)`,
				src:   "module github.com/photowey/hellocmd\n\ngo 1.18\n\nrequire (\n\tgithub.com/gobuffalo/packr/v2 v2.8.3\n\tgithub.com/urfave/cli v1.22.9\n)\n\nrequire (\n\tgithub.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect\n\tgithub.com/gobuffalo/logger v1.0.6 // indirect\n\tgithub.com/gobuffalo/packd v1.0.1 // indirect\n\tgithub.com/karrick/godirwalk v1.16.1 // indirect\n\tgithub.com/markbates/errx v1.1.0 // indirect\n\tgithub.com/markbates/oncer v1.0.0 // indirect\n\tgithub.com/markbates/safe v1.0.1 // indirect\n\tgithub.com/russross/blackfriday/v2 v2.0.1 // indirect\n\tgithub.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect\n\tgithub.com/sirupsen/logrus v1.8.1 // indirect\n\tgolang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect\n\tgolang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect\n)\n",
				temp:  "$name",
			},
			want: "github.com/photowey/hellocmd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegexpExtract(tt.args.regex, tt.args.src, tt.args.temp); got != tt.want {
				t.Errorf("RegexpExtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
