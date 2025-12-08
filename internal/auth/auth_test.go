package auth

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		name    string
		input   map[string]string
		want    string
		wantErr string
	}{
		"Correct Header": {
			input:   map[string]string{"Authorization": "ApiKey Test1234"},
			want:    "Test1234",
			wantErr: "not expecting an error",
		},
		"Wrong Authorization Header": {
			input:   map[string]string{"Authorization": "ApiKey"},
			want:    "",
			wantErr: "malformed authorization header",
		},
		"No Authorization Header": {
			input:   map[string]string{"host": "localhost:8080"},
			want:    "",
			wantErr: "no authorization header included",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			h := http.Header{}
			for k, v := range tc.input {
				h.Set(k, v)
			}
			got, errGot := GetAPIKey(h)
			if errGot != nil {
				if strings.Contains(errGot.Error(), tc.wantErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", errGot)
				return
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}
