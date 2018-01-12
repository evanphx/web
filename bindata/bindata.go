// Code generated by go-bindata.
// sources:
// index.html
// public/.agignore
// public/d3.v355.min.js
// public/elm-beta.js
// public/elm-beta.min.js
// public/elm.js
// public/elm.min.js
// public/fly_success
// public/fonts/FontAwesome.otf
// public/fonts/fontawesome-webfont.eot
// public/fonts/fontawesome-webfont.svg
// public/fonts/fontawesome-webfont.ttf
// public/fonts/fontawesome-webfont.woff
// public/fonts/fontawesome-webfont.woff2
// public/fonts/iosevka-regular.ttf
// public/graph.js
// public/images/apple_logo_grey_ic.svg
// public/images/apple_logo_white_ic.svg
// public/images/buckleup.svg
// public/images/concourse_logo_white.svg
// public/images/favicon-144x144.png
// public/images/favicon-aborted.png
// public/images/favicon-errored.png
// public/images/favicon-failed.png
// public/images/favicon-paused.png
// public/images/favicon-pending.png
// public/images/favicon-started.png
// public/images/favicon-succeeded.png
// public/images/favicon.png
// public/images/favicon.svg
// public/images/ic_aborted_brown.svg
// public/images/ic_close_white_24px.svg
// public/images/ic_error_orange.svg
// public/images/ic_failing_red.svg
// public/images/ic_pause_blue.svg
// public/images/ic_pending_grey.svg
// public/images/ic_running_green.svg
// public/images/ic_running_legend.svg
// public/images/ic_search_white_24px.svg
// public/images/linxus_logo_grey_ic.svg
// public/images/linxus_logo_white_ic.svg
// public/images/parachute_error_404.svg
// public/images/windows_logo_grey_ic.svg
// public/images/windows_logo_white_ic.svg
// public/index.js
// public/jquery-2.1.1.min.js
// public/jquery.pin.min.js
// public/main.css
// public/resources.js
// DO NOT EDIT!

package bindata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

