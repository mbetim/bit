package utils

import (
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
	var shellCmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		shellCmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		shellCmd = exec.Command("open", url)
	case "linux":
		shellCmd = exec.Command("xdg-open", url)
	default:
		return
	}

	shellCmd.Start()
}
