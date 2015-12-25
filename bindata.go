// Code generated by go-bindata.
// sources:
// bashenv/bash.bash
// bashenv/cmd.bash
// bashenv/fn.bash
// bashenv/plugn.bash
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	name    string
	size    int64
	mode    os.FileMode
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

var _bashenvBashBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x91\x5f\x6b\xdb\x30\x14\xc5\x9f\x97\x4f\x71\x22\x02\xde\x60\x89\x61\xe4\x29\x23\x81\x84\x6d\x64\x2f\xf1\x58\xc6\x68\x09\xa1\x28\xf2\x75\x2c\x2a\x4b\x45\x57\x26\xcd\xb7\xaf\xe4\x36\xc5\x4d\xfb\x64\xa3\x7b\xce\xef\xdc\x3f\x03\x5d\x61\xb7\x83\x18\xad\x96\xdb\xf5\xdd\xff\x9f\x7f\xb7\xbf\x37\xbf\x0a\x81\xb1\x09\x10\x53\x81\xfd\xfe\x3b\x42\x4d\x76\xf0\x89\x54\xed\x20\x86\x43\xdc\xba\xd6\x83\xcf\x1c\xa8\xc1\x4a\x72\x0d\xcd\x70\x6d\x80\xab\x50\xca\x40\x33\xf4\x58\xc5\x46\xf4\x9c\x7f\x0c\x49\x26\xb4\x0f\x47\x2f\x4b\x42\x70\xcf\xfe\x29\x9c\xc7\xd1\x53\x34\xfb\x49\xd4\x5f\x5a\xfa\xdc\x5a\xd9\xd0\x17\x81\xf9\x1c\xe2\x87\xf4\x27\x6d\xfb\x0d\x75\xdc\x97\x0f\x44\x61\x51\x6c\x71\xf3\x15\x6d\x4c\x58\xbb\x86\x0e\x9e\x4e\x29\x42\x5b\x0e\xd2\x18\x98\x88\xe7\xd0\x25\xce\xc4\xc5\x96\x01\x18\xa1\x93\x5e\x74\x87\x28\xc8\xae\xe9\xff\x62\x22\x64\x59\x42\x87\xc4\xcc\x29\xa8\x9c\x6b\x32\x86\x21\x6d\x09\x55\x4b\x7b\x24\x9c\xd3\x66\x62\xbe\xcf\x18\x5d\xf5\x3a\x88\xdb\xd2\xc5\x0a\xc6\x0a\xa2\x7b\xce\x5b\xf6\xb9\x71\x4a\x9a\xfc\xa0\x6d\x9e\xb2\xb1\x58\xf4\xf9\x22\x7b\x8b\x50\x75\xf2\xf3\x47\xce\x77\x5d\x2f\x61\xd3\x0e\xc8\x37\xda\x4a\x03\x26\x66\xed\x6c\xba\x57\xda\x60\x2c\xaa\xf8\x22\xfd\x39\x8d\x14\xe4\x3d\x81\xaa\x8a\x54\x98\x88\x57\x50\xa5\xe3\xf9\x1e\xe3\xd0\xdf\x06\xf1\xf7\x29\x00\x00\xff\xff\xbe\x15\xe5\x0d\x2d\x02\x00\x00")

func bashenvBashBashBytes() ([]byte, error) {
	return bindataRead(
		_bashenvBashBash,
		"bashenv/bash.bash",
	)
}

