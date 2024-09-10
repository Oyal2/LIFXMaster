//go:build darwin
// +build darwin

package beatdetector

import (
	"context"

	"github.com/oyal2/LIFXMaster/pkg/connection"
)

func RunAudio(ctx context.Context, c *connection.LIFXClient, targets ...uint64) {
}
