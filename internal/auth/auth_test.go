package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		key     string
		wantErr bool
	}{
		{
			name:    "Extract successful",
			header:  http.Header{"Authorization": []string{"ApiKey abcdefghijklmnop"}},
			key:     "abcdefghijklmnop",
			wantErr: false,
		},
		{
			name:    "Authorization header missing",
			header:  http.Header{},
			key:     "",
			wantErr: true,
		},
		{
			name:    "Wrong header format",
			header:  http.Header{"Authorization": []string{"ApiKeyabcdefghijklmnop"}},
			key:     "",
			wantErr: true,
		},
		{
			name:    "Missing key prefix",
			header:  http.Header{"Authorization": []string{"abcdefghijklmnop"}},
			key:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if key != tt.key {
				t.Errorf("GetAPIKey() = %v, want %v", key, tt.key)
			}
		})
	}
}
