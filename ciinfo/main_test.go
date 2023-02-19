package main

import (
	"io"
	"strings"
	"testing"

	"github.com/gkampitakis/ciinfo"
)

func mockExit(t *testing.T, expectedCode int) func(int) {
	t.Helper()
	return func(i int) {
		if i != expectedCode {
			t.Errorf("expected code: %d but got: %d\n", expectedCode, i)
		}
	}
}

type mockWriter struct {
	mockWrite func([]byte)
}

func (m mockWriter) Write(p []byte) (int, error) {
	m.mockWrite(p)
	return 0, nil
}

func TestCMD(t *testing.T) {
	t.Run("should be successful in case of pr", func(t *testing.T) {
		ciinfo.IsPr = true

		run(mockExit(t, 0), io.Discard, true, "")
	})

	t.Run("should exit with error in case not being a pr", func(t *testing.T) {
		ciinfo.IsPr = false
		ciinfo.IsCI = false

		run(mockExit(t, -1), io.Discard, true, "")
	})

	t.Run("should print json output", func(t *testing.T) {
		ciinfo.IsPr = true
		ciinfo.Name = "mock-ci"
		ciinfo.IsCI = true

		m := mockWriter{
			mockWrite: func(b []byte) {
				// clean up tabs,spaces and newlines
				chars := []string{" ", "\n", "\t"}
				v := string(b)
				for _, c := range chars {
					v = strings.ReplaceAll(v, c, "")
				}

				if v != `{"is_ci":true,"ci_name":"mock-ci","pull_request":true}` {
					t.Error("expected json format")
				}
			},
		}

		run(nil, m, false, "json")
	})

	t.Run("should print pretty output when not on ci", func(t *testing.T) {
		ciinfo.IsCI = false

		m := mockWriter{
			mockWrite: func(b []byte) {
				if string(b) != "Not running on CI.\n" {
					t.Error("expected not running on CI.")
				}
			},
		}

		run(nil, m, false, "pretty")
	})

	t.Run("should print pretty output", func(t *testing.T) {
		ciinfo.IsPr = true
		ciinfo.Name = "mock-ci"
		ciinfo.IsCI = true
		msgs := []string{
			"CI Name: mock-ci\n",
			"Is Pull Request.\n",
		}

		m := mockWriter{
			mockWrite: func(b []byte) {
				msg := msgs[0]
				msgs = msgs[1:]
				if string(b) != msg {
					t.Errorf("unexpected message: %s", string(b))
				}
			},
		}

		run(nil, m, false, "pretty")
	})

	t.Run("should print pretty output without name", func(t *testing.T) {
		ciinfo.IsPr = true
		ciinfo.Name = ""
		ciinfo.IsCI = true
		msgs := []string{
			"Running on CI.\n",
			"Is Pull Request.\n",
		}

		m := mockWriter{
			mockWrite: func(b []byte) {
				msg := msgs[0]
				msgs = msgs[1:]
				if string(b) != msg {
					t.Errorf("unexpected message: %s", string(b))
				}
			},
		}

		run(nil, m, false, "pretty")
	})

	t.Run("should be successful code", func(t *testing.T) {
		ciinfo.IsPr = true
		ciinfo.Name = "mock-ci"
		ciinfo.IsCI = true

		run(mockExit(t, 0), io.Discard, false, "")
	})

	t.Run("should be error code", func(t *testing.T) {
		ciinfo.IsCI = false

		run(mockExit(t, -1), io.Discard, false, "")
	})
}
