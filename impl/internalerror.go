package impl

type InternalError interface {
	error
	Message(string) string
}
