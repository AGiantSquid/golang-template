package service

import "testing"

// MockFormatter is a mock implementation of NameFormatter
type MockFormatter struct {
	FormatFunc func(string) string
	CallCount  int
}

// Format calls the mock FormatFunc and increments CallCount
func (m *MockFormatter) Format(name string) string {
	m.CallCount++
	return m.FormatFunc(name)
}

func TestGreeterWithDefaultFormatter(t *testing.T) {
	// Create a greeter with the default formatter
	formatter := &DefaultFormatter{}
	greeter := NewGreeter(formatter)

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty name", "", "Hello, World!"},
		{"With name", "Alice", "Hello, Alice!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := greeter.Greet(tt.input)
			if result != tt.expected {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGreeterWithMockFormatter(t *testing.T) {
	// Create a mock formatter that always returns "Mocked"
	mockFormatter := &MockFormatter{
		FormatFunc: func(string) string {
			return "Mocked"
		},
	}

	// Create a greeter with the mock formatter
	greeter := NewGreeter(mockFormatter)

	// Test the greeter
	result := greeter.Greet("Alice")
	expected := "Hello, Mocked!"

	// Check the result
	if result != expected {
		t.Errorf("Greet(%q) = %q, want %q", "Alice", result, expected)
	}

	// Check that the formatter was called
	if mockFormatter.CallCount != 1 {
		t.Errorf("Expected Format to be called once, got %d", mockFormatter.CallCount)
	}
}