func bashenvBashBash() (*asset, error) {
	bytes, err := bashenvBashBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bashenv/bash.bash", size: 557, mode: os.FileMode(420), modTime: time.Unix(1450979371, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bashenvCmdBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x41\x73\x9b\x3c\x10\x3d\xc3\xaf\xd8\x4f\x51\x32\xc9\x81\xe1\xc3\xa7\x0e\x1e\x77\x9c\x69\x7b\x6b\x7b\xc9\xd1\x38\x33\x2a\x88\xc0\x58\x16\x1e\x84\xdd\x74\x30\xff\xbd\xbb\x42\x60\x88\x7d\x6a\x2e\x46\xbb\xab\xdd\xf7\x9e\xde\xc6\xcf\x64\xaa\x44\x2d\x21\x78\x86\x2f\x3f\xbe\xbe\xf8\x7e\xba\xcf\x02\x55\x9a\xe6\xf1\x09\x5a\xdf\x1b\xd2\x99\x34\xe9\x8a\x7d\xc7\xb8\x01\x71\x12\xa5\x12\xbf\x94\x84\xb4\xda\xef\x85\xce\x0c\xbb\x14\x6a\xb3\x62\x3c\xc2\xc0\xd0\x27\xd8\xc9\x3f\x06\x18\xd7\x86\xc1\x19\x8c\xcc\x80\x99\x10\x4f\x71\x18\x32\xbf\xbb\xcc\xb3\x75\xf3\xa1\x63\xaf\xbc\xaa\x61\x07\xa5\xc6\x36\xed\x7f\x04\x73\xb3\xde\x76\x6c\x09\x59\xe5\x7b\x9e\x4c\x8b\x0a\x13\x3b\x02\x51\x69\x89\x43\xde\x6a\x79\x00\xf6\x4a\x43\xec\xcc\xaa\x6e\x66\x93\xb4\x9b\xf3\x2f\x6d\x83\x13\xc4\x1f\x7a\xca\xf7\x03\x9e\x6e\xe9\xf5\xcd\x66\x50\x31\xc8\x8f\x3a\x6d\xca\x4a\x83\xa0\x93\xd3\x6d\x22\x5b\xae\x2d\x55\x4c\xe3\x6f\xbb\x88\x03\x1e\x75\x98\x56\x55\x2a\x94\xd5\xc1\xa9\xa0\x09\x2e\x7f\x9c\x50\x79\x9a\xe3\xcd\x35\x9b\x2a\x10\x30\x78\xf8\x0c\x61\x26\x4f\xa1\x3e\x2a\x05\x0f\x0f\xbd\xaa\xda\xd1\xf2\x3d\xcb\x9b\x9e\x27\xe6\xad\x30\xe1\x1d\x7e\x05\x61\xc7\xb6\x2b\xdb\x6b\xce\x71\x54\xee\xc3\x0b\x39\xba\x7c\x81\x5d\xe5\x09\x11\x63\x90\x0a\x01\x71\xcd\x14\x49\x18\xa7\xdf\x04\xeb\xac\x41\xb4\x01\x1e\x41\xc2\x12\xbe\x4e\x50\x78\xdf\xeb\x9c\x73\xfa\x79\xd0\x3f\xbf\x83\x18\x59\x50\xd1\x88\x69\x00\x73\x11\x09\x93\x4b\x30\x45\x99\x37\x30\x84\xb1\x70\x16\x3f\x9f\xa1\xa9\x8f\x72\x48\x9b\x46\x34\x47\xb3\xfa\xdf\xf7\xca\x1c\x06\x55\x47\xb7\x5a\x19\x5f\x39\xc6\x13\x3e\xd3\x71\x09\x4d\x21\x35\x92\xe0\xed\x44\x3f\xac\x63\xdb\x0e\x6f\xaf\x49\x07\x65\x70\x0a\xb5\xdd\x6c\x30\x44\x39\xd8\x6e\xc7\x8b\xee\xbd\x7e\x56\x60\x8e\x69\x31\x38\x22\x06\x5b\x48\x79\x87\x6c\x41\x4f\xab\x86\x2e\x04\xeb\xba\x09\x7f\xcc\x75\x40\xba\xf6\x15\x4f\x74\x3f\x2f\x9d\x27\x46\x6b\x3c\x5f\xad\x6d\x6c\x0b\xd1\x55\x38\x73\xee\x2b\xd7\xc8\x59\xcb\x3b\xd4\xa5\x6e\x72\x60\x00\xf7\xc1\xe2\x93\x81\x7b\x93\xa0\xcd\x1c\xa9\xd9\xf8\x6b\x39\x7a\x38\xbd\xd7\x46\x40\xef\x65\x03\xbc\x67\xe8\x13\x54\xf7\xa2\x85\x54\x87\x5b\x7b\xf4\x52\x54\xbf\x0d\x50\x16\x61\x22\xe0\xbd\xb0\xbb\x44\xd0\x6f\x2d\x93\xa8\xdf\xc8\x0c\xf4\x0a\x83\x70\x14\x9a\x48\x07\xf8\x37\x30\x27\x92\xd6\x13\xf4\x81\x8b\x7e\x07\x4a\xa0\x02\x78\xa3\x2f\x9b\xd8\xab\xa5\x36\xe1\x3d\x31\x0b\xe9\x1f\x45\x1f\xd4\x18\x82\xf9\x19\xc2\x00\x9d\x3c\xb9\x6e\xf7\xfb\x86\x38\x0e\x89\x0e\x88\x97\x5b\xdf\x68\x34\x8f\x5b\x11\xbb\xfc\x24\xd2\x64\x33\x06\xb9\xac\x2a\xfe\xdf\x00\x00\x00\xff\xff\x4c\x2e\xf9\x29\xc7\x05\x00\x00")

