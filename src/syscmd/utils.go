package syscmd

import (
	"fmt"
	"io/ioutil"
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

func RunBash(bash string, variable string) []byte {
	bash = GetAbs(bash)

	v := fmt.Sprintf("%s %s", bash, variable)

	return RunSysCmd("/bin/bash", v)
}

func RunSysCmd(syscmd string, variables string) []byte {
	r := exec.Command("/bin/bash", "-c", fmt.Sprintf("%s %s", syscmd, variables))
	f, err := r.StdoutPipe()
	if err != nil {
		log.Fatalln("Cannot get stdout:", err)
	}

	var bs []byte

	if err = r.Start(); err == nil {
		bs, err = ioutil.ReadAll(f)
		if err != nil {
			log.Fatalln("RunSysCmd Read Error:", err)
		}
	} else {
		log.Fatalln("RunSysCmd Start Error:", err)
	}

	if err = r.Wait(); err != nil {
		log.Fatalln("RunSysCmd Wait Error:", err)
	}

	return bs
}
