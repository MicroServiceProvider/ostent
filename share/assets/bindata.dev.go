// +build !bin

package assets

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// cssIndexCss reads file data from disk. It returns an error on failure.
func cssIndexCss() (*asset, error) {
	path := filepath.Join(rootDir, "css/index.css")
	name := "css/index.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// cssIndexCssMap reads file data from disk. It returns an error on failure.
func cssIndexCssMap() (*asset, error) {
	path := filepath.Join(rootDir, "css/index.css.map")
	name := "css/index.css.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// faviconPng reads file data from disk. It returns an error on failure.
func faviconPng() (*asset, error) {
	path := filepath.Join(rootDir, "favicon.png")
	name := "favicon.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// jsRequirejs2117RequireMinJs reads file data from disk. It returns an error on failure.
func jsRequirejs2117RequireMinJs() (*asset, error) {
	path := filepath.Join(rootDir, "js/requirejs/2.1.17/require.min.js")
	name := "js/requirejs/2.1.17/require.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// jsSrcGenJscriptJs reads file data from disk. It returns an error on failure.
func jsSrcGenJscriptJs() (*asset, error) {
	path := filepath.Join(rootDir, "js/src/gen/jscript.js")
	name := "js/src/gen/jscript.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// jsSrcMilkBuildJs reads file data from disk. It returns an error on failure.
func jsSrcMilkBuildJs() (*asset, error) {
	path := filepath.Join(rootDir, "js/src/milk/build.js")
	name := "js/src/milk/build.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// jsSrcMilkIndexJs reads file data from disk. It returns an error on failure.
func jsSrcMilkIndexJs() (*asset, error) {
	path := filepath.Join(rootDir, "js/src/milk/index.js")
	name := "js/src/milk/index.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// jsSrcVendorBootstrap334CollapseBootstrapMinJs reads file data from disk. It returns an error on failure.
func jsSrcVendorBootstrap334CollapseBootstrapMinJs() (*asset, error) {
	path := filepath.Join(rootDir, "js/src/vendor/bootstrap/3.3.4-collapse/bootstrap.min.js")
	name := "js/src/vendor/bootstrap/3.3.4-collapse/bootstrap.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// jsSrcVendorHeadroom070HeadroomMinJs reads file data from disk. It returns an error on failure.
func jsSrcVendorHeadroom070HeadroomMinJs() (*asset, error) {
	path := filepath.Join(rootDir, "js/src/vendor/headroom/0.7.0/headroom.min.js")
	name := "js/src/vendor/headroom/0.7.0/headroom.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// jsSrcVendorJquery214Jquery214MinJs reads file data from disk. It returns an error on failure.
func jsSrcVendorJquery214Jquery214MinJs() (*asset, error) {
	path := filepath.Join(rootDir, "js/src/vendor/jquery/2.1.4/jquery-2.1.4.min.js")
	name := "js/src/vendor/jquery/2.1.4/jquery-2.1.4.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// jsSrcVendorReact0133ReactMinJs reads file data from disk. It returns an error on failure.
func jsSrcVendorReact0133ReactMinJs() (*asset, error) {
	path := filepath.Join(rootDir, "js/src/vendor/react/0.13.3/react.min.js")
	name := "js/src/vendor/react/0.13.3/react.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// jsSrcVendorRequirejsDomready201DomreadyJs reads file data from disk. It returns an error on failure.
func jsSrcVendorRequirejsDomready201DomreadyJs() (*asset, error) {
	path := filepath.Join(rootDir, "js/src/vendor/requirejs-domready/2.0.1/domReady.js")
	name := "js/src/vendor/requirejs-domready/2.0.1/domReady.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// robotsTxt reads file data from disk. It returns an error on failure.
func robotsTxt() (*asset, error) {
	path := filepath.Join(rootDir, "robots.txt")
	name := "robots.txt"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
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
	"css/index.css": cssIndexCss,
	"css/index.css.map": cssIndexCssMap,
	"favicon.png": faviconPng,
	"js/requirejs/2.1.17/require.min.js": jsRequirejs2117RequireMinJs,
	"js/src/gen/jscript.js": jsSrcGenJscriptJs,
	"js/src/milk/build.js": jsSrcMilkBuildJs,
	"js/src/milk/index.js": jsSrcMilkIndexJs,
	"js/src/vendor/bootstrap/3.3.4-collapse/bootstrap.min.js": jsSrcVendorBootstrap334CollapseBootstrapMinJs,
	"js/src/vendor/headroom/0.7.0/headroom.min.js": jsSrcVendorHeadroom070HeadroomMinJs,
	"js/src/vendor/jquery/2.1.4/jquery-2.1.4.min.js": jsSrcVendorJquery214Jquery214MinJs,
	"js/src/vendor/react/0.13.3/react.min.js": jsSrcVendorReact0133ReactMinJs,
	"js/src/vendor/requirejs-domready/2.0.1/domReady.js": jsSrcVendorRequirejsDomready201DomreadyJs,
	"robots.txt": robotsTxt,
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
	"css": &bintree{nil, map[string]*bintree{
		"index.css": &bintree{cssIndexCss, map[string]*bintree{
		}},
		"index.css.map": &bintree{cssIndexCssMap, map[string]*bintree{
		}},
	}},
	"favicon.png": &bintree{faviconPng, map[string]*bintree{
	}},
	"js": &bintree{nil, map[string]*bintree{
		"requirejs": &bintree{nil, map[string]*bintree{
			"2.1.17": &bintree{nil, map[string]*bintree{
				"require.min.js": &bintree{jsRequirejs2117RequireMinJs, map[string]*bintree{
				}},
			}},
		}},
		"src": &bintree{nil, map[string]*bintree{
			"gen": &bintree{nil, map[string]*bintree{
				"jscript.js": &bintree{jsSrcGenJscriptJs, map[string]*bintree{
				}},
			}},
			"milk": &bintree{nil, map[string]*bintree{
				"build.js": &bintree{jsSrcMilkBuildJs, map[string]*bintree{
				}},
				"index.js": &bintree{jsSrcMilkIndexJs, map[string]*bintree{
				}},
			}},
			"vendor": &bintree{nil, map[string]*bintree{
				"bootstrap": &bintree{nil, map[string]*bintree{
					"3.3.4-collapse": &bintree{nil, map[string]*bintree{
						"bootstrap.min.js": &bintree{jsSrcVendorBootstrap334CollapseBootstrapMinJs, map[string]*bintree{
						}},
					}},
				}},
				"headroom": &bintree{nil, map[string]*bintree{
					"0.7.0": &bintree{nil, map[string]*bintree{
						"headroom.min.js": &bintree{jsSrcVendorHeadroom070HeadroomMinJs, map[string]*bintree{
						}},
					}},
				}},
				"jquery": &bintree{nil, map[string]*bintree{
					"2.1.4": &bintree{nil, map[string]*bintree{
						"jquery-2.1.4.min.js": &bintree{jsSrcVendorJquery214Jquery214MinJs, map[string]*bintree{
						}},
					}},
				}},
				"react": &bintree{nil, map[string]*bintree{
					"0.13.3": &bintree{nil, map[string]*bintree{
						"react.min.js": &bintree{jsSrcVendorReact0133ReactMinJs, map[string]*bintree{
						}},
					}},
				}},
				"requirejs-domready": &bintree{nil, map[string]*bintree{
					"2.0.1": &bintree{nil, map[string]*bintree{
						"domReady.js": &bintree{jsSrcVendorRequirejsDomready201DomreadyJs, map[string]*bintree{
						}},
					}},
				}},
			}},
		}},
	}},
	"robots.txt": &bintree{robotsTxt, map[string]*bintree{
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, path.Join(name, child))
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
