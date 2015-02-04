package form

import "errors"

type Form struct {
	Errors map[string]error
}

func (f *Form) CheckIfEmpty(field, value, message string) bool {
	if value == "" {
		f.Errors[field] = errors.New(message)
		return true
	}

	return false
}

func (f *Form) Error(field string) string {
	if err, ok := f.Errors[field]; ok {
		return err.Error()
	}

	return ""
}

func (f *Form) HasErrors() bool {
	return len(f.Errors) > 0
}
