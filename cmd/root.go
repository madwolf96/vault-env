package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"vault-env/pkg/client"
	vaultenv "vault-env/pkg/vault-env"
)

var (
	mount   string
	rootCmd = &cobra.Command{
		Use:   "vault-env",
		Short: "vault-env is a tool to inject secrets from a vault into the environment",
	}

	ejectCmd = &cobra.Command{
		Use:   "eject",
		Short: "eject secrets from the environment",
		Run: func(cmd *cobra.Command, args []string) {
			client := client.ClientInit()
			ctx := context.Background()
			secret, err := client.KVv2(mount).Get(ctx, args[0])
			if err != nil {
				log.Fatalf("Unable to read the super secret password from the vault: %v", err)
			}
			vaultEnv := vaultenv.VaultEnv{}
			if err := vaultEnv.Eject(secret); err != nil {
				log.Fatalf("Unable to unset environment variables: %v", err)
			}
		},
	}

	injectCmd = &cobra.Command{
		Use:   "inject",
		Short: "inject secrets into the environment",
		Run: func(cmd *cobra.Command, args []string) {
			client := client.ClientInit()
			ctx := context.Background()
			secret, err := client.KVv2(mount).Get(ctx, args[0])
			if err != nil {
				log.Fatalf("Unable to read the super secret password from the vault: %v", err)
			}
			vaultEnv := vaultenv.VaultEnv{}
			if err := vaultEnv.Inject(secret); err != nil {
				log.Fatalf("Unable to export environment variables: %v", err)
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(ejectCmd)
	rootCmd.AddCommand(injectCmd)

	injectCmd.Flags().StringVarP(&mount, "mount", "m", "secret",
		"the mount point of the secret engine")
	ejectCmd.Flags().StringVarP(&mount, "mount", "m", "secret",
		"the mount point of the secret engine")
	err := ejectCmd.MarkFlagRequired("mount")
	if err != nil {
		log.Fatalf("Unable to mark mount flag as required: %v", err)
	}
	err = injectCmd.MarkFlagRequired("mount")
	if err != nil {
		log.Fatalf("Unable to mark mount flag as required: %v", err)
	}
}
