package facebook

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	types "../../types"
	utils "../../utils"
)

const name = "Facebook"
const locationBase = "https://graph.facebook.com"

var auth = utils.Config.Facebook
var client = getHttpClient()

// Facebook exported interface
type Facebook struct{}

// Name property
func (f Facebook) Name() string {
	return name
}

func (f Facebook) PostMessage(message string) bool {

	location := fmt.Sprintf("%s/%s/feed", locationBase, auth.Page.ID)

	content := url.Values{
		"message":      []string{message},
		"access_token": []string{auth.Page.Token}}

	response, err := http.PostForm(location, content)
	utils.HandleError(err)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		return true
	}

	utils.HandleResponse(response)
	return false

}

func (f Facebook) GetLastPost() types.Post {

	var content string
	timeline := getUserTimeline(1)

	if len(timeline.Data) > 0 {
		content = timeline.Data[0].Message
	} else {
		content = ""
	}

	return types.Post{
		Platform: name,
		Content:  content,
	}
}

func getUserTimeline(size int) types.FacebookTimeline {

	location := fmt.Sprintf("%s/%s/posts", locationBase, auth.Page.ID)
	request, err := http.NewRequest(http.MethodGet, location, nil)
	utils.HandleError(err)

	query := url.Values{
		"access_token": []string{auth.Page.Token}}

	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	utils.HandleError(err)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		timeline, err := types.UnmarshalFacebookTimeline(bodyBytes)
		utils.HandleError(err)

		timeline.Data = timeline.Data[:size]
		return timeline
	}

	return types.FacebookTimeline{}

}

func getHttpClient() http.Client {
	return http.Client{}
}
