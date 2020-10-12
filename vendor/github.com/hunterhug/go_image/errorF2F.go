package go_image

import "errors"

var (
	ExtNotSupportError = errors.New("ext of filename not support")
	FileNameError = errors.New("filename error")
	FileExistError = errors.New("File already exist error")
)
