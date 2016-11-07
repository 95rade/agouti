package agouti_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sclevine/agouti"
)

var _ = Describe("Capabilities", func() {
	var capabilities Capabilities
	proxyType := "manual"
	proxyAddress := "localhost:8080"

	BeforeEach(func() {
		capabilities = NewCapabilities("firstEnabled", "secondEnabled")

	})

	It("should successfully encode all provided options into JSON", func() {
		capabilities.Browser("some-browser").Version("v100").Platform("some-os")
		capabilities.With("withEnabled").Without("withoutDisabled")
		capabilities.Proxy(ProxyConfig{
			ProxyType: &proxyType,
			HttpProxy: &proxyAddress,
			SslProxy:  &proxyAddress,
		})
		Expect(capabilities.JSON()).To(MatchJSON(`{
			"browserName": "some-browser",
			"version": "v100",
			"platform": "some-os",
			"withEnabled": true,
			"withoutDisabled": false,
			"firstEnabled": true,
			"secondEnabled": true,
			"proxy": {
				"proxyType": "manual",
				"proxyAuthoconfigUrl": null,
				"ftpProxy": null,
				"httpProxy": "localhost:8080",
				"sslProxy": "localhost:8080",
				"socksProxy": null,
				"socksUsername": null,
				"socksPassword": null,
				"noProxy": null
			}
		}`))
	})

	Context("when the provided options cannot be converted to JSON", func() {
		It("should return an error", func() {
			capabilities["some-feature"] = func() {}
			_, err := capabilities.JSON()
			Expect(err).To(MatchError("json: unsupported type: func()"))
		})
	})
})
