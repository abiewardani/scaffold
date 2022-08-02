package storage

type Metadata struct {
	Identified bool `json:"identified"`
	Width      int  `json:"width"`
	Height     int  `json:"height"`
	Analyzed   bool
}

type ObjectOutput struct {
	ByteSize    int64
	Checksum    string
	Key         string
	Filename    string
	ContentType string
	Metadata    Metadata
}
