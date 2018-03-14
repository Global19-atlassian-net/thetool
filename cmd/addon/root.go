package addon

import "github.com/spf13/cobra"

func AddonCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "addon",
		Short: "manage addons for Gloo",
	}
	cmd.AddCommand(listCmd(), configureCmd())
	return cmd
}
