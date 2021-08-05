package client

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//https://ieftimov.com/post/testing-in-go-go-test/

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Suite")
}

var _ = Describe("Client Unit tests", func() {

	BeforeEach(func() {

	})

	Context("aaa", func() {
		It("bbb", func() {
			Expect(nil).To(BeNil())
		})
	})

	Context("ccc", func() {
		It("bbb", func() {
			Expect(nil).To(BeNil())
		})
	})

})
