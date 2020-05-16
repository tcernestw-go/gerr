package gerr

import "log"

// Gerr is an go error providing more functions
type Gerr struct {
	msg   string
	errs  []error // current errors
	cause *Gerr   // previous cause
	attms attachments
	tc    traceCollection // a collection of trace logs
}

// Error is an implement of golang error
func (gerr Gerr) Error() (msg string) {
	padding := "  "
	msg = "Failed:\n" + gerr.log(padding)
	for cause := gerr.cause; cause != nil; cause = cause.cause {
		msg += cause.log(padding)
	}
	return
}

// Wrap try to wrap the error into the cause of gerr
// if the cause is not a gerr, it will directly wrap as a gerr instead of adding into the cause
func Wrap(cause error, msg string) (gerr *Gerr) {
	if conv, ok := cause.(*Gerr); ok {
		gerr = &Gerr{}
		gerr.cause = conv
		log.Println("ok", ok)
	} else { // the cause is not a gerr
		gerr = &Gerr{
			errs: []error{cause},
			tc:   trace(1),
		}
	}
	gerr.msg = msg
	return
}

// AddErrs adds errors into gerr
func (gerr *Gerr) AddErrs(errs ...error) *Gerr {
	gerr.errs = append(gerr.errs, errs...)
	return gerr
}

// WithErrs replace errs into gerr
func (gerr *Gerr) WithErrs(errs ...error) *Gerr {
	gerr.errs = errs
	return gerr
}

// Errs returns only the errors in current gerr
func (gerr Gerr) Errs() (errs []error) {
	return gerr.errs
}

// Err returns only the first error of errors in current gerr
// it will returns nil if there is no error
func (gerr Gerr) Err() (err error) {
	if gerr.errs != nil {
		err = gerr.errs[0]
	}
	return
}

// AllCauseErrs returns all the errors in all the causes of gerr
func (gerr Gerr) AllCauseErrs() (errs []error) {
	for cause := gerr.cause; cause != nil; cause = cause.cause {
		errs = append(errs, cause.errs...)
	}
	return
}

// Attach allows user to attach anythings into gerr's attachment
func (gerr *Gerr) Attach(key string, attachment interface{}) *Gerr {
	gerr.attms.add(key, attachment)
	return gerr
}

// Attachment allows user to get attachment from current gerr by key
func (gerr *Gerr) Attachment(key string) interface{} {
	return gerr.attms.get(key)
}

// Trace enable the tracing of current gerr
func (gerr *Gerr) Trace() *Gerr {
	gerr.tc = trace(1)
	return gerr
}
