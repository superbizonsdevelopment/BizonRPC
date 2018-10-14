package bizonrpc

import (
	"os"
	"runtime"
)

func GetSystem() string {
	if runtime.GOOS == "darwin" {
		return "unix"
	}
	if runtime.GOOS == "unix" {
		return "unix"
	}
	if runtime.GOOS == "windows" {
		return "windows"
	}
	return ""
}

func GetTempPath() string {
	temp := os.Getenv("XDG_RUNTIME_DIR")
	if temp != "" {
		return temp
	}

	temp = os.Getenv("TMPDIR")
	if temp != "" {
		return temp
	}

	temp = os.Getenv("TMP")
	if temp != "" {
		return temp
	}

	temp = os.Getenv("TEMP")
	if temp != "" {
		return temp
	}

	return "/tmp"
}
