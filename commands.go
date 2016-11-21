package main

import (
	"github.com/urfave/cli"
)

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:      "deploy",
			Usage:     "Inicia o deploy da aplicação",
			ArgsUsage: "[aplicações...]",
			Before:    defineDockerHostCommand,
			Action:    deployAction,
		},
		{
			Name:   "update-images",
			Usage:  "Faz update das imagens a partir do mirror da Softplan",
			Before: defineDockerHostCommand,
			Action: forceUpdateImages,
		},
		{
			Name:            "docker",
			Usage:           "Executa comandos docker na maquina definida como Host",
			SkipFlagParsing: true,
			Before:          defineDockerHostCommand,
			Action:          executeDockerCmd,
		},
		{
			Name:            "compose",
			Usage:           "Executa comandos compose na maquina definida como Host",
			SkipFlagParsing: true,
			Before:          defineDockerHostCommand,
			Action:          executeDockerComposeCmd,
		},
		{
			Name:   "versions",
			Usage:  "Exibe as versões das aplicações",
			Action: showVersionsAction,
		},
		{
			Name:      "help",
			Aliases:   []string{"h"},
			Usage:     "Exibe a lista de commandos ou a ajuda de um comando",
			ArgsUsage: "[commando]",
			Action:    showHelpAction,
		},
	}
}
