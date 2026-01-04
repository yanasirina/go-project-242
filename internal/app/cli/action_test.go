package cli

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
)

func TestRunPathSizeAction(t *testing.T) {
	tests := []struct {
		name        string
		setupPath   func() (string, func(), error)
		args        []string
		expectError bool
		validate    func(*testing.T, string, error)
	}{
		{
			name: "valid file path",
			setupPath: func() (string, func(), error) {
				tempFile, err := os.CreateTemp(t.TempDir(), "test_cli_action")
				if err != nil {
					return "", nil, err
				}

				content := []byte("Test content for CLI action")

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
			args:        []string{"test"},
			expectError: false,
			validate: func(t *testing.T, output string, err error) {
				t.Helper()
				require.NoError(t, err)
				require.Contains(t, output, "B")
			},
		},
		{
			name: "no arguments provided",
			setupPath: func() (string, func(), error) {
				return "", func() {}, nil
			},
			args:        []string{"test"},
			expectError: true,
			validate: func(t *testing.T, output string, err error) {
				t.Helper()
				require.Error(t, err)
				require.Contains(t, err.Error(), "required arguments are missing")
			},
		},
		{
			name: "valid directory path",
			setupPath: func() (string, func(), error) {
				tempDir := t.TempDir()

				testFile := filepath.Join(tempDir, "test.txt")
				content := []byte("Test content for CLI action")

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
			args:        []string{"test"},
			expectError: false,
			validate: func(t *testing.T, output string, err error) {
				t.Helper()
				require.NoError(t, err)
				require.Contains(t, output, "B")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup, err := tt.setupPath()
			require.NoError(t, err)

			defer cleanup()

			args := slices.Clone(tt.args)
			args = append(args, path)

			cmd := &cli.Command{
				Name:   "size",
				Action: RunPathSizeAction,
			}

			old := os.Stdout
			r, w, err := os.Pipe()
			require.NoError(t, err)

			os.Stdout = w
			err = cmd.Run(context.Background(), args)

			w.Close() //nolint: errcheck

			os.Stdout = old

			var buf bytes.Buffer
			buf.ReadFrom(r) //nolint: errcheck

			output := buf.String()

			tt.validate(t, output, err)
		})
	}
}

func TestRunPathSizeActionWithFlags(t *testing.T) {
	tempDir := t.TempDir()

	defer os.RemoveAll(tempDir) //nolint:errcheck

	testFile := filepath.Join(tempDir, "test.txt")
	content := []byte("Test content for CLI flags")
	err := os.WriteFile(testFile, content, 0600)
	require.NoError(t, err)

	hiddenFile := filepath.Join(tempDir, ".hidden.txt")
	hiddenContent := []byte("Hidden file content")
	err = os.WriteFile(hiddenFile, hiddenContent, 0600)
	require.NoError(t, err)

	cmd := &cli.Command{
		Name: "size",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: HumanFlagName, Aliases: []string{"H"}},
			&cli.BoolFlag{Name: ShowAllFilesFlag, Aliases: []string{"a"}},
			&cli.BoolFlag{Name: RecursiveFlag, Aliases: []string{"r"}},
		},
		Action: RunPathSizeAction,
	}

	old := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = w
	osArgs := []string{"test", "--" + ShowAllFilesFlag, tempDir}

	err = cmd.Run(context.Background(), osArgs)
	require.NoError(t, err)

	w.Close() //nolint: errcheck

	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r) //nolint: errcheck
	output := buf.String()

	require.Contains(t, output, tempDir)
}
