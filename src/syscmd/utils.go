package syscmd

import (
	"log"
	"os/exec"
	"path/filepath"
)

func GetAbs(rel string) string {
	if !filepath.IsAbs(rel) {
		rel, _ = filepath.Abs(rel)
	}

	return rel
}

func RunBashCmd(bashCmd string) []byte {
	out, err := exec.Command("bash", "-c", bashCmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func RunSysCmd(sysCmd string) []byte {
	out, err := exec.Command(sysCmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}
