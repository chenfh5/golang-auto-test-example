package prd

import (
	"errors"
	"sync"
	"testing"

	"golang-auto-test-example/api"
	"golang-auto-test-example/api/mocks"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//go:generate go test ./... -v -race -count=1 -cover

func TestStorage_Bootstrap(t *testing.T) {
	type fields struct {
		s3Client api.ClientAPI
		mu       sync.RWMutex
		data     map[string]string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Success",
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Storage{
				s3Client: tt.fields.s3Client,
				mu:       tt.fields.mu,
				data:     tt.fields.data,
			}
			assert.Nil(t, p.s3Client)
			log.Info(p.s3Client)
			p.Bootstrap() // change the client after bootstrap
			assert.NotNil(t, p.s3Client)
			log.Info(p.s3Client)
		})
	}
}

func TestStorage_Get(t *testing.T) {
	type fields struct {
		s3Client api.ClientAPI
		mu       sync.RWMutex
		data     map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				data: map[string]string{"k1": "v1"},
			},

			args:    args{"k1"},
			want:    "v1",
			wantErr: false,
		},
		{
			name:    "Failure",
			fields:  fields{},
			args:    args{"your_key"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Storage{
				s3Client: tt.fields.s3Client,
				mu:       tt.fields.mu,
				data:     tt.fields.data,
			}
			got, err := p.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Set(t *testing.T) {
	mockValue := []byte{0, 128, 255}
	k, v := getKVfromByteArray(mockValue)

	mockS3 := &mocks.ClientAPI{}
	mockS3.On("GetObject", mock.Anything).Return(mockValue, nil).Once()
	mockS3.On("GetObject", mock.Anything).Return(nil, errors.New("error case")).Once()

	type fields struct {
		s3Client api.ClientAPI
		mu       sync.RWMutex
		data     map[string]string
	}
	type args struct {
		obj string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				s3Client: mockS3,
				data:     map[string]string{"k1": "v1"},
			},

			args:    args{"k1"},
			wantErr: false,
		},
		{
			name: "Failure",
			fields: fields{
				s3Client: mockS3,
			},
			args:    args{"your_key"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Storage{
				s3Client: tt.fields.s3Client,
				mu:       tt.fields.mu,
				data:     tt.fields.data,
			}
			if err := p.Set(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.True(t, len(p.data) == 2)
				assert.True(t, p.data["k1"] == "v1")
				assert.True(t, p.data[k] == v)
			}
		})
	}
	mockS3.AssertExpectations(t) // verify the mock execute times
}
