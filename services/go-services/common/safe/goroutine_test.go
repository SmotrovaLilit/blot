package safe

import (
	"context"
	"testing"
)

func TestGoContext(t *testing.T) {
	ctx := context.Background()
	GoContext(ctx, func() {
		panic("BOOM")
	})
}
