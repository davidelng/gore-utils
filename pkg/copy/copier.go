package copy

import (
	"fmt"
	"io"
	"os"
)

// Align the struct in decreasing order to prevent
// unnecessary padding in memory
type Config struct {
	Source      string
	Dest        string
	BufferSize  int64
	Verbose     bool
	Interactive bool
	DryRun      bool
}

type Copier struct {
	cfg Config
}

func (c *Copier) Run() error {
	return nil
}

func copyBuffer(srcPath, dstPath string) error {
	// Let's put here all of our errors for now
	var err error

	// Get info about the file
	info, err := os.Stat(srcPath)
	if err != nil {
		return err
	}
	totalSize := info.Size()

	src, err := os.Open(srcPath)
	defer src.Close()
	if err != nil {
		return err
	}

	dst, err := os.Create(dstPath)
	defer dst.Close()
	if err != nil {
		return err
	}

	// 32KB buffer
	buf := make([]byte, 32*1024)
	var totalWritten int64

	for {
		// Read into buffer
		n, err := src.Read(buf)
		if n > 0 {
			// Copy only the length of byte read
			// so we don't risk copying old or garbage data
			if _, err := dst.Write(buf[:n]); err != nil {
				return err
			}

			totalWritten += int64(n)
			percent := float64(totalWritten) / float64(totalSize) * 100
			fmt.Fprintf(os.Stdout, "\rCopying: %.2f%% complete", percent)

		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(os.Stdout, "\nDone.\n")

	return nil
}
