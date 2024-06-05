package vdf_test

import (
	"golang.org/x/crypto/sha3"
	"source.quilibrium.com/quilibrium/cermonyclient/vdf"
	nekrovdf "source.quilibrium.com/quilibrium/monorepo/nekryptology/pkg/vdf"
	"testing"
)

const intSizeBits = uint16(2048)

func getChallenge(seed string) [32]byte {
	return sha3.Sum256([]byte(seed))
}

func TestProveVerify(t *testing.T) {
	difficulty := uint64(10000)
	challenge := getChallenge("TestProveVerify")
	solution := vdf.WesolowskiSolve(intSizeBits, challenge[:], difficulty)
	isOk := vdf.WesolowskiVerify(intSizeBits, challenge[:], difficulty, solution)
	if !isOk {
		t.Fatalf("Verification failed")
	}
}

func TestProveRustVerifyNekro(t *testing.T) {
	difficulty := 100
	challenge := getChallenge("TestProveRustVerifyNekro")

	for i := 0; i < 100; i++ {
		solution := vdf.WesolowskiSolve(intSizeBits, challenge[:], uint64(difficulty))
		nekroVdf := nekrovdf.New(uint32(difficulty), challenge)
		isOk := nekroVdf.Verify([516]byte(solution))
		if !isOk {
			t.Fatalf("Verification failed")
		}
		challenge = sha3.Sum256(solution)
	}
}

func TestProveNekroVerifyRust(t *testing.T) {
	difficulty := 100
	challenge := getChallenge("TestProveNekroVerifyRust")

	for i := 0; i < 100; i++ {
		nekroVdf := nekrovdf.New(uint32(difficulty), challenge)
		nekroVdf.Execute()
		proof := nekroVdf.GetOutput()
		isOk := vdf.WesolowskiVerify(intSizeBits, challenge[:], uint64(difficulty), proof[:])
		if !isOk {
			t.Fatalf("Verification failed")
		}
		challenge = sha3.Sum256(proof[:])
	}
}
