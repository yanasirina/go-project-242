package handler

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathSizeHandler(t *testing.T) {
	tests := []struct {
		name         string
		setupPath    func() (string, func(), error)
		arguments    CommandArguments
		flags        CommandFlags
		expectError  bool
		validateSize func(*testing.T, int64)
	}{
		{
			name: "get size of regular file",
			setupPath: func() (string, func(), error) {
				tempFile, err := os.CreateTemp(t.TempDir(), "test_file")
				if err != nil {
					return "", nil, err
				}

				content := []byte("Hello, World!")

				_, err = tempFile.Write(content)
				if err != nil {
					tempFile.Close()           //nolint:errcheck
					os.Remove(tempFile.Name()) //nolint:errcheck

					return "", nil, err
				}

				tempFile.Close() //nolint:errcheck

				cleanup := func() {
					os.Remove(tempFile.Name()) //nolint:errcheck
				}

				return tempFile.Name(), cleanup, nil
			},
			arguments:   CommandArguments{Path: ""},
			flags:       CommandFlags{},
			expectError: false,
			validateSize: func(t *testing.T, size int64) {
				t.Helper()
				require.Equal(t, int64(13), size)
			},
		},
		{
			name: "get size of empty file",
			setupPath: func() (string, func(), error) {
				tempFile, err := os.CreateTemp(t.TempDir(), "test_file_empty")
				if err != nil {
					return "", nil, err
				}

				tempFile.Close() //nolint:errcheck

				cleanup := func() {
					os.Remove(tempFile.Name()) //nolint:errcheck
				}

				return tempFile.Name(), cleanup, nil
			},
			arguments:   CommandArguments{Path: ""},
			flags:       CommandFlags{},
			expectError: false,
			validateSize: func(t *testing.T, size int64) {
				t.Helper()
				require.Equal(t, int64(0), size)
			},
		},
		{
			name: "non-existent path",
			setupPath: func() (string, func(), error) {
				return "/non/existent/path", func() {}, nil
			},
			arguments:   CommandArguments{Path: ""},
			flags:       CommandFlags{},
			expectError: true,
			validateSize: func(t *testing.T, size int64) {
				t.Helper()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup, err := tt.setupPath()
			require.NoError(t, err)

			defer cleanup()

			handler := PathSizeHandler{
				Arguments: CommandArguments{Path: path},
				Flags:     tt.flags,
			}

			size, err := handler.GetPathSize()

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				tt.validateSize(t, size)
			}
		})
	}
}

func TestPathSizeHandler_Dirs(t *testing.T) {
	tests := []struct {
		name         string
		setupPath    func() (string, func(), error)
		arguments    CommandArguments
		flags        CommandFlags
		expectError  bool
		validateSize func(*testing.T, int64)
	}{
		{
			name: "get size of directory",
			setupPath: func() (string, func(), error) {
				tempDir := t.TempDir()

				testFile := filepath.Join(tempDir, "test.txt")
				content := []byte("Hello, World!")

				err := os.WriteFile(testFile, content, 0600)
				if err != nil {
					os.RemoveAll(tempDir) //nolint:errcheck

					return "", nil, err
				}

				cleanup := func() {
					os.RemoveAll(tempDir) //nolint:errcheck
				}

				return tempDir, cleanup, nil
			},
			arguments:   CommandArguments{Path: ""},
			flags:       CommandFlags{},
			expectError: false,
			validateSize: func(t *testing.T, size int64) {
				t.Helper()
				require.Equal(t, int64(13), size)
			},
		},
		{
			name: "get size of empty directory",
			setupPath: func() (string, func(), error) {
				tempDir := t.TempDir()

				cleanup := func() {
					os.RemoveAll(tempDir) //nolint:errcheck
				}

				return tempDir, cleanup, nil
			},
			arguments:   CommandArguments{Path: ""},
			flags:       CommandFlags{},
			expectError: false,
			validateSize: func(t *testing.T, size int64) {
				t.Helper()
				require.Equal(t, int64(0), size)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup, err := tt.setupPath()
			require.NoError(t, err)

			defer cleanup()

			handler := PathSizeHandler{
				Arguments: CommandArguments{Path: path},
				Flags:     tt.flags,
			}

			size, err := handler.GetPathSize()

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				tt.validateSize(t, size)
			}
		})
	}
}

