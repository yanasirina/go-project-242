package cli

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
)

func TestNewCLICommand(t *testing.T) {
	tests := []struct {
		name        string
		validateCmd func(*testing.T, *cli.Command)
	}{
		{
			name: "command has correct name",
			validateCmd: func(t *testing.T, cmd *cli.Command) {
				t.Helper()
				require.Equal(t, "hexlet-path-size", cmd.Name)
			},
		},
		{
			name: "command has non-empty usage",
			validateCmd: func(t *testing.T, cmd *cli.Command) {
				t.Helper()
				require.NotEmpty(t, cmd.Usage)
			},
		},
		{
			name: "command has correct number of flags",
			validateCmd: func(t *testing.T, cmd *cli.Command) {
				t.Helper()
				require.Len(t, cmd.Flags, 3)
			},
		},
		{
			name: "command has expected flags",
			validateCmd: func(t *testing.T, cmd *cli.Command) {
				t.Helper()

				flagNames := make(map[string]bool)

				for _, flag := range cmd.Flags {
					for _, name := range flag.Names() {
						flagNames[name] = true
					}
				}

				expectedFlags := []string{HumanFlagName, ShowAllFilesFlag, RecursiveFlag}
				for _, expectedFlag := range expectedFlags {
					require.True(t, flagNames[expectedFlag], "Command missing expected flag: %s", expectedFlag)
				}
			},
		},
		{
			name: "command has non-nil action",
			validateCmd: func(t *testing.T, cmd *cli.Command) {
				t.Helper()
				require.NotNil(t, cmd.Action)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCLICommand()
			tt.validateCmd(t, cmd)
		})
	}
}

func TestRunCMD(t *testing.T) {
	tests := []struct {
		name        string
		setupCmd    func() *cli.Command
		setupArgs   func() []string
		expectError bool
		cleanup     func()
	}{
		{
			name: "successful command execution",
			setupCmd: func() *cli.Command {
				return &cli.Command{
					Name:  "test-cmd",
					Usage: "test command",
					Action: func(ctx context.Context, c *cli.Command) error {
						return nil
					},
				}
			},
			setupArgs: func() []string {
				return []string{"test", "--help"}
			},
			expectError: false,
			cleanup:     func() {},
		},
		{
			name: "command with error action",
			setupCmd: func() *cli.Command {
				return &cli.Command{
					Name:  "test-cmd-error",
					Usage: "test command with error",
					Action: func(ctx context.Context, c *cli.Command) error {
						return ErrBadArguments
					},
				}
			},
			setupArgs: func() []string {
				return []string{"test", "arg1"}
			},
			expectError: true,
			cleanup:     func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origArgs := os.Args

			defer func() {
				os.Args = origArgs

				tt.cleanup()
			}()

			cmd := tt.setupCmd()
			os.Args = tt.setupArgs()

			err := RunCMD(cmd)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
