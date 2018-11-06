package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("Hello", func() {
    Describe("Something", func() {
        Context("Something", func() {
            It("True should be true!", func() {
                Expect(true).To(Equal(true))
            })
        })
    })
})
