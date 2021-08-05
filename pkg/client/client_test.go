package client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//https://ieftimov.com/post/testing-in-go-go-test/
// https://www.thegreatcodeadventure.com/mocking-http-requests-in-golang/
// https://stackoverflow.com/questions/57872522/how-to-mock-http-newrequest-to-return-an-error

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

	Context("ccc", func() {
		It("bbb", func() {
			Expect(nil).To(BeNil())
		})
	})

})
