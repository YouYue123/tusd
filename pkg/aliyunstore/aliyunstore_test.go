package aliyunstore_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tus/tusd/pkg/handler"
)

// go:generate mockgen -destination=./aliyun_mock_test.go -package=aliyunstore_test aliyunstore

const mockID = "123456789abcdefghijklmnopqrstuvwxyz"
const mockBucket = "bucket"
const mockSize = 1337
const mockReaderData = "helloworld"

var mockTusdInfoJson = fmt.Sprintf(`{"ID":"%s","Size":%d,"MetaData":{"foo":"bar"},"Storage":{"Bucket":"bucket","Key":"%s","Type":"aliyunstore"}}`, mockID, mockSize, mockID)
var mockTusdInfo = handler.FileInfo{
	ID:   mockID,
	Size: mockSize,
	MetaData: map[string]string{
		"foo": "bar",
	},
	Storage: map[string]string{
		"Type":   "aliyunstore",
		"Bucket": mockBucket,
		"Key":    mockID,
	},
}

func TestNewUpload(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	// assert := assert.New(t)
	// service := NewMockAliYunService(mockCtrl)
	// store := aliyunstore.New(mockBucket, service)
	// assert.Equal(store.Bucket, mockBucket)
	// data, err := json.Marshal(mockTusdInfo)
	// assert.Nil(err)
	// r := bytes.NewReader(data)

	// params := aliyunstore.AliYunObjectParams{
	// 	Bucket: store.Bucket,
	// 	ID:     fmt.Sprintf("%s.info", mockID),
	// }

	// ctx := context.Background()
	// print(ctx)
	// service.EXPECT().WriteObject(ctx, params, r).Return(int64(r.Len()), nil)
	// upload, err := store.NewUpload(context.Background(), mockTusdInfo)
	// assert.Nil(err)
	// assert.NotNil(upload)
}
