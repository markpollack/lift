package initializr_test

import (
	"path/filepath"
	"testing"

	. "github.com/cloudlift/lift/internal/testutils"
	"github.com/cloudlift/lift/pkg/initializr"
	. "github.com/onsi/gomega"
)

func TestInitializrNew(t *testing.T) {
	g := NewWithT(t)
	tempDir, tempDirRemove := TempDir(t, "initializr-new")
	request := initializr.InitializrRequest{
		Dependencies: "web",
		GroupID:      "com.foo.bar",
		ArtifactID:   "webtest",
	}
	err := initializr.CreateNewProject(request, tempDir)
	g.Expect(err).ToNot(HaveOccurred())

	pomFile := filepath.FromSlash(tempDir + "/pom.xml")
	g.Expect(pomFile).Should(BeARegularFile())
	contents := FileContents(t, pomFile)
	g.Expect(contents).Should(ContainSubstring("<artifactId>spring-boot-starter-web</artifactId>"))
	g.Expect(contents).Should(ContainSubstring("<groupId>com.foo.bar</groupId>"))
	g.Expect(contents).Should(ContainSubstring("<artifactId>webtest</artifactId>"))

	// Will remove the test directory only if all tests pass, otherwise leave it around for investigation.
	tempDirRemove()
}
