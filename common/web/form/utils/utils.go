package utils

type Validator interface {
	Validate()
	HasErrors() bool
}

func IsValid(v Validator) bool {
	v.Validate()
	if v.HasErrors() {
		return false
	}

	return true
}
