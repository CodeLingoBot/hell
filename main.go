package flow

import "fmt"

type ConfigError string

func ConfigErrorf(format string, args ...interface{}) error {
	return ConfigError(fmt.Sprintf(format, args...))
}

func (c ConfigError) Error() string {
	return string(c)
}

func IsConfigError(err error) bool {
	_, ok := err.(ConfigError)
	return ok
}
