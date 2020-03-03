// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/layouts/layout.html
// templates/layouts/mylayout.html
// templates/page1.html
// templates/partials/page1_partial1.html
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesLayoutsLayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xce\xc1\xa9\xc3\x30\x0c\x06\xe0\xf3\x33\x78\x07\xbd\x01\x8c\xc9\x5d\x78\x82\x9e\x4a\x17\x70\x6a\x51\x19\x94\xa4\x38\xca\xc1\x84\xec\x5e\xec\xba\x27\x49\xf0\x89\xff\x47\xd6\x45\x82\x35\xc8\x14\x53\x9b\x9a\x55\x28\xdc\x62\xdd\x0e\x45\xff\xbd\xac\xb1\x06\xfd\x4f\xcc\x5b\xaa\xc1\x9a\x3f\xe4\x29\x3c\x38\xef\x90\x77\x50\x26\x78\xc9\x36\x47\x01\x19\xaf\x3c\x75\x34\x17\xf0\x7d\xf9\x77\x0e\xee\xb4\x26\x2a\x5d\x3f\x8f\x52\x68\x55\x50\x5a\xde\x12\x95\x80\xa9\x10\x38\xd7\xec\x79\x42\xcd\x24\x09\xae\xab\x05\x8f\x40\xf4\xa3\xeb\x27\x00\x00\xff\xff\x68\xca\x16\xc2\xb4\x00\x00\x00")

func templatesLayoutsLayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLayoutsLayoutHtml,
		"templates/layouts/layout.html",
	)
}

func templatesLayoutsLayoutHtml() (*asset, error) {
	bytes, err := templatesLayoutsLayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/layouts/layout.html", size: 180, mode: os.FileMode(438), modTime: time.Unix(1565946441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLayoutsMylayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8f\x4d\x6a\xc5\x30\x0c\x84\xd7\x35\xf8\x0e\xd3\x03\x18\x93\xbd\xf1\x09\xba\x2a\xbd\x80\x53\xab\xc8\xe0\x9f\xe2\x28\x0b\x13\x72\xf7\x47\x9c\xbc\x95\x46\x62\x46\x7c\xe3\x58\x4a\xf6\x5a\x39\xa6\x10\xaf\x29\x49\x32\xf9\x32\xf0\x15\x46\xdb\xc5\xd9\xfb\xa0\x95\x56\xce\xbe\x4d\x6b\x8b\xc3\x6b\xf5\xe1\x78\xf1\x3f\x9c\x36\xa4\x0d\xc2\x84\x3c\x33\xf8\x6b\x7d\xae\xb6\x0c\x8b\x50\xe3\x14\x4d\x98\x3a\x7a\xdb\x85\x36\xb4\x9a\x87\xb3\xbc\xcc\x27\x6b\x87\x9d\xe2\xd3\x18\x7c\x53\x8d\x74\xc7\x7f\xf7\xde\xa9\x0a\x84\xca\x7f\x0e\x42\x60\xea\x04\x63\x2e\xef\x71\x60\x24\xca\x11\xe7\x79\x81\x3d\x40\xce\x3e\x75\x5e\x01\x00\x00\xff\xff\x64\xea\xc5\x1d\xd7\x00\x00\x00")

func templatesLayoutsMylayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLayoutsMylayoutHtml,
		"templates/layouts/mylayout.html",
	)
}

func templatesLayoutsMylayoutHtml() (*asset, error) {
	bytes, err := templatesLayoutsMylayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/layouts/mylayout.html", size: 215, mode: os.FileMode(438), modTime: time.Unix(1565946441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPage1Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x41\xaa\xc2\x30\x10\x00\xd0\xf5\x2f\xf4\x0e\xc3\xec\xbf\x25\x5b\x8d\x3d\x83\x37\x90\x69\x33\xa4\xa1\x63\x53\x26\x69\x40\x42\xee\x2e\xa2\xb8\x7c\xf0\xac\x0b\x05\x52\x7e\x0a\x5f\x71\xa2\x79\xf5\x1a\x8f\xcd\xfd\xcf\x51\xa2\x9e\x61\x12\x9a\xd7\x0b\xfc\x74\x30\x8e\x7d\xd7\x77\x7f\x76\x31\xe3\x8d\x3c\x83\x81\x5a\xc1\x2b\x73\x06\x0c\x1a\x12\x38\x2e\x2c\x71\x67\xc5\xd6\xec\xb0\x98\xcf\xaf\x15\x94\x37\xc7\x0a\xb8\x93\xe6\x40\x92\x86\x9d\x3c\x9b\xfb\x97\xe6\xb4\xe4\x87\x60\x6b\xef\x6e\x07\x17\xca\xd8\x77\xaf\x00\x00\x00\xff\xff\x47\x41\x4a\x5c\x9d\x00\x00\x00")

func templatesPage1HtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesPage1Html,
		"templates/page1.html",
	)
}

func templatesPage1Html() (*asset, error) {
	bytes, err := templatesPage1HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/page1.html", size: 157, mode: os.FileMode(438), modTime: time.Unix(1565946441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsPage1_partial1Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\xc9\x2c\x53\x28\x2e\xa9\xcc\x49\xb5\x55\x4a\x4a\x4c\xce\x4e\x2f\xca\x2f\xcd\x4b\xd1\x4d\xce\xcf\xc9\x2f\xb2\x52\x28\xcf\xc8\x2c\x49\xb5\x56\x80\xf2\x8a\x52\x53\x94\xec\x78\xb9\x38\x6d\x32\x0c\xed\x02\x12\xd3\x53\x15\x0c\xd5\x8b\x15\x02\x12\x8b\x4a\x32\x13\x73\x14\x0c\x6d\xf4\x33\x0c\xed\x78\xb9\x6c\xf4\x53\x32\xcb\xec\x78\xb9\x00\x01\x00\x00\xff\xff\xa2\xa6\x60\xb6\x59\x00\x00\x00")

func templatesPartialsPage1_partial1HtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsPage1_partial1Html,
		"templates/partials/page1_partial1.html",
	)
}

func templatesPartialsPage1_partial1Html() (*asset, error) {
	bytes, err := templatesPartialsPage1_partial1HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/page1_partial1.html", size: 89, mode: os.FileMode(438), modTime: time.Unix(1565946441, 0)}
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
	"templates/layouts/layout.html":          templatesLayoutsLayoutHtml,
	"templates/layouts/mylayout.html":        templatesLayoutsMylayoutHtml,
	"templates/page1.html":                   templatesPage1Html,
	"templates/partials/page1_partial1.html": templatesPartialsPage1_partial1Html,
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
	"templates": {nil, map[string]*bintree{
		"layouts": {nil, map[string]*bintree{
			"layout.html":   {templatesLayoutsLayoutHtml, map[string]*bintree{}},
			"mylayout.html": {templatesLayoutsMylayoutHtml, map[string]*bintree{}},
		}},
		"page1.html": {templatesPage1Html, map[string]*bintree{}},
		"partials": {nil, map[string]*bintree{
			"page1_partial1.html": {templatesPartialsPage1_partial1Html, map[string]*bintree{}},
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
