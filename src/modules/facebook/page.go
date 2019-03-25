package facebook

import (
	"fmt"
	"net/http"
)

const locationBase = "https://graph.facebook.com"

func PostMessage(message string, page string, token string) {
	location := fmt.Sprintf("%s/%s/feed?message=%s&access_token=%s", locationBase, page, message, token)
	response, err := http.Post(location, "", nil)

	if err != nil && response.StatusCode != 200 {
		fmt.Println(err)
	}
}
