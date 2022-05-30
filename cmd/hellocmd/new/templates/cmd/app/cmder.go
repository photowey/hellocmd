package cmder

import (
	"fmt"
	"os"

	"codeup.aliyun.com/uphicoo/gokit/stringz"
	"github.com/spf13/cobra"

	"uphicoo.com/uphicoo/project-template/internal/app"
)

const (
	DefaultConfigFile = "./config.toml"
)

var (
	conf string

	root = &cobra.Command{
		Use:   "project-template",
		Short: "项目模板服务",
		Long:  "项目模板服务",
		Run: func(cmd *cobra.Command, args []string) {
			// Do nothing
		},
	}
)

func init() {
	cobra.OnInitialize(startApp)
	// e.g.: app start -f ./config.toml
	root.PersistentFlags().StringVarP(&conf, "conf", "f", "", "toml格式配置文件路径")
	root.AddCommand(start)
}

// Run 启动 App
func Run() {
	if err := root.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startApp() {
	if conf == stringz.DefaultEmptyString {
		conf = DefaultConfigFile
	}

	// 1.读取配置文件-启动 App
	if err := app.Start(conf); err != nil {
		cobra.CheckErr(err)
	}
}
