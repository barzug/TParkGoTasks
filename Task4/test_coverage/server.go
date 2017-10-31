package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	CorrectAccessToken = "d41d8cd98f00b204e9800998ecf8427e"
)

var (
	filepath = "dataset.xml"
)

type XmlData struct {
	Users []User `xml:"row"`
}

func RespondJSON(w http.ResponseWriter, code int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}

func SearchServer(w http.ResponseWriter, r *http.Request) {

	accessToken := r.Header.Get("AccessToken")
	if accessToken != CorrectAccessToken {
		RespondJSON(w, http.StatusUnauthorized, nil)
		return
	}

	searchRequest, errRequestParam := getSearchRequest(r)
	if errRequestParam != "" {
		log.Print(errRequestParam)
		RespondJSON(w, http.StatusBadRequest, SearchErrorResponse{errRequestParam})
		return
	}

	fmt.Println(filepath)
	b, err := readFile(filepath)
	if err != nil {
		fmt.Printf("error: %v", err)
		RespondJSON(w, http.StatusBadRequest, SearchErrorResponse{})
		return
	}

	xmlData := make([]XmlData, 0)
	err = xml.Unmarshal(b, &xmlData)
	if err != nil {
		fmt.Printf("error: %v", err)
		RespondJSON(w, http.StatusBadRequest, SearchErrorResponse{})
		return
	}

	users := make([]User, 0)

	if searchRequest.Query != "" {
		for _, v := range xmlData[0].Users {
			if strings.Contains(v.Name, searchRequest.Query) || strings.Contains(v.About, searchRequest.Query) {
				users = append(users, v)
			}
		}
	} else {
		users = xmlData[0].Users
	}

	if searchRequest.OrderBy != 0 {
		switch searchRequest.OrderField {
		case "id":
			if searchRequest.OrderBy == 1 {
				sort.Sort(ById(users))
			} else if searchRequest.OrderBy == -1 {
				sort.Sort(sort.Reverse(ById(users)))
			}
		case "age":
			if searchRequest.OrderBy == 1 {
				sort.Sort(ByAge(users))
			} else if searchRequest.OrderBy == -1 {
				sort.Sort(sort.Reverse(ByAge(users)))
			}
		case "name":
			if searchRequest.OrderBy == 1 {
				sort.Sort(ByName(users))
			} else if searchRequest.OrderBy == -1 {
				sort.Sort(sort.Reverse(ByName(users)))
			}
		}
	}
	lastUserSearchIndex := searchRequest.Offset + searchRequest.Limit
	if lastUserSearchIndex != 0 {
		if len(users) > lastUserSearchIndex {
			users = users[searchRequest.Offset:lastUserSearchIndex]
		} else {
			users = users[searchRequest.Offset:]
		}
	}
	RespondJSON(w, http.StatusOK, users)
}

func readFile(filepath string) ([]byte, error) {
	xmlDataset, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer xmlDataset.Close()

	b, err := ioutil.ReadAll(xmlDataset)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func getSearchRequest(r *http.Request) (SearchRequest, string) {
	searchRequest := SearchRequest{}

	var err error

	limit := r.URL.Query().Get("limit")
	if limit != "" {
		searchRequest.Limit, err = strconv.Atoi(limit)
		if err != nil {
			return searchRequest, LimitParseError
		}
	}

	offset := r.URL.Query().Get("offset")
	if offset != "" {
		searchRequest.Offset, err = strconv.Atoi(offset)
		if err != nil {
			return searchRequest, OffsetParseError
		}
	}

	searchRequest.Query = r.URL.Query().Get("query")

	searchRequest.OrderField = r.URL.Query().Get("order_field")
	switch searchRequest.OrderField {
	case "id":
	case "age":
	case "name":
	case "":
		searchRequest.OrderField = "name"
	default:
		return searchRequest, ErrorBadOrderField
	}

	order_by := r.URL.Query().Get("order_by")
	if order_by != "" {
		searchRequest.OrderBy, err = strconv.Atoi(order_by)
		if err != nil {
			return searchRequest, OrderByParseError
		}
	}

	return searchRequest, ""
}

type ByAge []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ById []User

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }

type ByName []User

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return strings.Compare(a[i].Name, a[j].Name) < 0 }
