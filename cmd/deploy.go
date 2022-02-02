package cmd

import (
	"bloom/cloud"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
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
	s3Obj := cloud.S3Objectt{}
	s3client := cloud.GetClient(s3Obj).(*s3.S3)
	fmt.Print(s3client)
}

func checkFolderExists(out string) {
	_, err := os.Stat(out)
	if os.IsNotExist(err) {
		log.Fatalf("Build folder '%s' does not exist.", out)
	}
	log.Printf("Build folder '%s' . Moving on ...", out)
}
func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVarP(&out, "out", "O", "out", "The output directory with the generated files")
}
