package api

type MysqlClient struct {
	id      int64
	conf    string
	session string
}

// NewClient ...
func NewMysqlClient(conf string) (ClientAPI, error) {
	return &MysqlClient{conf: conf}, nil
}

// GetObject ...
func (sc *MysqlClient) GetObject(key string) ([]byte, error) {
	return nil, nil // TODO
}

// PutObject ...
func (sc *MysqlClient) PutObject(key string, payload []byte) error {
	return nil // TODO

}

// DeleteObject ...
func (sc *MysqlClient) DeleteObject(key string) error {

	return nil // TODO
}

// ListObjects ...
func (sc *MysqlClient) ListObjects(prefix string) (interface{}, error) {
	return nil, nil // TODO
}

// ExistObject ...
func (sc *MysqlClient) ExistObject(key string) bool {
	return false // TODO
}
