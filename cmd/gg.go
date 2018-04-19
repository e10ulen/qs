// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"
	"os/exec"
	"time"
	"github.com/spiegel-im-spiegel/logf"
	"github.com/spf13/cobra"
	"github.com/e10ulen/sandbox/lib"
)

const DateFormat = "2006/01/02 15:04"
// ggCmd represents the gg command
var ggCmd = &cobra.Command{
	Use:   "gg",
	Short: "Commit and Push",
	Long: `Add & Commit message & push Automatic execution`,
	Run: func(cmd *cobra.Command, args []string) {
		logf.SetOutput(os.Stdout)
		logf.SetMinLevel(logf.INFO)
		tm := time.Now()
		dir, err := os.Getwd()
		if err == nil {
			logf.Debugf("ディレクトリ取得の確認:" + dir + "\n")
		}
		logf.Debug("時間の確認:" + tm.Format(DateFormat) + "\n")
		add, err := exec.Command("git", "add", "--all").CombinedOutput()
		if err != nil {
			logf.Errorf("%v\n", err)
		}
		cmt, err := exec.Command("git", "commit", "-m", "[Commit]" + tm.Format(DateFormat) + " " + lib.ScanLine()).CombinedOutput()
		if err != nil {
			logf.Errorf("%v\n", err)
		}
		push, err := exec.Command("git", "push", "-u").CombinedOutput()
		if err != nil {
			logf.Errorf("%v\n", err)
		}
		if add != nil {
			logf.Printf(string(add))
		}
		if cmt != nil {
			logf.Printf(string(cmt))
		}
		if push != nil {
			logf.Printf(string(push))
		}
	},
}

func init() {
	rootCmd.AddCommand(ggCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ggCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ggCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
