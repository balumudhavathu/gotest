package adder_test

import(
	ad "adder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Sutite", func() {
	 Context("Test API", func() {
	    It("Should pass", func() {
	       sum := ad.Add(2, 3)
	       Expect(sum).To(Equal(5))
	    })
	 })
})

