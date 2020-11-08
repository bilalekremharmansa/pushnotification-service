package cli

import (
	"fmt"
	"os"
	"log"

	"github.com/spf13/cobra"

	"bilalekrem.com/pushnotification-service/internal/config"
    "bilalekrem.com/pushnotification-service/internal/push"
    "bilalekrem.com/pushnotification-service/internal/push/firebaseadminsdk"

	"bilalekrem.com/pushnotification-service/api/rest"
)

var (
	cfgFile     string

	// --
	token     string
	title     string
	body     string
	serviceAccountFilePath     string

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

    initSendPushCommand(rootCmd)
}

func initSendPushCommand(rootCommand *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "send",
        Short: "send a push notification",
        Run: func(cmd *cobra.Command, args []string) {
            service := firebaseadminsdk.NewWithServiceAccount(serviceAccountFilePath)

            notification := push.NewNotification(title, body)

            err := service.Send(notification, token)
            if err != nil {
                log.Fatalf("Sending push notification failed: [%v]", err)
            }
        },
    }

    cmd.PersistentFlags().StringVar(&token, "token", "", "push notification token to send push notification server")
    cmd.PersistentFlags().StringVar(&title, "title", "", "title of push notification")
    cmd.PersistentFlags().StringVar(&body, "body", "", "body of push notification")
    cmd.PersistentFlags().StringVar(&serviceAccountFilePath, "service-account-file", "", "path of service account file")

    rootCmd.AddCommand(cmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
    log.Println(cfgFile)

    config.InitConfig(cfgFile)
}