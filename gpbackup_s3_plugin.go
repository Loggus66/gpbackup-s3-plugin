package main

import (
	"fmt"
	"os"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/Loggus66/gpbackup-s3-plugin/s3plugin"
	"github.com/urfave/cli"
)

func main() {
	gplog.InitializeLogging("gpbackup_s3_plugin", "")
	app := cli.NewApp()
	cli.VersionFlag = cli.BoolTFlag{
		Name:  "version",
		Usage: "print version of gpbackup_s3_plugin",
	}
	app.Version = s3plugin.Version
	app.Usage = ""
	app.UsageText = "Not supported as a standalone utility. " +
		"This plugin must be used in conjunction with gpbackup and gprestore."

	app.Commands = []cli.Command{
		{
			Name:   "setup_plugin_for_backup",
			Action: s3plugin.SetupPluginForBackup,
			Before: buildBeforeFunc(3, 4),
		},
		{
			Name:   "setup_plugin_for_restore",
			Action: s3plugin.SetupPluginForRestore,
			Before: buildBeforeFunc(3, 4),
		},
		{
			Name:   "cleanup_plugin_for_backup",
			Action: s3plugin.CleanupPlugin,
			Before: buildBeforeFunc(3, 4),
		},
		{
			Name:   "cleanup_plugin_for_restore",
			Action: s3plugin.CleanupPlugin,
			Before: buildBeforeFunc(3, 4),
		},
		{
			Name:   "backup_file",
			Action: s3plugin.BackupFile,
			Before: buildBeforeFunc(2),
		},
		{
			Name:   "backup_directory",
			Action: s3plugin.BackupDirectory,
			Before: buildBeforeFunc(2, 3),
			Hidden: true,
		},
		{
			Name:   "backup_directory_parallel",
			Action: s3plugin.BackupDirectoryParallel,
			Before: buildBeforeFunc(2, 3),
			Hidden: true,
		},
		{
			Name:   "restore_file",
			Action: s3plugin.RestoreFile,
			Before: buildBeforeFunc(2),
		},
		{
			Name:   "restore_directory",
			Action: s3plugin.RestoreDirectory,
			Before: buildBeforeFunc(2, 3),
			Hidden: true,
		},
		{
			Name:   "restore_directory_parallel",
			Action: s3plugin.RestoreDirectoryParallel,
			Before: buildBeforeFunc(2, 3),
			Hidden: true,
		},
		{
			Name:   "backup_data",
			Action: s3plugin.BackupData,
			Before: buildBeforeFunc(2),
		},
		{
			Name:   "restore_data",
			Action: s3plugin.RestoreData,
			Before: buildBeforeFunc(2),
		},
		{
			Name:   "plugin_api_version",
			Action: s3plugin.GetAPIVersion,
			Before: buildBeforeFunc(0),
		},
		{
			Name:   "delete_backup",
			Action: s3plugin.DeleteBackup,
			Before: buildBeforeFunc(2),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		gplog.Error(err.Error())
		os.Exit(1)
	}
}

func buildBeforeFunc(expectedNArgs ...int) (beforeFunc cli.BeforeFunc) {
	beforeFunc = func(context *cli.Context) error {
		actualNArg := context.NArg()
		argMatched := false
		for _, expectedNArg := range expectedNArgs {
			if actualNArg == expectedNArg {
				argMatched = true
				break
			}
		}
		if !argMatched {
			return fmt.Errorf("ERROR: Invalid number of arguments to plugin command. "+
				"Expected %v arguments. Got %d arguments", expectedNArgs, actualNArg)
		} else {
			return nil
		}

	}
	return beforeFunc
}
