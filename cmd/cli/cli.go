package cli

import (
	"fmt"
	"os"
	"log"
	"sync"

	"github.com/spf13/cobra"

	"bilalekrem.com/pushnotification-service/internal/config"
    "bilalekrem.com/pushnotification-service/internal/push"
    "bilalekrem.com/pushnotification-service/internal/push/firebaseadminsdk"

	"bilalekrem.com/pushnotification-service/api/rest"
	"bilalekrem.com/pushnotification-service/api/grpc"
)

var (
	cfgFile     string

	// --
	token     string
	title     string
	body     string
	serviceAccountFilePath     string

	rootCmd = &cobra.Command{
		Use:   "pns",
		Short: "push notification service",
		Long: ``,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	serverCmd := &cobra.Command{
        Use:   "server",
        Short: "run server(s) (default [rest, gRPC] enabled)",
        Run: func(cmd *cobra.Command, args []string) {
            log.Println("Starting servers")

            initConfig()
            startServers()
        },
    }

    serverCmd.
        PersistentFlags().
        StringVar(&cfgFile, "config", "", fmt.Sprintf("config file path (default: '%s')", config.GetDefaultConfigPath()))
    rootCmd.AddCommand(serverCmd)

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

    cmd.PersistentFlags().StringVar(&serviceAccountFilePath, "service-account-file",
            config.GetDefaultServiceAccountFilePath(), "service account file path (json format)")

    rootCmd.AddCommand(cmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
    if cfgFile != "" {
        config.InitConfig(cfgFile)
        return
    }

    defaultConfigFilePath := config.GetDefaultConfigPath()
    if fileExists(defaultConfigFilePath) {
        config.InitConfig(defaultConfigFilePath)
        return
    }

    config.InitDefaultConfig()
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func startServers() {
    serversConfig := config.GetAppConfig().GetServersConfig()

    // --

    var wg sync.WaitGroup

    // -- rest server
    wg.Add(1)
    go func() {
        defer wg.Done()

        restConfig := serversConfig.GetRestConfig()
        restServer := rest.NewRestServerWithConfig(restConfig)
        restServer.Start()
    }()

    // -- gRPC server
    wg.Add(1)
    go func() {
        defer wg.Done()

        gRPCConfig := serversConfig.GetGRPCConfig()
        gRPCServer := grpc.NewGRPCServerWithConfig(gRPCConfig)
        gRPCServer.Start()
    }()

    wg.Wait()
}
