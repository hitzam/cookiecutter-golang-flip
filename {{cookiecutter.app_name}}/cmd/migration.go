package cmd

import (
	"fmt"
	"os"
	"time"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	goCoreLog "gitlab.com/flip-id/go-core/helpers/log"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/config"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/appcontext"
)

var migrateUpCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Up DB {{ cookiecutter.app_name }}",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.GetConfig()
		appCtx := appcontext.NewAppContext(c)
		mSource := getMigrateSource()

		doMigrate(appCtx, mSource, appcontext.DBDialectMysql, migrate.Up)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migratedown",
	Short: "Migrate Up DB {{ cookiecutter.app_name }}",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.GetConfig()
		appCtx := appcontext.NewAppContext(c)
		mSource := getMigrateSource()

		doMigrate(appCtx, mSource, appcontext.DBDialectMysql, migrate.Down)
	},
}

var migrateNewCmd = &cobra.Command{
	Use:   "migratenew [migration name]",
	Short: "Create new migration file",
	Long:  `Create new migration file on folder migrations/sql with timestamp as prefix`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mDir := "migrations/sql/"

		createMigrationFile(mDir, args[0])
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrateNewCmd)
}

func getMigrateSource() migrate.FileMigrationSource {
	source := migrate.FileMigrationSource{
		Dir: "migrations/sql",
	}

	return source
}

func doMigrate(appCtx *appcontext.AppContext, mSource migrate.FileMigrationSource, dbDialect string, direction migrate.MigrationDirection) error {
	logger := goCoreLog.GetLogger(nil)
	db, err := appCtx.GetDBInstance(dbDialect)
	if err != nil {
		logger.FormatLog("GetDBInstance", err, dbDialect).Error("Error connection to DB")
		return err
	}

	sqlDB, _ := db.DB()

	defer sqlDB.Close()

	total, err := migrate.Exec(sqlDB, dbDialect, mSource, direction)
	if err != nil {
		logger.FormatLog("Migrate", err, dbDialect).Error("Fail migration")
		return err
	}
	logger.Infof("Migrate Success, total migrated: %d", total)
	return nil
}

func createMigrationFile(mDir string, mName string) error {
	logger := goCoreLog.GetLogger(nil)
	var migrationContent = `-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
`
	filename := fmt.Sprintf("%d_%s.sql", time.Now().Unix(), mName)
	filepath := fmt.Sprintf("%s%s", mDir, filename)

	f, err := os.Create(filepath)
	if err != nil {
		logger.FormatLog("create file", err, filepath).Error("Error create migration file")
		return err
	}
	defer f.Close()

	f.WriteString(migrationContent)
	f.Sync()

	logger.Info("New migration file has been created")
	return nil
}
