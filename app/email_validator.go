package app

import (
	"gopkg.in/webdeskltd/dadata.v2"
)

type emailValidator struct {
	Token  string
	Secret string
}

func (ev *emailValidator) validEmail(email string) (bool, error) {
	checker := dadata.NewDaData(ev.Token, ev.Secret)

	result, err := checker.CleanEmails(email)
	if err != nil {
		return false, err
	}

	for _, e := range result {
		if e.QualityCode == 3 {
			return false, nil
		}
	}

	return true, nil
}
