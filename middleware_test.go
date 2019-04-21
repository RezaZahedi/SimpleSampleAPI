package main

import (
	"errors"
	"testing"
)

// I couldn't still figure out how to test middleware stuff.
// logging can't be tested as it uses time.Now() function.
func TestLogging(t *testing.T) {

}

// I couldn't still figure out how to test middleware stuff.
func TestSplitting(t *testing.T) {

}

// I couldn't still figure out how to test middleware stuff.
func TestChain(t *testing.T) {

}

func TestUrlParser(t *testing.T) {
	tests := []struct {
		url     string
		totLen  int
		index   int
		isInt   bool
		want    int
		wantErr error
	}{
		{"/book/2", 3, 3, true, 2, nil},
		{"/book/2/another", 3, 3, true, 0, errors.New("Bad URL")},
		{"/reza/3", 3, 4, true, 0, errors.New("Bad URL")},
		{"/rez/za/3/adf/1", 6, 4, true, 3, nil},
		{"/rez/za/3/adf/1", 6, 5, true, 0, errors.New("Bad URL")},
	}
	t.Log("Given the need to test the urlParser function.")
	for i, tt := range tests {
		t.Logf("\tTest: %d\tWhen parsing URL: <%s> with supposed length of <%d> and wanting the value at index <%d>",
			i, tt.url, tt.totLen, tt.index)
		{
			id, err := urlParser(tt.url, tt.totLen, tt.index, tt.isInt)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("\t%s\tShould be able to return the correct error", Failed)
				} else {
					t.Logf("\t%s\tShould be able to return the correct error", Succeed)
				}
			} else {
				t.Logf("\t%s\tShould be able to return the correct error", Succeed)
			}
			t.Log(" \t\t", "want:", tt.wantErr)
			t.Log(" \t\t", "got :", err)

			if id.(int) != tt.want {
				t.Errorf("\t%s\tShould be able to return the correct value from the url", Failed)
			} else {
				t.Logf("\t%s\tShould be able to return the correct value from the url", Succeed)
			}
			t.Log(" \t\t", "want:", tt.want)
			t.Log(" \t\t", "got :", id)
		}
	}
}
