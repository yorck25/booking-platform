package common

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func Download(wg *sync.WaitGroup, url string) {
	if wg != nil {
		defer wg.Done()
	}

	cmd := exec.Command("curl", "-fsSL", "-O", url)
	if err := cmd.Run(); err != nil {
		fmt.Printf("failed to download %s: %v\n", url, err)
	}
}

func DownloadSingle(url string) {
	Download(nil, url)
}

func GetURLBody(url string) (string, error) {
	cmd := exec.Command("curl", "-fsSL", url)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("curl failed: %v: %s", err, stderr.String())
	}

	return strings.TrimSpace(out.String()), nil
}
