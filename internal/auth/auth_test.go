package auth

import (
	"errors"
	"net/http"
	"testing"
)

func sameErrorMessage(err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}

	if err1 == nil || err2 == nil {
		return false
	}

	return err1.Error() == err2.Error()
}

func TestGetAPIKey(t *testing.T) {
	// Arrange: define test cases
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr error
	}{
		{
			name:    "no auth header",
			headers: http.Header{},
			wantKey: "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"some token"},
			},
			wantKey: "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "valid header",
			headers: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			wantKey: "abc123",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act: execute function
			gotKey, gotErr := GetAPIKey(tt.headers)

			// Assert: check if results match expectations
			if gotKey != tt.wantKey {
				t.Errorf("expected key: %v, got: %v", tt.wantKey, gotKey)
			}

			if !sameErrorMessage(gotErr, tt.wantErr) {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, gotErr)
			}
		})
	}

	// // Arrange: no headers at all
	// headers := http.Header{}

	// want := struct {
	// 	key string
	// 	err error
	// }{
	// 	key: "",
	// 	err: ErrNoAuthHeaderIncluded,
	// }

	// // Act
	// gotKey, gotErr := GetAPIKey(headers)

	// got := struct {
	// 	key string
	// 	err error
	// }{
	// 	key: gotKey,
	// 	err: gotErr,
	// }

	// // Assert
	// if got.key != want.key {
	// 	t.Errorf("expected key %v, got %v", want.key, got.key)
	// }

	// if !errors.Is(got.err, want.err) {
	// 	t.Errorf("expected error: %v, got: %v", got.err, want.err)
	// }
}
