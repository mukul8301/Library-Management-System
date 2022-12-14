// package category
package user_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mukul1234567/Library-Management-System/user"
	usermock "github.com/mukul1234567/Library-Management-System/user/mocks"
	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	fmt.Println("Expected Code :", expected, "Actual Code : ", actual)
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}

}

func makeHTTPCall(handler http.HandlerFunc, method, path, body string) (rr *httptest.ResponseRecorder) {
	request := []byte(body)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(request))
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return
}

// Create:
func TestSuccessfullCreate(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Success "))

	rr := makeHTTPCall(user.Create(cs), http.MethodPost, "/users", `{
		"id":"67",
		"first_name": "Rushikesh",
		"last_name": "Markad",
		"gender": "Male",
		"age": 22,
		"address": "Pune",
		"email": "markaaa@gmail.com",
		"password": "ertikol@123",
		"mob_no": "6985749623",
		"role":"user"
	}`)
	fmt.Println("Account Success")
	checkResponseCode(t, http.StatusCreated, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenInvalidRequestBody(t *testing.T) {
	cs := &usermock.Service{}
	// cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("HEllo"))
	rr := makeHTTPCall(user.Create(cs), http.MethodPost, "/users", `{
		"id":"67",
		"first_name": "Rushikesh",
		"last_name": "Markad",
		"gender": "Male",
		"age": 22,
		"address": "Pune",
		"email": "markaaa@gmail.com",
		"password": "ertikol@123",
		"mob_no": 6985749623,
		"role":"user"
	}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenEmptyName(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Empty Name"))

	rr := makeHTTPCall(user.Create(cs), http.MethodPost, "/users", `{
		"id":"67",
		"first_name": "Rushikesh",
		"last_name": "",
		"gender": "Male",
		"age": 22,
		"address": "Pune",
		"email": "markaaa@gmail.com",
		"password": "ertikol@123",
		"mob_no": "6985749623",
		"role":"user"
	}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenInternalError(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Internal Error"))

	rr := makeHTTPCall(user.Create(cs), http.MethodPost, "/users", `{
		"id":"67",
		"first_name": "Rushikesh",
		"last_name": "Markad",
		"gender": "Male",
		"age": 22,
		"address": "Pune",
		"email": "markaaa@gmail.com",
		"password": "ertikol@123",
		"mob_no": "6985749623",
		"role":"user"
	}`)

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

// List :
func TestSuccessfullList(t *testing.T) {
	cs := &usermock.Service{}
	var resp user.ListResponse
	cs.On("List", mock.Anything).Return(resp, nil, errors.New("Internal Error"))

	rr := makeHTTPCall(user.List(cs), http.MethodGet, "/users", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

//How to test
// func TestListWhenNoUsers(t *testing.T) {
// 	cs := &usermock.Service{}
// 	var resp user.ListResponse
// 	cs.On("List", mock.Anything).Return(resp, errors.New("Internal Error"))

// 	rr := makeHTTPCall(user.List(cs), http.MethodGet, "/users", "")

// 	checkResponseCode(t, http.StatusNotFound, rr.Code)
// 	cs.AssertExpectations(t)
// }

//Test
func TestListInternalError(t *testing.T) {
	cs := &usermock.Service{}
	var resp user.ListResponse
	cs.On("List", mock.Anything).Return(resp, errors.New("Internal Error"))

	rr := makeHTTPCall(user.List(cs), http.MethodGet, "/users", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

//FindById
//not bad reqe
//not find err
func TestSuccessfullFindByID(t *testing.T) {
	cs := &usermock.Service{}
	var resp user.FindByIDResponse
	cs.On("FindByID", mock.Anything, mock.Anything).Return(resp, nil)

	rr := makeHTTPCall(user.FindByID(cs), http.MethodGet, "/users/8c0fbdd5-0dda-45ca-bb92-38e528c06817", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

//Test
// func TestFindByIDWhenIDNotExist(t *testing.T) {
// 	cs := &usermock.Service{}
// 	var resp user.FindByIDResponse
// 	cs.On("FindByID", mock.Anything, mock.Anything).Return(resp, errors.New("Error"))

// 	rr := makeHTTPCall(user.FindByID(cs), http.MethodGet, "/users/06a655e7-1dd6-49a8-ba26-b2c954f47eah", "")

// 	checkResponseCode(t, http.StatusNotFound, rr.Code)
// 	cs.AssertExpectations(t)
// }

func TestFindByIdWhenInternalError(t *testing.T) {
	cs := &usermock.Service{}
	var resp user.FindByIDResponse
	cs.On("FindByID", mock.Anything, mock.Anything).Return(resp, errors.New("InternalError"))

	rr := makeHTTPCall(user.FindByID(cs), http.MethodGet, "/users/857485695", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

//DeleteByID
func TestSuccessfullDeleteByID(t *testing.T) {
	cs := &usermock.Service{}

	cs.On("DeleteByID", mock.Anything, mock.Anything).Return(nil)

	rr := makeHTTPCall(user.DeleteByID(cs), http.MethodDelete, "/users/ 06a655e7-1dd6-49a8-ba26-b2c954f47eaf", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

// func TestDeleteByIDWhenIDNotExist(t *testing.T) {
// 	cs := &usermock.Service{}
// 	cs.On("DeleteByID", mock.Anything, mock.Anything).Return(nil)

// 	rr := makeHTTPCall(user.DeleteByID(cs), http.MethodDelete, "/users/ 18a655e7-1dd6-49a8-ba26-b2c954f47fag", "")

// 	checkResponseCode(t, http.StatusNotFound, rr.Code)
// 	cs.AssertExpectations(t)
// }

func TestDeleteByIDWhenInternalError(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("DeleteByID", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

	rr := makeHTTPCall(user.DeleteByID(cs), http.MethodDelete, "/users/ 06a655e7-1dd6-49a8-ba26-b2c954f47eaf", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

// func TestSuccessfullUpdate(t *testing.T) {
// 	cs := &usermock.Service{}
// 	cs.On("update", mock.Anything, mock.Anything).Return(nil)

// 	rr := makeHTTPCall(user.Update(cs), http.MethodPut, "/users", `{"id":"1", "name":"sports"}`)

// 	checkResponseCode(t, http.StatusOK, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestUpdateWhenInvalidRequestBody(t *testing.T) {
// 	cs := &CategoryServiceMock{}

// 	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1", "name":"sports",}`)

// 	checkResponseCode(t, http.StatusBadRequest, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestUpdateWhenEmptyID(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("update", mock.Anything, mock.Anything).Return(errEmptyID)

// 	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"name":"Sports"}`)

// 	checkResponseCode(t, http.StatusBadRequest, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestUpdateWhenEmptyName(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("update", mock.Anything, mock.Anything).Return(errEmptyName)

// 	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1"}`)

// 	checkResponseCode(t, http.StatusBadRequest, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestUpdateWhenInternalError(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("update", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

// 	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1"}`)

// 	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
// 	cs.AssertExpectations(t)
// }
