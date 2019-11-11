package fingerprint

import (
	"fmt"
)

type Fingerprint struct {
	Hash  uint
	Where int
}

func (f *Fingerprint) ToString() string {
	return fmt.Sprintf("[%d,%d]", f.Hash, f.Where)
}
