package utils

import (
	"fmt"
	"regexp"
	"time"
)

// IsValidDate checks if the date is of the format YYY-MM-DD and if the date is valid
// The method returns an error if the date is not valid
func IsValidDate(date string) error {
	dateFormatRx, err := regexp.Compile(dateFormatRegex)
	currentDate := time.Now()
	if err != nil {
		return Error{ErrStr: regexCompileErr, ErrMsg: regexCompileErrStr}.NewError()
	}
	if !dateFormatRx.MatchString(date) {
		return Error{ErrStr: dateFormatErr, ErrMsg: dateFormatErrStr}.NewError()
	}
	d, err := time.Parse(dateFormatLayout, date)
	if err != nil {
		return Error{ErrStr: invalidDateErr, ErrMsg: fmt.Sprintf(invalidDateErrStr, err)}.NewError()
	}
	if d.Format(dateFormatLayout) > currentDate.Format(dateFormatLayout) {
		return Error{ErrStr: invalidDateErr, ErrMsg: greaterThanCurrentDateErrStr}.NewError()
	}
	return nil
}

// IsValidYear checks if the year is between 1960 and the current year
// The method returns an error if the year is not valid
func IsValidYear(year int) error {
	currentYear := time.Now().Year()
	if year < 1990 || year > currentYear {
		return Error{ErrStr: invalidYearErr, ErrMsg: fmt.Sprintf(invalidYearErrStr, currentYear)}.NewError()
	}
	return nil
}
