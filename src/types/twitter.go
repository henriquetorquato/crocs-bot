package types

import "encoding/json"

type TwitterTimeline []Tweet

func UnmarshalTwitterTimeline(data []byte) (TwitterTimeline, error) {
	var r TwitterTimeline
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwitterTimeline) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Tweet struct {
	Coordinates          *Coordinates     `json:"coordinates"`
	Truncated            bool             `json:"truncated"`
	CreatedAt            string           `json:"created_at"`
	Favorited            bool             `json:"favorited"`
	IDStr                string           `json:"id_str"`
	InReplyToUserIDStr   interface{}      `json:"in_reply_to_user_id_str"`
	Entities             TimelineEntities `json:"entities"`
	Text                 string           `json:"text"`
	Contributors         interface{}      `json:"contributors"`
	ID                   float64          `json:"id"`
	RetweetCount         int64            `json:"retweet_count"`
	InReplyToStatusIDStr interface{}      `json:"in_reply_to_status_id_str"`
	Geo                  *Coordinates     `json:"geo"`
	Retweeted            bool             `json:"retweeted"`
	InReplyToUserID      interface{}      `json:"in_reply_to_user_id"`
	Place                *Place           `json:"place"`
	Source               string           `json:"source"`
	User                 User             `json:"user"`
	InReplyToScreenName  interface{}      `json:"in_reply_to_screen_name"`
	InReplyToStatusID    interface{}      `json:"in_reply_to_status_id"`
	PossiblySensitive    *bool            `json:"possibly_sensitive,omitempty"`
}

type Coordinates struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

type TimelineEntities struct {
	Urls         []URLElement  `json:"urls"`
	Hashtags     []interface{} `json:"hashtags"`
	UserMentions []UserMention `json:"user_mentions"`
}

type URLElement struct {
	ExpandedURL *string `json:"expanded_url"`
	URL         string  `json:"url"`
	Indices     []int64 `json:"indices"`
	DisplayURL  *string `json:"display_url"`
}

type UserMention struct {
	Name       string  `json:"name"`
	IDStr      string  `json:"id_str"`
	ID         int64   `json:"id"`
	Indices    []int64 `json:"indices"`
	ScreenName string  `json:"screen_name"`
}

type Place struct {
	Name        string      `json:"name"`
	CountryCode string      `json:"country_code"`
	Country     string      `json:"country"`
	Attributes  Attributes  `json:"attributes"`
	URL         string      `json:"url"`
	ID          string      `json:"id"`
	BoundingBox BoundingBox `json:"bounding_box"`
	FullName    string      `json:"full_name"`
	PlaceType   string      `json:"place_type"`
}

type Attributes struct {
}

type BoundingBox struct {
	Coordinates [][][]float64 `json:"coordinates"`
	Type        string        `json:"type"`
}

type User struct {
	Name                           string       `json:"name"`
	ProfileSidebarFillColor        string       `json:"profile_sidebar_fill_color"`
	ProfileBackgroundTile          bool         `json:"profile_background_tile"`
	ProfileSidebarBorderColor      string       `json:"profile_sidebar_border_color"`
	ProfileImageURL                string       `json:"profile_image_url"`
	CreatedAt                      string       `json:"created_at"`
	Location                       string       `json:"location"`
	FollowRequestSent              bool         `json:"follow_request_sent"`
	IDStr                          string       `json:"id_str"`
	IsTranslator                   bool         `json:"is_translator"`
	ProfileLinkColor               string       `json:"profile_link_color"`
	Entities                       UserEntities `json:"entities"`
	DefaultProfile                 bool         `json:"default_profile"`
	URL                            string       `json:"url"`
	ContributorsEnabled            bool         `json:"contributors_enabled"`
	FavouritesCount                int64        `json:"favourites_count"`
	UTCOffset                      *int64       `json:"utc_offset"`
	ProfileImageURLHTTPS           string       `json:"profile_image_url_https"`
	ID                             int64        `json:"id"`
	ListedCount                    int64        `json:"listed_count"`
	ProfileUseBackgroundImage      bool         `json:"profile_use_background_image"`
	ProfileTextColor               string       `json:"profile_text_color"`
	FollowersCount                 int64        `json:"followers_count"`
	Lang                           string       `json:"lang"`
	Protected                      bool         `json:"protected"`
	GeoEnabled                     bool         `json:"geo_enabled"`
	Notifications                  bool         `json:"notifications"`
	Description                    string       `json:"description"`
	ProfileBackgroundColor         string       `json:"profile_background_color"`
	Verified                       bool         `json:"verified"`
	TimeZone                       *string      `json:"time_zone"`
	ProfileBackgroundImageURLHTTPS string       `json:"profile_background_image_url_https"`
	StatusesCount                  int64        `json:"statuses_count"`
	ProfileBackgroundImageURL      string       `json:"profile_background_image_url"`
	DefaultProfileImage            bool         `json:"default_profile_image"`
	FriendsCount                   int64        `json:"friends_count"`
	Following                      bool         `json:"following"`
	ShowAllInlineMedia             bool         `json:"show_all_inline_media"`
	ScreenName                     string       `json:"screen_name"`
}

type UserEntities struct {
	URL         DescriptionClass  `json:"url"`
	Description *DescriptionClass `json:"description"`
}

type DescriptionClass struct {
	Urls []URLElement `json:"urls"`
}
