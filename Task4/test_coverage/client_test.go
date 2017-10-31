package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClientCorrectRequestWithLimit(t *testing.T) {
	expectedResult := &SearchResponse{
		Users: []User{
			{
				Id:     0,
				Name:   "Boyd",
				Age:    22,
				About:  "Nulla cillum enim voluptate consequat laborum esse excepteur occaecat commodo nostrud excepteur ut cupidatat. Occaecat minim incididunt ut proident ad sint nostrud ad laborum sint pariatur. Ut nulla commodo dolore officia. Consequat anim eiusmod amet commodo eiusmod deserunt culpa. Ea sit dolore nostrud cillum proident nisi mollit est Lorem pariatur. Lorem aute officia deserunt dolor nisi aliqua consequat nulla nostrud ipsum irure id deserunt dolore. Minim reprehenderit nulla exercitation labore ipsum.\n",
				Gender: "male",
			},
			{
				Id:     1,
				Name:   "Hilda",
				Age:    21,
				About:  "Sit commodo consectetur minim amet ex. Elit aute mollit fugiat labore sint ipsum dolor cupidatat qui reprehenderit. Eu nisi in exercitation culpa sint aliqua nulla nulla proident eu. Nisi reprehenderit anim cupidatat dolor incididunt laboris mollit magna commodo ex. Cupidatat sit id aliqua amet nisi et voluptate voluptate commodo ex eiusmod et nulla velit.\n",
				Gender: "female",
			},
		},
		NextPage: true,
	}

	request := SearchRequest{
		Limit:      2,
		Offset:     0,
		Query:      "",
		OrderField: "",
		OrderBy:    0,
	}

	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if err != nil {
		t.Error("Unexpected error")
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Wrong Response: got %+v, want %+v", result, expectedResult)
	}
}

func TestClientCorrectRequestWithLimitAndOffset(t *testing.T) {
	expectedResult := &SearchResponse{
		Users: []User{
			{
				Id:     1,
				Name:   "Hilda",
				Age:    21,
				About:  "Sit commodo consectetur minim amet ex. Elit aute mollit fugiat labore sint ipsum dolor cupidatat qui reprehenderit. Eu nisi in exercitation culpa sint aliqua nulla nulla proident eu. Nisi reprehenderit anim cupidatat dolor incididunt laboris mollit magna commodo ex. Cupidatat sit id aliqua amet nisi et voluptate voluptate commodo ex eiusmod et nulla velit.\n",
				Gender: "female",
			},
		},
		NextPage: true,
	}

	request := SearchRequest{
		Limit:      1,
		Offset:     1,
		Query:      "",
		OrderField: "",
		OrderBy:    0,
	}

	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if err != nil {
		t.Error("Unexpected error")
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Wrong Response: got %+v, want %+v", result, expectedResult)
	}
}

func TestClientCorrectRequestWithBigLimit(t *testing.T) {
	request := SearchRequest{
		Limit:      100,
		Offset:     0,
		Query:      "",
		OrderField: "",
		OrderBy:    0,
	}

	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if err != nil {
		t.Error("Unexpected error")
	}

	if len(result.Users) != 25 {
		t.Errorf("Wrong number of users: got %+v, want %+v", len(result.Users), 25)
	}
}

func TestClientIncorrectOffset(t *testing.T) {
	request := SearchRequest{
		Limit:      1,
		Offset:     -1,
		Query:      "",
		OrderField: "",
		OrderBy:    0,
	}

	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if result != nil || err == nil {
		t.Errorf("Wrong Response: expected Error")
	}
}

func TestClientIncorrectLimit(t *testing.T) {
	request := SearchRequest{
		Limit:      -1,
		Offset:     1,
		Query:      "",
		OrderField: "",
		OrderBy:    0,
	}

	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if result != nil || err == nil {
		t.Errorf("Wrong Response: expected Error")
	}
}

func TestClientUnathorized(t *testing.T) {
	request := SearchRequest{
		Limit:      1,
		Offset:     1,
		Query:      "",
		OrderField: "",
		OrderBy:    0,
	}

	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	client := SearchClient{
		AccessToken: "",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if result != nil || err == nil {
		t.Errorf("Wrong Response: expected Error")
	}
}

func TestClientIncorrectOrderField(t *testing.T) {
	request := SearchRequest{
		Limit:      1,
		Offset:     1,
		Query:      "",
		OrderField: "incorrectField",
		OrderBy:    0,
	}

	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if result != nil || err == nil {
		t.Errorf("Wrong Response: expected Error")
	}
}

func TestClientIncorrectOrderUrl(t *testing.T) {
	request := SearchRequest{
		Limit:      1,
		Offset:     1,
		Query:      "",
		OrderField: "incorrectField",
		OrderBy:    0,
	}

	httptest.NewServer(http.HandlerFunc(SearchServer))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         "",
	}

	result, err := client.FindUsers(request)

	if result != nil || err == nil {
		t.Errorf("Wrong Response: expected Error")
	}
}

func TestClientTimeout(t *testing.T) {
	request := SearchRequest{
		Limit:      1,
		Offset:     1,
		Query:      "",
		OrderField: "incorrectField",
		OrderBy:    0,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for {
			fmt.Print("")
		}
	}))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if result != nil || err == nil {
		t.Errorf("Wrong Response: expected Error")
	}
}

func TestClientInternalServerErr(t *testing.T) {
	request := SearchRequest{
		Limit:      1,
		Offset:     1,
		Query:      "",
		OrderField: "incorrectField",
		OrderBy:    0,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RespondJSON(w, http.StatusInternalServerError, SearchErrorResponse{})
	}))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if result != nil || err == nil {
		t.Errorf("Wrong Response: expected Error")
	}
}

func TestClientIncorrectErrJSONFormat(t *testing.T) {
	request := SearchRequest{
		Limit:      1,
		Offset:     1,
		Query:      "",
		OrderField: "incorrectField",
		OrderBy:    0,
	}

	type IncorrectErrorResponse struct {
		Error int
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RespondJSON(w, http.StatusBadRequest, IncorrectErrorResponse{Error: 1})
	}))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if result != nil || err == nil {
		t.Errorf("Wrong Response: expected Error")
	}
}

func TestClientUnknownBadRequestErr(t *testing.T) {
	request := SearchRequest{
		Limit:      1,
		Offset:     1,
		Query:      "",
		OrderField: "incorrectField",
		OrderBy:    0,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RespondJSON(w, http.StatusBadRequest, SearchErrorResponse{})
	}))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if result != nil || err == nil {
		t.Errorf("Wrong Response: expected Error")
	}
}

func TestClientIncorrectResultJSONFormat(t *testing.T) {
	request := SearchRequest{
		Limit:      1,
		Offset:     1,
		Query:      "",
		OrderField: "incorrectField",
		OrderBy:    0,
	}

	type IncorrectErrorResponse struct {
		Error int
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RespondJSON(w, http.StatusOK, IncorrectErrorResponse{Error: 1})
	}))

	client := SearchClient{
		AccessToken: "d41d8cd98f00b204e9800998ecf8427e",
		URL:         ts.URL,
	}

	result, err := client.FindUsers(request)

	if result != nil || err == nil {
		t.Errorf("Wrong Response: expected Error")
	}
}

