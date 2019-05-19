package twitter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	types "../../types"
	utils "../../utils"
	"github.com/mrjones/oauth"
)

const name = "Twitter"
const locationBase = "https://api.twitter.com/1.1"

// Twitter exported interface
type Twitter struct{}

// Name property
func (f Twitter) Name() string {
	return name
}

var client = getHttpClient()

func (t Twitter) PostMessage(message string) bool {

	location := fmt.Sprintf("%s/statuses/update.json", locationBase)
	content := url.Values{"status": []string{message}}

	response, err := client.PostForm(location, content)
	utils.HandleError(err)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		return true
	}

	utils.HandleResponse(response)
	return false

}

func (t Twitter) GetLastPost() types.Post {
	timeline := getUserTimeline(1)
	return types.Post{
		Platform: name,
		Content:  timeline[0].Text,
	}
}

func getUserTimeline(size int) types.TwitterTimeline {

	location := fmt.Sprintf("%s/statuses/home_timeline.json", locationBase)
	request, err := http.NewRequest(http.MethodGet, location, nil)

	query := url.Values{"count": []string{string(size)}}
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	utils.HandleError(err)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		timeline, err := types.UnmarshalTwitterTimeline(bodyBytes)
		utils.HandleError(err)
		return timeline
	}

	utils.HandleResponse(response)
	return types.TwitterTimeline{}

}

func getHttpClient() http.Client {

	auth := utils.Config.Twitter

	consumer := oauth.NewConsumer(
		auth.Consumer.Key,
		auth.Consumer.Secret,
		oauth.ServiceProvider{})

	token := oauth.AccessToken{
		Token:  auth.Access.Token,
		Secret: auth.Access.Secret}

	client, err := consumer.MakeHttpClient(&token)
	utils.HandleError(err)

	return *client

}