func bashenvCmdBashBytes() ([]byte, error) {
	return bindataRead(
		_bashenvCmdBash,
		"bashenv/cmd.bash",
	)
}

func bashenvCmdBash() (*asset, error) {
	bytes, err := bashenvCmdBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bashenv/cmd.bash", size: 1479, mode: os.FileMode(420), modTime: time.Unix(1450380776, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bashenvFnBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x51\x41\x6e\xc2\x30\x10\x3c\xc7\xaf\x18\xad\x22\x01\xaa\xa2\x08\xae\x34\x3d\x56\xea\x1b\x28\x07\xcb\xac\x89\xd5\xd4\x89\x6c\x03\xaa\x28\x7f\xef\x1a\xd2\x92\x43\x55\x55\xb9\x64\x77\x66\x67\x76\xd6\xca\xfa\x4a\x87\x7d\x9c\x2f\x70\x56\xc5\x8e\x4d\xa7\x03\x63\xc7\xd1\x34\xf4\xe2\xe3\xc0\x26\x41\xc3\x1e\xbc\x49\xae\xf7\xb3\x08\x21\x1f\xde\xd9\xa7\x48\xaa\xe8\x7a\xa3\xbb\xdc\xe9\x9c\xe7\xa6\x9c\xa7\x8f\x81\x51\x2e\xf1\x89\x7d\xe0\x01\xdf\x6a\x63\x59\x1d\x41\x53\x03\x12\xa0\x65\xbd\x43\xb5\x5c\xa8\x82\x4d\xdb\xa3\x62\x50\x79\x1e\x05\xeb\x1a\x35\xbd\x7a\xba\x64\xa2\x3e\xbd\xa1\x7a\x6e\x30\xab\x9b\xfa\x3c\x04\xe7\x13\xe8\x91\xca\x25\x3d\xd1\x65\x26\x78\x0a\xc8\x5c\xc8\xa7\x2e\x2a\xa7\xca\x16\xff\x4e\x95\xa1\xe0\x86\x5c\x51\x1e\xc8\x44\xf9\xe1\xa3\xe4\xa3\x5f\x82\x45\x33\x59\x9e\xd6\xb8\x6e\x5f\xe6\xfe\xe8\xee\xbc\xed\xff\x70\x8f\x13\x7b\xba\x73\xac\x6f\x72\x26\xc4\xb6\x3f\xc5\xfe\x10\x0c\x4b\xbd\xa2\xf1\x3a\x54\x5a\x8f\x72\x3e\xbe\x18\xa4\x5a\xfc\x40\xb8\x01\xd7\xcd\x26\x80\x2a\x9c\xc5\x66\x23\xa3\x77\x49\xc2\x76\xbb\x46\x6a\xd9\xab\xa2\xb8\x25\x13\x5d\x39\xa1\x76\x1d\x2a\x8f\x87\x95\xf4\x6f\xc3\xd6\x49\x9c\xaf\x00\x00\x00\xff\xff\x5c\xaf\x12\xe0\x23\x02\x00\x00")

