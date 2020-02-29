package utils

import (
	"fmt"
	"regexp"
	"time"
)

const (
	dateFormatRegex  = `(^([0-9]{4})-([0-9]{2})-([0-9]{2})$)`
	dateFormatLayout = "2006-01-02"

	regexCompileErrMsg              = "Unable to compile regex"
	regexCompileErrDetail           = "there was an error compiling the regex"
	dateFormatErrMsg                = "Invalid date format"
	dateFormatErrDetail             = "date should be of the format YYYY-MM-DD"
	invalidDateErrMsg               = "Data is not valid"
	invalidDateErrDetail            = "there was an error while parsing the date : %v"
	greaterThanCurrentDateErrDetail = "date should not be greater than current date"
	invalidYearErrMsg               = "Year is not valid"
	invalidYearErrDetail            = "year should be between 1990 and %d"
)

// IsValidDate checks if the provided date is of the format YYYY-MM-DD and if the date is valid
// The method returns a boolean to indicate if the data is valid or not and an error if the date is not valid
func IsValidDate(date string) (bool, error) {
	currentDate := time.Now()
	dateFormatRx, err := regexp.Compile(dateFormatRegex)
	if err != nil {
		return false, Error{Message: regexCompileErrMsg, Detail: regexCompileErrDetail}.NewError()
	}
	if !dateFormatRx.MatchString(date) {
		return false, Error{Message: dateFormatErrMsg, Detail: dateFormatErrDetail}.NewError()
	}
	d, err := time.Parse(dateFormatLayout, date)
	if err != nil {
		return false, Error{Message: invalidDateErrMsg, Detail: fmt.Sprintf(invalidDateErrDetail, err)}.NewError()
	}
	if d.Format(dateFormatLayout) > currentDate.Format(dateFormatLayout) {
		return false, Error{Message: invalidDateErrMsg, Detail: greaterThanCurrentDateErrDetail}.NewError()
	}
	return true, nil
}

// IsValidYear checks if the year is between 1990 and the current year
// The method returns a true or a false indicating if the year is valid or not and an error if the year is not valid
func IsValidYear(year int) (bool, error) {
	currentYear := time.Now().Year()
	if year < 1990 || year > currentYear {
		return false, Error{Message: invalidYearErrMsg, Detail: fmt.Sprintf(invalidYearErrDetail, currentYear)}.NewError()
	}
	return true, nil
}
