// Code generated by go-bindata.
// sources:
// plugins/api/schema.graphql
// plugins/api/static/index.html
// DO NOT EDIT!

package assets

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

var _pluginsApiSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xb1\x6e\xc3\x30\x10\x43\x77\x7d\x05\x83\x2c\x09\xd0\x2f\xd0\x9c\x25\x1d\x5a\xb4\x49\xa7\x20\x83\xe2\xb0\xb6\x0a\x5b\x72\xef\xe4\x02\x46\xd1\x7f\x2f\x4e\xf6\x92\x49\x27\x8a\xe4\x13\x4e\x9b\xd0\x07\xc1\x39\x0e\x74\xeb\xfc\x7c\x7a\x7d\x71\x4e\x9b\x8e\x43\xc0\xaf\x03\xbe\x27\xca\xec\xf1\x66\x87\xfb\x73\x6e\x8b\x73\xc7\x45\x45\x99\x47\x3e\x41\x38\x0a\x95\xa9\x28\x42\xdf\x23\x7f\xa2\x74\x04\x53\x91\x19\x63\x8e\xa6\xc7\x54\x32\xf2\x24\xc8\xb7\x2f\x36\x05\xad\x84\xb1\x73\x16\x5f\x8a\x2b\x69\x8b\x77\x16\x89\xfc\x21\x34\xa6\xb6\x27\x26\xa5\xe0\x36\xe3\x78\x70\xa8\x97\x5d\xbc\x7b\x1c\x0f\x7b\x8f\x0f\xa5\x3c\x66\x8c\x6d\x1e\x5d\xbd\xba\xdb\x7b\x5c\xcc\x77\xdd\x2c\x1f\xaf\x99\x0a\xb5\xa9\x32\x97\xbe\x8d\x03\x38\x84\xd8\x7b\x9c\x8a\xc4\xd4\x9a\x30\x52\x86\xa8\x1a\x73\x52\x8f\xcb\xaa\x5f\xed\xa5\x11\x86\xc2\xbb\xaf\x8b\xb3\xee\xff\x00\x00\x00\xff\xff\x90\x47\x9c\xa7\x48\x01\x00\x00")

func pluginsApiSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsApiSchemaGraphql,
		"plugins/api/schema.graphql",
	)
}

