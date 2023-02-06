package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Example Suite")
}

const installationID = 2

var _ = BeforeSuite(func() {
	// Mock get token
	gock.New("https://api.github.com").
		Post(fmt.Sprintf("/app/installations/%d/access_tokens", installationID)).
		Reply(200).JSON(map[string]interface{}{
		"token": "test",
		"permissions": map[string]interface{}{
			"issues": "write",
		}})
	//
	gock.New("https://api.github.com").
		Post("/repos/foo/bar/issues/1/comments").
		AddMatcher(func(r1 *http.Request, r2 *gock.Request) (bool, error) {
			data := make(map[string]string)
			if err := json.NewDecoder(r1.Body).Decode(&data); err != nil {
				return false, err
			}
			return data["body"] == "Thanks!", nil
		}).
		Reply(200)
})

var _ = AfterSuite(func() {
	gock.Off()
})
