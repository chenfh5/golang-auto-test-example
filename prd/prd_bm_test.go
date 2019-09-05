package prd

import (
	"math/rand"
	"testing"

	"golang-auto-test-example/api/mocks"

	"github.com/stretchr/testify/mock"
)

//go:generate go test -v -bench=. -run=none -cpuprofile cpu.prof -memprofile mem.prof -count=1 -cover
//go:generate go tool pprof -http=:8080 cpu.prof
//go:generate go tool pprof -http=:8081 mem.prof

func BenchmarkGet(b *testing.B) {
	b.StopTimer()

	p := &Storage{
		data: map[string]string{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
			"k4": "v4",
		},
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := p.Get("your_key")
		if err != nil {
			//
		}
	}
} // BenchmarkGet-12       	 5000000	       263 ns/op

func BenchmarkSet(b *testing.B) {
	b.StopTimer()
	mockS3 := &mocks.ClientAPI{}
	mockS3.On("GetObject", mock.Anything).Return([]byte{
		byte(rand.Intn(256)), byte(rand.Intn(256)),
		byte(rand.Intn(256)), byte(rand.Intn(256)),
		byte(rand.Intn(256)), byte(rand.Intn(256)),
		byte(rand.Intn(256)), byte(rand.Intn(256)),
		byte(rand.Intn(256)), byte(rand.Intn(256)),
	}, nil).Maybe()

	p := &Storage{
		s3Client: mockS3,
		data: map[string]string{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
			"k4": "v4",
		},
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := p.Set("your_s3_bucket_file")
		if err != nil {
			panic("error set")
		}
	}
} // BenchmarkSet-12    	  100000	     15614 ns/op (BenchmarkSet在12个proc的情况下1秒内被运行了100000次，平均每次是13331ns)
