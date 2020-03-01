package utils

import (
	"log"
	"testing"
)

const (
	testExecErr  = "There was an error testing the function : %v"
	errExpMsg    = "Expected an error, but got nil"
	errNotExpMsg = "Did not expect an error, but got %v"
	testUser     = "testUser"
	testPassword = "test123"
)

func TestBasicAuth_Check(t *testing.T) {

	ba := BasicAuth{}

	r := Request{
		Url:    "http://localhost",
		Method: "GET",
		Auth: Auth{
			Username: "",
			Password: "",
		},
	}

	if err := r.NewRequest(); err != nil {
		log.Fatalf(testExecErr, err)
	}

	if err := ba.Check(r.Request); err == nil {
		t.Error(errExpMsg)
	}

	r = Request{
		Url:    "http://localhost",
		Method: "GET",
		Auth: Auth{
			Username: testUser,
			Password: testPassword,
		},
	}

	if err := r.NewRequest(); err != nil {
		log.Fatalf(testExecErr, err)
	}

	if err := ba.Check(r.Request); err != nil {
		t.Errorf(errNotExpMsg, err)
	}

}

func TestBasicAuth_Get(t *testing.T) {

	ba := BasicAuth{}

	r := Request{
		Url:    "http://localhost",
		Method: "GET",
	}

	if err := r.NewRequest(); err != nil {
		log.Fatalf(testExecErr, err)
	}

	// check basic auth
	if err := ba.Get(r.Request); err == nil {
		t.Error("Test no basic auth : ", errExpMsg)
	}

	// check base64 decode error
	r.Request.Header.Set("Authorization", "Basic dGVzdDp=0ZXN0OnRlc3Q=")
	if err := ba.Get(r.Request); err == nil {
		t.Error("Test base64 decode error : ", errExpMsg)
	}

	// test01 incorrect basic auth
	r.Request.Header.Set("Authorization", "dGVzdA==")
	if err := ba.Get(r.Request); err == nil {
		t.Error("Test01 incorrect basic auth : ", errExpMsg)
	}

	// test02 incorrect basic auth
	r.Request.Header.Set("Authorization", "Basic dGVzdA==")
	if err := ba.Get(r.Request); err == nil {
		t.Error("Test02 incorrect basic auth : ", errExpMsg)
	}

	r.Request.SetBasicAuth(testUser, testPassword)

	if err := ba.Get(r.Request); err != nil {
		log.Fatalf(testExecErr, err)
	}

	if ba.Username != testUser || ba.Password != testPassword {
		t.Errorf("Test Failed!, expected: %v, got: %v", BasicAuth{testUser, testPassword}, ba)
	}
}
