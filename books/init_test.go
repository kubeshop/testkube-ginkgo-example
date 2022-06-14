package books

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestkubeGinkgoExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestkubeGinkgoExample Suite")
}
