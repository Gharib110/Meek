//go:build linux
// +build linux

// This file is compiled only on linux. It contains paths used by the linux
// browser bundle.
// http://golang.org/pkg/go/build/#hdr-Build_Constraints

package main

import (
	"os/exec"
	"syscall"
)

const (
	firefoxPath                  = "./firefox"
	firefoxProfilePath           = "TorBrowser/Data/Browser/profile.meek-http-helper"
	torDataDirFirefoxProfilePath = ""
	profileTemplatePath          = ""
	// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Native_manifests#Linux
	helperNativeManifestDir    = "TorBrowser/Data/Browser/.mozilla/native-messaging-hosts"
	helperNativeExecutablePath = "TorBrowser/Tor/PluggableTransports/meek-http-helper"
)

func osSpecificCommandSetup(cmd *exec.Cmd) {
	// Extra insurance against stray child processes: send SIGTERM when this
	// process terminates. Only works on Linux.
	cmd.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGTERM}
}

func installHelperNativeManifest() error {
	return writeNativeManifestToFile(helperNativeManifestDir, helperNativeExecutablePath)
}

func uninstallHelperNativeManifest() error {
	// Nothing to do here: the host manifest file is written inside the
	// browser directory, so we assume we don't have to clean it up.
	return nil
}
