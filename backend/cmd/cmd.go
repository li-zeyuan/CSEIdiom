package cmd

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/li-zeyuan/CSEIdiom/backend/config"
	"github.com/li-zeyuan/CSEIdiom/backend/router"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	configFile string
)

func NewCmd() *cobra.Command {
	serveCmd := cobra.Command{
		Use:   "run",
		Short: "run",
		Long:  "Run service.",
		Run:   run,
	}
	serveCmd.Flags().StringVar(&configFile, "config", "../../../sun/governmentexam/backend_config.yaml", "config file")
	return &serveCmd
}

func run(cmd *cobra.Command, args []string) {
	err := config.LoadConfigFile(configFile)
	if err != nil {
		log.Fatal("Load config file failed, exit!", zap.String("file", configFile), zap.Error(err))
		return
	}

	err = mylogger.Init(&config.AppCfg.Logging)
	if err != nil {
		log.Fatal("init log error: ", err)
		return
	}

	err = mysqlstore.New(&config.AppCfg.Mysql)
	if err != nil {
		log.Fatal("init mysql error: ", err)
		return
	}

	go func() {
		mylogger.Info("government_exam server listen address", zap.String("address", config.AppCfg.ListenAddress))
		s := &http.Server{
			Addr:    config.AppCfg.ListenAddress,
			Handler: router.New(),
		}
		if err = s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			mylogger.Fatal("government_exam server listen fail: ", zap.Error(err))
		}

		defer func() {
			if err = s.Close(); err != nil {
				mylogger.Error("Error closing the server: ", zap.Error(err))
			}

		}()
	}()

	var sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-sigChan:
			mysqlstore.Close()

			mylogger.Info("Received SIGTERM, gracefully exit...")
			return
		}
	}
}
