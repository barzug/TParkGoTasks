package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestServerWithoutAccessToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()

	SearchServer(w, req)

	if status := w.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestServerIncorrectAccessTokezn(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "qwerty")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestServerIncorrectRequestParamLimit(t *testing.T) {

	searcherParams := url.Values{}

	searcherParams.Add("limit", "string")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestServerIncorrectRequestParamOffset(t *testing.T) {

	searcherParams := url.Values{}

	searcherParams.Add("offset", "string")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestServerIncorrectRequestParamOrderField(t *testing.T) {

	searcherParams := url.Values{}

	searcherParams.Add("order_field", "noNameOrIdOrAge")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestServerIncorrectFilePath(t *testing.T) {
	searcherParams := url.Values{}

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	filepath = "IncorrectFilePath"

	SearchServer(w, req)

	filepath = "dataset.xml"

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestServerIncorrectFileData(t *testing.T) {
	searcherParams := url.Values{}

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	filepath = "server_test.go"

	SearchServer(w, req)

	filepath = "dataset.xml"

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestServerIncorrectRequestParamOrderBy(t *testing.T) {

	searcherParams := url.Values{}

	searcherParams.Add("order_by", "string")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestServerCorrectRequestWithLimit(t *testing.T) {
	response := `[{"Id":0,"Name":"Boyd","Age":22,"About":"Nulla cillum enim voluptate consequat laborum esse excepteur occaecat commodo nostrud excepteur ut cupidatat. Occaecat minim incididunt ut proident ad sint nostrud ad laborum sint pariatur. Ut nulla commodo dolore officia. Consequat anim eiusmod amet commodo eiusmod deserunt culpa. Ea sit dolore nostrud cillum proident nisi mollit est Lorem pariatur. Lorem aute officia deserunt dolor nisi aliqua consequat nulla nostrud ipsum irure id deserunt dolore. Minim reprehenderit nulla exercitation labore ipsum.\n","Gender":"male"},{"Id":1,"Name":"Hilda","Age":21,"About":"Sit commodo consectetur minim amet ex. Elit aute mollit fugiat labore sint ipsum dolor cupidatat qui reprehenderit. Eu nisi in exercitation culpa sint aliqua nulla nulla proident eu. Nisi reprehenderit anim cupidatat dolor incididunt laboris mollit magna commodo ex. Cupidatat sit id aliqua amet nisi et voluptate voluptate commodo ex eiusmod et nulla velit.\n","Gender":"female"}]`

	searcherParams := url.Values{}

	searcherParams.Add("limit", "2")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)
	if strings.Compare(bodyStr, response) != 0 {
		t.Errorf("Wrong Response: resp: %+v want %+v", response, bodyStr)
	}
}

func TestServerCorrectRequestWithSearch(t *testing.T) {
	response := `[{"Id":34,"Name":"Kane","Age":34,"About":"Lorem proident sint minim anim commodo cillum. Eiusmod velit culpa commodo anim consectetur consectetur sint sint labore. Mollit consequat consectetur magna nulla veniam commodo eu ut et. Ut adipisicing qui ex consectetur officia sint ut fugiat ex velit cupidatat fugiat nisi non. Dolor minim mollit aliquip veniam nostrud. Magna eu aliqua Lorem aliquip.\n","Gender":"male"}]`

	searcherParams := url.Values{}

	searcherParams.Add("limit", "2")
	searcherParams.Add("query", "Kane")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)
	if strings.Compare(bodyStr, response) != 0 {
		t.Errorf("Wrong Response: got %+v, want %+v", response, bodyStr)
	}
}

func TestServerCorrectRequestWithOrderFieldByNameAsc(t *testing.T) {
	response := `[{"Id":15,"Name":"Allison","Age":21,"About":"Labore excepteur voluptate velit occaecat est nisi minim. Laborum ea et irure nostrud enim sit incididunt reprehenderit id est nostrud eu. Ullamco sint nisi voluptate cillum nostrud aliquip et minim. Enim duis esse do aute qui officia ipsum ut occaecat deserunt. Pariatur pariatur nisi do ad dolore reprehenderit et et enim esse dolor qui. Excepteur ullamco adipisicing qui adipisicing tempor minim aliquip.\n","Gender":"male"},{"Id":16,"Name":"Annie","Age":35,"About":"Consequat fugiat veniam commodo nisi nostrud culpa pariatur. Aliquip velit adipisicing dolor et nostrud. Eu nostrud officia velit eiusmod ullamco duis eiusmod ad non do quis.\n","Gender":"female"},{"Id":19,"Name":"Bell","Age":26,"About":"Nulla voluptate nostrud nostrud do ut tempor et quis non aliqua cillum in duis. Sit ipsum sit ut non proident exercitation. Quis consequat laboris deserunt adipisicing eiusmod non cillum magna.\n","Gender":"male"}]`

	searcherParams := url.Values{}

	searcherParams.Add("limit", "3")
	searcherParams.Add("order_field", "name")
	searcherParams.Add("order_by", "1")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)
	if strings.Compare(bodyStr, response) != 0 {
		t.Errorf("Wrong Response: got %+v, want %+v", response, bodyStr)
	}
}

func TestServerCorrectRequestWithOrderFieldByNameDesc(t *testing.T) {
	response := `[{"Id":13,"Name":"Whitley","Age":40,"About":"Consectetur dolore anim veniam aliqua deserunt officia eu. Et ullamco commodo ad officia duis ex incididunt proident consequat nostrud proident quis tempor. Sunt magna ad excepteur eu sint aliqua eiusmod deserunt proident. Do labore est dolore voluptate ullamco est dolore excepteur magna duis quis. Quis laborum deserunt ipsum velit occaecat est laborum enim aute. Officia dolore sit voluptate quis mollit veniam. Laborum nisi ullamco nisi sit nulla cillum et id nisi.\n","Gender":"male"},{"Id":33,"Name":"Twila","Age":36,"About":"Sint non sunt adipisicing sit laborum cillum magna nisi exercitation. Dolore officia esse dolore officia ea adipisicing amet ea nostrud elit cupidatat laboris. Proident culpa ullamco aute incididunt aute. Laboris et nulla incididunt consequat pariatur enim dolor incididunt adipisicing enim fugiat tempor ullamco. Amet est ullamco officia consectetur cupidatat non sunt laborum nisi in ex. Quis labore quis ipsum est nisi ex officia reprehenderit ad adipisicing fugiat. Labore fugiat ea dolore exercitation sint duis aliqua.\n","Gender":"female"}]`

	searcherParams := url.Values{}

	searcherParams.Add("limit", "2")
	searcherParams.Add("order_field", "name")
	searcherParams.Add("order_by", "-1")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)
	if strings.Compare(bodyStr, response) != 0 {
		t.Errorf("Wrong Response: got %+v, want %+v", response, bodyStr)
	}
}

func TestServerCorrectRequestWithOrderFieldByIdAsk(t *testing.T) {
	response := `[{"Id":0,"Name":"Boyd","Age":22,"About":"Nulla cillum enim voluptate consequat laborum esse excepteur occaecat commodo nostrud excepteur ut cupidatat. Occaecat minim incididunt ut proident ad sint nostrud ad laborum sint pariatur. Ut nulla commodo dolore officia. Consequat anim eiusmod amet commodo eiusmod deserunt culpa. Ea sit dolore nostrud cillum proident nisi mollit est Lorem pariatur. Lorem aute officia deserunt dolor nisi aliqua consequat nulla nostrud ipsum irure id deserunt dolore. Minim reprehenderit nulla exercitation labore ipsum.\n","Gender":"male"},{"Id":1,"Name":"Hilda","Age":21,"About":"Sit commodo consectetur minim amet ex. Elit aute mollit fugiat labore sint ipsum dolor cupidatat qui reprehenderit. Eu nisi in exercitation culpa sint aliqua nulla nulla proident eu. Nisi reprehenderit anim cupidatat dolor incididunt laboris mollit magna commodo ex. Cupidatat sit id aliqua amet nisi et voluptate voluptate commodo ex eiusmod et nulla velit.\n","Gender":"female"}]`

	searcherParams := url.Values{}

	searcherParams.Add("limit", "2")
	searcherParams.Add("order_field", "id")
	searcherParams.Add("order_by", "1")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)
	if strings.Compare(bodyStr, response) != 0 {
		t.Errorf("Wrong Response: got %+v, want %+v", response, bodyStr)
	}
}

func TestServerCorrectRequestWithOrderFieldByIdDesk(t *testing.T) {
	response := `[{"Id":34,"Name":"Kane","Age":34,"About":"Lorem proident sint minim anim commodo cillum. Eiusmod velit culpa commodo anim consectetur consectetur sint sint labore. Mollit consequat consectetur magna nulla veniam commodo eu ut et. Ut adipisicing qui ex consectetur officia sint ut fugiat ex velit cupidatat fugiat nisi non. Dolor minim mollit aliquip veniam nostrud. Magna eu aliqua Lorem aliquip.\n","Gender":"male"},{"Id":33,"Name":"Twila","Age":36,"About":"Sint non sunt adipisicing sit laborum cillum magna nisi exercitation. Dolore officia esse dolore officia ea adipisicing amet ea nostrud elit cupidatat laboris. Proident culpa ullamco aute incididunt aute. Laboris et nulla incididunt consequat pariatur enim dolor incididunt adipisicing enim fugiat tempor ullamco. Amet est ullamco officia consectetur cupidatat non sunt laborum nisi in ex. Quis labore quis ipsum est nisi ex officia reprehenderit ad adipisicing fugiat. Labore fugiat ea dolore exercitation sint duis aliqua.\n","Gender":"female"}]`

	searcherParams := url.Values{}

	searcherParams.Add("limit", "2")
	searcherParams.Add("order_field", "id")
	searcherParams.Add("order_by", "-1")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)
	if strings.Compare(bodyStr, response) != 0 {
		t.Errorf("Wrong Response: got %+v, want %+v", response, bodyStr)
	}
}

