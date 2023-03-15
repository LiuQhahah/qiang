package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

type FileMode uint32
type fileItem struct {
	name     string    // base name of the file
	size     int64     // length in bytes for regular files; system-dependent for others
	mode     FileMode  // file mode bits
	modeTime time.Time // modification time
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolP("all", "a", false, "show all file info")

}

var lsCmd = &cobra.Command{
	Use: "ls",
	Run: func(cmd *cobra.Command, args []string) {

		files, err := os.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}

		allFlags, _ := cmd.Flags().GetBool("all")
		if allFlags {
			for _, file := range files {
				fileInfo, _ := file.Info()
				Item := fileItem{
					name:     fileInfo.Name(),
					size:     fileInfo.Size(),
					mode:     FileMode(fileInfo.Mode()),
					modeTime: fileInfo.ModTime(),
				}

				fmt.Println(Item)
			}
		} else {
			for _, file := range files {
				fmt.Println(file.Name())
			}
		}

	},
}
