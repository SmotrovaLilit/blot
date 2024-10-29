package gameset

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStatus_CanPlayCard(t *testing.T) {
	tests := []struct {
		name   string
		status Status
		want   bool
	}{
		{
			name:   "should return true when status is playing",
			status: StatusPlaying,
			want:   true,
		},
		{
			name:   "should return false when status is waited for players",
			status: StatusWaitedForPlayers,
			want:   false,
		},
		{
			name:   "should return false when status is ready to start",
			status: StatusReadyToStart,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.status.CanPlayCard())
		})
	}
}
