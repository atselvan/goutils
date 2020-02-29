package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	fileNotFoundErrMsg              = "File not found"
	fileNotFoundErrDetail           = "the file %s was not found"
	fileReadErrMsg                  = "Unable to read file"
	fileCreateErrMsg                = "Unable to create a new file"
	fileOpenErrMsg                  = "Unable to open file"
	fileWriteErrMsg                 = "Unable to write to the file"
	JsonMarshalErrMsg               = "JSON Marshal Error"
	JsonUnmarshalErrMsg             = "JSON Unmarshal Error"
	YamlMarshalErrMsg               = "YAML Marshal Error"
	YamlUnmarshalErrMsg             = "YAML Unmarshal Error"
)

// ReadYamlFile reads a yaml file and puts the contents into the out variables
// out variable should be a pointer to a valid struct
// The method returns and error if reading a file or the unmarshal process fails
func ReadYamlFile(filePath string, out interface{}) error {
	data, err := ReadFile(filePath)
	if err != nil 	{
		return err
	}
	err = yaml.Unmarshal(data, out)
	if err != nil {
		return Error{Message: YamlUnmarshalErrMsg, Detail: err.Error()}.NewError()
	}
	return err
}

// WriteYamlFile encodes the data from a input interface into yaml format
// and writes the data into a file
// The in interface should be an address to a valid struct
// The method returns an error if there is an error with the yaml encode
// or with writing to the file
func WriteYamlFile(filePath string, in interface{}) error {
	data, err := yaml.Marshal(in)
	if err != nil {
		return Error{Message: YamlMarshalErrMsg, Detail: err.Error()}.NewError()
	}
	err = WriteFile(filePath, data)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// CreateFile creates a new file
// The method returns an error if there was an issue with creating an new file
func CreateFile(filePath string) (*os.File, error) {
	f, err := os.Create(filePath)
	if err != nil {
		return f, Error{Message: fileCreateErrMsg, Detail: err.Error()}.NewError()
	} else {
		return f, nil
	}
}

// OpenFile opens a file
// The method returns an error if there is an issue with opening the file
func OpenFile(filePath string) (*os.File, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return f, Error{Message: fileOpenErrMsg, Detail: err.Error()}.NewError()
	} else {
		return f, nil
	}
}

// ReadFile checks if a file exists and if it does tries to reads the contents of the
// file and returns the data back
// The method returns an error the file does not exist or if there was an error in reading the contents of the file
func ReadFile(filePath string) ([]byte, error) {
	if !FileExists(filePath) {
		return nil, Error{Message: fileNotFoundErrMsg, Detail: fmt.Sprintf(fileNotFoundErrDetail, filePath)}.NewError()
	} else if data, err := ioutil.ReadFile(filePath); err != nil {
		return nil, Error{Message: fileReadErrMsg, Detail: err.Error()}.NewError()
	} else {
		return data, err
	}
}

// WriteFile creates a new file if the file does not exists and writes data into the file
// The method returns an error if there was an issue creating a new file
// or while writing data into the file
func WriteFile(filePath string, data []byte) error {
	var (
		err error
	)

	if !FileExists(filePath) {
		_, err = CreateFile(filePath)
		if err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return Error{Message: fileWriteErrMsg, Detail: err.Error()}.NewError()
	}

	return nil
}

// FileExists checks if a file exists and returns an error if the file was not found
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	} else {
		return true
	}
}
