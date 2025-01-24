package format

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

// PNG generates a PNG file by converting DOT output using Graphviz.
func PNG(w io.Writer, cfg quickrisk.Config) error {
	// Create a temporary file to store the DOT content
	tmpDOT, err := os.CreateTemp("", "qr_*.dot")
	if err != nil {
		return fmt.Errorf("failed to create temp DOT file: %w", err)
	}
	defer os.Remove(tmpDOT.Name()) // Clean up temporary file

	// Write DOT content to the temporary file
	err = func() error {
		defer tmpDOT.Close()
		DOT(tmpDOT, cfg)
		return nil
	}()
	if err != nil {
		return fmt.Errorf("failed to write DOT content: %w", err)
	}

	// Run the Graphviz command to generate the PNG file
	cmd := exec.Command("dot", "-Tpng", tmpDOT.Name())
	cmd.Stdout = w         // Pipe output directly to the provided writer
	cmd.Stderr = os.Stderr // Redirect errors to standard error for debugging

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute Graphviz command: %w", err)
	}

	return nil
}
