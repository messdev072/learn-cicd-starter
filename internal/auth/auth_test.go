package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
		err   error
	}{
		"no auth header": {
			input: http.Header{},
			want:  "",
			err:   ErrNoAuthHeaderIncluded,
		},
		"valid auth header": {
			input: http.Header{"Authorization": []string{"ApiKey abc123"}},
			want:  "abc123",
			err:   nil,
		},
		"invalid auth header format": {
			input: http.Header{"Authorization": []string{"InvalidFormat abc123"}},
			want:  "",
			err:   errors.New("malformed authorization header"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if (err != nil) != (tc.err != nil) {
				t.Errorf("GetAPIKey() error = %v, want error %v", err, tc.err)
				return
			}
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("GetAPIKey() error = %v, want error %v", err, tc.err)
				return
			}
			if got != tc.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tc.want)
			}
		})
	}
}
