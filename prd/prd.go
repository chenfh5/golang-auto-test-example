package prd

import (
	"fmt"
	"sync"

	"golang-auto-test-example/api"
)

//go:generate gotests -all -w prd.go

type Storage struct {
	s3Client api.ClientAPI // 这里是interface类型
	mu       sync.RWMutex
	data     map[string]string
}

// init
func (p *Storage) Bootstrap() {
	client, err := api.NewS3Client("your_s3_conf") // here is s3, can be mysql or es as well
	if err != nil {
		panic(fmt.Sprintf("cannot init s3 client error: %+v", err))
	}
	p.s3Client = client
}

func (p *Storage) Get(key string) (string, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if v, ok := p.data[key]; ok {
		return v, nil
	}
	return "", fmt.Errorf("key: %+v is not exist in data map", key)
}

func (p *Storage) Set(obj string) error {
	b, err := p.s3Client.GetObject(obj)
	if err != nil {
		return fmt.Errorf("cannot get s3 object: %s because: %s", obj, err.Error())
	}

	k, v := getKVfromByteArray(b)
	p.mu.Lock()
	defer p.mu.Unlock()
	p.data[k] = v
	return nil
}

func getKVfromByteArray(b []byte) (string, string) {
	n := len(b)
	key := string(b[0 : n/2])
	value := string(b[n/2:])
	return key, value
}
