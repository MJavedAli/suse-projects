package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rancher API Test", func() {
	config := TestConfig{
		RancherURL: "https://localhost:8443",
		Username:   "username",
		Password:   "password",
	}


	It("should log in to Rancher and get the token", func() {
		token, err := loginAndGetToken(config)
		Expect(err).ToNot(HaveOccurred())
		Expect(token).ToNot(BeEmpty())

		By("Logging the token")
		println("Login successful. Token:", token)
	})
})
