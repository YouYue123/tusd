// Package aliyunstore provides an AliYun(also known as Alibaba Cloud) storage based backend.
//
// AliYunStore is a storage backend that uses the GCSAPI interface in order to store uploads
// on AliYun Object Storage Service(OSS). Uploads will be represented by two files in OSS.
// the data file will be stored as an extensionless object [uid] and the JSON info file will stored as [uid].info.
// In order to store uploads on OSS, make sure to specify the appropriate AliYun service
// account file path in the ALIYUN_SERVICE_ACCOUNT_FILE environment variable.

package aliyunstore

import (
	"context"
	"io"
	"strings"

	"github.com/tus/tusd/pkg/handler"
)

type AliYunStore struct {
	// Specifies the bucket that uploads will be stored in
	Bucket string

	// ObjectPrefix is prepended to the name of each object that is created.
	// It can be used to create a pseudo-directory structure in the bucket,
	// e.g. "path/to/my/uploads".
	ObjectPrefix string

	// Service specifies an interface used to communicate with the AliYun
	// object storage service backend. Implementation can be seen in aliyunservice.go file.
	Service AliYunService
}

// New constructs a new AliYun storage backend using the supplied AliYun bucket name
// and service object.
func New(bucket string, service AliYunService) AliYunStore {
	return AliYunStore{
		Bucket:  bucket,
		Service: service,
	}
}

func (store AliYunStore) UseIn(composer *handler.StoreComposer) {
	composer.UseCore(store)
	composer.UseTerminater(store)
}

func (store AliYunStore) NewUpload(ctx context.Context, info handler.FileInfo) (handler.Upload, error) {
	// if info.ID == "" {
	// 	info.ID = uid.Uid()
	// }

	// info.Storage = map[string]string{
	// 	"Type":   "aliyunstore",
	// 	"Bucket": store.Bucket,
	// 	"Key":    store.keyWithPrefix(info.ID),
	// }

	// err := store.writeInfo(ctx, store.keyWithPrefix(info.ID), info)
	// if err != nil {
	// 	return &aliYunUpload{info.ID, &store}, err
	// }

	return &aliYunUpload{info.ID, &store}, nil
}

func (store AliYunStore) GetUpload(ctx context.Context, id string) (handler.Upload, error) {
	return &aliYunUpload{id, &store}, nil
}

func (store AliYunStore) AsTerminatableUpload(upload handler.Upload) handler.TerminatableUpload {
	return upload.(*aliYunUpload)
}

func (store AliYunStore) writeInfo(ctx context.Context, id string, info handler.FileInfo) error {
	// data, err := json.Marshal(info)
	// if err != nil {
	// 	return err
	// }

	// r := bytes.NewReader(data)

	// i := fmt.Sprintf("%s.info", id)
	// params := GCSObjectParams{
	// 	Bucket: store.Bucket,
	// 	ID:     i,
	// }

	// _, err = store.Service.WriteObject(ctx, params, r)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (store AliYunStore) keyWithPrefix(key string) string {
	prefix := store.ObjectPrefix
	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	return prefix + key
}

type aliYunUpload struct {
	id    string
	store *AliYunStore
}

func (upload aliYunUpload) WriteChunk(ctx context.Context, offset int64, src io.Reader) (int64, error) {
	// id := upload.id
	// store := upload.store

	// prefix := fmt.Sprintf("%s_", store.keyWithPrefix(id))
	// filterParams := GCSFilterParams{
	// 	Bucket: store.Bucket,
	// 	Prefix: prefix,
	// }

	// names, err := store.Service.FilterObjects(ctx, filterParams)
	// if err != nil {
	// 	return 0, err
	// }

	// maxIdx := -1

	// for _, name := range names {
	// 	split := strings.Split(name, "_")
	// 	idx, err := strconv.Atoi(split[len(split)-1])
	// 	if err != nil {
	// 		return 0, err
	// 	}

	// 	if idx > maxIdx {
	// 		maxIdx = idx
	// 	}
	// }

	// cid := fmt.Sprintf("%s_%d", store.keyWithPrefix(id), maxIdx+1)
	// objectParams := GCSObjectParams{
	// 	Bucket: store.Bucket,
	// 	ID:     cid,
	// }

	// n, err := store.Service.WriteObject(ctx, objectParams, src)
	// if err != nil {
	// 	return 0, err
	// }

	return 0, nil
}

