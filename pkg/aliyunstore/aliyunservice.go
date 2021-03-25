package aliyunstore

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliYunObjectParams struct {
	// Bucket specifies the GCS bucket that the object resides in.
	Bucket string

	// ID specifies the ID of the GCS object.
	ID string
}

// AliYunService is an interface composed of all the necessary AliYun
// operations that are required to enable the tus protocol
// to work with AliYun's Object Storage Services(OSS).
type AliYunService interface{}

// AliYunClientManager holds the https://www.alibabacloud.com/product/oss client
// as well as its associated context.

type AliYunClientManager struct {
	Client *oss.Client
}

func NewAliYunClientManager() (*AliYunClientManager, error) {
	client, err := oss.New("", "", "")
	if err != nil {
		return nil, err
	}
	clientManager := &AliYunClientManager{
		Client: client,
	}
	return clientManager, nil
}

func (clientManager AliYunClientManager) ReadObject() (AliYunReader, error) {
	return nil, nil
}

type AliYunReader interface {
	Close() error
	ContentType() string
	Read(p []byte) (int, error)
	Remain() int64
	Size() int64
}
