package modules

import (
	"fmt"
	"net/http"
	"net/url"

	utils "../utils"
)

const facebookLocationBase = "https://graph.facebook.com"

type Facebook struct{}

func (f Facebook) Name() string {
	return "Facebook"
}

func (f Facebook) PostMessage(message string) bool {

	auth := utils.Config.Facebook

	location := fmt.Sprintf("%s/%s/feed", facebookLocationBase, auth.Page.ID)

	content := url.Values{
		"message":      []string{message},
		"access_token": []string{auth.Page.Token}}

	response, err := http.PostForm(location, content)
	utils.HandleError(err)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		return true
	}

	return false

}
