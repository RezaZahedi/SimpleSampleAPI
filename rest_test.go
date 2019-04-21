package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	bf := bF.Reset()
	r := httptest.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()

	bf.ReadAll(w, r)

	wantStatus := http.StatusOK
	t.Log("Given the need to test the ReadAll endpoint.")
	{
		if w.Code != wantStatus {
			t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d]", Failed, wantStatus, w.Code)
		}
		t.Logf("\t%s\tShould receive a status code of %d for the response.", Succeed, wantStatus)

		if wantStatus == http.StatusOK {
			books := make(map[int]Book)

			if err := json.NewDecoder(w.Body).Decode(&books); err != nil {
				t.Fatalf("\t%s\tShould be able to decode the response.", Failed)
			}
			t.Logf("\t%s\tShould be able to decode the response.", Succeed)

			wantName := "Revelation"
			if books[1].Name == wantName {
				t.Logf("\t%s\tShould have \"%s\" for Name in the response.", Succeed, wantName)
			} else {
				t.Errorf("\t%s\tShould have \"%s\" for Name in the response.", Failed, wantName)
			}
			t.Log(" \t\t", "want:", wantName)
			t.Log(" \t\t", "got :", books[1].Name)
		}
	}
}

func TestMethodMux(t *testing.T) {
	bf := bF.Reset()
	tests := []struct {
		id         int
		method     string
		target     string
		wantStatus int
	}{
		{1, "GET", "/book/1", http.StatusOK},
		{2, "POST", "/book/2", http.StatusInternalServerError},
		{1, "DELETE", "/book/1", http.StatusSeeOther},
		{3, "PATCH", "/book/3", http.StatusMethodNotAllowed},
	}
	t.Log("Given the need to test the MethodMux handler.")
	for i, tt := range tests {
		t.Logf("\tTest: %d\tWhen requesting with method: <%s> to target: <%s> for status code of <%d>", i, tt.method, tt.target, tt.wantStatus)
		{
			r := httptest.NewRequest(tt.method, tt.target, nil)
			ctx := context.WithValue(r.Context(), ContextKeyID, tt.id)
			w := httptest.NewRecorder()

			bf.MethodMux(w, r.WithContext(ctx))
			if w.Code != tt.wantStatus {
				t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d]", Failed, tt.wantStatus, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code of %d for the response.", Succeed, tt.wantStatus)
		}
	}
}

func TestReadOne(t *testing.T) {
	bf := bF.Reset()
	tests := []struct {
		id         int
		target     string
		wantStatus int
	}{
		{1, "/book/1", http.StatusOK},
		{2, "/book/2", http.StatusNotFound},
	}
	t.Log("Given the need to test the ReadOne endpoint.")

	for i, tt := range tests {
		t.Logf("\tTest: %d\tWhen requesting book id = <%d> for status code of <%d>", i, tt.id, tt.wantStatus)
		{
			r := httptest.NewRequest("GET", tt.target, nil)
			ctx := context.WithValue(r.Context(), ContextKeyID, tt.id)
			w := httptest.NewRecorder()

			bf.ReadOne(w, r.WithContext(ctx))
			if w.Code != tt.wantStatus {
				t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d]", Failed, tt.wantStatus, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code of %d for the response.", Succeed, tt.wantStatus)

		}
	}
}

func TestCreateOne(t *testing.T) {
	bf := bF.Reset()
	tests := []struct {
		id         int
		target     string
		body       string
		wantStatus int
	}{
		{2, "/book/2", `{"Id":2, "Name":"Reza"}`, http.StatusSeeOther},
		{1, "/book/1", `{"Id":1, "Name":"Reza"}`, http.StatusBadRequest},
		{1, "/book/2", `{"Id":1, "Name":"Reza"}`, http.StatusBadRequest},
	}
	t.Log("Given the need to test the CreateOne endpoint.")

	for i, tt := range tests {
		t.Logf("\tTest: %d\tWhen requesting the creation of book with id = <%d> for status code of <%d>", i, tt.id, tt.wantStatus)
		{
			r := httptest.NewRequest("POST", tt.target, strings.NewReader(tt.body))
			ctx := context.WithValue(r.Context(), ContextKeyID, tt.id)
			w := httptest.NewRecorder()

			bf.CreateOne(w, r.WithContext(ctx))
			if w.Code != tt.wantStatus {
				t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d]", Failed, tt.wantStatus, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code of %d for the response.", Succeed, tt.wantStatus)
		}
		bf = bf.Reset()
	}

}

func TestDeleteOne(t *testing.T) {
	bf := bF.Reset()
	tests := []struct {
		id         int
		target     string
		wantStatus int
	}{
		{1, "/book/1", http.StatusSeeOther},
		{2, "/book/2", http.StatusBadRequest},
	}
	t.Log("Given the need to test the DeleteOne endpoint.")

	for i, tt := range tests {
		t.Logf("\tTest: %d\tWhen requesting the deletion of book with id = <%d> for status code of <%d>", i, tt.id, tt.wantStatus)
		{
			r := httptest.NewRequest("DELETE", tt.target, nil)
			ctx := context.WithValue(r.Context(), ContextKeyID, tt.id)
			w := httptest.NewRecorder()

			bf.DeleteOne(w, r.WithContext(ctx))
			if w.Code != tt.wantStatus {
				t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d]", Failed, tt.wantStatus, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code of %d for the response.", Succeed, tt.wantStatus)
		}
		bf = bf.Reset()
	}

}

func TestQuery(t *testing.T) {
	bf := bF.Reset()
	tests := []struct {
		name       string
		want       []Book
		wantStatus int
	}{
		{"Revelation", []Book{{1, "Revelation", "des R"}}, http.StatusOK},
		{"somethingElse", []Book{}, http.StatusNotFound},
	}
	t.Log("Given the need to test the Query endpoint.")

	for i, tt := range tests {
		t.Logf("\tTest: %d\tWhen querying for name = <%s> and expecting status code of <%d>", i, tt.name, tt.wantStatus)
		{
			path := "/book?name="
			r := httptest.NewRequest("GET", path+tt.name, nil)
			w := httptest.NewRecorder()

			bf.Query(w, r)
			if w.Code != tt.wantStatus {
				t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d]", Failed, tt.wantStatus, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code of %d for the response.", Succeed, tt.wantStatus)

			if tt.wantStatus == http.StatusOK {
				books := make([]Book, 0)
				if err := json.NewDecoder(w.Body).Decode(&books); err != nil {
					t.Fatalf("\t%s\tShould be able to decode the response.", Failed)
				}
				t.Logf("\t%s\tShould be able to decode the response.", Succeed)

				if books[0] != tt.want[0] {
					t.Errorf("\t%s\tShould have \"%v\" for the response body.", Failed, tt.want[0])
				} else {
					t.Logf("\t%s\tShould have \"%v\" for the response body.", Succeed, tt.want[0])
				}
				t.Log(" \t\t", "want:", tt.want)
				t.Log(" \t\t", "got :", books)
			}

		}
	}
}
