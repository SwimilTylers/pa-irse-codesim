package syscmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

func Dump(from string) []byte {
	if !filepath.IsAbs(from) {
		from, _ = filepath.Abs(from)
	}

	r := exec.Command("/bin/bash", "-c", fmt.Sprintf("/bin/bash src/syscmd/llvm-dump.sh %s", from))
	f, err := r.StdoutPipe()
	if err != nil {
		log.Fatalln("Cannot get stdout:", err)
	}

	var bs []byte

	if err = r.Start(); err == nil {
		bs, err = ioutil.ReadAll(f)
		if err != nil {
			log.Fatalln("Dump Read Error:", err)
		}
	} else {
		log.Fatalln("Dump Start Error:", err)
	}

	if err = r.Wait(); err != nil {
		log.Fatalln("Dump Wait Error:", err)
	}

	return bs
}
