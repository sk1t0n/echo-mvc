package cli

import "testing"

func Test_isDirectory(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		path string
		want bool
	}{
		{".", ".", true},
		{"..", "..", true},
		{"./..", "./..", true},
		{"../../../", "../../../", true},
		{"/", "/", true},
		{"dir1", "dir1", true},
		{"../dir1", "../dir1", true},
		{"./dir-1/dir_2", "./dir-1/dir_2", true},
		{"/dir1/../../dir2", "/dir1/../../dir2", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isDirectory(tt.path)
			if got != tt.want {
				t.Errorf("isDirectory(%s) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func Test_validateArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"no args", []string{}, true},
		{"no command", []string{"echo-mvc"}, true},
		{"no directory", []string{"echo-mvc", "new"}, true},
		{"invalid command", []string{"echo-mvc", "New", "dir"}, true},
		{"invalid directory", []string{"echo-mvc", "new", "dir1$"}, true},
		{"valid args", []string{"echo-mvc", "new", "dir"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := validateArgs(tt.args)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("validateArgs(%v) failed: %v", tt.args, gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatalf("validateArgs(%v) succeeded unexpectedly", tt.args)
			}
		})
	}
}
