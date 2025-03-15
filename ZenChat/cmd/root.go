// Package cmd /*
package cmd

import (
	"errors"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/common/sdk"
	"github.com/mazezen/zenchat/internel/router"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	c       string
	rootCmd = &cobra.Command{
		Use:   "ZenChat",
		Short: "golang IM application",
		Long:  `use when you need to create an new IM application`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Args: func(cmd *cobra.Command, args []string) error {
			if c == "" {
				return errors.New("require at least one arg")
			}
			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			sdk.ParseConfig(c)
			bootLogger()
			//vm.NewLogger()
			bootDb()
			bootRedis()

			errs := make(chan error, 1)
			if err := bootHttpServer(); err != nil {
				errs <- err
			}

			select {
			case err := <-errs:
				if err != nil {
					panic(err)
				}
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func bootLogger() {
	itools.NewLogger(
		itools.WithLoggerFilePath(sdk.GetConf().Logger.FilePath),
		itools.WithLoggerMaxSize(sdk.GetConf().Logger.MaxSize),
		itools.WithLoggerMaxAge(sdk.GetConf().Logger.MaxAge),
		itools.WithLoggerLocalTime(),
		itools.WithLoggerCompress(),
	).Start()
}

func bootDb() {
	itools.NewXrmEngine(
		itools.WithXrmHost(sdk.GetConf().Db.Host),
		itools.WithXrmPort(sdk.GetConf().Db.Port),
		itools.WithXrmUsername(sdk.GetConf().Db.Username),
		itools.WithXrmPassword(sdk.GetConf().Db.Password),
		itools.WithXrmDatabase(sdk.GetConf().Db.Database),
		itools.WithXrmCharset(sdk.GetConf().Db.Charset),
		itools.WithXrmMaxIdleConn(sdk.GetConf().Db.MaxIdleConn),
		itools.WithXrmMaxOpenConn(sdk.GetConf().Db.MaxOpenConn),
		itools.WithXrmShowSql(sdk.GetConf().Db.ShowSql)).Connect()
}

func bootRedis() {
	itools.NewRedisClient(
		itools.WithRedisAddress(sdk.GetConf().Redis.Addr),
		itools.WithRedisPassword(sdk.GetConf().Redis.Password),
		itools.WithRedisDB(sdk.GetConf().Redis.DB),
	).Connect()
}

func bootHttpServer() error {
	return router.RunHttpServer()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&c, "config", "c", "./config.yaml", "config file")
}