func pluginsApiSchemaGraphql() (*asset, error) {
	bytes, err := pluginsApiSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/api/schema.graphql", size: 328, mode: os.FileMode(420), modTime: time.Unix(1527188991, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsApiStaticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xeb\x6f\xdb\x36\x10\xff\xde\xbf\xe2\xe6\x6c\x90\x5d\xd8\x94\xd3\xf5\x01\xa8\x76\x86\xb6\x49\xbb\x16\x69\xd3\xc6\x19\x8a\x7d\x2b\x4d\x9e\x2c\xa6\x14\xa9\x1e\x29\x3b\x6a\x91\xff\x7d\xa0\x64\xc9\x8a\x91\x0c\xd9\xb0\x41\x1f\x4c\xde\xe3\x77\x0f\xde\xc3\xb3\x9f\x8e\xcf\x5e\x5d\xfc\xf9\xf1\x04\x32\x9f\xeb\xa3\x07\xb3\xe6\x07\x60\x96\x21\x97\xe1\x00\x30\x73\xbe\xd2\xd8\x9c\x01\x96\x56\x56\xf0\x63\x7b\x01\xc8\x50\xad\x32\x9f\xc0\xe1\x74\xfa\xcb\xf3\x8e\x9a\x73\x5a\x29\x93\xc0\x74\x47\xda\x28\xe9\xb3\x7d\x39\xbb\x46\x4a\xb5\xdd\x24\x90\x29\x29\xd1\xb4\x9c\xeb\xed\xef\xc1\x8a\x78\x91\xa9\x6f\xfa\x76\x8b\xeb\x6c\x5f\x81\x5d\x6e\xfc\xc4\xdb\xaf\x68\x7a\x1a\x4b\x2e\xbe\xae\xc8\x96\x46\x26\xa0\x95\x41\x4e\x93\x15\x71\xa9\xd0\xf8\xe1\x41\xfa\x2c\x7c\x63\x38\xc0\x47\xe1\x1b\xed\x9c\x5b\x5a\x92\x48\x93\xa5\xf5\xde\xe6\x09\x1c\x16\x57\xe0\xac\x56\x12\x0e\xe4\x34\x7c\x3b\xc9\xd4\x1a\x3f\x49\x79\xae\x74\x95\x80\xab\x9c\xc7\x7c\x0c\x13\x5e\x14\x1a\x27\xed\x35\x5a\x70\x03\xaf\x89\x1b\xa1\x9c\xb0\xd1\x18\x22\xb6\x78\xfd\x61\x71\xac\x5c\xa1\x79\x35\x39\xc7\x55\xa9\x39\x05\xfa\x02\x57\x16\xe1\x8f\xb7\xd1\x18\xea\x63\x47\xfa\xfc\x31\xb0\x7f\x47\xbd\x46\xaf\x04\x87\x0f\x58\x62\x34\x86\xac\x25\x8c\x21\x3a\x2d\x85\x92\x1c\xde\x10\x37\x32\xf0\x38\x29\xae\xc7\xe0\xb8\x71\x13\x87\xa4\xd2\x9d\xd3\x05\x97\x52\x99\x55\x02\xcf\x8a\x2b\x38\x7c\x5c\x5c\xc1\xd3\xe2\x6a\x2f\x26\xa7\xbe\x63\x52\x33\x6f\x26\x7a\x16\xf7\x6a\x62\xa6\x95\xf9\x0a\x84\x7a\x3e\xa8\xa9\x2e\x43\xf4\x03\xc8\x08\xd3\xf9\x20\xf3\xbe\x70\x49\x1c\x0b\x69\x2e\x1d\x13\xda\x96\x32\xd5\x9c\x90\x09\x9b\xc7\xfc\x92\x5f\xc5\x5a\x2d\x5d\xdc\xbe\x73\x3c\x65\x87\x53\xf6\xa8\xbb\x33\xe1\xdc\x00\xe2\xdb\x0a\x31\x7e\x08\x67\x6b\x24\x52\x12\x1d\x3c\x8c\xdb\x02\x68\x35\x27\xc2\x1a\xcf\x95\x41\x02\xb6\x0e\x69\x58\x6a\x9c\xa0\x54\xde\xd2\x2d\xc5\xf4\xf4\xe9\xdf\x87\xe8\x04\xa9\xc2\x83\x23\x71\xef\x90\x52\xf4\x22\x8b\x1f\xb1\x29\xfb\xb5\x39\xb3\x5c\x19\x76\xe9\x06\x47\xb3\xb8\x81\xfb\xf7\xd8\x84\x5c\xf8\xf8\xf0\x09\x7b\xc2\x1e\x37\x97\xff\x15\x7c\x22\x6d\xfe\x1f\x1a\xb8\xf3\xb1\xf7\xe1\x67\x71\x3b\x85\x66\x61\xec\x6c\x2d\x4a\xb5\x06\xa1\xb9\x73\xf3\x41\xd7\xed\x83\xa3\x77\x9f\x2f\xe0\xa2\x6e\xfc\x99\x32\x45\xe9\x41\xc9\x3e\x1f\x0a\xcd\x05\x66\x56\x4b\xa4\xf9\x60\x27\xbc\xb2\xe8\x20\x43\xc2\x60\x59\xaa\x75\xcf\x46\x00\x68\x5d\x1b\x1c\x9d\x5a\x1e\xda\x85\x31\xd6\x97\xeb\xa7\x62\xcd\x09\x1c\x72\x12\x19\xcc\x61\xa3\x8c\xb4\x1b\xa6\xad\xe0\x5e\x59\xc3\x1a\xc6\xf3\x4e\xb0\xe0\xc4\x73\xf4\x48\x0e\xe6\xf0\xe3\xba\x61\x48\x2b\xca\x1c\x8d\x67\x2b\xf4\x27\x1a\xc3\xf1\x65\xf5\x56\x0e\xa3\x2e\x8e\x68\xc4\xd6\x5c\x97\x08\x73\x08\xd0\x7a\xe1\x2d\xf1\x15\x06\x85\xb7\x1e\xf3\x61\xd4\x3a\x9c\x5c\x6e\xfc\x45\xa3\xf1\xfc\x41\x0d\x9e\x96\x46\x04\x57\xa0\x16\xf9\x74\xfa\x3a\xd4\x24\xd2\x70\x7b\xfd\x18\x1c\x72\xa3\xae\x37\x84\x35\xce\x43\x8b\x02\xf3\x7f\xe0\x5c\xdb\x48\x1a\x3d\x84\x07\xdc\x06\xd9\x35\x5d\xf4\x42\x08\x2c\x7c\x94\x40\x14\x66\xa4\x6a\x52\x14\x5f\x3a\x6b\xa2\xf1\x4e\xea\x95\x35\x1e\x8d\x9f\x5c\x54\x05\xde\x2a\xdb\xf6\x6b\x6b\x4f\xa5\x30\x6c\x1d\x1e\xf5\xec\xdd\xc8\x94\xbb\x3b\x53\xe3\x2e\xdc\xde\x16\xb8\x2d\x80\xfb\x86\x70\x9f\x20\x6e\x48\xbf\x28\x7d\x66\x49\x7d\xaf\xf9\x51\xb2\x4b\xff\x6f\xf0\xe5\x25\x72\x42\x82\x9f\x7f\xb4\xc4\xeb\x2f\x90\x80\x29\xb5\xee\x10\xae\xf7\xd7\x21\xa1\x2f\xc9\x40\x3d\x7e\x86\x51\xfc\xad\x44\xaa\xa2\x71\x2f\x92\x1c\x7d\x66\x65\x02\x51\x61\x9d\xef\xf9\xb2\x8d\x7a\xdc\x5b\x86\xb2\x4a\xe0\xdd\xe2\xec\x03\x73\x9e\x94\x59\xa9\xb4\xda\x2b\x9d\x9d\xb0\x20\x94\x68\xbc\xe2\xda\x25\x10\x29\x23\x74\x19\xb6\x51\xeb\xdc\x88\xf9\x0c\xcd\xb0\x2b\xc8\x21\xa1\x2b\xac\x71\xd8\x7f\xb4\xad\xeb\x2d\x8b\x79\xbc\xf2\xc3\xee\x5d\xee\xc6\x78\x69\x65\xd5\xc7\xf1\x54\xdd\x78\xb9\x2d\x6e\x1d\x49\xc1\xc9\xe1\x4d\xcd\xdd\xc3\x5f\x83\xe0\x5e\x64\x30\x44\x22\x4b\xa3\xdb\x40\xfa\x9a\x3d\xc5\xce\xc7\x86\xd6\xdc\xe3\x18\xce\xd1\x48\x24\x98\xbd\xa9\x2b\xef\xd3\x29\xc4\x47\xa0\x8c\xb7\xe0\x33\xac\x13\xcc\x5a\xc9\x05\x62\x4d\x3c\x3f\x79\x71\xfc\xfe\x04\x94\xa9\x6f\xde\x16\xa0\x71\x8d\x1a\x6c\x0a\x3e\x53\x0e\x72\x2b\x4b\x1d\x18\xa0\x91\x93\x81\xdc\x12\x02\x5f\xda\xd2\xb7\x48\x99\xdd\x40\x65\x4b\x10\xdc\x80\x28\x9d\xb7\xb9\xfa\x8e\xd0\x79\xb0\xac\xa0\x20\xbb\x56\x61\xb4\x81\x54\x69\x8a\x84\xc6\x43\xdd\xc6\x0e\x2c\xb5\x30\xe1\xbf\x42\xc8\x33\xd7\x20\x32\xa5\x25\x60\x33\x01\x5c\xe3\xf2\x79\xd8\x12\xc7\x67\xef\x19\xd5\x21\x0e\xb7\x19\xa8\xc9\x4c\x10\x72\x8f\xdb\x91\x31\x6c\x4d\xf7\xab\x30\x6d\x46\x51\xb2\x37\x9a\x76\x05\xd3\x9e\xee\x9c\x41\x6d\x33\x47\xa3\x5a\x72\x9b\xf9\x9b\xbb\xa4\x59\x21\xb3\xb8\xf9\x8b\xfb\x57\x00\x00\x00\xff\xff\x45\xab\x31\x54\xfa\x0a\x00\x00")

func pluginsApiStaticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsApiStaticIndexHtml,
		"plugins/api/static/index.html",
	)
}

func pluginsApiStaticIndexHtml() (*asset, error) {
	bytes, err := pluginsApiStaticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/api/static/index.html", size: 2810, mode: os.FileMode(420), modTime: time.Unix(1527188991, 0)}
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
	"plugins/api/schema.graphql": pluginsApiSchemaGraphql,
	"plugins/api/static/index.html": pluginsApiStaticIndexHtml,
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
	"plugins": &bintree{nil, map[string]*bintree{
		"api": &bintree{nil, map[string]*bintree{
			"schema.graphql": &bintree{pluginsApiSchemaGraphql, map[string]*bintree{}},
			"static": &bintree{nil, map[string]*bintree{
				"index.html": &bintree{pluginsApiStaticIndexHtml, map[string]*bintree{}},
			}},
		}},
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

