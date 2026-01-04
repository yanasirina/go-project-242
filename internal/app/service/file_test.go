package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileGetSize(t *testing.T) {
	tests := []struct {
		name         string
		content      []byte
		expectedSize int64
		expectError  bool
	}{
		{
			name:         "non-empty file",
			content:      []byte("Hello, World!"),
			expectedSize: 13,
			expectError:  false,
		},
		{
			name:         "empty file",
			content:      []byte(""),
			expectedSize: 0,
			expectError:  false,
		},
		{
			name:         "file with special characters",
			content:      []byte("Hello\nWorld\t!"),
			expectedSize: 13,
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFile, err := os.CreateTemp(t.TempDir(), "test_file")
			require.NoError(t, err)

			defer os.Remove(tempFile.Name())
			defer tempFile.Close()

			_, err = tempFile.Write(tt.content)
			require.NoError(t, err)

			fileInfo, err := tempFile.Stat()
			require.NoError(t, err)

			file := NewFile(fileInfo)
			size, err := file.GetSize()

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedSize, size)
			}
		})
	}
}