//nolint:gocyclo
func TestPathSizeHandler_SymLinks(t *testing.T) {
	tests := []struct {
		name         string
		setupPath    func() (string, func(), error)
		arguments    CommandArguments
		flags        CommandFlags
		expectError  bool
		validateSize func(*testing.T, int64)
	}{
		{
			name: "symbolic link to file",
			setupPath: func() (string, func(), error) {
				tempFile, err := os.CreateTemp(t.TempDir(), "target_file")
				if err != nil {
					return "", nil, err
				}

				content := []byte("Hello, Symbolic Link!")

				_, err = tempFile.Write(content)
				if err != nil {
					tempFile.Close()
					os.Remove(tempFile.Name())

					return "", nil, err
				}

				tempFile.Close()

				tempDir := t.TempDir()
				symlinkPath := filepath.Join(tempDir, "symlink_to_file")

				err = os.Symlink(tempFile.Name(), symlinkPath)
				if err != nil {
					os.Remove(tempFile.Name())

					return "", nil, err
				}

				cleanup := func() {
					os.Remove(symlinkPath)
					os.Remove(tempFile.Name())
				}

				return symlinkPath, cleanup, nil
			},
			arguments:   CommandArguments{Path: ""},
			flags:       CommandFlags{},
			expectError: false,
			validateSize: func(t *testing.T, size int64) {
				t.Helper()
				require.Positive(t, size, "Symbolic link size should be greater than 0 (length of target path)")
			},
		},
		{
			name: "symbolic link to directory",
			setupPath: func() (string, func(), error) {
				tempDir := t.TempDir()
				testFile := filepath.Join(tempDir, "test.txt")
				content := []byte("Hello from linked directory!")

				err := os.WriteFile(testFile, content, 0600)
				if err != nil {
					return "", nil, err
				}

				parentDir := t.TempDir()

				symlinkPath := filepath.Join(parentDir, "symlink_to_dir")

				err = os.Symlink(tempDir, symlinkPath)
				if err != nil {
					os.RemoveAll(tempDir)

					return "", nil, err
				}

				cleanup := func() {
					os.Remove(symlinkPath)
					os.RemoveAll(tempDir)
				}

				return symlinkPath, cleanup, nil
			},
			arguments:   CommandArguments{Path: ""},
			flags:       CommandFlags{},
			expectError: false,
			validateSize: func(t *testing.T, size int64) {
				t.Helper()
				require.Positive(t, size, "Symbolic link size should be greater than 0 (length of target path)")
			},
		},
		{
			name: "broken symbolic link",
			setupPath: func() (string, func(), error) {
				tempDir := t.TempDir()

				brokenSymlinkPath := filepath.Join(tempDir, "broken_symlink")
				targetPath := filepath.Join(tempDir, "non_existent_target")

				err := os.Symlink(targetPath, brokenSymlinkPath)
				if err != nil {
					return "", nil, err
				}

				cleanup := func() {
					os.Remove(brokenSymlinkPath)
				}

				return brokenSymlinkPath, cleanup, nil
			},
			arguments:   CommandArguments{Path: ""},
			flags:       CommandFlags{},
			expectError: false,
			validateSize: func(t *testing.T, size int64) {
				t.Helper()
				require.Positive(t, size, "Broken symbolic link size should be greater than 0 (length of target path)")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup, err := tt.setupPath()
			require.NoError(t, err)

			defer cleanup()

			handler := PathSizeHandler{
				Arguments: CommandArguments{Path: path},
				Flags:     tt.flags,
			}

			size, err := handler.GetPathSize()

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				tt.validateSize(t, size)
			}
		})
	}
}

//nolint:gocyclo
func TestPathSizeHandlerGetFormatedSize(t *testing.T) {
	tests := []struct {
		name            string
		setupPath       func() (string, func(), error)
		flags           CommandFlags
		expectedPattern string
		expectError     bool
	}{
		{
			name: "formatted size without human flag",
			setupPath: func() (string, func(), error) {
				tempFile, err := os.CreateTemp(t.TempDir(), "test_file_formatted")
				if err != nil {
					return "", nil, err
				}

				content := []byte("Test content for formatting")

				_, err = tempFile.Write(content)
				if err != nil {
					tempFile.Close()           //nolint:errcheck
					os.Remove(tempFile.Name()) //nolint:errcheck

					return "", nil, err
				}

				tempFile.Close() //nolint:errcheck

				cleanup := func() {
					os.Remove(tempFile.Name()) //nolint:errcheck
				}

				return tempFile.Name(), cleanup, nil
			},
			flags:           CommandFlags{HumanizeSize: false},
			expectedPattern: "B",
			expectError:     false,
		},
		{
			name: "formatted size with human flag",
			setupPath: func() (string, func(), error) {
				tempFile, err := os.CreateTemp(t.TempDir(), "test_file_human")
				if err != nil {
					return "", nil, err
				}

				content := []byte("Test content")

				_, err = tempFile.Write(content)
				if err != nil {
					tempFile.Close()           //nolint:errcheck
					os.Remove(tempFile.Name()) //nolint:errcheck

					return "", nil, err
				}

				tempFile.Close() //nolint:errcheck

				cleanup := func() {
					os.Remove(tempFile.Name()) //nolint:errcheck
				}

				return tempFile.Name(), cleanup, nil
			},
			flags:           CommandFlags{HumanizeSize: true},
			expectedPattern: "B",
			expectError:     false,
		},
		{
			name: "formatted size for zero-byte file with human flag",
			setupPath: func() (string, func(), error) {
				tempFile, err := os.CreateTemp(t.TempDir(), "test_file_zero")
				if err != nil {
					return "", nil, err
				}

				tempFile.Close() //nolint:errcheck

				cleanup := func() {
					os.Remove(tempFile.Name()) //nolint:errcheck
				}

				return tempFile.Name(), cleanup, nil
			},
			flags:           CommandFlags{HumanizeSize: true},
			expectedPattern: "0B",
			expectError:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup, err := tt.setupPath()
			require.NoError(t, err)

			defer cleanup()

			handler := PathSizeHandler{
				Arguments: CommandArguments{Path: path},
				Flags:     tt.flags,
			}

			result, err := handler.GetFormatedSize()

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Contains(t, result, tt.expectedPattern)
			}
		})
	}
}
