package vdf_test

import (
	"source.quilibrium.com/quilibrium/cermonyclient/vdf"
	"testing"
)

const intSizeBits = uint16(2048)
const difficulty = uint64(10000)

func TestProveVerify(t *testing.T) {
	solution := vdf.WesolowskiSolve(intSizeBits, []byte{0x01, 0x02, 0x03}, difficulty)
	isOk := vdf.WesolowskiVerify(intSizeBits, []byte{0x01, 0x02, 0x03}, difficulty, solution)
	if !isOk {
		t.Errorf("WesolowskiVerify failed")
	}
}