func bashenvFnBashBytes() ([]byte, error) {
	return bindataRead(
		_bashenvFnBash,
		"bashenv/fn.bash",
	)
}

func bashenvFnBash() (*asset, error) {
	bytes, err := bashenvFnBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bashenv/fn.bash", size: 547, mode: os.FileMode(420), modTime: time.Unix(1450380776, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bashenvPlugnBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x57\xdd\x4f\xdc\x38\x10\x7f\xde\xfd\x2b\xe6\xa2\x1c\x82\x93\xc2\x16\x4e\xf7\x52\xae\xd5\xa1\x2b\xa2\x48\x88\x43\x7c\xdc\x4b\xaf\x42\x26\x71\x36\x16\x8e\x1d\xc5\xce\x02\x2d\xfd\xdf\x6f\xfc\x91\xac\xb3\x78\x17\xda\xbe\x40\x98\x19\x8f\x7f\x9e\x8f\xdf\x0c\xd3\x05\x6d\x15\x93\x62\x7b\x07\xbe\x4e\x27\x05\xcd\x39\x69\x29\x14\x54\xe5\xef\x92\xcb\x4a\xde\x83\xd7\x27\xd3\x09\xcd\x2b\x09\x49\xc3\xbb\xb9\x78\x0b\xe9\xd7\xf3\xd3\xeb\xe3\xb3\x9b\x7f\x8f\x2e\x2e\x4f\xfe\x39\x7b\x9b\x15\x74\xf1\x2d\x99\x7e\x9b\x4e\x99\x50\x9a\x70\x1e\x73\x78\xe2\x54\x40\x40\xd0\x7b\x30\x9e\x98\x80\xb2\x95\x35\x4a\x8e\x99\x86\xeb\x8b\xd3\x64\x79\xa6\x6b\xf9\xbb\x24\xdd\x4b\x40\x90\x9a\xe2\xd7\x3e\xea\xf2\x02\x92\xd4\xdc\x7c\x72\x76\x73\x7e\x78\xf5\x71\x46\x16\x84\x71\x72\xcb\x29\x2a\xe7\xe8\x22\xe7\x52\x50\xb4\xc1\xc3\x09\xa4\xfb\xf6\x44\x06\xef\x61\x86\xf8\x66\xa2\xe3\xdc\x40\xec\xc4\x06\x90\x17\xb4\x96\x0b\x3a\x46\xd7\x5f\xe2\xa5\x2a\x40\xe9\x24\x16\xe8\x74\xd2\xd6\x90\xb5\xe5\x0a\x44\x2a\xcc\xd1\x62\x96\x3a\xd3\x75\x66\xc3\x25\x4b\x43\x03\xb5\x29\x88\xa6\x31\x9c\xd7\x56\xd3\xe3\x24\xa2\x00\xd9\x68\xcc\x14\x3e\xeb\x11\x1a\x14\x69\x09\xb9\xac\x6b\xa6\x67\x9a\xcc\x67\xb7\x2d\x11\x79\x15\x07\xee\xed\x34\x53\x95\x8f\xf3\xa7\x4f\xf0\x0b\x64\x6b\x83\x3d\x40\x84\xcf\x9f\x61\x6b\x0b\x5c\x69\x9c\x3b\x28\xdb\x5e\xb9\x03\x42\x6a\xf0\x91\xa6\x45\x62\x0d\x1f\x30\x47\x7b\x9b\xf2\x18\x84\x09\x41\x64\x5f\xd0\x6e\x89\xae\xbf\xcf\xc2\x4b\xb7\x4d\xc2\xd5\x63\x7d\x2b\x39\xcb\xb3\x96\x96\xf0\xf1\xe8\xf0\xc3\x4e\x0c\x13\xc6\x43\xd0\xc2\x84\x24\xdd\xfe\x13\x76\x67\xbb\xee\x92\x9b\xa5\xeb\x9d\x25\xbe\x37\xbe\x94\x2a\x9a\xdf\xc9\x4e\x43\x4d\x94\xa6\x2d\x6c\x85\x55\x64\x2d\x1a\xfc\x8a\x88\x87\x83\xd9\x8a\x92\x95\x06\x79\x26\x9e\xbd\xe9\x00\x74\x45\xc5\x74\x62\x8f\x97\x54\xe7\x15\x64\x19\x66\x4d\xad\x38\x18\xbb\x0f\x9c\xac\xda\xb9\xc7\x8f\x6e\x79\x1f\x7d\xf6\x60\xfb\x2c\x79\xae\xf2\x0a\x5b\x59\x41\xf8\x02\x97\x48\x09\x5c\xd1\x17\x3d\xa0\x5d\xc9\xe2\x8d\xc8\x99\xd2\xb1\xda\x3e\x45\x39\x18\xa6\xe0\x32\x27\x3c\xe8\x3a\x55\x61\x91\x43\xa6\xc0\x78\x98\x73\x79\x8b\xbe\x65\x0b\x0d\xd1\x15\x56\x1a\xac\x29\xa9\xdf\x0e\xa0\x90\x88\x33\xf4\x86\x85\xbe\x7d\x4b\x14\x35\xe4\x02\xa9\x71\xb0\x93\x0c\x26\x58\xb2\xba\x53\xc6\xc4\xe4\x8b\xc2\xa6\x7e\x0e\xcb\xcd\xab\xe0\xe9\xc9\xfd\x5d\x30\x65\x05\x81\x6b\xcf\xa7\xc6\x77\x4e\xb4\xbb\x79\xe6\x3c\xed\x6a\x59\x73\x78\x02\xf3\x2b\x9b\x53\xed\xc8\xd6\x74\x59\xd2\xb3\x70\xe0\xc8\x45\xea\xbb\xbc\x98\x23\x2d\xb3\x2c\xe1\x3c\x35\x2d\x13\x1a\x89\x08\xe0\xd7\x6c\xff\x8d\xc2\x9f\x7f\x98\x1f\x7b\xe6\x53\xfd\x67\x8e\x0c\x8d\x9e\xa4\x3d\x06\xfc\x74\xf1\x31\x5f\xc6\xa5\x21\x15\x64\xdd\x21\x3b\xdd\x32\x3b\x98\x63\xdd\xb2\xf9\x9c\xb6\xb1\x34\x5f\x39\x95\x82\x4a\xca\x3b\x93\xbf\x3e\x7e\xcf\x69\xd6\x58\x58\xae\x3a\x00\x55\xb1\x52\xaf\x2f\x05\x97\x94\xd5\x62\xe8\x93\xd6\x97\x02\x5d\x60\x08\x4d\xf4\xa4\x28\xd9\x3c\xa3\x0f\x8d\x6c\x31\x8e\x61\x49\xb8\x2a\xc6\x38\x01\x4c\x2c\x17\x3d\x0c\xe1\x98\xa5\x06\x4f\xcf\x46\x23\x21\xda\xfc\xb5\x21\x20\x13\x1f\x8e\x6c\x8e\xad\x71\x4f\x1e\x21\x38\x84\xc1\x72\x30\x63\xb1\x3a\xb2\x1a\x1c\x94\x03\x3d\x46\x27\x50\x7d\x57\xb0\x16\xb2\x26\x3e\x83\xd0\x80\x0b\xc8\x4a\xf5\x32\xb1\xbf\x38\xc3\x28\x97\xa4\xe8\x9f\x61\xb0\xfb\x62\x8f\x81\xff\xe0\x54\x3f\x8f\x1e\x27\xe7\xf7\x03\xf3\x39\xc6\x5e\x88\x61\x3b\xc6\x16\xf1\x45\xe3\x0c\xbb\x96\x68\xb7\xf1\xc4\x46\x65\xb8\x8c\x60\xe7\x8d\xd1\x38\x07\xb6\x03\x93\x51\x0b\x06\x61\x35\x0e\x92\x00\x96\x2b\xbd\x68\xca\x5d\x51\xbe\x1e\xdc\x6b\x21\xf9\x6a\x4f\xc2\x45\xc3\xa3\x51\xf1\x20\x5d\xfe\x60\x90\x00\xdb\xac\x33\x9f\xbf\xf7\x0b\xa4\x9d\x7e\x56\x3a\xc0\x51\x74\x23\xe8\x68\xf0\x98\x60\x51\xa0\x27\x28\x67\x84\xb3\x2f\x58\x6c\xc8\x27\x75\xa3\x1f\x7b\xe0\x86\x23\xd7\xd6\x18\x2a\xb4\xec\x70\xee\xae\x07\xf2\x8a\xf2\x5c\x63\x10\x2e\xa9\x6e\x56\x2e\x61\x16\x1e\xa4\x99\x9e\x22\x84\xba\x4a\x61\xf6\xd9\x37\x4a\x76\x6d\x1e\xed\xb1\x4b\xab\xb1\xab\x1c\x8e\x6c\x05\x86\x0c\xad\xf5\xca\xe2\xfa\x33\xbc\xd9\x33\x61\x19\x30\xe1\x73\xca\xdc\x5d\x2e\x6a\xee\xb5\x0e\x34\xbc\x70\x64\xe3\x14\xa9\x09\xf3\xff\xa3\x98\x6a\xc9\xa8\xc4\x9d\xa4\xa1\x25\xc6\xf5\xc0\xec\x54\x49\x7a\x75\x71\xf8\xf7\x51\x7f\xad\xb5\x79\x98\x62\x3f\xd4\x45\x5f\xed\x7e\x7c\x8d\x64\x7e\x39\x1d\xc9\x86\x7f\x0e\xc6\x52\xbb\xcb\x8c\x44\x66\x79\x19\x09\x3c\xb5\x8f\x64\x2e\x7e\x23\x91\xe7\xc9\x91\xcc\xe7\xb5\x0f\x55\x12\x2a\x33\xa1\x7c\xd3\x0d\x3b\xd6\x6a\x0f\x06\x9e\x96\x64\x17\x13\xbb\xbf\x62\x1a\xb5\x72\xc0\x34\x98\x0f\x20\xde\x9f\x24\xfd\x88\xfa\x3f\x00\x00\xff\xff\x0a\xb0\xa1\x84\x31\x0e\x00\x00")

func bashenvPlugnBashBytes() ([]byte, error) {
	return bindataRead(
		_bashenvPlugnBash,
		"bashenv/plugn.bash",
	)
}

func bashenvPlugnBash() (*asset, error) {
	bytes, err := bashenvPlugnBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bashenv/plugn.bash", size: 3633, mode: os.FileMode(420), modTime: time.Unix(1450979821, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"bashenv/bash.bash": bashenvBashBash,
	"bashenv/cmd.bash": bashenvCmdBash,
	"bashenv/fn.bash": bashenvFnBash,
	"bashenv/plugn.bash": bashenvPlugnBash,
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"bashenv": &bintree{nil, map[string]*bintree{
		"bash.bash": &bintree{bashenvBashBash, map[string]*bintree{}},
		"cmd.bash": &bintree{bashenvCmdBash, map[string]*bintree{}},
		"fn.bash": &bintree{bashenvFnBash, map[string]*bintree{}},
		"plugn.bash": &bintree{bashenvPlugnBash, map[string]*bintree{}},
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

