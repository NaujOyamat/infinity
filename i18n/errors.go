package i18n

// LanguageFileNotFoundError occurs when not language files are found in the directory
type LanguageFileNotFoundError struct{}

// Error return error message
func (e *LanguageFileNotFoundError) Error() string {
	return "not language files found in the directory"
}

// LanguageNotFoundError occurs when not language is found in memory
type LanguageNotFoundError struct{}

// Error return error message
func (e *LanguageNotFoundError) Error() string {
	return "not language is found in memory"
}

// CurrentLanguageTagNotExistsError occurs when current language tag not exists in memory
type CurrentLanguageTagNotExistsError struct{}

// Error return error message
func (e *CurrentLanguageTagNotExistsError) Error() string {
	return "language tag not exists in memory"
}

// LanguagesNotInitializedError occurs when the languages are not initialized
type LanguagesNotInitializedError struct{}

// Error return error message
func (e *LanguagesNotInitializedError) Error() string {
	return "languages are not initialized"
}

// InvalidSpecificationError occurs when the specification is not a structure
// pointer or one of its properties is not of type string
type InvalidSpecificationError struct{}

// Error return error message
func (e *InvalidSpecificationError) Error() string {
	return "specification must be a struct pointer and all its properties of type string"
}
