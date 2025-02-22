package fileutils

type FileInfo struct {
	Name             string
	RawSize          int64
	HumanizeSize     string
	Modified         int64
	HumanizeModified string
	LinkPath         string
}
