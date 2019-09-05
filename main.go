package main

import (
	"fmt"

	"golang-auto-test-example/prd"

	log "github.com/sirupsen/logrus"
)

func main() {
	s := prd.Storage{}
	s.Bootstrap()

	err := s.Set("your_s3_bucket_file")
	if err != nil {
		log.Info(fmt.Sprintf("set error:= %+v", err))
	}

	key := "your_key"
	val, err := s.Get(key)
	if err != nil {
		log.Info(fmt.Sprintf("get key: %+v error: %+v", key, err))
	} else {
		fmt.Println(fmt.Sprintf("key: %+v -> value: %+v", key, val))
	}
}