func (upload aliYunUpload) GetInfo(ctx context.Context) (handler.FileInfo, error) {
	// id := upload.id
	// store := upload.store

	info := handler.FileInfo{}
	// i := fmt.Sprintf("%s.info", store.keyWithPrefix(id))

	// params := GCSObjectParams{
	// 	Bucket: store.Bucket,
	// 	ID:     i,
	// }

	// r, err := store.Service.ReadObject(ctx, params)
	// if err != nil {
	// 	if err == storage.ErrObjectNotExist {
	// 		return info, handler.ErrNotFound
	// 	}
	// 	return info, err
	// }

	// buf := make([]byte, r.Size())
	// _, err = r.Read(buf)
	// if err != nil {
	// 	return info, err
	// }

	// if err := json.Unmarshal(buf, &info); err != nil {
	// 	return info, err
	// }

	// prefix := store.keyWithPrefix(id)
	// filterParams := GCSFilterParams{
	// 	Bucket: store.Bucket,
	// 	Prefix: prefix,
	// }

	// names, err := store.Service.FilterObjects(ctx, filterParams)
	// if err != nil {
	// 	return info, err
	// }

	// var offset int64 = 0
	// var firstError error = nil
	// var wg sync.WaitGroup

	// sem := make(chan struct{}, CONCURRENT_SIZE_REQUESTS)
	// errChan := make(chan error)
	// ctxCancel, cancel := context.WithCancel(ctx)
	// defer cancel()

	// go func() {
	// 	for err := range errChan {
	// 		if err != context.Canceled && firstError == nil {
	// 			firstError = err
	// 			cancel()
	// 		}
	// 	}
	// }()

	// for _, name := range names {
	// 	sem <- struct{}{}
	// 	wg.Add(1)
	// 	params = GCSObjectParams{
	// 		Bucket: store.Bucket,
	// 		ID:     name,
	// 	}

	// 	go func(params GCSObjectParams) {
	// 		defer func() {
	// 			<-sem
	// 			wg.Done()
	// 		}()

	// 		size, err := store.Service.GetObjectSize(ctxCancel, params)

	// 		if err != nil {
	// 			errChan <- err
	// 			return
	// 		}

	// 		atomic.AddInt64(&offset, size)
	// 	}(params)
	// }

	// wg.Wait()
	// close(errChan)

	// if firstError != nil {
	// 	return info, firstError
	// }

	// info.Offset = offset
	// err = store.writeInfo(ctx, store.keyWithPrefix(id), info)
	// if err != nil {
	// 	return info, err
	// }

	return info, nil
}

func (upload aliYunUpload) FinishUpload(ctx context.Context) error {
	// id := upload.id
	// store := upload.store

	// prefix := fmt.Sprintf("%s_", store.keyWithPrefix(id))
	// filterParams := GCSFilterParams{
	// 	Bucket: store.Bucket,
	// 	Prefix: prefix,
	// }

	// names, err := store.Service.FilterObjects(ctx, filterParams)
	// if err != nil {
	// 	return err
	// }

	// composeParams := GCSComposeParams{
	// 	Bucket:      store.Bucket,
	// 	Destination: store.keyWithPrefix(id),
	// 	Sources:     names,
	// }

	// err = store.Service.ComposeObjects(ctx, composeParams)
	// if err != nil {
	// 	return err
	// }

	// err = store.Service.DeleteObjectsWithFilter(ctx, filterParams)
	// if err != nil {
	// 	return err
	// }

	// info, err := upload.GetInfo(ctx)
	// if err != nil {
	// 	return err
	// }

	// objectParams := GCSObjectParams{
	// 	Bucket: store.Bucket,
	// 	ID:     store.keyWithPrefix(id),
	// }

	// err = store.Service.SetObjectMetadata(ctx, objectParams, info.MetaData)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (upload aliYunUpload) Terminate(ctx context.Context) error {
	// id := upload.id
	// store := upload.store

	// filterParams := GCSFilterParams{
	// 	Bucket: store.Bucket,
	// 	Prefix: store.keyWithPrefix(id),
	// }

	// err := store.Service.DeleteObjectsWithFilter(ctx, filterParams)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (upload aliYunUpload) GetReader(ctx context.Context) (io.Reader, error) {
	// id := upload.id
	// store := upload.store

	// params := GCSObjectParams{
	// 	Bucket: store.Bucket,
	// 	ID:     store.keyWithPrefix(id),
	// }

	// r, err := store.Service.ReadObject(ctx, params)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
