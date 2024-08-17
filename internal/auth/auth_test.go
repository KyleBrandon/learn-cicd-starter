package auth

import "testing"

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header map[string][]string
		error  bool
		expect string
	}{
		"happy":          {header: map[string][]string{"Authorization": {"ApiKey 12345"}}, error: false, expect: "12345"},
		"no_auth_header": {header: map[string][]string{"uthorization": {"ApiKy 12345"}}, error: true, expect: "no authorization header included"},
		"bad_key_name":   {header: map[string][]string{"Authorization": {"ApiKy 12345"}}, error: true, expect: "malformed authorization header"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.header)
			if tc.error {
				if err.Error() != tc.expect {
					t.Fatalf("expected: %v, got: %v", tc.expect, err.Error())
				}
			} else {
				if got != tc.expect {
					t.Fatalf("expected: %v, got: %v", tc.expect, got)
				}
			}

		})
	}

}