func TestServerCorrectRequestWithOrderFieldByAgeAsk(t *testing.T) {
	response := `[{"Id":1,"Name":"Hilda","Age":21,"About":"Sit commodo consectetur minim amet ex. Elit aute mollit fugiat labore sint ipsum dolor cupidatat qui reprehenderit. Eu nisi in exercitation culpa sint aliqua nulla nulla proident eu. Nisi reprehenderit anim cupidatat dolor incididunt laboris mollit magna commodo ex. Cupidatat sit id aliqua amet nisi et voluptate voluptate commodo ex eiusmod et nulla velit.\n","Gender":"female"},{"Id":15,"Name":"Allison","Age":21,"About":"Labore excepteur voluptate velit occaecat est nisi minim. Laborum ea et irure nostrud enim sit incididunt reprehenderit id est nostrud eu. Ullamco sint nisi voluptate cillum nostrud aliquip et minim. Enim duis esse do aute qui officia ipsum ut occaecat deserunt. Pariatur pariatur nisi do ad dolore reprehenderit et et enim esse dolor qui. Excepteur ullamco adipisicing qui adipisicing tempor minim aliquip.\n","Gender":"male"}]`

	searcherParams := url.Values{}

	searcherParams.Add("limit", "2")
	searcherParams.Add("order_field", "age")
	searcherParams.Add("order_by", "1")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)
	if strings.Compare(bodyStr, response) != 0 {
		t.Errorf("Wrong Response: got %+v, want %+v", response, bodyStr)
	}
}

