package initializr_test

import (
	"testing"

	"github.com/cloudlift/lift/pkg/initializr"
	. "github.com/cloudlift/lift/pkg/testutils"
	. "github.com/onsi/gomega"
)

func TestInitializrNew(t *testing.T) {
	g := NewGomegaWithT(t)
	request := initializr.InitializrRequest{
		Dependencies: "web",
		GroupId:      "com.foo.bar",
		ArtifactId:   "webtest",
	}
	resp, err := initializr.Generate(request)
	g.Expect(err).ToNot(HaveOccurred())

	tempDir, tempDirRemove := TempDir(t, "initializr-new")

	initializr.Unpack(resp.Contents, tempDir)
	pomFile := tempDir + "/pom.xml"
	g.Expect(pomFile).Should(BeARegularFile())
	contents := FileContents(t, pomFile)
	g.Expect(contents).Should(ContainSubstring("<artifactId>spring-boot-starter-web</artifactId>"))
	g.Expect(contents).Should(ContainSubstring("<groupId>com.foo.bar</groupId>"))
	g.Expect(contents).Should(ContainSubstring("<artifactId>webtest</artifactId>"))

	// Will remove the test directory only if all tests pass, otherwise leave it around for investigation.
	tempDirRemove()
}
