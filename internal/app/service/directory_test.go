package service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

//nolint:gocyclo
func TestDirectoryGetSize(t *testing.T) {
	tests := []struct {
		name               string
		setupDir           func(tempDir string) error
		includeHiddenFiles bool
		recursive          bool
		expectedSize       int64
		expectError        bool
	}{
		{
			name: "directory with single file",
			setupDir: func(tempDir string) error {
				content := []byte("Hello, World!")

				return os.WriteFile(filepath.Join(tempDir, "test.txt"), content, 0600)
			},
			includeHiddenFiles: false,
			recursive:          false,
			expectedSize:       13,
			expectError:        false,
		},
		{
			name:               "empty directory",
			setupDir:           func(tempDir string) error { return nil },
			includeHiddenFiles: false,
			recursive:          false,
			expectedSize:       0,
			expectError:        false,
		},
		{
			name: "directory with multiple files",
			setupDir: func(tempDir string) error {
				content1 := []byte("Hello, World!")
				content2 := []byte("Go testing")

				err := os.WriteFile(filepath.Join(tempDir, "file1.txt"), content1, 0600)
				if err != nil {
					return err
				}

				return os.WriteFile(filepath.Join(tempDir, "file2.txt"), content2, 0600)
			},
			includeHiddenFiles: false,
			recursive:          false,
			expectedSize:       23,
			expectError:        false,
		},
		{
			name: "directory with hidden files excluded",
			setupDir: func(tempDir string) error {
				regularContent := []byte("regular file content")
				hiddenContent := []byte("hidden file content")

				err := os.WriteFile(filepath.Join(tempDir, "regular.txt"), regularContent, 0600)
				if err != nil {
					return err
				}

				return os.WriteFile(filepath.Join(tempDir, ".hidden.txt"), hiddenContent, 0600)
			},
			includeHiddenFiles: false,
			recursive:          false,
			expectedSize:       20,
			expectError:        false,
		},
		{
			name: "directory with hidden files included",
			setupDir: func(tempDir string) error {
				regularContent := []byte("regular file content")
				hiddenContent := []byte("hidden file content")

				err := os.WriteFile(filepath.Join(tempDir, "regular.txt"), regularContent, 0600)
				if err != nil {
					return err
				}

				return os.WriteFile(filepath.Join(tempDir, ".hidden.txt"), hiddenContent, 0600)
			},
			includeHiddenFiles: true,
			recursive:          false,
			expectedSize:       39,
			expectError:        false,
		},
		{
			name: "non-recursive directory with subdirectory",
			setupDir: func(tempDir string) error {
				subDir := filepath.Join(tempDir, "subdir")

				err := os.Mkdir(subDir, 0755)
				if err != nil {
					return err
				}

				rootContent := []byte("root file content")
				subContent := []byte("subdirectory file content")

				err = os.WriteFile(filepath.Join(tempDir, "root.txt"), rootContent, 0600)
				if err != nil {
					return err
				}

				return os.WriteFile(filepath.Join(subDir, "sub.txt"), subContent, 0600)
			},
			includeHiddenFiles: false,
			recursive:          false,
			expectedSize:       17,
			expectError:        false,
		},
		{
			name: "recursive directory with subdirectory",
			setupDir: func(tempDir string) error {
				subDir := filepath.Join(tempDir, "subdir")

				err := os.Mkdir(subDir, 0755)
				if err != nil {
					return err
				}

				rootContent := []byte("root file content")
				subContent := []byte("subdirectory file content")

				err = os.WriteFile(filepath.Join(tempDir, "root.txt"), rootContent, 0600)
				if err != nil {
					return err
				}

				return os.WriteFile(filepath.Join(subDir, "sub.txt"), subContent, 0600)
			},
			includeHiddenFiles: false,
			recursive:          true,
			expectedSize:       42,
			expectError:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()

			defer os.RemoveAll(tempDir) //nolint:errcheck

			err := tt.setupDir(tempDir)
			require.NoError(t, err)

			dir := NewDirectory(tempDir, tt.includeHiddenFiles, tt.recursive)
			size, err := dir.GetSize()

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedSize, size)
			}
		})
	}
}
