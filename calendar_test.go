package ical

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestCalendar_Decode(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		title  string
		input  *Calendar
		expect string
	}{
		{
			title: "",
			input: nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			file := hoge(t, tt.expect)
			b := &bytes.Buffer{}
			tt.input.Decode(b)
			if !bytes.Equal(b.Bytes(), file) {
				t.Fatal("result is not expected")
			}
		})
	}
}

func hoge(t *testing.T, path string) []byte {
	t.Helper()
	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	res, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	return res
}
