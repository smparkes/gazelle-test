package parquet

import (
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/xitongsys/parquet-go/source"
	"golang.org/x/net/context"
)

type gcsReader struct {
	objHandle *storage.ObjectHandle
	client    *storage.Client
	ctx       context.Context
	length    int64
	offset    int64

	Reader io.ReadCloser
}

func newGcsReader(cli *storage.Client, bucket string, obj string) (*gcsReader, error) {
	ctx := context.Background()
	objHandle := cli.Bucket(bucket).Object(obj)
	oa, err := objHandle.Attrs(ctx)
	if err != nil {
		return nil, err
	}

	reader := &gcsReader{
		objHandle: objHandle,
		client:    cli,
		ctx:       ctx,
		length:    oa.Size,
		offset:    0,
	}

	return reader, nil
}

func NewGcsParquetFile(cli *storage.Client, bucket string, obj string) (source.ParquetFile, error) {
	gcsReader, err := newGcsReader(cli, bucket, obj)
	if err != nil {
		return nil, err
	}
	return gcsReader, nil
}

func (r *gcsReader) Create(name string) (source.ParquetFile, error) {
	return nil, fmt.Errorf("create not supported for gcs reader interface")
}

func (r *gcsReader) Open(name string) (source.ParquetFile, error) {
	return &gcsReader{
		objHandle: r.objHandle,
		client:    r.client,
		ctx:       context.Background(),
		length:    r.length,
		offset:    0,
	}, nil
}

func (r *gcsReader) Seek(offset int64, whence int) (int64, error) {
	// move
	switch whence {
	case 0: // from the beginning of the file
		r.offset = offset
	case 1: // from the current position
		r.offset += offset
	case 2: // from the end of the file
		r.offset = r.length + offset
	}
	return r.offset, nil
}

func (r *gcsReader) Read(b []byte) (int, error) {
	return 0, nil
}

func (r *gcsReader) Write(b []byte) (n int, err error) {
	return 0, fmt.Errorf("write not supported for gcs reader interface")
}

func (r *gcsReader) Close() error {
	return nil
}
