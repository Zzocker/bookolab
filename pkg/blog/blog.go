// blog is generic logger used for this project

package blog

// Logger represents logger
// it just interface
// implantation TODO
type Logger interface {
	Infof(string, ...string)
	Debugf(string, ...string)
	Errorf(string, ...string)
	WithFields(fields map[string]string) Logger
}

// New Create blog logger
func New() Logger {
	// Implement me TODO
	return nil
}
