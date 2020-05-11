package cmd

import (
	"fmt"
	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"os"
)

var templateRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "delete", "destroy"},
	Short:   "Remove a template",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			fmt.Printf("Unable to create a Civo API Client: %s\n", aurora.Red(err))
			os.Exit(1)
		}

		if utility.AskForConfirmDelete("template") == nil {
			template, err := client.GetTemplateByCode(args[0])
			if err != nil {
				fmt.Printf("Unable to find the template for your search: %s\n", aurora.Red(err))
				os.Exit(1)
			}

			_, err = client.DeleteTemplate(template.ID)

			ow := utility.NewOutputWriterWithMap(map[string]string{"ID": template.ID, "Name": template.Name})

			switch outputFormat {
			case "json":
				ow.WriteSingleObjectJSON()
			case "custom":
				ow.WriteCustomOutput(outputFields)
			default:
				fmt.Printf("The template called %s with ID %s was delete\n", aurora.Green(template.Name), aurora.Green(template.ID))
			}
		} else {
			fmt.Println("Operation aborted.")
		}
	},
}