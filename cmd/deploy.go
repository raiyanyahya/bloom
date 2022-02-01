package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your website",
	Long:  "Deploy your website on bloom.",
	Run:   deploy,
}

func deploy(cmd *cobra.Command, args []string) {
	fmt.Println("deploy", out)
	checkFolderExists(out)
}

func checkFolderExists(out string) {
	folderInfo, err := os.Stat(out)
	if os.IsNotExist(err) {
		log.Fatalf("Build folder '%s' does not exist.", out)
	}
	log.Println(folderInfo)
}
func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVarP(&out, "out", "O", "out", "The output directory with the generated files")
}
