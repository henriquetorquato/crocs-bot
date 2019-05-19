package types

type Platform interface {
	Name() string
	PostMessage(message string) bool
	GetLastPost() Post
}

type Post struct {
	Platform string
	Content  string
}
