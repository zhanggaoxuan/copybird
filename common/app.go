package common

import (
	"github.com/copybird/copybird/core"
	"github.com/copybird/copybird/modules/backup/compress/gzip"
	"github.com/copybird/copybird/modules/backup/compress/lz4"
	"github.com/copybird/copybird/modules/backup/encrypt/aesgcm"
	"github.com/copybird/copybird/modules/backup/input/mongodb"
	"github.com/copybird/copybird/modules/backup/input/mysql"
	postgres "github.com/copybird/copybird/modules/backup/input/postgresql"
	"github.com/copybird/copybird/modules/backup/output/gcp"
	"github.com/copybird/copybird/modules/backup/output/http"
	"github.com/copybird/copybird/modules/backup/output/local"
	"github.com/copybird/copybird/modules/backup/output/s3"
	"github.com/copybird/copybird/modules/backup/output/scp"
	"github.com/copybird/copybird/operator"
	"github.com/spf13/cobra"
	"log"
	//"log"
	//"github.com/spf13/cobra"
)

type App struct {
	cmmRoot     *cobra.Command
	cmdBackup   *cobra.Command
	cmdOperator *cobra.Command
	vars        map[string]interface{}
}

func NewApp() *App {
	return &App{
		vars: make(map[string]interface{}),
	}
}

func (a *App) Run() error {
	a.registerModules()
	var rootCmd = &cobra.Command{Use: "copybird"}
	a.cmdBackup = &cobra.Command{
		Use:   "backup",
		Short: "Start new backup",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run:   cmdCallback(a.DoBackup),
	}
	a.cmdOperator = &cobra.Command{
		Use:   "operator",
		Short: "Start Kubernetes operator",
		Run: func(cmd *cobra.Command, args []string) {
			operator.Run()
		},
	}
	rootCmd.AddCommand(a.cmdBackup)
	rootCmd.AddCommand(a.cmdOperator)
	a.Setup()
	return rootCmd.Execute()
}

func (a *App) DoBackup() error {
	return nil
}

func cmdCallback(f func() error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := f()
		if err != nil {
			log.Printf("cmd err: %s", err)
		}
	}
}

func (a *App) registerModules() {
	core.RegisterModule(&mysql.BackupInputMysql{})
	core.RegisterModule(&postgres.BackupInputPostgresql{})
	core.RegisterModule(&mongodb.BackupInputMongodb{})
	core.RegisterModule(&gzip.BackupCompressGzip{})
	core.RegisterModule(&lz4.BackupCompressLz4{})
	core.RegisterModule(&aesgcm.BackupEncryptAesgcm{})
	core.RegisterModule(&gcp.BackupOutputGcp{})
	core.RegisterModule(&http.BackupOutputHttp{})
	core.RegisterModule(&local.BackupOutputLocal{})
	core.RegisterModule(&s3.BackupOutputS3{})
	core.RegisterModule(&scp.BackupOutputScp{})
}