// indexHtml reads file data from disk. It returns an error on failure.
func indexHtml() (*asset, error) {
	path := filepath.Join(rootDir, "index.html")
	name := "index.html"
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

// publicAgignore reads file data from disk. It returns an error on failure.
func publicAgignore() (*asset, error) {
	path := filepath.Join(rootDir, "public/.agignore")
	name := "public/.agignore"
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

// publicD3V355MinJs reads file data from disk. It returns an error on failure.
func publicD3V355MinJs() (*asset, error) {
	path := filepath.Join(rootDir, "public/d3.v355.min.js")
	name := "public/d3.v355.min.js"
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

// publicElmBetaJs reads file data from disk. It returns an error on failure.
func publicElmBetaJs() (*asset, error) {
	path := filepath.Join(rootDir, "public/elm-beta.js")
	name := "public/elm-beta.js"
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

// publicElmBetaMinJs reads file data from disk. It returns an error on failure.
func publicElmBetaMinJs() (*asset, error) {
	path := filepath.Join(rootDir, "public/elm-beta.min.js")
	name := "public/elm-beta.min.js"
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

// publicElmJs reads file data from disk. It returns an error on failure.
func publicElmJs() (*asset, error) {
	path := filepath.Join(rootDir, "public/elm.js")
	name := "public/elm.js"
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

// publicElmMinJs reads file data from disk. It returns an error on failure.
func publicElmMinJs() (*asset, error) {
	path := filepath.Join(rootDir, "public/elm.min.js")
	name := "public/elm.min.js"
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

// publicFly_success reads file data from disk. It returns an error on failure.
func publicFly_success() (*asset, error) {
	path := filepath.Join(rootDir, "public/fly_success")
	name := "public/fly_success"
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

// publicFontsFontawesomeOtf reads file data from disk. It returns an error on failure.
func publicFontsFontawesomeOtf() (*asset, error) {
	path := filepath.Join(rootDir, "public/fonts/FontAwesome.otf")
	name := "public/fonts/FontAwesome.otf"
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

// publicFontsFontawesomeWebfontEot reads file data from disk. It returns an error on failure.
func publicFontsFontawesomeWebfontEot() (*asset, error) {
	path := filepath.Join(rootDir, "public/fonts/fontawesome-webfont.eot")
	name := "public/fonts/fontawesome-webfont.eot"
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

// publicFontsFontawesomeWebfontSvg reads file data from disk. It returns an error on failure.
func publicFontsFontawesomeWebfontSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/fonts/fontawesome-webfont.svg")
	name := "public/fonts/fontawesome-webfont.svg"
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

// publicFontsFontawesomeWebfontTtf reads file data from disk. It returns an error on failure.
func publicFontsFontawesomeWebfontTtf() (*asset, error) {
	path := filepath.Join(rootDir, "public/fonts/fontawesome-webfont.ttf")
	name := "public/fonts/fontawesome-webfont.ttf"
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

// publicFontsFontawesomeWebfontWoff reads file data from disk. It returns an error on failure.
func publicFontsFontawesomeWebfontWoff() (*asset, error) {
	path := filepath.Join(rootDir, "public/fonts/fontawesome-webfont.woff")
	name := "public/fonts/fontawesome-webfont.woff"
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

// publicFontsFontawesomeWebfontWoff2 reads file data from disk. It returns an error on failure.
func publicFontsFontawesomeWebfontWoff2() (*asset, error) {
	path := filepath.Join(rootDir, "public/fonts/fontawesome-webfont.woff2")
	name := "public/fonts/fontawesome-webfont.woff2"
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

// publicFontsIosevkaRegularTtf reads file data from disk. It returns an error on failure.
func publicFontsIosevkaRegularTtf() (*asset, error) {
	path := filepath.Join(rootDir, "public/fonts/iosevka-regular.ttf")
	name := "public/fonts/iosevka-regular.ttf"
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

// publicGraphJs reads file data from disk. It returns an error on failure.
func publicGraphJs() (*asset, error) {
	path := filepath.Join(rootDir, "public/graph.js")
	name := "public/graph.js"
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

// publicImagesApple_logo_grey_icSvg reads file data from disk. It returns an error on failure.
func publicImagesApple_logo_grey_icSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/apple_logo_grey_ic.svg")
	name := "public/images/apple_logo_grey_ic.svg"
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

// publicImagesApple_logo_white_icSvg reads file data from disk. It returns an error on failure.
func publicImagesApple_logo_white_icSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/apple_logo_white_ic.svg")
	name := "public/images/apple_logo_white_ic.svg"
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

// publicImagesBuckleupSvg reads file data from disk. It returns an error on failure.
func publicImagesBuckleupSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/buckleup.svg")
	name := "public/images/buckleup.svg"
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

// publicImagesConcourse_logo_whiteSvg reads file data from disk. It returns an error on failure.
func publicImagesConcourse_logo_whiteSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/concourse_logo_white.svg")
	name := "public/images/concourse_logo_white.svg"
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

// publicImagesFavicon144x144Png reads file data from disk. It returns an error on failure.
func publicImagesFavicon144x144Png() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/favicon-144x144.png")
	name := "public/images/favicon-144x144.png"
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

// publicImagesFaviconAbortedPng reads file data from disk. It returns an error on failure.
func publicImagesFaviconAbortedPng() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/favicon-aborted.png")
	name := "public/images/favicon-aborted.png"
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

// publicImagesFaviconErroredPng reads file data from disk. It returns an error on failure.
func publicImagesFaviconErroredPng() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/favicon-errored.png")
	name := "public/images/favicon-errored.png"
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

// publicImagesFaviconFailedPng reads file data from disk. It returns an error on failure.
func publicImagesFaviconFailedPng() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/favicon-failed.png")
	name := "public/images/favicon-failed.png"
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

// publicImagesFaviconPausedPng reads file data from disk. It returns an error on failure.
func publicImagesFaviconPausedPng() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/favicon-paused.png")
	name := "public/images/favicon-paused.png"
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

// publicImagesFaviconPendingPng reads file data from disk. It returns an error on failure.
func publicImagesFaviconPendingPng() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/favicon-pending.png")
	name := "public/images/favicon-pending.png"
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

// publicImagesFaviconStartedPng reads file data from disk. It returns an error on failure.
func publicImagesFaviconStartedPng() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/favicon-started.png")
	name := "public/images/favicon-started.png"
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

// publicImagesFaviconSucceededPng reads file data from disk. It returns an error on failure.
func publicImagesFaviconSucceededPng() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/favicon-succeeded.png")
	name := "public/images/favicon-succeeded.png"
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

// publicImagesFaviconPng reads file data from disk. It returns an error on failure.
func publicImagesFaviconPng() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/favicon.png")
	name := "public/images/favicon.png"
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

// publicImagesFaviconSvg reads file data from disk. It returns an error on failure.
func publicImagesFaviconSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/favicon.svg")
	name := "public/images/favicon.svg"
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

// publicImagesIc_aborted_brownSvg reads file data from disk. It returns an error on failure.
func publicImagesIc_aborted_brownSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/ic_aborted_brown.svg")
	name := "public/images/ic_aborted_brown.svg"
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

// publicImagesIc_close_white_24pxSvg reads file data from disk. It returns an error on failure.
func publicImagesIc_close_white_24pxSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/ic_close_white_24px.svg")
	name := "public/images/ic_close_white_24px.svg"
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

// publicImagesIc_error_orangeSvg reads file data from disk. It returns an error on failure.
func publicImagesIc_error_orangeSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/ic_error_orange.svg")
	name := "public/images/ic_error_orange.svg"
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

// publicImagesIc_failing_redSvg reads file data from disk. It returns an error on failure.
func publicImagesIc_failing_redSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/ic_failing_red.svg")
	name := "public/images/ic_failing_red.svg"
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

// publicImagesIc_pause_blueSvg reads file data from disk. It returns an error on failure.
func publicImagesIc_pause_blueSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/ic_pause_blue.svg")
	name := "public/images/ic_pause_blue.svg"
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

// publicImagesIc_pending_greySvg reads file data from disk. It returns an error on failure.
func publicImagesIc_pending_greySvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/ic_pending_grey.svg")
	name := "public/images/ic_pending_grey.svg"
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

// publicImagesIc_running_greenSvg reads file data from disk. It returns an error on failure.
func publicImagesIc_running_greenSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/ic_running_green.svg")
	name := "public/images/ic_running_green.svg"
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

// publicImagesIc_running_legendSvg reads file data from disk. It returns an error on failure.
func publicImagesIc_running_legendSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/ic_running_legend.svg")
	name := "public/images/ic_running_legend.svg"
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

// publicImagesIc_search_white_24pxSvg reads file data from disk. It returns an error on failure.
func publicImagesIc_search_white_24pxSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/ic_search_white_24px.svg")
	name := "public/images/ic_search_white_24px.svg"
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

// publicImagesLinxus_logo_grey_icSvg reads file data from disk. It returns an error on failure.
func publicImagesLinxus_logo_grey_icSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/linxus_logo_grey_ic.svg")
	name := "public/images/linxus_logo_grey_ic.svg"
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

// publicImagesLinxus_logo_white_icSvg reads file data from disk. It returns an error on failure.
func publicImagesLinxus_logo_white_icSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/linxus_logo_white_ic.svg")
	name := "public/images/linxus_logo_white_ic.svg"
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

// publicImagesParachute_error_404Svg reads file data from disk. It returns an error on failure.
func publicImagesParachute_error_404Svg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/parachute_error_404.svg")
	name := "public/images/parachute_error_404.svg"
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

// publicImagesWindows_logo_grey_icSvg reads file data from disk. It returns an error on failure.
func publicImagesWindows_logo_grey_icSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/windows_logo_grey_ic.svg")
	name := "public/images/windows_logo_grey_ic.svg"
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

// publicImagesWindows_logo_white_icSvg reads file data from disk. It returns an error on failure.
func publicImagesWindows_logo_white_icSvg() (*asset, error) {
	path := filepath.Join(rootDir, "public/images/windows_logo_white_ic.svg")
	name := "public/images/windows_logo_white_ic.svg"
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

// publicIndexJs reads file data from disk. It returns an error on failure.
func publicIndexJs() (*asset, error) {
	path := filepath.Join(rootDir, "public/index.js")
	name := "public/index.js"
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

// publicJquery211MinJs reads file data from disk. It returns an error on failure.
func publicJquery211MinJs() (*asset, error) {
	path := filepath.Join(rootDir, "public/jquery-2.1.1.min.js")
	name := "public/jquery-2.1.1.min.js"
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

// publicJqueryPinMinJs reads file data from disk. It returns an error on failure.
func publicJqueryPinMinJs() (*asset, error) {
	path := filepath.Join(rootDir, "public/jquery.pin.min.js")
	name := "public/jquery.pin.min.js"
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

// publicMainCss reads file data from disk. It returns an error on failure.
func publicMainCss() (*asset, error) {
	path := filepath.Join(rootDir, "public/main.css")
	name := "public/main.css"
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

// publicResourcesJs reads file data from disk. It returns an error on failure.
func publicResourcesJs() (*asset, error) {
	path := filepath.Join(rootDir, "public/resources.js")
	name := "public/resources.js"
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
	"index.html":                              indexHtml,
	"public/.agignore":                        publicAgignore,
	"public/d3.v355.min.js":                   publicD3V355MinJs,
	"public/elm-beta.js":                      publicElmBetaJs,
	"public/elm-beta.min.js":                  publicElmBetaMinJs,
	"public/elm.js":                           publicElmJs,
	"public/elm.min.js":                       publicElmMinJs,
	"public/fly_success":                      publicFly_success,
	"public/fonts/FontAwesome.otf":            publicFontsFontawesomeOtf,
	"public/fonts/fontawesome-webfont.eot":    publicFontsFontawesomeWebfontEot,
	"public/fonts/fontawesome-webfont.svg":    publicFontsFontawesomeWebfontSvg,
	"public/fonts/fontawesome-webfont.ttf":    publicFontsFontawesomeWebfontTtf,
	"public/fonts/fontawesome-webfont.woff":   publicFontsFontawesomeWebfontWoff,
	"public/fonts/fontawesome-webfont.woff2":  publicFontsFontawesomeWebfontWoff2,
	"public/fonts/iosevka-regular.ttf":        publicFontsIosevkaRegularTtf,
	"public/graph.js":                         publicGraphJs,
	"public/images/apple_logo_grey_ic.svg":    publicImagesApple_logo_grey_icSvg,
	"public/images/apple_logo_white_ic.svg":   publicImagesApple_logo_white_icSvg,
	"public/images/buckleup.svg":              publicImagesBuckleupSvg,
	"public/images/concourse_logo_white.svg":  publicImagesConcourse_logo_whiteSvg,
	"public/images/favicon-144x144.png":       publicImagesFavicon144x144Png,
	"public/images/favicon-aborted.png":       publicImagesFaviconAbortedPng,
	"public/images/favicon-errored.png":       publicImagesFaviconErroredPng,
	"public/images/favicon-failed.png":        publicImagesFaviconFailedPng,
	"public/images/favicon-paused.png":        publicImagesFaviconPausedPng,
	"public/images/favicon-pending.png":       publicImagesFaviconPendingPng,
	"public/images/favicon-started.png":       publicImagesFaviconStartedPng,
	"public/images/favicon-succeeded.png":     publicImagesFaviconSucceededPng,
	"public/images/favicon.png":               publicImagesFaviconPng,
	"public/images/favicon.svg":               publicImagesFaviconSvg,
	"public/images/ic_aborted_brown.svg":      publicImagesIc_aborted_brownSvg,
	"public/images/ic_close_white_24px.svg":   publicImagesIc_close_white_24pxSvg,
	"public/images/ic_error_orange.svg":       publicImagesIc_error_orangeSvg,
	"public/images/ic_failing_red.svg":        publicImagesIc_failing_redSvg,
	"public/images/ic_pause_blue.svg":         publicImagesIc_pause_blueSvg,
	"public/images/ic_pending_grey.svg":       publicImagesIc_pending_greySvg,
	"public/images/ic_running_green.svg":      publicImagesIc_running_greenSvg,
	"public/images/ic_running_legend.svg":     publicImagesIc_running_legendSvg,
	"public/images/ic_search_white_24px.svg":  publicImagesIc_search_white_24pxSvg,
	"public/images/linxus_logo_grey_ic.svg":   publicImagesLinxus_logo_grey_icSvg,
	"public/images/linxus_logo_white_ic.svg":  publicImagesLinxus_logo_white_icSvg,
	"public/images/parachute_error_404.svg":   publicImagesParachute_error_404Svg,
	"public/images/windows_logo_grey_ic.svg":  publicImagesWindows_logo_grey_icSvg,
	"public/images/windows_logo_white_ic.svg": publicImagesWindows_logo_white_icSvg,
	"public/index.js":                         publicIndexJs,
	"public/jquery-2.1.1.min.js":              publicJquery211MinJs,
	"public/jquery.pin.min.js":                publicJqueryPinMinJs,
	"public/main.css":                         publicMainCss,
	"public/resources.js":                     publicResourcesJs,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"public": &bintree{nil, map[string]*bintree{
		".agignore":       &bintree{publicAgignore, map[string]*bintree{}},
		"d3.v355.min.js":  &bintree{publicD3V355MinJs, map[string]*bintree{}},
		"elm-beta.js":     &bintree{publicElmBetaJs, map[string]*bintree{}},
		"elm-beta.min.js": &bintree{publicElmBetaMinJs, map[string]*bintree{}},
		"elm.js":          &bintree{publicElmJs, map[string]*bintree{}},
		"elm.min.js":      &bintree{publicElmMinJs, map[string]*bintree{}},
		"fly_success":     &bintree{publicFly_success, map[string]*bintree{}},
		"fonts": &bintree{nil, map[string]*bintree{
			"FontAwesome.otf":           &bintree{publicFontsFontawesomeOtf, map[string]*bintree{}},
			"fontawesome-webfont.eot":   &bintree{publicFontsFontawesomeWebfontEot, map[string]*bintree{}},
			"fontawesome-webfont.svg":   &bintree{publicFontsFontawesomeWebfontSvg, map[string]*bintree{}},
			"fontawesome-webfont.ttf":   &bintree{publicFontsFontawesomeWebfontTtf, map[string]*bintree{}},
			"fontawesome-webfont.woff":  &bintree{publicFontsFontawesomeWebfontWoff, map[string]*bintree{}},
			"fontawesome-webfont.woff2": &bintree{publicFontsFontawesomeWebfontWoff2, map[string]*bintree{}},
			"iosevka-regular.ttf":       &bintree{publicFontsIosevkaRegularTtf, map[string]*bintree{}},
		}},
		"graph.js": &bintree{publicGraphJs, map[string]*bintree{}},
		"images": &bintree{nil, map[string]*bintree{
			"apple_logo_grey_ic.svg":    &bintree{publicImagesApple_logo_grey_icSvg, map[string]*bintree{}},
			"apple_logo_white_ic.svg":   &bintree{publicImagesApple_logo_white_icSvg, map[string]*bintree{}},
			"buckleup.svg":              &bintree{publicImagesBuckleupSvg, map[string]*bintree{}},
			"concourse_logo_white.svg":  &bintree{publicImagesConcourse_logo_whiteSvg, map[string]*bintree{}},
			"favicon-144x144.png":       &bintree{publicImagesFavicon144x144Png, map[string]*bintree{}},
			"favicon-aborted.png":       &bintree{publicImagesFaviconAbortedPng, map[string]*bintree{}},
			"favicon-errored.png":       &bintree{publicImagesFaviconErroredPng, map[string]*bintree{}},
			"favicon-failed.png":        &bintree{publicImagesFaviconFailedPng, map[string]*bintree{}},
			"favicon-paused.png":        &bintree{publicImagesFaviconPausedPng, map[string]*bintree{}},
			"favicon-pending.png":       &bintree{publicImagesFaviconPendingPng, map[string]*bintree{}},
			"favicon-started.png":       &bintree{publicImagesFaviconStartedPng, map[string]*bintree{}},
			"favicon-succeeded.png":     &bintree{publicImagesFaviconSucceededPng, map[string]*bintree{}},
			"favicon.png":               &bintree{publicImagesFaviconPng, map[string]*bintree{}},
			"favicon.svg":               &bintree{publicImagesFaviconSvg, map[string]*bintree{}},
			"ic_aborted_brown.svg":      &bintree{publicImagesIc_aborted_brownSvg, map[string]*bintree{}},
			"ic_close_white_24px.svg":   &bintree{publicImagesIc_close_white_24pxSvg, map[string]*bintree{}},
			"ic_error_orange.svg":       &bintree{publicImagesIc_error_orangeSvg, map[string]*bintree{}},
			"ic_failing_red.svg":        &bintree{publicImagesIc_failing_redSvg, map[string]*bintree{}},
			"ic_pause_blue.svg":         &bintree{publicImagesIc_pause_blueSvg, map[string]*bintree{}},
			"ic_pending_grey.svg":       &bintree{publicImagesIc_pending_greySvg, map[string]*bintree{}},
			"ic_running_green.svg":      &bintree{publicImagesIc_running_greenSvg, map[string]*bintree{}},
			"ic_running_legend.svg":     &bintree{publicImagesIc_running_legendSvg, map[string]*bintree{}},
			"ic_search_white_24px.svg":  &bintree{publicImagesIc_search_white_24pxSvg, map[string]*bintree{}},
			"linxus_logo_grey_ic.svg":   &bintree{publicImagesLinxus_logo_grey_icSvg, map[string]*bintree{}},
			"linxus_logo_white_ic.svg":  &bintree{publicImagesLinxus_logo_white_icSvg, map[string]*bintree{}},
			"parachute_error_404.svg":   &bintree{publicImagesParachute_error_404Svg, map[string]*bintree{}},
			"windows_logo_grey_ic.svg":  &bintree{publicImagesWindows_logo_grey_icSvg, map[string]*bintree{}},
			"windows_logo_white_ic.svg": &bintree{publicImagesWindows_logo_white_icSvg, map[string]*bintree{}},
		}},
		"index.js":            &bintree{publicIndexJs, map[string]*bintree{}},
		"jquery-2.1.1.min.js": &bintree{publicJquery211MinJs, map[string]*bintree{}},
		"jquery.pin.min.js":   &bintree{publicJqueryPinMinJs, map[string]*bintree{}},
		"main.css":            &bintree{publicMainCss, map[string]*bintree{}},
		"resources.js":        &bintree{publicResourcesJs, map[string]*bintree{}},
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
