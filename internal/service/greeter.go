package service

// NameFormatter is an interface for formatting names
type NameFormatter interface {
	Format(name string) string
}

// DefaultFormatter is the default implementation of NameFormatter
type DefaultFormatter struct{}

// Format formats a name with the default formatting
func (f *DefaultFormatter) Format(name string) string {
	if name == "" {
		return "World"
	}
	return name
}

// Greeter provides greeting functionality
type Greeter struct {
	formatter NameFormatter
}

// NewGreeter creates a new Greeter with the given formatter
func NewGreeter(formatter NameFormatter) *Greeter {
	return &Greeter{formatter: formatter}
}

// Greet returns a greeting for the given name
func (g *Greeter) Greet(name string) string {
	formattedName := g.formatter.Format(name)
	return "Hello, " + formattedName + "!"
}
