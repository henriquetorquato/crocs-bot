package modules

import (
	"fmt"
	"net/url"

	utils "../utils"
	"github.com/mrjones/oauth"
)

const twitterLocationBase = "https://api.twitter.com/1.1"

type Twitter struct{}

func (f Twitter) Name() string {
	return "Twitter"
}

func (t Twitter) PostMessage(message string) bool {

	auth := utils.Config.Twitter

	location := fmt.Sprintf("%s/statuses/update.json", twitterLocationBase)
	content := url.Values{"status": []string{message}}

	consumer := oauth.NewConsumer(
		auth.Consumer.Key,
		auth.Consumer.Secret,
		oauth.ServiceProvider{})

	token := oauth.AccessToken{
		Token:  auth.Access.Token,
		Secret: auth.Access.Secret}

	client, err := consumer.MakeHttpClient(&token)
	utils.HandleError(err)

	response, err := client.PostForm(location, content)
	utils.HandleError(err)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		return true
	}

	utils.HandleResponse(response)
	return false

}
