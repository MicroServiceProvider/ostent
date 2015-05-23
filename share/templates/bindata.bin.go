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
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
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

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1d\x69\x6f\xdb\x46\xf6\x7b\x7e\x05\x57\x4d\x17\x4d\x51\x49\x9b\x6c\xbb\x2d\x52\xdb\x80\x63\xc9\x89\x50\x1f\x82\x25\xbb\xdb\x4f\x05\x25\x8e\xa4\x89\x29\x92\x25\x47\x72\x5c\xc1\xff\x7d\xdf\x9b\x83\xe7\x8c\x4c\x52\x87\x9d\x6c\x0a\x14\xa1\x38\xf3\xe6\xcd\xbb\x0f\x92\xe3\x83\x7f\x74\x2e\x4f\x86\x7f\xf4\xbb\xd6\x8c\xcd\xdd\x17\x47\x07\xfc\x1f\xcb\x82\x0b\x62\x3b\x70\x81\x97\x73\xc2\x6c\x6b\x3c\xb3\xc3\x88\xb0\xc3\xc6\x82\x4d\x9a\xbf\x34\xd2\x43\x33\xc6\x82\x26\xf9\x6b\x41\x97\x87\x8d\xff\x36\xaf\x8f\x9b\x27\xfe\x3c\xb0\x19\x1d\xb9\xa4\x61\x8d\x7d\x8f\x11\x0f\xe0\x7a\xdd\x43\xe2\x4c\x49\x06\xd2\xb3\xe7\xe4\xb0\xb1\xa4\xe4\x2e\xf0\x43\x96\x9a\x7c\x47\x1d\x36\x3b\x74\xc8\x92\x8e\x49\x93\xff\xf8\xc1\xa2\x1e\x65\xd4\x76\x9b\xd1\xd8\x76\xc9\xe1\x6b\xb5\x10\xa3\xcc\x25\xfc\x1a\x7e\xad\x56\xad\x8e\xcd\xec\xd6\x07\x3f\x62\xb8\xf8\xc3\x83\x05\x57\xb0\xe4\x41\x3b\x99\x77\x74\xe0\x52\xef\xd6\x0a\x89\x7b\xd8\xa0\x80\xb2\x61\xb1\xfb\x00\xf6\x41\xe7\xf6\x94\xb4\x03\x6f\xda\xb0\x66\x21\x99\x1c\x36\xda\x13\x7b\x89\x13\x5a\x78\xaf\x00\x1a\xb1\x7b\x97\x44\x33\x42\x98\x5a\x80\x91\x4f\xac\x3d\x8e\xa2\x18\x1e\xae\xdb\xd4\x73\xc8\xa7\x16\xde\x95\x2b\x44\xe3\x90\x06\x2c\x0d\xf2\xd1\x5e\xda\xe2\x6e\x23\xcf\x67\x2b\x0a\xc7\xb0\xd0\xc7\xa8\x1d\x22\x87\x43\x02\x57\x6f\x5a\xaf\x5b\xaf\x7f\x56\x37\x5a\x73\xea\xb5\x3e\x02\x4e\x07\x08\x6f\xce\x6d\xea\x89\xf9\xab\x15\x9d\x58\xad\xe1\xf1\xfb\xf7\xdd\xce\x88\x7a\x0f\x0f\x30\x4f\x6e\x46\x40\xac\x56\xc4\x8d\x80\x43\x80\xa1\x3d\xa7\xee\xad\x1c\xe4\x03\x9e\xf3\xf0\xd0\x50\x4c\x3d\x68\x8b\xcd\xc9\xfd\xb7\xa5\x6a\x1c\x1d\x8c\x7c\xe7\x5e\xde\xf4\xec\xa5\x35\x76\xed\x28\x3a\x6c\xc0\xe5\xc8\x0e\x2d\xf1\x4f\xd3\x21\x13\x7b\xe1\x32\xf5\x73\x42\x3f\x11\xa7\xc9\xfc\xa0\x61\x85\x3e\x88\x11\x67\xd3\x29\x28\x0b\x88\x21\xc6\xe7\xd0\x78\x31\xd4\x08\x20\x89\x00\xa4\xbb\xa0\x8e\x9a\x93\x9b\xe5\x84\x7e\xe0\xf8\x77\x9e\xc2\x82\x3b\x24\x61\x32\x99\x6f\x76\xc1\x98\xef\x65\x37\x09\x1b\x99\x4e\x5d\x02\x7a\xe7\xba\x76\x10\x11\x47\x49\x52\x4c\x96\x3c\x15\x93\x70\x2f\x62\x96\xba\x6d\x87\x53\x14\xd4\x37\x72\xad\x78\x38\x85\x96\xcb\x3b\xb0\x63\xb4\x51\xd8\xf4\x3d\xf7\x3e\x3b\x05\x26\x0d\xc5\x3e\x12\x66\x00\xcf\x01\x6c\xcd\x4a\xa8\x97\x4d\x40\x5b\x58\xea\xd9\x40\xb6\x05\x13\x33\x42\xb0\x2d\xea\x24\xe2\x1a\x85\xb6\x07\x2c\xcf\x09\x51\x09\x45\xb2\x55\x4e\xca\x48\x42\x4d\x6d\x58\x76\x48\xed\x26\xf9\x04\xd8\x1d\x02\x2b\xb3\x70\x41\x62\xeb\xcb\x0b\x02\x15\x06\xd1\xcf\xa4\x73\x28\x90\x51\xf4\x1e\x07\x6d\x00\xca\xd3\x65\x67\x48\x5a\xb8\x05\x02\xe6\xc4\x5b\x28\xfd\x16\xd7\x7c\x9b\xae\x3d\x22\xae\x4b\x9c\xd1\x7d\x9e\x05\x39\x0c\x2e\x2d\x2c\x29\x35\x5a\x2e\x1a\x84\x24\x02\xa7\x96\x31\x9b\x18\x3c\x5c\x78\x1e\xf5\xa6\x07\x6d\x97\x16\xd7\x2d\x01\x0f\x42\x2a\x48\x04\xbc\x54\x7e\x5e\x8a\x5f\x9d\xde\x60\x78\xd5\x7b\x87\xec\xb2\x8b\xaa\x51\xd8\x46\x7b\xe1\x66\x38\x98\xe1\x71\x22\xa6\xbc\x59\x25\x2e\x41\xdc\xb0\xd6\xd8\x5d\x5a\x2c\xe8\x9a\xe4\x54\xf8\x47\xc3\x6b\x44\x46\x83\x46\xde\x35\x68\x68\x8e\x29\xee\xf5\x91\x58\x1d\x83\x73\xac\x2b\x72\x57\x91\xa7\x66\xb8\xb6\x0e\xb3\x95\xb3\x86\x4a\x06\x50\x94\x54\x8c\x15\xb0\x15\x46\x53\x74\x9d\x1d\xeb\x74\x5e\x23\xa6\x22\x9b\xeb\x68\xbf\x6e\x37\x1b\xaa\x3f\x5f\x62\x11\x30\x3a\x27\x05\x01\xc5\xcb\x97\x5c\x46\xba\x2b\xb1\x9a\xc6\x4f\x69\xad\x22\xc3\xcf\x6b\x0e\xaa\x33\x0c\xc9\x53\xed\x0e\xb3\x06\xa2\x9f\x59\x34\x23\x93\xc6\xab\xcb\x90\x4e\x67\xcc\xec\x6a\xd6\xa9\x7d\x6a\x6d\x97\x46\xac\x49\x3d\x48\x83\x88\xd1\xaa\xd4\xd2\x06\x9e\x0a\xf7\xfc\x0d\x26\x01\x7a\xee\xe1\x50\x35\x8e\x3d\x86\x2b\x88\x4c\xa8\x82\x68\xbb\x98\x96\x53\x13\xa6\xa5\x3d\x05\x67\xcf\x76\xa1\x0a\x45\x8b\xb1\xec\x31\xa3\xcb\x62\x88\xcb\x45\xe0\x25\x09\x23\xd4\x7e\x43\x0c\xae\x13\x76\xb1\x1a\x88\xde\xb6\xdb\x77\x77\x77\x2d\x88\xa3\x21\xfc\xdf\x1a\xfb\xf3\xb6\xc8\xc3\x21\x6b\x75\x89\x1d\x91\xa8\xed\xda\x8c\x44\xba\xa8\x72\x39\x18\x76\x2f\x86\xba\x50\xb2\x99\xb3\x51\xb4\xee\xc4\xe3\xc8\xc5\xb7\xe1\x72\xb6\xc3\x46\xbe\x9a\x72\x43\x37\xdd\xab\x41\xef\xf2\x62\x3f\x7e\x28\x17\xce\x53\x3f\xe1\x07\xf8\x0a\x59\x31\xa8\x88\xc4\x6b\x01\x7d\xb6\x6f\x05\xb6\xe3\x40\x2a\x83\xf5\x42\x44\x96\x40\xf7\x7d\x34\xd3\x96\x09\xa1\x7f\x67\x28\x0d\x46\xae\x3f\xbe\xc5\x04\xbf\x39\x77\x9a\x6f\xd4\x85\x3f\x99\x40\x91\xd5\x7c\x9d\xcd\x19\x52\x60\xa0\xd9\xc4\x7d\x5d\xf0\x96\xa8\x56\xf9\x3c\x44\x28\x9b\xc5\x21\xb8\xda\xc0\x8e\xad\x11\x83\x4c\x19\x51\x37\xac\x89\x1f\x82\x5e\x2d\xe6\xf3\xfb\x39\x99\x03\x85\x13\x3a\x1d\xcf\xc8\xf8\x76\xe4\x7f\xe2\x83\xcd\x08\xab\xc9\x6f\xe2\xc1\x35\x09\x19\xae\x2a\x14\x53\x2c\xce\x6b\x3c\xf2\x97\xd5\x98\xd8\x2e\x66\x48\xdf\x7d\x8c\xa0\xc4\x11\x32\x3f\x71\x29\x30\xac\xf5\x81\x3a\x44\xac\x7b\xde\x3d\x7f\x05\xf5\xb0\x70\x0c\xaa\xc6\x53\x6e\x0b\xb0\x6b\x0c\xe3\x9c\xcc\xfd\xf0\x5e\x67\x8a\xd4\x0b\x16\x4c\xb8\x12\x03\x65\x79\x2e\x25\x23\xa2\xc4\x4a\x7e\x73\x8b\x9d\x51\xc7\x21\x5e\xec\x4a\xf8\x2f\x8d\xf2\x21\xab\x5f\xe8\xb6\xb2\x6d\x74\x5c\x1b\x8c\xc9\x5b\x22\x2c\x85\x38\x16\x85\x58\xb1\x84\x24\xe2\xad\x0a\xec\xb9\xaa\x3b\x8d\x15\x94\x64\xae\xf0\xcc\xfc\x90\xfe\x8d\x66\xe2\x36\xf9\xed\x91\x1f\x72\x75\xf0\xa1\xd6\x9a\xf3\x5b\x7a\xbf\x92\xd2\x6d\x9c\xd4\x9c\x86\xfe\x22\x68\xa2\x81\x11\xc7\xe0\x3d\x32\x66\x04\x9a\xc7\x41\xac\xf8\xaa\x19\xcd\x73\xf1\x41\x14\x7c\x86\x50\x5b\xb0\x1f\x58\x88\x2f\xa6\x3a\x04\x59\x7b\xca\x99\x86\x69\xcd\x44\x01\x72\x62\x36\xcf\xff\xc0\xd9\xad\xd3\xa5\xb2\x5b\xe5\xb2\x86\xfc\xe7\x11\xb3\x1b\xfc\x7e\xdc\x37\x5b\x1c\x48\x9d\x44\x77\xb6\x21\x07\xaa\x4b\x19\xb1\x70\xcd\xc7\x88\xd3\x66\xf8\x6b\x87\x6a\x29\xd0\x6a\x75\x47\xd9\x2c\xcb\x98\x2b\x32\x81\x18\x38\x03\x13\x78\x78\x58\xad\x18\x99\x07\x18\xc1\xac\x06\xb0\x16\x1c\x7f\xf4\xf6\x6d\x28\x26\x34\xac\x16\x4e\xe0\x3c\x33\x6f\xaa\x8d\x5b\x29\x5b\xab\x68\x6a\xf9\x47\x6c\xbc\xa2\x75\x57\xb5\x6b\xde\x5b\x91\xa8\x9a\xcc\xc6\x7e\xa9\xbe\x38\xd3\xb0\x49\x5c\xfc\x09\x90\x02\x10\xd9\xa5\x69\xc3\x54\x63\x47\xee\x96\xb6\x34\xd7\xc5\xd5\x7f\xef\x37\x8e\xd2\x89\x39\x8c\xaa\xb1\x5d\x45\xd1\xde\xa9\xd9\xa4\xe9\x64\x4d\x93\x44\xae\x34\xb4\x47\xbd\xd3\xd6\x10\x9b\xd1\xfa\x7e\x49\x2e\xb2\x16\x69\xfd\xa2\x03\x6b\x2c\xbe\x3a\x71\x95\xcb\x66\x2b\x61\xf5\x73\x8d\xa1\x40\x33\xf5\x0c\x15\x68\x5a\xa8\xfb\x0e\x96\xeb\xad\xab\xcb\x0b\x49\x74\x63\xc2\xba\x1c\x1a\xe1\x0f\xa7\x8c\x7d\xd5\xa6\x2c\x67\x97\x62\x0f\xd8\xfd\xe8\x9d\xf2\xde\xde\x33\x0f\xa0\xb8\xcb\xdd\xc5\xcf\x42\x2f\x09\xc3\x93\x56\x29\xd3\xc5\x33\x9d\x34\x23\xd8\xf0\x78\x16\x3f\x21\x19\xa1\x4d\xc7\x4d\xd3\xd3\xe1\xf1\xbb\x41\x8b\x9e\xf6\x8f\x4f\x7e\xeb\x0e\x07\xad\x6b\xea\x31\x9d\x69\x8a\x75\xed\x44\xec\x81\x3d\xbe\x25\xcc\x6c\x14\x7d\x31\xae\xaf\x6e\x4d\xf5\x6d\xbd\xbd\x77\xaf\xae\x2e\xaf\x2a\x6c\x9d\x84\xa1\x1f\x9a\x77\xde\xe5\xc3\x5b\xd9\xb8\x6a\xf9\xac\xdf\xff\xbb\x3f\x86\xdd\x0a\xdb\x1f\xdd\x33\x62\xde\xfd\x3b\x1c\xad\xbe\xf9\x62\x5f\x61\x5b\x39\x5b\xa2\x2a\x29\xf6\x00\x2b\x54\xc6\x9e\x76\x3e\x5a\x75\x7c\xa5\xf3\x4f\x22\x66\xf3\x61\xf0\x4e\x86\x00\x23\xd9\x8e\x81\x7d\x8d\xca\xcb\xc0\xbf\xa1\x81\xa4\x52\xc7\x98\xe2\x9a\x09\x64\x0c\xff\x67\xdd\x3c\x32\xe1\xbd\xd4\xf5\x0a\xac\x4f\x59\xd3\xae\x38\x2f\x51\x94\x63\xfc\x23\xd6\x9d\xe1\xbb\xa0\xb6\x36\xdb\x05\xf8\x16\xb8\x2e\x4c\xb4\x02\xd3\x13\x17\xb0\x2b\x9e\x0b\x0c\xe5\x58\xbe\xde\x21\x65\x38\xce\x29\xad\xcd\x70\x0e\x5d\x9b\xdf\x5b\xac\x96\xde\xec\xb7\x5a\x1a\x07\x0b\x73\xb9\x14\x0f\xee\xaa\x5e\x3a\xe9\x5f\x9b\x0b\x26\xc0\xae\x11\x39\x80\x94\x28\x8c\x34\x64\x7d\xd1\x95\x51\x22\xa9\x3a\xa5\x91\x10\xc3\xd7\x96\x63\xd9\x96\xa3\x56\x33\xf3\x0a\xf0\x5c\xab\x28\x69\x73\xa6\x32\x6a\xbb\xb4\x19\xeb\x28\xd8\xc5\xe7\x50\x48\xf1\x6d\x3e\xef\x4e\x24\x0a\xac\x9a\xd1\x57\x35\xf7\x38\xc8\x02\xaa\x9a\x01\x16\x20\x9f\x41\x6c\xdd\x73\x27\xd2\x59\xd3\x89\x74\x76\xdc\x89\xec\xac\xe9\x44\x3a\xe5\x3a\x91\x9d\x2a\x9d\xc8\x22\xad\x5f\x74\xbc\x75\x36\xe9\x44\x76\xbe\x76\x22\x3f\xf3\x4e\x64\x67\x6d\x27\x52\x67\x5f\xb5\x29\x33\x46\xd0\xce\x67\xd1\x89\xec\x3c\xc3\x4e\xa4\xb3\xbe\x9b\xd7\x11\xc5\xa7\x73\xda\xbb\xb8\xec\x94\x6f\x87\x39\x40\x98\xef\xac\xe9\x87\xf5\xf8\xf0\x26\xdd\x3c\xa7\x54\x37\x2f\xde\x7f\xa5\x6e\x9e\xf3\xb9\x75\xf3\x62\x7e\xa7\xb8\x63\x68\x6e\xe8\x24\x6a\xea\x6e\x74\xea\x75\x37\xf2\x28\xd6\xb6\x37\xca\x6a\x58\x2a\xf5\x52\xd4\xd6\xcc\xbf\x14\xf8\xc6\x1d\x25\x27\xd7\x51\x2a\xc1\xf4\xc7\x3b\x4a\x9b\xf2\xbc\x44\x47\xa9\xa4\x51\x64\x38\xbe\x49\x47\xc9\xd9\x7f\x47\x29\xfb\x2e\x9d\xfa\x51\xe6\x83\x99\x3a\x6f\xcb\xfd\x92\x7b\x5b\x6e\xcf\x7d\xab\x20\x32\xe7\xd6\x6a\x6c\x57\xb9\x75\x7f\x60\xce\xad\x75\xef\x12\x1f\xf5\x43\x7f\x4c\xa2\x48\xe7\x3f\x0b\x89\x74\x91\xb0\x2f\x3a\x91\x8e\x65\x55\x27\x91\xe6\x82\xa8\x97\x48\x8b\x97\xd4\x4d\x2d\x2b\x0b\x13\x2c\xdd\x3b\xf1\x1a\xf5\xde\x42\x92\xd4\x1f\x6c\x23\x49\x7a\xfe\x69\xbf\xb9\xa3\x66\x7a\x05\x5f\x2c\xba\xfb\x62\xc0\x05\xf3\xb4\x4c\x15\x81\x51\x13\xfb\x03\xcf\x67\x1d\x32\x0e\x89\xcd\x2b\x81\x75\x45\x41\x10\x21\x8e\x2d\x52\xd9\xac\x44\xe0\xdc\x0f\x49\x4d\x02\x93\xb2\x67\x3d\x7d\x88\x62\x77\x85\x4f\x7f\x10\xb8\x8b\x68\x08\x76\xb9\x9b\xb2\x67\x97\x2d\xba\x20\xaa\xe8\xde\x2a\x3a\xb6\x38\x67\x09\xea\xa6\x2b\x41\xf4\x35\x51\xd9\x4d\xa2\xb2\x9c\x9a\x13\x15\x35\xb6\xab\x44\xe5\xe6\xbd\x39\x51\xd1\x7d\x8a\x74\x74\x23\xbe\x42\xb2\xa6\xae\x3f\xc2\x8f\xe9\x99\xcd\x16\x65\x92\x96\x22\x91\x5f\x74\xd2\x12\xcb\xad\x4e\xd2\xc2\x85\xf2\x05\x24\x2d\x37\xef\xff\xdf\x93\x16\xd3\xd7\x7c\x62\xd1\x9d\x24\x2d\xcf\x2d\xb2\x2d\xab\xda\x40\x45\xed\x8f\x23\xdb\x72\x5a\x33\xb2\x2d\xa7\xcf\x20\xb2\xad\x3d\xd5\x43\x01\x2c\xed\xd0\x42\x76\x59\x87\x96\xcc\x7d\x1e\x1e\x7e\xcd\x1f\xaf\x21\x4f\xd5\xc0\x83\x36\xf8\x99\x2c\xab\x55\xfb\xfb\x17\xdf\xb7\xd1\xd2\x04\xc1\x9a\x67\x6f\x2e\x0d\x1a\x38\xe1\xe5\x89\xf5\xf6\xd0\xf2\x43\xab\x75\x72\x76\x3c\x18\x5c\x1c\x9f\x77\xad\x06\x97\x1e\x0c\x4b\x7d\x05\xcc\xbd\xce\x31\x63\x21\x04\x0e\x0b\x41\x1e\x1e\x36\x74\xab\x47\x4a\x31\x00\xbd\xc0\x76\x81\xc7\x2b\xa0\x34\x38\x4a\x9c\x20\x9c\x84\x30\x43\xd8\xc1\xa9\x1f\xca\x2d\xc4\x3b\x40\x1a\x70\xe8\xfc\xf7\x01\x9e\xac\x22\x86\x5f\x88\x34\x31\x9b\x14\xaa\xe5\xca\x70\x26\x58\x40\x12\x11\x71\xe6\x08\x3f\x77\x79\xd3\xbd\xba\xea\x75\xba\x78\x47\xb8\x2f\x7e\x0a\xca\x23\xbc\x5b\xad\x20\x6a\x4e\x89\xf5\x92\xfe\x60\xbd\x1c\x63\xa6\x0d\x73\xa5\xea\xf7\xaf\x5b\x67\x34\xc2\xfd\xb1\x70\xb5\xba\x25\xf7\x82\x0f\xc1\xa2\x09\xa8\x47\xf7\xcd\x8b\x86\x00\x69\x5d\x08\x7a\x0e\x98\xa3\x2c\x2a\xf1\xe7\x96\xe7\xdf\x85\xea\x6b\x20\x20\x30\x86\x38\x68\x33\xc7\x08\xd5\x78\x91\x32\x22\xc9\x49\x05\x7b\x1d\x91\xf0\x04\x01\x38\xd6\xcc\xb2\x38\x94\xb1\x95\xa3\xda\x58\x06\xf7\x91\x44\x62\xe5\xb1\xc0\xd0\x96\x90\xf4\x1c\x97\x18\x48\xc1\x21\x03\x96\x36\x0b\x53\xaa\x92\xfd\xa7\x8c\xe2\x08\xaf\x82\x62\xc5\x8b\x78\xc7\xf8\xe3\xb5\x15\x31\xb0\x59\x92\x78\x2c\x20\x2a\x75\x64\x12\x60\x56\x3e\x82\xcd\xd2\xde\x23\xfe\x85\x03\x8f\x68\x01\x9f\x86\xa2\xca\x1c\xc2\xb2\xf0\x68\x26\xf8\x1f\x7d\x9b\xf3\x79\x75\xb0\x80\xa8\x76\x8f\x04\x45\xb5\x09\x96\xb6\xe4\x2a\xbf\xab\xce\x20\x62\xc9\x21\x44\x8f\xbc\x9d\xc0\xfd\x80\x88\x11\x2c\xf1\xb1\x5c\x80\x55\x3c\x8a\x6a\xfb\x96\x76\x2b\xa7\xca\xad\x7c\x18\x9e\x9f\x9d\x5e\x5e\x59\x98\x06\x3d\xee\xad\x63\x8f\x03\xe5\xf8\x6d\xe2\x6d\x3a\xa7\x1c\xbd\xce\xe3\xa8\x76\xb6\xf0\x3a\x0e\x0d\xf9\x21\x37\x62\x81\x56\x87\x86\x17\xfc\x48\x9b\x9c\x01\xe6\xdc\xce\x1a\x0e\x72\xff\xfc\x1d\xfe\x63\x01\x4d\x80\xf0\xf5\x1b\xc4\xd9\x14\x48\x1d\xb2\xcc\xa0\x23\x4b\x44\xf7\x4a\xeb\xbf\x36\xc0\xf9\x9f\x34\x4a\x2d\x85\x7a\x94\x05\x3f\x03\xdc\xe7\x50\xc7\x4b\x9b\xba\xd5\x40\xc0\x24\x41\x55\xfe\xe9\x8d\xa2\xe0\xd7\x83\x68\x11\x24\x2e\xc9\x4c\xc9\x22\x22\x01\x09\xc7\x90\x23\x01\x3d\xc9\x0f\x24\x29\x5e\xb4\x2f\xee\x71\x02\xd4\xb2\xa5\x7c\xa6\xda\xd8\xd0\x67\x76\x9a\x96\x8d\x1c\x60\xf6\xf1\xc6\xe3\xc9\xc5\x26\x4e\x52\x1a\xd1\xcb\xb3\x44\xcf\xcf\xa8\x77\x8b\xe1\x23\x71\x2d\xf9\x43\xbd\x8e\x12\xc7\x61\xa9\x87\x9c\xb0\xc9\x33\x30\x91\xd3\x41\xeb\x03\xfc\xc4\x04\xc7\x4a\xc5\x12\x39\xa4\xe2\x95\x65\xa5\x7d\x4f\x87\x1f\x36\x27\xfc\x93\x06\xc6\x0e\x09\xd3\x02\x72\x46\xe7\x92\xce\xcc\x21\x51\x26\x57\x59\x81\x9e\xf3\xbe\x91\x1e\x18\xd2\xd3\x73\xee\x2f\x3c\x46\x9c\x7c\x3c\x55\x30\xbb\xa2\x47\x5b\xa2\x1e\xe5\xe8\x39\xbe\x39\xee\x9d\x29\x92\xb2\x7b\x13\x43\xd9\x60\xcf\x97\xe0\x86\x2a\xc5\xa3\x87\x49\x91\xb4\x67\x7a\xae\x07\xdd\x4e\x2c\xa1\xec\xde\xf8\x50\x2e\x43\xe2\x4b\xa0\x17\x51\xe4\x68\x61\xd2\x12\xda\x33\x3d\xc3\xcb\xe1\xb1\x41\x3e\x62\x48\x23\x1f\xee\x7c\xf4\xf2\x91\x30\xdb\x90\x4f\xdc\x6d\xd8\x20\x17\xc8\x44\xf0\x6d\x24\x04\xf2\xc1\xfb\x93\x65\x04\x02\xbf\x3e\x25\x90\xef\x14\xec\x39\x27\x90\x58\xf7\x9a\x14\x18\x70\xd6\xca\x0a\x7a\x93\x90\x90\x8a\x20\x8b\x5d\xa4\x05\xb8\xea\xc6\x79\x81\x78\x2d\x69\x7b\x89\x41\xe6\x45\x93\xcd\xeb\x23\x63\x40\x94\x21\xb9\x64\xf8\x54\x01\xaf\xba\xf7\x13\xa1\xa5\x3a\x1c\x77\xe1\xd5\xc1\x84\xa7\xdc\x56\x65\x93\x75\x3f\x9b\xbb\x33\xba\xef\xfa\x86\x4e\x12\x5f\xd6\x33\x57\x37\x34\x53\xdd\x48\x13\xa7\x93\x56\xc6\x83\xd5\x75\x59\x54\xd5\x14\xd9\x75\xcb\x7b\x0d\x00\xe8\x10\x97\xd9\x3d\xaf\x32\xc8\xe5\x82\x55\x81\xa9\x86\x21\xbb\xf8\x46\x76\x4f\x73\x05\xc1\xf6\xda\x22\x3d\xb0\xdb\x70\x62\x1b\x4d\xbd\xd8\x54\xb0\xe4\x7b\x62\xef\x7a\xc3\x81\x05\x4e\xd3\x8a\xf0\xd1\x4b\xfa\x78\xe0\x9e\xb7\xbe\xd5\x70\x90\x3b\xc6\x6c\x74\xd0\x4e\xdf\x39\xc2\x73\x01\x6b\x76\x3c\x1e\xdf\x1c\x48\xe5\xc9\x77\xc7\xd0\x0b\x59\xfc\x5d\x39\x6b\xee\x3b\x0b\xd7\xb7\x7e\x7c\xbf\x01\x03\xdf\xe5\xb6\xf8\xed\x8f\xef\x77\xbe\xc7\xca\x7c\xac\xba\xc9\x4d\x1c\x33\xdd\x6e\x9a\x19\x7f\x31\xfc\x44\x8e\x59\x60\xd7\x7b\x66\xf9\x29\xf4\xae\x5c\xb3\x5c\xfe\xab\x6f\x7e\x5c\x39\x9e\x93\x73\x36\xf8\x65\x95\x21\xaf\xb3\xdb\x2d\xf8\x37\x93\xe3\xdd\x13\x7a\xbd\x4b\x2d\x83\x7c\x1b\x8e\xd3\xe0\x2c\xb7\x81\x7e\x33\x97\x98\xf6\x61\xdb\xf0\x89\xea\xf0\x8a\x27\x72\x8a\x12\xbd\xde\x2b\xaa\x83\x39\x76\xe5\x16\xd5\xfa\x5f\xfd\x62\x09\x05\xf9\xea\x18\x9f\x89\x63\x14\x39\xdd\xd3\xba\x47\xf3\x1e\x9e\x85\x93\xcc\x38\xb5\xcd\xbd\xe4\x9c\xcc\xb7\xff\x0e\x04\x2c\x9a\x38\xc2\xf3\xee\xb9\xce\x05\xe2\xb1\x96\xc2\xf9\xdd\x52\xfc\x1b\x1e\x08\xd4\xfa\x8d\xe2\xb6\x73\xde\x2f\x35\x52\xd2\x4f\x20\xc0\x69\x95\x26\x1d\x02\x6c\xfb\xc9\x9d\x5c\x73\x93\x06\x1d\x2e\xb1\xd5\xe7\x76\xf1\x89\xa0\x7b\x7d\x71\x21\x65\x42\x28\x95\x2a\xf3\xd7\x74\xd1\xf6\xd0\x3e\x53\xc6\xb1\xb9\x99\x05\xd1\xde\x92\x90\x20\xf4\xc7\x89\xf5\xf5\x07\x7c\xaf\x3a\x0b\x0c\x54\xf6\x11\x50\xb4\x3f\x84\x6b\xf5\x7b\x1d\xd4\x01\x6f\xba\xe6\x6d\x24\xa9\x9f\xf8\x1c\x47\xc1\x94\x33\xb4\x18\xe2\x3a\x07\x21\x87\x53\xff\x55\xcf\x79\x82\x08\xec\x2f\xd4\xd2\x22\x7f\xe3\xdb\x32\x25\x13\xa0\x84\xb6\x90\xfa\x21\x65\xf7\x15\xc1\x2e\xe8\xb8\xa4\xf3\x89\x41\x06\xf4\xef\xaa\x20\x57\x24\xa2\x0e\xf1\xd6\x64\x5b\xe8\x79\x54\x1f\xfc\x48\x81\x0d\xe5\x5f\x1d\x59\xff\x94\xa5\x86\x18\x7e\xe4\x62\xc0\x8c\x73\x9d\x18\xf2\x79\xe8\x46\x0e\x4d\x7d\xb8\x50\xe7\x15\x84\x37\xbb\x78\x05\x41\xdf\xd5\xcf\x3e\xb3\xef\x0f\x80\x27\xda\x47\xc2\x62\x44\xf7\x44\x18\xee\xc7\x0f\x84\x75\x20\xfb\x7c\x20\x9c\xa7\xe6\xda\x48\xcd\xb5\x81\x9a\x6b\x33\x35\xd7\xbb\xa3\x46\xff\x3a\x45\x81\x9a\x41\xf7\x4a\xf7\x78\x5b\x8e\x68\x9e\x6e\xe3\x7d\xed\xd3\x7a\x05\xb2\xc7\x97\x0f\x0a\x9a\x76\xd5\x33\x69\x1a\x8c\x68\x35\x2d\xa1\x45\x0b\xf2\x94\x9a\x76\xd1\x3b\xe9\xea\x65\xc3\x47\x34\xb2\xb9\xe8\x19\xa9\x11\x20\x4f\x28\x9b\x9b\xde\xd5\x50\x4f\x0d\x1f\xd1\x50\x83\xf7\x0d\x9a\x26\x40\x9e\x90\x9a\xab\xee\xc0\xa0\x69\x38\xa2\xd3\x34\xb8\x6f\xf2\x02\x1c\x64\x97\x9a\x96\x8e\x8c\x62\x5e\x96\x9a\x61\xef\xdc\xa0\x69\x7c\x44\xf7\x8e\x0b\xdc\x37\xc8\x46\x80\xec\x44\x36\xe5\x7c\x1a\xc6\x42\x83\xdd\x1c\xeb\xa9\x39\xb9\x3c\x3f\x3f\xbe\xe8\x14\xdf\x10\x53\x20\xcf\xe5\x85\x1d\x99\x59\x6f\x9e\xa2\xa7\x6a\x39\x58\x4a\xd2\x2d\x58\x93\xf9\xdc\xe0\xc6\x76\x17\x90\x3f\xa9\xd7\x91\x2b\x60\x58\x4e\x79\x87\x93\x67\x23\xa1\x05\x69\x38\x7e\xd7\x22\x6e\xc5\x99\x98\xef\xe2\xaa\x87\x8d\x9f\xe2\xa4\x4d\xfe\x95\x26\xf1\x69\x1e\x3f\x12\xda\x98\x3c\x95\xd8\x41\x5c\x86\x6c\x5c\x67\xcc\xed\xf1\x2c\xc9\x87\xe4\xfe\xce\xe1\x26\x62\xd4\xd5\x1b\xf2\x6f\x9c\xc9\xa2\x63\xb1\xe0\x29\x22\xae\xd2\xba\xe6\x05\x81\xae\x18\xa8\x96\x80\xfe\x9c\x20\xc9\x2f\xff\xca\x5c\x70\x60\xa1\x8d\xd3\x2e\xec\x5c\x5e\xac\x9d\xd5\x0f\xfd\x25\x64\xdd\xe1\xba\x99\x55\xf6\xfc\x4b\xb2\x67\xfc\xde\x92\xa4\x37\x2d\x2f\x07\x78\x7f\x2d\x01\x55\x10\xfe\xf4\xaf\x04\xe3\x84\xba\x24\xb0\xd9\x4c\x83\x54\x0a\x14\x67\xfc\x89\x53\x1e\x4d\xd9\x4b\x28\x9f\x3c\x76\xa9\x44\x15\x6c\xfc\x96\x2c\x6d\x07\xce\x2b\xc3\x27\x88\x39\x7b\xb3\x32\x4b\xaf\x9b\x1f\xbb\x92\xba\xf5\x88\xfa\xdc\x2c\x2e\x38\xc4\xc7\x6b\xf2\x1b\xb4\x2d\x34\x5b\x68\xbe\x1d\x92\x0c\x61\xe1\x65\x1c\x0c\xa4\xda\x1a\x27\x70\xe5\x33\x8e\x3a\x34\x24\x63\xc6\xff\xde\xd7\x96\x7a\x2b\xb1\x36\x6c\xe0\xb9\xd5\xa7\xa6\xb0\x40\xea\xa3\x50\xfe\x59\x99\xfc\x2c\x34\x75\xdd\x8c\xe6\x96\x04\x10\x37\x04\xcf\x33\x7f\xc3\x38\x35\xdb\x76\x1c\xf5\x67\xf8\x8e\xe4\x27\xaf\x71\x64\xcb\x7d\x76\xcc\xbf\x5c\xc5\x4f\xd6\x43\xdf\x8d\x51\xf0\x19\xe9\x3f\xfe\xdd\xb0\x80\x13\x63\x32\xf3\x5d\x10\x83\x38\x99\x48\x7c\x62\x2a\xf5\x3d\x75\x7e\x92\xb8\x1f\x1f\x97\xd4\x27\x21\xf5\x51\xd9\x1b\xd6\x12\x83\x0f\x07\x16\x37\x63\xdd\x14\xc4\xf0\x6f\x0e\x15\xe7\xfe\x17\x00\x00\xff\xff\xfa\xc4\x6f\x10\xc0\x7d\x00\x00")

func index_html_bytes() ([]byte, error) {
	return bindata_read(
		_index_html,
		"index.html",
	)
}

func index_html() (*asset, error) {
	bytes, err := index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "index.html", size: 32192, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
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
	"index.html": index_html,
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

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"index.html": &_bintree_t{index_html, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

