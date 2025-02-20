//go:build darwin
// +build darwin

// This file is compiled only on mac. It contains paths used by the mac
// browser bundle.
// http://golang.org/pkg/go/build/#hdr-Build_Constraints

package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

const (
	// During startup of meek-client-torbrowser, the browser profile is
	// created and maintained under a meek-specific directory by making a
	// recursive copy of everything under profileTemplatePath (see
	// https://bugs.torproject.org/18904).
	// If the TOR_BROWSER_TOR_DATA_DIR env var is set, the path for the
	// meek-specific profile directory is constructed by appending
	// torDataDirFirefoxProfilePath to TOR_BROWSER_TOR_DATA_DIR. Otherwise,
	// firefoxProfilePath (a relative path) is used.
	firefoxPath                  = "../firefox"
	torDataDirFirefoxProfilePath = "PluggableTransports/profile.meek-http-helper"
	firefoxProfilePath           = "../../../../TorBrowser-Data/Tor/PluggableTransports/profile.meek-http-helper"
	profileTemplatePath          = "../../Resources/TorBrowser/Tor/PluggableTransports/template-profile.meek-http-helper"
	helperNativeExecutablePath   = "../Tor/PluggableTransports/meek-http-helper"
)

func osSpecificCommandSetup(cmd *exec.Cmd) {
	// nothing
}

func installHelperNativeManifest() error {
	var homeDir string
	torDataDir := os.Getenv("TOR_BROWSER_TOR_DATA_DIR")
	if torDataDir != "" {
		homeDir = filepath.Join(torDataDir, "..", "Browser")
	} else {
		homeDir = "../../../../TorBrowser-Data/Browser"
	}
	// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Native_manifests#Mac_OS_X
	return writeNativeManifestToFile(filepath.Join(homeDir, "Mozilla", "NativeMessagingHosts"), helperNativeExecutablePath)
}

func uninstallHelperNativeManifest() error {
	// Nothing to do here: the host manifest file is written inside the
	// browser directory, so we assume we don't have to clean it up.
	return nil
}
