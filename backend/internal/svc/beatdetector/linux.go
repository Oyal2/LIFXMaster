//go:build linux
// +build linux

package beatdetector

import (
	"context"

	"github.com/oyal2/LIFXMaster/pkg/connection"
)

func RunAudio(ctx context.Context, c *connection.LIFXClient, targets ...uint64) {
}
