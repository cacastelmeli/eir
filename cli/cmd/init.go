package cmd

import (
	"github.com/urfave/cli/v2"
)

func InitCommand() *cli.Command {
	return &cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initialize gexarch (this will override previous configuration)",
		Action:  initCommandAction,
	}
}

func initCommandAction(ctx *cli.Context) error {
	/*typesPath := ctx.Args().Get(0)

	if typesPath == "" {
		return errors.New("missing types-path argument")
	}

	conf := &config.ProcessorConfig{
		CliConfig: &config.CliConfig{
			TypesPath: typesPath,
		},
		ModulePath: util.ParseModfile().Module.Mod.Path,
	}

	workingDirectory, err := os.Getwd()
	util.PanicIfError(err)

	templateProcessor := processor.NewTemplateProcessor(conf)
	templateProcessor.ProcessTemplate("init", workingDirectory)*/

	return nil
}
