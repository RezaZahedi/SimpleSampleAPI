package main

import (
	"sync"
	"testing"
)

var rwmutex = sync.RWMutex{}
var bF *BookFetcher

func (bf *BookFetcher) Reset() *BookFetcher {
	bf = NewBookFetcher(map[int]Book{1: {1, "Revelation", "des R"}}, &rwmutex)
	return bf
}
func TestReadBooks(t *testing.T) {
	bf := bF.Reset()
	want := `{"1":{"id":1,"name":"Revelation","desc":"des R"}}`

	t.Log("Given the need to test the ReadBooks function.")
	{
		got, err := ReadBooks(bf)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to marshal to json: %v", Failed, err)
		} else {
			t.Logf("\t%s\tShould be able to marshal to json", Succeed)
		}

		if string(got) != want {
			t.Errorf("\t%s\tShould be able to return the correct json output.", Failed)
		} else {
			t.Logf("\t%s\tShould be able to return the correct json output.", Succeed)
		}
		t.Log(" \t\t", "want:", want)
		t.Log(" \t\t", "got :", string(got))
	}
}

func TestReadBook(t *testing.T) {
	bf := bF.Reset()
	tests := []struct {
		id      int
		want    string
		wantErr error
	}{
		{1, `{"id":1,"name":"Revelation","desc":"des R"}`, nil},
		{10, ``, &BookError{"Book Not Found"}},
	}

	t.Log("Given the need to test the ReadBook function.")

	for i, tt := range tests {
		t.Logf("\tTest: %d\tWhen checking book id = <%d> for output of <%s> and error of <%v>",
			i, tt.id, tt.want, tt.wantErr)
		{
			got, err := ReadBook(bf, tt.id)
			if _, ok := err.(*BookError); ok {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("\t%s\tShould be able to return the correct error", Failed)
				} else {
					t.Logf("\t%s\tShould be able to return the correct error", Succeed)
				}
			} else if err != tt.wantErr {
				t.Errorf("\t%s\tShould be able to return the correct error", Failed)
			} else {
				t.Logf("\t%s\tShould be able to return the correct error", Succeed)
			}
			t.Log(" \t\t", "want:", tt.wantErr)
			t.Log(" \t\t", "got :", err)

			if string(got) != tt.want {
				t.Errorf("\t%s\tShould be able to return the correct json output.", Failed)
			} else {
				t.Logf("\t%s\tShould be able to return the correct json output.", Succeed)
			}
			t.Log(" \t\t", "want:", tt.want)
			t.Log(" \t\t", "got :", string(got))
		}
	}
}

func TestCreateBook(t *testing.T) {
	bf := bF.Reset()
	tests := []struct {
		id      int
		book    Book
		wantErr error
	}{
		{2, Book{2, "Reza", "desR"}, nil},
		{1, Book{1, "Reza,", "desR"}, &BookError{"Bad Request: Book Already Exists"}},
		{2, Book{3, "Reza", "desR"}, &BookError{"Bad Request: Missmatching IDs"}},
	}

	t.Log("Given the need to test the CreateBook function.")
	for i, tt := range tests {
		t.Logf("\tTest: %d\tWhen Creating book id = <%d> with values of <%+v> and error of <%v>",
			i, tt.id, tt.book, tt.wantErr)
		{
			err := CreateBook(bf, tt.id, tt.book)
			if _, ok := err.(*BookError); ok {
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
			bf = bf.Reset()
		}
	}
}

func TestDeleteBook(t *testing.T) {
	bf := bF.Reset()
	tests := []struct {
		id      int
		wantErr error
	}{
		{1, nil},
		{2, &BookError{"Bad Request: Book Not Found"}},
	}
	t.Log("Given the need to test the DeleteBook function.")
	{
		for i, tt := range tests {
			t.Logf("\tTest: %d\tWhen Deleting book id = <%d> for the error of <%v>",
				i, tt.id, tt.wantErr)
			{
				err := DeleteBook(bf, tt.id)
				if _, ok := err.(*BookError); ok {
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
				bf = bf.Reset()
			}
		}
	}
}

func TestQueryBook(t *testing.T) {
	bf := bF.Reset()
	tests := []struct {
		name    string
		want    string
		wantErr error
	}{
		{"Revelation", `[{"id":1,"name":"Revelation","desc":"des R"}]`, nil},
		{"Reza", ``, &BookError{"Found No Book Named: Reza"}},
	}
	t.Log("Given the need to test the QueryBook function.")
	{
		for i, tt := range tests {
			t.Logf("\tTest: %d\tWhen querying books with name = <%s> for output of <%s> and error of <%v>",
				i, tt.name, tt.want, tt.wantErr)
			{
				got, err := QueryBook(bf, tt.name)
				if _, ok := err.(*BookError); ok {
					if err.Error() != tt.wantErr.Error() {
						t.Errorf("\t%s\tShould be able to return the correct error", Failed)
					} else {
						t.Logf("\t%s\tShould be able to return the correct error", Succeed)
					}
				} else if err != tt.wantErr {
					t.Errorf("\t%s\tShould be able to return the correct error", Failed)
				} else {
					t.Logf("\t%s\tShould be able to return the correct error", Succeed)
				}
				t.Log(" \t\t", "want:", tt.wantErr)
				t.Log(" \t\t", "got :", err)

				if string(got) != tt.want {
					t.Errorf("\t%s\tShould be able to return the correct json output.", Failed)
				} else {
					t.Logf("\t%s\tShould be able to return the correct json output.", Succeed)
				}
				t.Log(" \t\t", "want:", tt.want)
				t.Log(" \t\t", "got :", string(got))
			}
		}
	}
}
