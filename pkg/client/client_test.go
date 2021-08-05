package client

import (
	"github.com/andrewrobinson/accountapi/pkg/client/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
)

//https://ieftimov.com/post/testing-in-go-go-test/
// https://www.thegreatcodeadventure.com/mocking-http-requests-in-golang/
// https://stackoverflow.com/questions/57872522/how-to-mock-http-newrequest-to-return-an-error

var _ = Describe("Client Unit tests", func() {

	id := uuid.FromStringOrNil("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	BeforeEach(func() {

	})

	//this is just to get coverage on doGet, doPost and doDelete where httpClient.Do returns an error
	Context("MockClientThatErrorsOnDo", func() {

		accountClient := NewTestingAccountRestClient("http://localhost", &mocks.MockClientThatErrorsOnDo{})

		It("on doGet via a Fetch call", func() {
			_, err := accountClient.Fetch(id)
			Expect(err.Error()).To(Equal("MOCK ERROR ON DO"))
		})
		It("on doPost via a Create call", func() {
			_, err := accountClient.Create(buildAccountDataForCreate(id))
			Expect(err.Error()).To(Equal("MOCK ERROR ON DO"))
		})
		It("on doDelete via a Delete call", func() {
			err := accountClient.Delete(id, int64(0))
			Expect(err.Error()).To(Equal("MOCK ERROR ON DO"))
		})

	})

})
