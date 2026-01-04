package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var errOriginalError = errors.New("original error")

func TestWrap(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		msg         string
		expectNil   bool
		validateErr func(*testing.T, error)
	}{
		{
			name:      "Wrap with nil error",
			err:       nil,
			msg:       "wrapped message",
			expectNil: true,
			validateErr: func(t *testing.T, err error) {
				t.Helper()
				require.NoError(t, err)
			},
		},
		{
			name:      "Wrap with non-nil error",
			err:       errOriginalError,
			msg:       "wrapped message",
			expectNil: false,
			validateErr: func(t *testing.T, err error) {
				t.Helper()
				require.Error(t, err)
				require.Contains(t, err.Error(), "wrapped message")
				require.Contains(t, err.Error(), "original error")
			},
		},
		{
			name:      "Wrap with empty message",
			err:       errOriginalError,
			msg:       "",
			expectNil: false,
			validateErr: func(t *testing.T, err error) {
				t.Helper()
				require.Error(t, err)
				require.Contains(t, err.Error(), ": original error")
			},
		},
		{
			name:      "Wrap with empty error",
			err:       nil,
			msg:       "",
			expectNil: true,
			validateErr: func(t *testing.T, err error) {
				t.Helper()
				require.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Wrap(tt.err, tt.msg)
			tt.validateErr(t, result)
		})
	}
}

func TestWrapf(t *testing.T) {
	tests := []struct {
		name        string
		testFunc    func() error
		expectNil   bool
		validateErr func(*testing.T, error)
	}{
		{
			name: "Wrapf with nil error",
			testFunc: func() error {
				return Wrapf(nil, "wrapped message with %d", 42)
			},
			expectNil: true,
			validateErr: func(t *testing.T, err error) {
				t.Helper()
				require.NoError(t, err)
			},
		},
		{
			name: "Wrapf with non-nil error",
			testFunc: func() error {
				return Wrapf(errOriginalError, "wrapped message with %d", 42)
			},
			expectNil: false,
			validateErr: func(t *testing.T, err error) {
				t.Helper()
				require.Error(t, err)
				require.Contains(t, err.Error(), "wrapped message with 42")
				require.Contains(t, err.Error(), "original error")
			},
		},
		{
			name: "Wrapf with multiple args",
			testFunc: func() error {
				return Wrapf(errOriginalError, "error %s with code %d", "test", 123)
			},
			expectNil: false,
			validateErr: func(t *testing.T, err error) {
				t.Helper()
				require.Error(t, err)
				require.Contains(t, err.Error(), "error test with code 123")
				require.Contains(t, err.Error(), "original error")
			},
		},
		{
			name: "Wrapf with empty format",
			testFunc: func() error {
				return Wrapf(errOriginalError, "")
			},
			expectNil: false,
			validateErr: func(t *testing.T, err error) {
				t.Helper()
				require.Error(t, err)
				require.Contains(t, err.Error(), ": original error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.testFunc()
			tt.validateErr(t, result)
		})
	}
}
