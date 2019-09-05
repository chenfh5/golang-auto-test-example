package api

type ESClient struct {
	id      int64
	conf    string
	session string
}

// NewClient ...
func NewEsClient(conf string) (ClientAPI, error) {
	return &ESClient{conf: conf}, nil
}

// GetObject ...
func (p *ESClient) GetObject(key string) ([]byte, error) {
	return nil, nil // TODO
}

// PutObject ...
func (p *ESClient) PutObject(key string, payload []byte) error {
	return nil // TODO

}

// DeleteObject ...
func (p *ESClient) DeleteObject(key string) error {

	return nil // TODO
}

// ListObjects ...
func (p *ESClient) ListObjects(prefix string) (interface{}, error) {
	return nil, nil // TODO
}

// ExistObject ...
func (p *ESClient) ExistObject(key string) bool {
	return false // TODO
}
