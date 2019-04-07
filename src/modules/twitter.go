package modules

import (
	"fmt"
	"net/url"

	utils "../utils"
	handler "../utils/error"
	"github.com/mrjones/oauth"
)

const twitterLocationBase = "https://api.twitter.com/1.1"

type Twitter struct{}

func (t Twitter) PostMessage(message string) {

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
	handler.HandleError(err)

	response, err := client.PostForm(location, content)
	handler.HandleError(err)

	defer response.Body.Close()

	if response.StatusCode == 200 {
		fmt.Println("Twitter")
	}

}
