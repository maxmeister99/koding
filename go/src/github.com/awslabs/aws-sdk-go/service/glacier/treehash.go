package glacier

import (
	"crypto/sha256"
	"io"
)

const bufsize = 1024 * 1024

type Hash struct {
	TreeHash   []byte
	LinearHash []byte
}

func ComputeHashes(r io.ReadSeeker) Hash {
	r.Seek(0, 0)       // Read the whole stream
	defer r.Seek(0, 0) // Rewind stream at end

	buf := make([]byte, bufsize)
	hashes := [][]byte{}
	hsh := sha256.New()

	for {
		// Build leaf nodes in 1MB chunks
		n, err := io.ReadAtLeast(r, buf, bufsize)
		if n == 0 {
			break
		}

		tmpHash := sha256.Sum256(buf[:n])
		hashes = append(hashes, tmpHash[:])
		hsh.Write(buf[:n]) // Track linear hash while we're at it

		if err != nil {
			break // This is the last chunk
		}
	}

	return Hash{
		LinearHash: hsh.Sum(nil),
		TreeHash:   buildHashTree(hashes),
	}
}

func buildHashTree(hashes [][]byte) []byte {
	if hashes == nil || len(hashes) == 0 {
		return nil
	}

	for len(hashes) > 1 {
		tmpHashes := [][]byte{}

		for i := 0; i < len(hashes); i += 2 {
			if i+1 <= len(hashes)-1 {
				tmpHash := append(append([]byte{}, hashes[i]...), hashes[i+1]...)
				tmpSum := sha256.Sum256(tmpHash)
				tmpHashes = append(tmpHashes, tmpSum[:])
			} else {
				tmpHashes = append(tmpHashes, hashes[i])
			}
		}

		hashes = tmpHashes
	}

	return hashes[0]
}
