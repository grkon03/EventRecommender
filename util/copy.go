package util

import (
	"bytes"
	"encoding/gob"
)

func DeepCopy(src any, dst any) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buf).Decode(dst)
}
