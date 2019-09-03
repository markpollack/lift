package commands

import (
	"bytes"
	"github.com/onsi/gomega/format"
	"os/exec"
	"testing"

	. "github.com/cloudlift/lift/pkg/testutils"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestLiftCliCommands(t *testing.T) {

	//Setup: build binary to be tested
	g := NewGomegaWithT(t)
	binPath, err := gexec.Build("github.com/cloudlift/lift/cmd/lift")
	g.Expect(err).ToNot(HaveOccurred())

	t.Run("`lift` command with no subcommands or flags prints command usage", func(t *testing.T) {
		expectedOutput :=
			`lift is a tool for enriching your application so it can be deployed to multiple cloud platforms with minimal effort.

Usage:
  lift [command]

Available Commands:
  help        Help about any command
  initializr  Initializr commands
  platform    Platform commands

Flags:
  -h, --help   help for lift

Use "lift [command] --help" for more information about a command.
`
		g := NewWithT(t)
		cmd := exec.Command(binPath)
		buffer := bytes.NewBuffer(nil)
		sess, _ := gexec.Start(cmd, buffer, buffer)
		g.Eventually(sess).Should(gexec.Exit(0))
		g.Expect(buffer.String()).To(Equal(expectedOutput))
	})

	t.Run("`lift platform list` command with no flags prints a table of platforms", func(t *testing.T) {
		expectedOutput :=
			`+-----------------------+--------+--------------+---------+-------------+
| [1;42m        NAME         [0m | [1;42mALIAS [0m | [1;42m    TYPE    [0m | [1;42mPROFILE[0m | [1;42mDESCRIPTION[0m |
+-----------------------+--------+--------------+---------+-------------+`
		g := NewWithT(t)
		cmd := exec.Command(binPath, "platform", "list")
		buffer := bytes.NewBuffer(nil)
		sess, _ := gexec.Start(cmd, buffer, buffer)
		g.Eventually(sess).Should(gexec.Exit(0))
		g.Expect(buffer.String()).To(ContainSubstring(expectedOutput))
	})

	tempDir, tempDirRemove := TempDir(t, "initializr-new")
	t.Run( "`lift initializr new` invoke the Initializr service", func(t *testing.T) {
		expectedOutput :=
		`Invoking Initializr service at https://start.spring.io
Initializr zip file extracted to ` + tempDir + "\n"

		g := NewWithT(t)
		format.TruncatedDiff = false;
		cmd := exec.Command(binPath, "initializr", "new", "--artifactId", "com.foo.bar", "--groupId", "mygroup", "--path", tempDir)
		buffer := bytes.NewBuffer(nil)
		sess, _ := gexec.Start(cmd, buffer, buffer)
		g.Eventually(sess).Should(gexec.Exit(0))
		g.Expect(buffer.String()).To(Equal(expectedOutput))
		tempDirRemove()
	})

	//Cleanup for subtests
	gexec.CleanupBuildArtifacts()

}
