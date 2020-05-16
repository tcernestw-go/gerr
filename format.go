package gerr

func (gerr Gerr) log(padding string) (msg string) {
	if gerr.msg != "" {
		msg = padding + "Message: " + gerr.msg + "\n"
	}
	msg += gerr.logErrs(padding)
	msg += gerr.tc.log(padding + padding)
	attmLog := gerr.attms.log(padding + padding)
	if attmLog != "" {
		msg += padding + "Attachments:\n" + attmLog
	}
	return
}

func (gerr Gerr) logErrs(padding string) (msg string) {
	length := len(gerr.errs)
	for i, err := range gerr.errs {
		msg += err.Error()
		if i < length-1 {
			msg += "; "
		}
	}
	if msg != "" {
		msg = padding + "Errors: " + msg + "\n"
	}
	return
}
