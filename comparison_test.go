package gerr

import (
	"errors"
	"testing"
)

func Test_copyErrs(t *testing.T) {
	var (
		err1     = errors.New("error 1")
		err2     = errors.New("error 2")
		err3     = errors.New("error 3")
		errs     = []error{err1, err2, err3}
		copyErrs = copyErrs(errs)
		expected = []error{err1, err2, err3}
	)
	copyErrs[1] = nil
	for i := range errs {
		if expected[i] != errs[i] {
			t.Errorf("copyErrs() faile to copy errors, errs: %v, expected: %v", errs, expected)
			break
		}
	}
}

func Test_containAnyErrs(t *testing.T) {
	var (
		err1  = errors.New("Error 1")
		err2  = errors.New("Error 2")
		err3  = errors.New("Error 3")
		err4  = errors.New("Error 4")
		err5  = errors.New("Error 5")
		err6  = errors.New("Error 6")
		tests = []struct {
			name        string
			errs        []error
			checks      []error
			wantContain bool
		}{
			{name: "Exactly Same", errs: []error{err1, err2, err3}, checks: []error{err1, err2, err3}, wantContain: true},
			{name: "Contain All", errs: []error{err1, err2, err3, err4, err5}, checks: []error{err1, err2, err3}, wantContain: true},
			{name: "Only Contain Some", errs: []error{err1, err2}, checks: []error{err1, err2, err3}, wantContain: true},
			{name: "Contain Some With Others", errs: []error{err1, err6}, checks: []error{err1, err2, err3}, wantContain: true},
			{name: "Does Not Contain", errs: []error{err4, err5}, checks: []error{err1, err2, err3}, wantContain: false},
			{name: "Errs Nil, Checks Not Empty", errs: nil, checks: []error{err1, err2, err3}, wantContain: false},
			{name: "Errs Empty, Checks Not Empty", errs: []error{}, checks: []error{err1, err2, err3}, wantContain: false},
			{name: "Errs Not Empty, Checks Nil", errs: []error{err1, err2, err3}, checks: nil, wantContain: false},
			{name: "Errs Not Empty, Checks Empty", errs: []error{err1, err2, err3}, checks: []error{}, wantContain: false},
			{name: "Both Nil", errs: nil, checks: nil, wantContain: true},
			{name: "Both Empty", errs: []error{}, checks: []error{}, wantContain: true},
			{name: "Errs Nil, Checks Empty", errs: nil, checks: []error{}, wantContain: true},
			{name: "Errs Empty, Checks Nil", errs: []error{}, checks: nil, wantContain: true},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotContain := containAnyErrs(tt.errs, tt.checks); gotContain != tt.wantContain {
				t.Errorf("containAnyErrs() = %v, want %v", gotContain, tt.wantContain)
			}
		})
	}
}

func Test_containAllErrs(t *testing.T) {
	var (
		err1  = errors.New("Error 1")
		err2  = errors.New("Error 2")
		err3  = errors.New("Error 3")
		err4  = errors.New("Error 4")
		err5  = errors.New("Error 5")
		err6  = errors.New("Error 6")
		tests = []struct {
			name        string
			errs        []error
			checks      []error
			wantContain bool
		}{
			{name: "Exactly Same", errs: []error{err1, err2, err3}, checks: []error{err1, err2, err3}, wantContain: true},
			{name: "Contain All", errs: []error{err1, err2, err3, err4, err5}, checks: []error{err1, err2, err3}, wantContain: true},
			{name: "Only Contain Some", errs: []error{err1, err2}, checks: []error{err1, err2, err3}, wantContain: false},
			{name: "Contain Some With Others", errs: []error{err1, err6}, checks: []error{err1, err2, err3}, wantContain: false},
			{name: "Does Not Contain", errs: []error{err4, err5}, checks: []error{err1, err2, err3}, wantContain: false},
			{name: "Errs Nil, Checks Not Empty", errs: nil, checks: []error{err1, err2, err3}, wantContain: false},
			{name: "Errs Empty, Checks Not Empty", errs: []error{}, checks: []error{err1, err2, err3}, wantContain: false},
			{name: "Errs Not Empty, Checks Nil", errs: []error{err1, err2, err3}, checks: nil, wantContain: false},
			{name: "Errs Not Empty, Checks Empty", errs: []error{err1, err2, err3}, checks: []error{}, wantContain: false},
			{name: "Both Nil", errs: nil, checks: nil, wantContain: true},
			{name: "Both Empty", errs: []error{}, checks: []error{}, wantContain: true},
			{name: "Errs Nil, Checks Empty", errs: nil, checks: []error{}, wantContain: true},
			{name: "Errs Empty, Checks Nil", errs: []error{}, checks: nil, wantContain: true},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotContain := containAllErrs(tt.errs, tt.checks); gotContain != tt.wantContain {
				t.Errorf("containAllErrs() = %v, want %v", gotContain, tt.wantContain)
			}
		})
	}
}
