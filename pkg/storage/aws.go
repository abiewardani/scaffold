package storage

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"image"
	"io"
	"mime/multipart"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	secureRandom "gitlab.com/abiewardani/scaffold/pkg/secure_random"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var ObjStorageAws StorageAws

type StorageAws interface {
	GetObject(keyObject string) (urlPath string)
	Upload(file multipart.File, fileHeader *multipart.FileHeader) (*ObjectOutput, error)
}

type storageAwsCtx struct {
	Session *session.Session
	AwsCnf  StorageAwsConfig
}

type StorageAwsConfig struct {
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	Region             string
	Bucket             string
	ExpiredTime        int
}

func NewStorageAWS(config *StorageAwsConfig) (StorageAws, error) {
	storageAws := storageAwsCtx{
		AwsCnf: *config,
	}

	var err error

	storageAws.Session, err = session.NewSession(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.AwsAccessKeyID, config.AwsSecretAccessKey, ""),
	})

	if err != nil {
		return nil, err
	}

	ObjStorageAws = &storageAws
	return &storageAws, nil
}

func (c *storageAwsCtx) GetObject(keyObject string) (urlPath string) {
	params := &s3.GetObjectInput{
		Bucket: aws.String(c.AwsCnf.Bucket),
		Key:    aws.String(keyObject),
	}

	svc := s3.New(c.Session)
	req, _ := svc.GetObjectRequest(params)
	url, err := req.Presign(time.Duration(c.AwsCnf.ExpiredTime) * time.Minute) // Set link expiration time
	if err != nil {
		return ``
	}

	return url
}

func (c *storageAwsCtx) Upload(file multipart.File, fileHeader *multipart.FileHeader) (*ObjectOutput, error) {
	// Generate Key
	key, err := secureRandom.GenerateUniqueSecureToken(20)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	contentType := fileHeader.Header.Get("Content-Type")
	res := ObjectOutput{
		Key:         key,
		Filename:    fileHeader.Filename,
		ContentType: contentType,
		ByteSize:    fileHeader.Size,
	}

	// Calculate Checksum File
	checksumFile, err := ComputeCheckSum(file)
	if err != nil {
		return nil, err
	}
	res.Checksum = checksumFile

	// Upload To AWS
	file.Seek(0, 0)
	_, err = s3.New(c.Session).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(c.AwsCnf.Bucket),
		Key:                aws.String(key),
		ACL:                aws.String("public-read"),
		Body:               file,
		ContentLength:      aws.Int64(fileHeader.Size),
		ContentType:        aws.String(fileHeader.Header.Get("Content-Type")),
		ContentDisposition: aws.String(`inline`),
	})

	if err != nil {
		return nil, err
	}

	// Build Metadata
	metadata := Metadata{
		Identified: true,
		Analyzed:   true,
	}
	if strings.HasPrefix(contentType, "image/") {
		file.Seek(0, 0)
		m, _, err := image.DecodeConfig(file)
		if err != nil {
			return nil, err
		}

		metadata.Width = m.Width
		metadata.Height = m.Height
	}
	res.Metadata = metadata

	return &res, nil
}

func ComputeCheckSum(file multipart.File) (checksum string, err error) {
	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		return ``, err
	}

	hashMd5 := hash.Sum(nil)
	h := sha256.Sum256(hashMd5)
	return base64.StdEncoding.EncodeToString(h[:]), nil
}
