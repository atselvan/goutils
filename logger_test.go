package utils

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestLogger_EnableDebug(t *testing.T) {
	l := Logger{
		Request: nil,
		Message: "",
	}
	if l.EnableDebug() {
		t.Errorf("Expected 'false' but got '%t'", l.EnableDebug())
	}

	os.Setenv("DEBUG_LOGS", "true")

	if !l.EnableDebug() {
		t.Errorf("Expected 'true' but got '%t'", l.EnableDebug())
	}

	os.Setenv("DEBUG_LOGS", "false")

	if l.EnableDebug() {
		t.Errorf("Expected 'true' but got '%t'", l.EnableDebug())
	}

	os.Setenv("DEBUG_LOGS", "something")

	if l.EnableDebug() {
		t.Errorf("Expected 'true' but got '%t'", l.EnableDebug())
	}

}

func TestLogger_Info(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg := "this is a information log"
	Logger{Message: msg}.Info()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if !strings.Contains(string(out), "INFO:") || !strings.Contains(string(out), msg) {
		t.Error("Did not get the expected information log")
	}
}

func TestLogger_Warn(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg := "this is a warning log"
	Logger{Message: msg}.Warn()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if !strings.Contains(string(out), "WARN:") || !strings.Contains(string(out), msg) {
		t.Error("Did not get the expected warning log")
	}
}

func TestLogger_Error(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg := "this is a error log"
	Logger{Message: msg}.Error()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if !strings.Contains(string(out), "ERROR:") || !strings.Contains(string(out), msg) {
		t.Error("Did not get the expected error log")
	}
}

