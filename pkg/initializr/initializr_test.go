package initializr_test

import (
	"os"
	"testing"

	"github.com/cloudlift/lift/pkg/initializr"
	. "github.com/onsi/gomega"
)

func TestGenerate(t *testing.T) {
	g := NewGomegaWithT(t)
	request := initializr.InitializrRequest{}
	request.ArtifactId = "webdemo"
	resp, err := initializr.Generate(request)
	g.Expect(err).ToNot(HaveOccurred())

	path := "/tmp/liftziptests"
	os.RemoveAll(path)
	initializr.Unpack(resp.Contents, path)
	g.Expect("/tmp/liftziptests/pom.xml").Should(BeARegularFile())

}
