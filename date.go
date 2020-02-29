package utils

import (
	"fmt"
	"regexp"
	"time"
)

// IsValidDate checks if the provided date is of the format YYYY-MM-DD and if the date is valid
// The method returns an error if the date is not valid
func IsValidDate(date string) error {
	dateFormatRx, err := regexp.Compile(dateFormatRegex)
	currentDate := time.Now()
	if err != nil {
		return Error{Message: regexCompileErrMsg, Detail: regexCompileErrDetail}.NewError()
	}
	if !dateFormatRx.MatchString(date) {
		return Error{Message: dateFormatErrMsg, Detail: dateFormatErrDetail}.NewError()
	}
	d, err := time.Parse(dateFormatLayout, date)
	if err != nil {
		return Error{Message: invalidDateErrMsg, Detail: fmt.Sprintf(invalidDateErrDetail, err)}.NewError()
	}
	if d.Format(dateFormatLayout) > currentDate.Format(dateFormatLayout) {
		return Error{Message: invalidDateErrMsg, Detail: greaterThanCurrentDateErrDetail}.NewError()
	}
	return nil
}

// IsValidYear checks if the year is between 1960 and the current year
// The method returns an error if the year is not valid
func IsValidYear(year int) error {
	currentYear := time.Now().Year()
	if year < 1990 || year > currentYear {
		return Error{Message: invalidYearErrMsg, Detail: fmt.Sprintf(invalidYearErrDetail, currentYear)}.NewError()
	}
	return nil
}
