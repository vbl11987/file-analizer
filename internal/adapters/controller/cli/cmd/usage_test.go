package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

type mockManagerProcessor struct {
	ExecuteFn func(path string) error
}

func (m mockManagerProcessor) Execute(path string) error {
	return m.ExecuteFn(path)
}

type mockLogger struct{}

func (m mockLogger) Info(messageFormat string, v ...interface{}) {
	fmt.Println(messageFormat, v)
}
func (m mockLogger) Fatal(message string, err ...error) {
	fmt.Println(message, err)
}

func Test_UsageCommand(t *testing.T) {
	p := mockManagerProcessor{}

	t.Run("Regular execution of the usage command", func(t *testing.T) {
		p.ExecuteFn = func(path string) error { return nil }
		cmd := UsageCommand(&p, mockLogger{})
		b := bytes.NewBufferString("fake/path/to/")
		cmd.SetOut(b)
		cmd.SetArgs([]string{b.String()})
		err := cmd.Execute()
		if err != nil {
			t.Errorf("usage command failed: %v", err)
		}
	})

	t.Run("Failed execution of the usage command", func(t *testing.T) {
		p.ExecuteFn = func(path string) error { return fmt.Errorf("fake error") }
		cmd := UsageCommand(&p, mockLogger{})
		b := bytes.NewBufferString("fake/path/to/")
		cmd.SetOut(b)
		cmd.SetArgs([]string{b.String()})
		err := cmd.Execute()
		if err != nil {
			t.Errorf("error in usage command failed: %v", err)
		}
	})
}
