package fileutils

type FileInfo struct {
	Name             string
	Executable       bool
	RawSize          int64
	HumanizeSize     string
	Modified         int64
	HumanizeModified string
	IsLink           bool
	LinkPath         string
}
