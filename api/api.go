package api

//go:generate mockery -name=ClientAPI

type ClientAPI interface {
	GetObject(string) ([]byte, error)
	PutObject(string, []byte) error
	DeleteObject(string) error
	ListObjects(string) (interface{}, error)
	ExistObject(string) bool
}
