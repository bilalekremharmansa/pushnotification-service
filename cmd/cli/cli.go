package cli

import (
	"fmt"
	"os"
	"log"

	"github.com/spf13/cobra"

	"bilalekrem.com/pushnotification-service/internal/config"

	"bilalekrem.com/pushnotification-service/api/rest"
)

var (
	cfgFile     string

	rootCmd = &cobra.Command{
		Use:   "pushnotification-service",
		Short: "push notification service",
		Long: ``,
	}

)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	runCmd := &cobra.Command{
        Use:   "run",
        Short: "run service",
        Run: func(cmd *cobra.Command, args []string) {
            log.Println("Starting server")
            server := rest.NewRestServerWithConfig()
            server.Start()
        },
    }
    runCmd.PersistentFlags().StringVar(&cfgFile, "config", "/tmp/config.yaml", "config file (default is /tmp/config.yaml)")
    rootCmd.AddCommand(runCmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
    log.Println(cfgFile)

    config.InitConfig(cfgFile)
}