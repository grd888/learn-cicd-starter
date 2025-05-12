package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "Valid API Key",
			headers: http.Header{"Authorization": []string{"ApiKey test-key"}},
			want:    "test-key",
			wantErr: nil,
		},
		{
			name:    "Missing Authorization Header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "Malformed Header - No ApiKey Prefix",
			headers: http.Header{"Authorization": []string{"Bearer test-key"}},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name:    "Empty Authorization Header",
			headers: http.Header{"Authorization": []string{""}},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) ||
				(err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
