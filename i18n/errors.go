package i18n

// NotFoundLanguagesError occurs when no language files are found in the directory
type NotFoundLanguagesError struct{}

// Error return error message
func (e *NotFoundLanguagesError) Error() string {
	return "not language files found in the directory"
}
