package gerr

// ContainAllErrs check if gerr contains all the errors provided
// the checking including the causes
func (gerr Gerr) ContainAllErrs(errs ...error) (contain bool) {
	return containAllErrs(append(gerr.errs, gerr.AllCauseErrs()...), errs)
}

// ContainAnyErrs check if gerr contains any of the errors provided
// the checking including the causes
func (gerr Gerr) ContainAnyErrs(errs ...error) (contain bool) {
	return containAnyErrs(append(gerr.errs, gerr.AllCauseErrs()...), errs)
}

// IsAllErrs check if this gerr is all the errors provided
// it will not check the causes
func (gerr Gerr) IsAllErrs(errs ...error) (is bool) {
	return containAllErrs(gerr.errs, errs)
}

// IsAnyErrs check if this gerr is any the errors provided
// it will not check the causes
func (gerr Gerr) IsAnyErrs(errs ...error) (is bool) {
	return containAnyErrs(gerr.errs, errs)
}

// CauseContainAllErrs check if the causes of gerr contains all the errors provided
func (gerr Gerr) CauseContainAllErrs(errs ...error) (contain bool) {
	return containAllErrs(gerr.AllCauseErrs(), errs)
}

// CauseContainAnyErrs check if the causes of gerr contains any of the errors provided
func (gerr Gerr) CauseContainAnyErrs(errs ...error) (contain bool) {
	return containAnyErrs(gerr.AllCauseErrs(), errs)
}

func containAllErrs(errs []error, checks []error) (contain bool) {
	errs = copyErrs(errs)
	checks = copyErrs(checks)

	if len(errs) == 0 && len(checks) == 0 {
		return true
	} else if len(checks) == 0 {
		return false
	}

	for _, err := range errs {
		for i, check := range checks {
			if err == check { // take it away if they are same error
				checks = append(checks[:i], checks[i+1:]...)
				break
			}
		}
		if len(checks) == 0 { // already contains all errors
			return true
		}
	}
	return
}

func containAnyErrs(errs []error, checks []error) (contain bool) {
	errs = copyErrs(errs)
	checks = copyErrs(checks)

	if len(errs) == 0 && len(checks) == 0 {
		return true
	}

	for _, err := range errs {
		for _, check := range checks {
			if err == check { // return true when there is one error same
				return true
			}
		}
	}
	return
}

func copyErrs(errs []error) (copyErrs []error) {
	copyErrs = make([]error, len(errs))
	copy(copyErrs, errs)
	return
}