func TestServerCorrectRequestWithOrderFieldByAgeDesk(t *testing.T) {
	response := `[{"Id":32,"Name":"Christy","Age":40,"About":"Incididunt culpa dolore laborum cupidatat consequat. Aliquip cupidatat pariatur sit consectetur laboris labore anim labore. Est sint ut ipsum dolor ipsum nisi tempor in tempor aliqua. Aliquip labore cillum est consequat anim officia non reprehenderit ex duis elit. Amet aliqua eu ad velit incididunt ad ut magna. Culpa dolore qui anim consequat commodo aute.\n","Gender":"female"},{"Id":13,"Name":"Whitley","Age":40,"About":"Consectetur dolore anim veniam aliqua deserunt officia eu. Et ullamco commodo ad officia duis ex incididunt proident consequat nostrud proident quis tempor. Sunt magna ad excepteur eu sint aliqua eiusmod deserunt proident. Do labore est dolore voluptate ullamco est dolore excepteur magna duis quis. Quis laborum deserunt ipsum velit occaecat est laborum enim aute. Officia dolore sit voluptate quis mollit veniam. Laborum nisi ullamco nisi sit nulla cillum et id nisi.\n","Gender":"male"}]`

	searcherParams := url.Values{}

	searcherParams.Add("limit", "2")
	searcherParams.Add("order_field", "age")
	searcherParams.Add("order_by", "-1")

	req := httptest.NewRequest("GET", "/?"+searcherParams.Encode(), nil)

	w := httptest.NewRecorder()

	req.Header.Add("AccessToken", "d41d8cd98f00b204e9800998ecf8427e")

	SearchServer(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	bodyStr := string(body)
	if strings.Compare(bodyStr, response) != 0 {
		t.Errorf("Wrong Response: got %+v, want %+v", response, bodyStr)
	}
}

func TestRespondJsonErrorData(t *testing.T) {
	w := httptest.NewRecorder()

	RespondJSON(w, http.StatusOK, make(chan int))

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("SearchServer returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
