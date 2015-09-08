// Code generated by go-bindata.
// sources:
// index.html
// DO NOT EDIT!

// +build bin

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x5c\x7b\x73\xdb\x36\x12\xff\x3f\x9f\x02\xe5\xa5\xf7\x47\xa7\x94\x6a\xc7\xe9\xe5\x52\xc9\x37\x8e\xe5\xb6\x9c\xc6\x8e\xc6\x96\x73\xd7\xe9\x74\x6e\x20\x12\x92\x50\x53\x24\x0b\x40\xb2\x55\x9d\xbe\xfb\x2d\x1e\x7c\x4a\xa4\x1e\x64\x92\xaa\xf1\x8c\x25\x12\xc0\xfe\xb0\x0b\xec\xfe\xf0\x20\xa8\xce\x17\xbd\x77\x97\x83\x9f\xfb\x57\x68\x22\xa6\xfe\xb3\xf3\x8e\xfa\x42\x08\x2e\x08\xf6\xe0\x42\x5e\x4e\x89\xc0\xc8\x9d\x60\xc6\x89\xe8\x5a\x33\x31\xb2\x5f\x59\x26\x4b\x50\xe1\x13\x75\x0d\x77\xcb\x65\xab\x87\x05\x6e\xfd\x78\xb3\x5a\xa1\x90\x0b\x12\x88\x4e\x3b\x2d\x61\x80\x26\x42\x44\x36\xf9\x7d\x46\xe7\x5d\xeb\x3f\xf6\xfd\x85\x7d\x19\x4e\x23\x2c\xe8\xd0\x27\x16\x72\xc3\x40\x8a\x75\x2d\xe7\xaa\x4b\xbc\x31\xb1\xb2\x92\x01\x9e\x92\xae\x35\xa7\xe4\x31\x0a\x99\xc8\x14\x7e\xa4\x9e\x98\x74\x3d\x32\xa7\x2e\xb1\xd5\xcd\xd7\x88\x06\x54\x50\xec\xdb\xdc\xc5\x3e\xe9\x9e\x18\xa0\xe5\xf2\xc7\xc1\xf5\x5b\x64\x75\xbe\xb0\x56\x2b\xdb\xfe\x85\x8e\x90\x73\xf5\xeb\x79\xc7\xa7\xc1\x03\x62\xc4\xef\x5a\x7c\x02\xd8\xee\x4c\x20\x0a\xf0\x16\x9a\x30\x32\xea\x5a\xed\x11\x9e\xcb\xfb\x16\x7c\x6c\x42\xfa\x85\x04\x1e\x1d\xfd\x6a\xdb\x59\x24\x0d\xc0\xe9\x1f\x84\x77\xad\x17\xa7\x4f\x2f\x4e\x8b\x70\x3c\xbe\xb0\x5f\x9c\xb6\xa2\x60\x6c\x21\xb1\x88\xc0\x44\x3a\xc5\x63\xd2\x96\x09\xc6\xfc\x14\x14\x47\x91\x4f\x6c\x11\xce\xdc\x89\x9d\xab\xe0\xe4\xf4\x9b\x27\xf8\x2f\xaf\x02\x32\x5b\x7b\x42\x9e\x9d\x3d\xc1\x7f\x05\xe4\xd9\xd9\xbe\x90\x2f\x4f\x9f\xe0\xbf\x02\xf2\xe5\xe9\xbe\x90\xaf\xc0\xf0\x57\x55\x86\xbf\x2a\x31\x9c\x8b\x85\x4f\xf8\x84\x10\x11\x37\xbc\x20\x4f\xa2\xed\x72\x9e\x80\xc1\x75\x9b\x06\x1e\x79\x6a\xc9\x54\x83\xc0\x5d\x46\x23\x91\x15\xf9\x0d\xcf\xb1\x4e\xb5\x8a\x81\x82\x38\x73\x01\xe8\x37\xde\x66\xd2\xe9\x19\x81\xab\xd3\xd6\x49\xeb\xe4\x55\x9c\xd0\x9a\xd2\xa0\xf5\x1b\xd4\xe9\x41\xf0\xd8\x53\x4c\x03\x5d\x7e\xb9\x04\xef\x6c\x0d\x2e\x7e\xf8\xe1\xaa\x37\xa4\xc1\x6a\x05\xe5\x8c\x32\x5a\x62\xb9\x24\x3e\x27\xab\x15\xd4\xd0\x9e\x52\xff\xc1\x64\xaa\x8c\xc0\x5b\xad\xac\x38\x30\x3b\x6d\xad\x9c\xd1\xbf\x6d\x62\xfb\xbc\x33\x0c\xbd\x85\x49\x0c\xf0\x1c\xb9\x3e\xe6\xd0\xa2\x70\x39\xc4\x0c\xe9\x2f\xdb\x23\x23\x3c\xf3\x45\x7c\xcb\x05\xc4\xaa\x0b\x1d\x11\x59\x88\x85\x10\x5a\xb2\x38\x1d\x43\x22\xf4\x49\x52\xa1\x47\x13\x34\x19\xa5\x60\x13\x61\xf6\xc8\x9f\x51\x2f\x2e\x53\x28\x65\xd0\xa5\x66\x84\x41\x64\xfb\xf6\xd4\xb3\x4f\x50\x84\x3d\x8f\x06\x63\xdb\x27\x23\xa0\x91\xb8\x0b\x62\xf9\xe1\x4c\x88\x30\x28\x40\x88\x70\x3c\xf6\x89\x84\xf0\x71\xc4\x89\x17\xf7\xad\x2e\x6c\x5a\x59\x17\x92\xca\xe9\x52\x71\x32\x66\x63\xd9\x75\x7f\x33\x58\x49\x76\xa6\x5a\xe5\x01\x11\x4e\xaa\xe5\xcc\x0e\x03\x7f\x91\x2f\x02\x85\x06\x5a\x8f\xb4\x75\xa0\x17\x40\xac\x02\x49\x79\x2b\x54\xbb\x06\xf5\xa7\x91\x6c\xeb\x46\xcc\x75\x02\x46\xd4\xeb\x5a\x13\x68\xd9\x7c\x3f\x0c\x19\x0e\x3c\x48\xa3\xd1\x49\x1a\xef\xd0\x19\x72\x3c\x80\xf2\x30\x3e\x48\x3e\x47\xd9\x41\xa3\xd0\xcc\xd9\xac\x4e\x1b\xe7\x6a\x6d\x83\xef\x14\x3c\x49\xaa\x51\xec\xb7\xd4\x09\x75\x02\xaa\xe8\xd8\xf3\xce\xcc\xcf\xd8\x10\x17\x85\xaf\x62\xf7\xfb\x34\x2e\x87\x5d\x41\xe7\x64\xbd\xf1\xb0\x31\x58\x0e\x76\xfc\x75\xbb\xfd\xf8\xf8\xd8\x02\x8b\x19\xfc\xb7\xdc\x70\xda\xd6\xa3\x23\x30\x80\x4f\x30\x27\xbc\xed\x63\x41\xb8\xf8\x97\x3b\x8d\xba\xb1\xd5\xef\xaf\x6e\xef\x9c\x77\x6b\xad\xa2\xf0\xe3\xc1\x15\xaf\xf7\x9a\x4f\x4b\x95\x8d\x03\x04\x18\x0b\x65\xae\x81\x72\xd8\x98\x06\x32\xbc\xd0\x88\x32\x2e\x54\xea\xba\x4d\x6e\xe8\x91\x02\x94\x4c\x02\x3e\x5d\x6b\x00\xa5\x62\xd2\xb5\x3d\xe7\x6e\x70\xeb\xbc\x91\x9d\x28\x05\x1a\xd6\xb9\xae\xa2\x34\xfa\x7b\x30\xe4\xd1\x77\x3a\x2c\xa4\x13\xd1\x68\xbd\x58\xc6\x1d\x9d\xbe\xb4\x64\x2d\x3a\x8c\x2d\x7f\x42\x0b\x67\x6b\x16\xce\xaa\x2d\xbc\x3f\x36\x0b\x7d\x5c\xb4\xd0\xc7\x95\x16\xbe\xbd\x68\xc0\xc2\xf6\xcc\xaf\xe2\xa4\xcc\x2d\xdc\x80\x11\x66\xa4\xdd\x3e\x36\xe6\xca\xb0\xf0\xb1\x38\x66\x4a\x03\xa7\x64\x9a\x65\x37\x39\x5c\x9e\xc5\xe3\x66\x38\x1a\xc1\x0c\xc4\x3e\xc9\xb1\xdb\x72\x29\xc8\x34\x92\x44\x83\x2c\x18\xd3\xa1\x56\xfe\xfa\xb5\xbe\xf8\x2f\xb4\x03\xf1\x15\x62\x4b\xb6\xcb\x66\x6e\xf5\x46\xc5\x0a\x5f\xd6\xab\x50\x02\x6e\xa8\x2f\xd7\x6e\xdb\x1b\xc2\x8d\x66\x0d\x37\x84\x42\xac\x68\x08\xda\x74\x43\xd0\x46\x1a\x22\xe2\x45\xb5\x4e\xbe\xa9\xa7\x97\x44\xac\xaf\xd7\x7c\xdc\xb4\x5e\x12\x71\x9b\x5e\xc9\xcd\x96\xf9\x7a\x2c\x30\x87\x19\xaf\xa4\x06\xd4\x8d\xc7\xad\xd5\xea\xbb\xe2\xc4\xd9\xcc\x97\xe5\x14\x5a\x2d\x97\x97\xcb\xf6\x57\xcf\xbe\x6a\xaf\x56\xcb\xa5\xd6\x2f\xa3\xf0\x10\xc3\xfa\xd5\x93\xab\x0d\x39\x79\x85\x45\x62\x3c\x29\x00\xf8\x1f\xe1\x5b\xae\x91\x83\x4b\x9f\xba\x0f\xdd\xa5\x98\x50\xde\x9a\xc0\x7c\xc9\x27\x2a\x65\x05\x3a\xa8\x26\x5b\xad\x64\xf1\x4b\x7d\x09\x35\x3f\x4b\xe6\x6f\x69\x01\xeb\xc9\xc7\x43\xe2\x23\xfd\x15\xcf\xd7\x93\x65\x6a\xeb\x8d\x54\x24\x43\x77\xe7\xd2\xc0\x01\x34\x83\x9e\x55\x9d\x9b\xd5\xc2\x6e\xb6\x30\x3a\x9e\x34\x69\x4c\xa2\x0a\xaa\x6d\xd6\xae\xc6\x18\x4f\x82\x28\x07\x87\xe5\x96\x2c\xf3\x48\xc5\x04\xc9\x3b\x70\x46\xe9\x59\x00\x2f\x3f\xf4\xf2\x6a\xb9\x84\x99\xec\x98\xa0\xe7\xf4\x6b\xf4\xdc\x0d\x19\x41\xaf\xbb\x48\x0f\x23\x97\xfd\xfb\xd6\x5b\xca\x65\x4b\x0a\x06\xba\x5f\x08\xc1\x7e\x22\x0b\x24\x49\xc9\x06\xbc\xe1\xc2\xbe\xb1\xb4\x50\xeb\x46\xd9\x7c\xde\x11\x5e\x1c\x0d\x6a\x10\x54\x0d\x8a\xd4\x65\x10\x3e\x32\x1c\x25\x16\x26\x62\x9d\xb6\xf0\x4a\x45\x93\x45\x69\xa1\xf5\x66\x9c\x44\x84\xb9\x30\x4d\xd4\xd3\x39\xbd\xc0\x49\x53\xbb\x71\x05\xf7\x9c\x30\xa5\x5a\xae\x5a\x9d\x9a\x19\x22\xcf\x3f\xac\x16\x77\x0b\xbe\xae\x84\x4a\xfc\x78\x3a\xfc\x1b\x53\xb1\xae\x84\x4e\x6d\x5a\x0b\x9b\x06\x73\xc2\x92\x85\xe7\x06\x6d\x1c\x88\x9f\x75\x6d\x74\xea\x46\x6d\xda\x82\x65\xfc\x3f\xff\xb5\x4b\x34\x08\x2c\xb7\xe0\x64\xa1\xe7\x51\xea\xe2\x7d\xcc\xf0\x54\x76\x83\xca\x06\x8b\x54\xfc\xde\x84\xc1\x1f\x84\x85\xe8\x79\x24\x63\x20\x40\x96\xce\x55\x9f\xf6\x24\x04\xd3\x2c\x94\x2c\xba\xec\x09\xf5\x3c\x12\x58\x71\x00\x64\xb6\x16\x41\xe5\x98\xb0\xc5\x24\x4b\xe5\xc9\x9d\xcc\x28\x6b\x65\x95\x2f\x3d\xf5\xcb\x7d\x04\xc0\xab\xf6\x2a\x2f\x1d\x60\x2f\x01\xd9\x47\x79\x81\xb6\xb1\x53\x19\x16\x6f\xbe\x88\x74\xf7\xa5\x6a\xc4\x8b\x79\x4a\x8f\x78\x22\x1d\x82\x54\x63\xef\xc3\x78\xde\x68\x4f\xc2\xf3\x28\x7f\x48\x3d\xa1\xf7\xfd\x66\xae\xf3\x46\x86\xea\x3c\xca\xe4\x92\xde\xd2\x82\xad\x1e\x65\x37\x70\xab\x7b\x1d\x21\x19\x2d\x99\x78\xc8\x30\x5e\xb2\x49\xa0\xf7\x06\xc0\xff\xb4\x3c\x99\x27\xf2\x3a\x02\xf2\xc9\xb1\xe3\x1f\x02\x9d\x51\x2d\x0b\x1d\x27\x57\x46\xf8\x06\xb2\xce\x6d\xa9\xa8\xec\x31\xc3\x0b\x5f\x96\x96\x91\x60\xf6\x3c\x9c\x00\xd6\x15\xb0\xc6\x66\x24\x59\xc3\x24\x55\x3b\x32\x35\x1f\xd8\x28\xce\xbb\x98\x63\xea\xe7\x95\x2a\x5a\xbb\xa6\x18\x1a\x8e\x33\xa4\xb2\x91\x65\x14\x36\x44\x4f\x5f\x27\xc5\x8d\xb1\x97\x2d\x80\xf8\xe5\xba\x2d\xb3\x0c\xe8\x97\x1b\x4d\xba\xcf\x97\xf8\x80\xad\x3d\x93\x5b\x80\x9b\x34\xf4\x4a\x5a\xfb\xde\x64\x7d\x40\x9d\x44\x28\xb0\xbf\xae\x94\xca\x2c\xd1\x6a\x20\x45\x32\x6a\xd5\x22\x7b\x6f\x74\x20\xd7\xf7\x46\x4d\x50\xfd\x46\xb6\xd3\x3b\xbf\x24\x98\x4d\xe5\x34\xd3\x42\xef\xd5\x3e\x3d\x68\xa7\xaa\x7d\x40\x27\xc8\xea\xa9\x07\x3c\x50\x9b\x52\xfc\x00\x8c\x53\x64\x5d\x87\xb3\x40\xc8\x4d\xe1\x83\x41\x5e\x20\x4b\x05\x24\x40\x64\xd8\xff\x40\xb0\x33\x64\xdd\xcb\x10\x6a\x02\xeb\xa5\xc2\xf2\x1a\xc1\xfa\x16\x59\xca\xe3\x8a\x60\xb5\x46\x32\x33\xfe\xd4\x1f\xc8\xe8\xbe\x03\x19\x1d\xa5\x4e\xee\x94\x0c\x63\x34\x1e\xc6\xcc\x18\x46\x47\xad\x43\x87\xaf\xac\xa8\x0e\xf0\x24\xe5\xc3\xf0\x8a\xea\x2b\xf4\xe6\xe7\xc1\xd5\x1d\x72\x6e\xd0\x34\xf4\x66\x7e\x88\xce\x7e\xc8\x72\x0c\xa8\xf0\x66\x21\x08\x77\x82\x02\xc5\xe8\x2a\x0c\xd2\x1b\x2a\xb8\x84\x80\x71\x02\x71\xe2\x86\x81\x57\xc0\xe8\x11\x5f\x60\x59\xaa\x88\xf3\xe1\x2d\x7b\x77\x3f\xa8\x36\xed\xdd\x4c\x6c\xb5\x4d\x82\x6c\x37\x6e\x0d\xe9\x43\x5a\x17\x61\xf7\x81\xe8\x66\x2f\x33\xaf\xaf\x8b\x54\xf7\x5d\x3f\xc5\xa9\xb6\xb0\x04\xed\x63\xd8\x58\xd5\x87\x46\xad\x2d\xbd\xd8\xcf\x20\xed\x64\xe6\x47\xed\x4b\xc2\x58\xc8\x2a\xbb\xf2\x4a\x95\xa8\xee\xc9\xab\x04\xa5\xda\xc2\xcd\x58\x1f\xc1\xbe\xaa\x6e\xd4\x4a\x6d\xe9\xc5\xab\x14\x67\x17\x13\xcb\xfb\xb0\xd6\x44\x88\x1e\x3a\x11\x72\x1a\x99\x08\xe5\xd6\xbc\x0e\xcc\x4b\xd8\x08\xbb\x64\xeb\x52\x13\x05\x21\x9b\xca\xb1\x19\xc6\xa3\xd0\x97\x6d\x22\x47\xa4\x6f\x33\x2b\x50\xc5\x74\x9d\xe1\x76\x52\x37\x35\x77\xda\xc3\x54\xf8\xeb\xa2\xe4\xe6\x4e\x52\x65\xa1\x5f\x72\xb2\xe8\x7f\x28\x8e\xcf\x14\xa4\x92\x97\xb6\x69\x50\x1d\xee\x65\x4a\x18\xf7\x4a\x61\xaa\x22\x6a\x9b\x0a\x95\xbe\xba\x51\x83\xc6\x16\xff\xb4\xa9\x29\xd3\x94\x4c\x0f\xd8\xed\x04\xa9\x34\x28\xae\xaf\xae\x37\xcf\x9c\xa0\x94\x99\x3a\x3d\x50\x68\x18\x25\xd6\xfa\x89\x4a\xcd\x0c\x0d\x25\xf3\xa0\x34\x67\x97\x4d\x34\x23\xf0\xbd\x59\x1a\xef\x22\x50\xde\x94\xd9\xd5\xb0\x82\xcd\xae\x41\xf5\xe3\x43\x53\xdf\x0e\x2b\xc0\xbc\x82\x8d\x2e\xce\x00\xf0\x40\x52\xba\x26\xd3\x4f\xba\x13\x27\xfb\x69\xcf\x9d\x3b\x6f\x9f\xf2\xaa\x99\x1b\x0b\xad\x38\x22\xea\xc7\x56\xf2\xd0\xb0\xb4\xc7\xe4\xa3\x30\xf9\x67\xfa\xac\x1f\x72\x2a\xcf\x8d\xa4\xdb\xa7\xd0\x47\x0a\x05\xa9\xcf\xe4\x61\x47\xb2\xf8\xd0\x00\x5b\xc5\xd5\xf9\x29\x1a\x8c\xad\x78\xe9\x21\x3f\x3a\x58\x89\xca\xc7\x34\x03\xb3\xc4\x93\x52\xab\xec\x73\x19\x23\xae\x08\x0f\x0d\x45\x60\x0f\xfd\xd0\x7d\x48\xc6\x63\xfd\x05\x8c\x58\xaa\xc1\xe4\xcc\x92\x1f\x72\xdb\x89\x06\xa3\xd0\x4a\x76\xae\x41\x1a\x4a\xa4\xec\xa8\x55\x32\x27\x57\xa4\x69\xe6\x41\x9d\x5a\x62\x65\x1d\x7b\xbd\x0e\x9d\x5d\xf4\x65\x93\x9e\x98\x2c\xc9\xc9\x78\x91\x36\x0a\xd8\x7a\x44\xc7\xe9\xb3\xf5\x34\x9e\xcd\x43\xc8\xd3\x92\x33\x69\x55\x93\x26\x53\x1c\xa6\x28\x78\x91\x3d\x7b\x60\x44\xd5\x03\xb3\xb5\xe3\x4c\xda\x12\x6f\xc3\xd1\x83\xf5\xc3\x03\x86\x34\x33\xfc\xb9\xa6\x9f\xec\xa7\x31\x0b\x67\x11\xd2\x9a\xe9\x9b\xf8\x58\x9e\xbe\x79\x96\x53\x60\x43\x34\x14\x1e\x53\xa2\xb7\x84\xf3\x5e\xea\x26\x1e\x92\xd5\x28\x97\x48\x9d\x72\x0b\x90\x8e\x5b\x74\x1d\x32\xb2\x0d\xa9\xda\xec\xb5\xc7\xc6\x56\x49\xd1\x94\x07\x3e\x66\xe7\xdf\x02\x7d\xec\xd9\xf7\x41\xeb\x62\xc8\x43\x7f\x26\xc8\x9f\xde\x09\x6e\x32\x5c\x51\xcf\x09\xb6\x21\x35\xed\x04\x92\x54\x34\x89\xab\x9b\x2d\x0f\x57\x34\x7b\xa4\xe7\x09\xf6\x25\x7e\x6f\x74\x10\xef\xab\xad\xd4\x83\x69\x3f\x27\xbd\x2b\xeb\x83\x50\x53\xa4\xaf\xea\xaf\xe2\xfc\x9e\x7c\x74\x34\xe3\x78\x4c\x6a\x52\x7f\x66\xcb\xf9\xb8\x99\xbf\x37\x3a\x1e\xe2\x07\x5d\x9b\xe1\xfd\x12\xa0\xcf\x8a\xf6\xc1\x83\x8f\x8f\xf5\x55\xd8\x35\x41\xfa\x25\x40\x9f\x90\xf3\xe3\xa7\x5f\x87\x53\x3e\x3d\x8c\xf2\x9d\x5a\x94\xef\x1c\x42\xf9\x4e\x73\x94\xef\x6c\xa3\xfc\x64\x03\x8b\xd7\xa4\x7c\xe7\xaf\x42\xf9\xce\x11\x51\xbe\xd3\x14\xe5\x97\x00\x7d\x56\x94\xef\x1c\x23\xe5\x3b\x4d\x51\x7e\x09\xd0\x27\xa4\x7c\x5a\x9b\xf2\xe5\xe1\xf8\x43\x38\x5f\xef\xc9\x1d\x4c\xfa\x79\xf1\x5d\x59\x5f\x4a\x35\x45\xfb\x5a\x83\x2a\xde\x87\x12\x21\x5b\xd4\xe4\xfc\xec\xde\xe5\x71\x93\x3e\x58\x72\x3c\xac\x2f\x95\x6d\x86\xf6\xcb\x90\x3e\x2b\xde\x97\x5e\x7c\x7c\xc4\xaf\x63\xaf\x09\xe6\x2f\x43\xfa\x84\xd4\x9f\x3c\x4e\x39\x9c\xfb\x23\x7e\x10\xf5\xf7\x79\x1d\xe6\xcf\x49\xef\x4a\xfc\x20\xd4\x14\xef\xab\xfa\xab\x68\xbf\xcf\x42\x98\xea\xf3\xda\xb3\x7d\x55\xd1\x5f\x81\xf8\xfb\xfc\x78\x78\x1f\x74\x6d\x86\xf6\x4b\x80\x3e\x2b\xd6\x07\x0f\x3e\x3e\xd2\x57\x61\xd7\x04\xe7\x97\x00\x7d\x42\xca\x8f\x78\x5d\xc6\x9f\x8f\x0f\x62\xfc\xf7\xe3\x3a\x8c\x9f\x93\xde\x95\xf1\x41\xa8\x29\xc6\x57\xf5\x57\x31\xfe\x7b\x3c\x66\x38\x10\x35\xf9\x5e\x55\xf3\x57\xe0\xfb\xf7\xe3\xe3\xe1\x7b\xd0\xb5\x19\xbe\x2f\x01\xfa\xac\xf8\x1e\x3c\xf8\xf8\xf8\x5e\x85\x5d\x13\x7c\x5f\x02\xf4\x09\xf9\x7e\x3e\xae\xc3\xf7\x7c\xcf\x73\x71\x11\xcc\x7b\x33\xc3\xc2\xdd\xe6\x33\x71\x11\x37\x47\xe2\x22\x2a\x4f\xc4\x49\xa1\x56\xdf\xe9\x55\xbc\x03\x6c\x8e\x94\xc9\x37\x7e\xd2\xd2\xbb\x1c\x42\x4b\x24\xee\x0b\x12\xcf\x12\x97\x55\xb9\xe6\x8d\xde\x3d\x00\xfb\x8c\x86\x8c\x8a\xc5\x9e\x62\x37\xd4\xdd\xf1\xbc\x5e\x22\x72\x47\xff\xd8\x57\xe4\x96\x70\xea\xa9\xc3\x7b\x65\x62\xf2\x68\x1f\x61\x56\xbe\x19\x06\xb4\xf0\x1a\x46\x3e\xbb\xf0\x96\x46\xad\x93\x7c\xf1\x3c\x64\xef\x83\x7c\x99\x25\xd1\xc7\x7c\xcd\xaa\xcf\xd5\x6b\x56\xe0\x78\xb5\xdf\x21\x92\x50\xa7\xc8\xba\x6f\x08\xea\x05\x40\xdd\x5d\xdd\x1e\xf8\xd6\x96\x44\x38\x03\xbb\x6e\x1b\xd1\xe5\x25\xb2\x6e\x9c\x46\x90\xbe\x45\xd6\x7b\xe7\x76\xd0\x08\xd6\x3f\x90\x75\x7b\x75\xd7\x08\xd4\x2b\x64\x0d\x9c\xeb\xab\x18\xcb\x44\xd1\x81\x60\xff\x44\xd6\xe5\xbb\xeb\xeb\x8b\x9b\x9e\xee\xbc\x5a\x47\x36\x0d\x57\xd7\x3f\xb1\x99\x39\x0b\x0c\x50\xc5\x57\xf5\xd5\x0f\x4e\x6c\x7a\x65\xb6\xa5\xe3\x4d\x5d\xc4\xef\xb5\xee\x51\xeb\x7c\xac\xde\x9e\xb0\xf4\x60\xf1\x40\x16\xf2\xb7\x51\x74\x52\x4c\x5f\xb9\x13\xfd\x2f\x13\xea\x32\xbf\x08\xa6\xe7\xdf\xea\x38\x7a\x29\x47\xed\xa0\x45\x32\xe0\x99\x11\x6d\x8a\xdd\x49\xca\x4e\xa6\x96\x6b\x48\x94\x72\x9b\x87\xb7\xb9\x2e\x64\xc6\xb8\xd9\x4c\x0d\x72\x12\xa7\x75\x7f\x9f\x19\xe5\x12\x6e\xcd\x64\x6d\xa2\x5e\x95\xbd\xf6\x82\x5c\x3e\xbb\xcf\xc2\x39\x10\x3e\xab\x28\x72\x27\xb0\xa8\x82\x30\xa6\x8d\xa8\x2f\xd7\x79\x62\xb2\x8d\xe8\x77\x68\x4b\x0e\x93\x06\x77\xb2\xcb\xf4\x81\x8e\x10\xf9\x1d\xad\xf7\xa4\x87\x04\x9b\x91\x92\xf0\x2a\x78\x0e\xca\x21\x56\x95\x4f\x02\xe5\xd0\xc1\x2b\x9e\x54\xed\x3d\x78\x65\xd6\x77\x8d\xbd\x1a\xd3\x2b\x1c\x13\x4f\xb3\xa4\xd3\x94\x66\xc6\x2e\x53\x5a\x40\x39\x4c\x69\x6e\x8f\x32\xe2\x0a\xf5\x6c\xab\xa1\x33\xe7\x89\xbf\xd4\xa0\xb0\x84\x6b\x73\x3f\xc2\xf3\x16\x52\xb6\xfe\x10\x8f\x21\x2f\x59\x76\xbf\xdf\xff\x29\xb0\xbc\xec\xfd\x49\x96\x31\xcd\xcf\x76\x02\xf6\x85\x4f\xc7\xc1\xa5\x4e\xd7\xac\xb6\x51\xc7\xdc\x2f\x02\xc9\x74\x23\x12\xcf\xfd\xf3\x5a\x16\x08\x5a\xfe\x88\x10\x66\x44\x64\x65\x54\x23\x66\x5f\x0f\xc3\xa6\x5d\x27\x89\x81\xff\x0f\x00\x00\xff\xff\x16\x52\xf2\x0f\xec\x58\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 22764, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"index.html": indexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.html": &bintree{indexHtml, map[string]*bintree{
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, filepath.Join(name, child))
                if err != nil {
                        return err
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

