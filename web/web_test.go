package web_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/airconduct/go-probot/web"
)

var _ = Describe("Web", func() {
	It("Index.html should be existing", func() {
		_, err := web.Static().Open("index.html")
		Expect(err).Should(BeNil())
	})
	It("Handler should be registered", func() {
		mux := http.NewServeMux()
		web.RegisterHandler(mux)
		server := httptest.NewServer(mux)

		resp, err := http.Get(server.URL)
		Expect(err).Should(BeNil())
		Expect(resp.StatusCode).Should(Equal(404))

		resp, err = http.Get(server.URL + "/dashboard")
		Expect(err).Should(BeNil())
		Expect(resp.StatusCode).Should(Equal(200))

		resp, err = http.Get(server.URL + "/dashboard/about")
		Expect(err).Should(BeNil())
		Expect(resp.StatusCode).Should(Equal(200))

		resp, err = http.Get(server.URL + "/assets/")
		Expect(err).Should(BeNil())
		Expect(resp.StatusCode).Should(Equal(200))
	})
})
