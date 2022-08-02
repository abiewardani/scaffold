package domain

import (
	"time"

	"gitlab.com/abiewardani/scaffold/pkg/storage"
)

var (
	ActiveStorageCompany = ``
	ActiveStorageUser    = ``
)

type ActiveStorageAttachment struct {
	ID                int                `gorm:"column:id;primary_key"`
	Name              string             `gorm:"column:name"`
	RecordType        string             `gorm:"column:record_type"`
	RecordID          string             `gorm:"column:record_id"`
	BlobID            string             `gorm:"column:blob_id"`
	ActiveStorageBlob *ActiveStorageBlob `gorm:"foreignKey:BlobID;references:ID"`
	CreatedAt         time.Time          `gorm:"column:created_at"`
	StorageUrl        string             `gorm:"-"`
}

func (c *ActiveStorageAttachment) TableName() string {
	return `active_storage_attachments`
}

type ActiveStorageBlob struct {
	ID          int       `gorm:"column:id;primary_key"`
	Key         string    `gorm:"column:key"`
	FileName    string    `gorm:"column:file_name"`
	ContentType string    `gorm:"column:content_type"`
	Metadata    string    `gorm:"column:metadadata"`
	ByteSize    int       `gorm:"column:byte_size"`
	Checksum    string    `gorm:"column:checksum"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	ServiceName string    `gorm:"column:service_name"`
}

func (c *ActiveStorageBlob) TableName() string {
	return `active_storage_blobs`
}

type ActiveStorageParamsView struct {
	RecordType string
	RecordID   string
}

func (c *ActiveStorageAttachment) GetStorageUrl() string {
	c.StorageUrl = storage.ObjStorageAws.GetObject(c.ActiveStorageBlob.Key)
	return c.StorageUrl
}
