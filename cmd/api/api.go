package main

import (
	"fmt"
	"os"
	"time"

	"github.com/karasunokami/otus-calendar/api"
	"github.com/spf13/cobra"
)

var (
	apiParamHost           string
	apiParamPort           string
	apiParamReadTimeout    uint8
	apiParamWriteTimeout   uint8
	apiParamMaxHeaderBytes int

	apiServerCmd = &cobra.Command{
		Use:   "api-server",
		Short: "Run REST API server",
		Run: func(cmd *cobra.Command, args []string) {
			config := &api.ServerConfig{
				Host:           apiParamHost,
				Port:           apiParamPort,
				ReadTimeout:    time.Duration(apiParamReadTimeout) * time.Second,
				WriteTimeout:   time.Duration(apiParamWriteTimeout) * time.Second,
				MaxHeaderBytes: apiParamMaxHeaderBytes,
			}

			fmt.Println("Starting API server...")
			apiServer, err := api.InitializeServer()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			apiServer.Start(config)
		},
	}
)

func init() {
	apiServerCmd.PersistentFlags().StringVar(&apiParamHost, "host", "0.0.0.0", "Listening host")
	apiServerCmd.PersistentFlags().StringVar(&apiParamPort, "port", "8080", "Listening port")
	apiServerCmd.PersistentFlags().Uint8Var(&apiParamReadTimeout, "read-timeout", 10, "Read timeout in seconds")
	apiServerCmd.PersistentFlags().Uint8Var(&apiParamWriteTimeout, "write-timeout", 10, "Write timeout in seconds")
	apiServerCmd.PersistentFlags().IntVar(&apiParamMaxHeaderBytes, "max-header-size", 1<<20, "Max header size in bytes")
}
