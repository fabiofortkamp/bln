/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// define possible error conditions
const (
	ErrLinkNameExists     = 1
	ErrLinkToNotExist     = 2
	ErrLinkCreationFailed = 3
	ErrParsingFlags       = 4
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bln",
	Short: "A better interface for ln",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		linkName, err := cmd.Flags().GetString("link-name")
		if err != nil {
			fmt.Println(err)
			os.Exit(ErrParsingFlags)
		}

		// if linkName already exists, abort and print an error message
		if _, err = os.Stat(linkName); err == nil {
			fmt.Printf("ERROR: File %s already exists\n", linkName)
			// exit with status 1
			os.Exit(ErrLinkNameExists)
		}
		linkTo, err := cmd.Flags().GetString("link-to")
		if err != nil {
			fmt.Println(err)
		}

		// Check if linkTo exists
		if _, err = os.Stat(linkTo); os.IsNotExist(err) {
			fmt.Printf("ERROR: File %s does not exist\n", linkTo)
			os.Exit(ErrLinkToNotExist)
		}

		// Get the symolic flag value
		symbolic, err := cmd.Flags().GetBool("symbolic")
		if err != nil {
			fmt.Println(err)
			os.Exit(ErrParsingFlags)
		}

		// Create the link
		if symbolic {
			err = os.Symlink(linkTo, linkName)
			if err != nil {
				fmt.Println(err)
				os.Exit(ErrLinkCreationFailed)
			}
		} else {
			// in this case (hard link), check if linlTo is a directory
			// and if so, print an error message and exit
			fileInfo, err := os.Stat(linkTo)
			if err != nil {
				fmt.Println(err)
				os.Exit(ErrLinkCreationFailed)
			}
			if fileInfo.IsDir() {
				fmt.Printf("ERROR: %s is a directory (not permitted for hard links)\n", linkTo)
				os.Exit(ErrLinkCreationFailed)
			}
			err = os.Link(linkTo, linkName)
			if err != nil {
				fmt.Println(err)
				os.Exit(ErrLinkCreationFailed)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// create required flags for --link-name and --link-to
	var linkName string
	var linkTo string
	rootCmd.Flags().StringVarP(&linkName, "link-name", "l", "", "Link name (will be created)")
	rootCmd.Flags().StringVarP(&linkTo, "link-to", "t", "", "File to point the link to (must exist)")
	rootCmd.MarkFlagRequired("link-name")
	rootCmd.MarkFlagRequired("link-to")

	var symbolic bool
	rootCmd.Flags().BoolVarP(&symbolic, "symbolic", "s", false, "Symbolic link")
}
