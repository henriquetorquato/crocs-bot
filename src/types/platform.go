package types

type Platform interface {
	Name() string
	PostMessage(message string) bool
}
