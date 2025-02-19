package piscator

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "piscator",
	Short: "piscator is a CLT for cloning GitHub repos",
	Long: `Embark on a grand voyage across the GitHub seas! Set sail to create
a magnificent directory, inspired by the name of a fearless sailor or a
mighty pirate. With each collection you reel in, a new Git repository shall
be forged, like a sturdy ship ready to conquer the code oceans. Prepare
yourself to navigate through the user's or organization's treasures,
uncovering hidden gems and secret code islands. Will you include private
repositories, like mysterious hidden coves? Or perhaps venture into the
realm of forked repositories, tracing the footsteps of fellow sailors? As
you reel in the collections, a legendary repos.json file shall be forged,
capturing the essence of your brave expedition. Choose the winds of
verbosity to whisper tales of each step or keep silent like a true sailor.
Raise the anchor, set your course, and let the adventure begin!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To see available commands run help")
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your piscator command '%s'", err)
		os.Exit(1)
	}

	return nil
}
