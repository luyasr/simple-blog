// Package cmd /*
package cmd

import (
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simple-blog",
	Short: "一个简单的博客系统",
	Long:  "一个简单的博客系统",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.OnInitialize(initIoc)
	rootCmd.AddCommand(startCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initIoc() {
	// 初始化Controller
	if err := ioc.Controller().Init(); err != nil {
		panic(err)
	}

	// 初始化Handler
	if err := ioc.HttpHandler().Init(); err != nil {
		panic(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.simple-blog.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
