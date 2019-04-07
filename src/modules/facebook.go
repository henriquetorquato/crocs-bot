package modules

import (
	"fmt"
	"net/http"
	"net/url"

	utils "../utils"
	handler "../utils/error"
)

const facebookLocationBase = "https://graph.facebook.com"

type Facebook struct{}

func (f Facebook) PostMessage(message string) {

	auth := utils.Config.Facebook

	location := fmt.Sprintf("%s/%s/feed", facebookLocationBase, auth.Page.ID)

	content := url.Values{
		"message":      []string{message},
		"access_token": []string{auth.Page.Token}}

	response, err := http.PostForm(location, content)
	handler.HandleError(err)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		fmt.Println("Facebook")
	}

}
