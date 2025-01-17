package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Used to modify configuration items in the [base] group of the configuration file",
	Long:  `Used to modify configuration items in the [base] group of the configuration file

Format:
  ./coscli config set [flags]

Example:
  ./coscli config set -t example-token`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		setConfigItem(cmd)
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)

	configSetCmd.Flags().StringP("secret_id", "i", "", "Set secret id")
	configSetCmd.Flags().StringP("secret_key", "k", "", "Set secret key")
	configSetCmd.Flags().StringP("session_token", "t", "", "Set session token")
}

func setConfigItem(cmd *cobra.Command) {
	flag := false
	secretID, _ := cmd.Flags().GetString("secret_id")
	secretKey, _ := cmd.Flags().GetString("secret_key")
	sessionToken, _ := cmd.Flags().GetString("session_token")

	if secretID != "" {
		flag = true
		if secretID == "@" {
			config.Base.SecretID = ""
		} else {
			config.Base.SecretID = secretID
		}
	}
	if secretKey != "" {
		flag = true
		if secretKey == "@" {
			config.Base.SecretKey = ""
		} else {
			config.Base.SecretKey = secretKey
		}
	}
	if sessionToken != "" {
		flag = true
		if sessionToken == "@" {
			config.Base.SessionToken = ""
		} else {
			config.Base.SessionToken = sessionToken
		}
	}
	if !flag {
		_, _ = fmt.Fprintln(os.Stderr, "Enter at least one configuration item to be modified!")
		fmt.Println(cmd.UsageString())
		os.Exit(1)
	}

	viper.Set("cos.base", config.Base)
	if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Modify successfully!")
}
