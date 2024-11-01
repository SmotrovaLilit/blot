package game

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStatus_SetBet(t *testing.T) {
	testCases := []struct {
		Name          string
		ShouldFail    bool
		Status        Status
		ExpectedError error
	}{
		{
			Name:       "should fail when status is not betting",
			ShouldFail: true,
			ExpectedError: ErrGameNotReadyToSetBet{
				Status: StatusPlaying.String(),
			},
			Status: StatusPlaying,
		},
		{
			Name:   "should set bet",
			Status: StatusBetting,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := tt.Status.SetBet()
			if tt.ShouldFail {
				require.Error(t, err)
				require.Equal(t, tt.ExpectedError, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, StatusPlaying, got)
		})
	}
}
