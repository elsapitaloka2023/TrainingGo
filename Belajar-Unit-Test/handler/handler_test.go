package handler_test

import (
	"TrainingGo/Belajar-Unit-Test/handler"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestCaseQuestion struct {
	name           string
	questionA      int
	questionB      int
	expectedResult int
}

func TestCalculator(t *testing.T) {
	testCases := []TestCaseQuestion{
		{questionA: 1, questionB: 2, expectedResult: 3},
		{questionA: 2, questionB: 2, expectedResult: 4},
		{questionA: 3, questionB: 2, expectedResult: 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := handler.Calculator(tc.questionA, tc.questionB)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}

func TestGetHelloMessage(t *testing.T) {
	expectedOutput := "Halo dari Gin!"
	actualOutput := handler.GetHelloMessage()
	assert.Equal(t, expectedOutput, actualOutput, "The message should be '%s'", expectedOutput)
}

func TestRootHandler(t *testing.T) {
	// Create a new response recorder
	// This recorder will act as the target of our http request
	// We can inspect it later to see if the handler did what we wanted
	recorder := httptest.NewRecorder()

	// Create a fake request. We don't really need to fill in the details
	// because our handler will not depend on them. But we do need to pass
	// something in so that the router can match the request to a route
	request, err := http.NewRequest(http.MethodGet, "/root", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a Gin context from the recorder and request. This context
	// will contain the response from the handler
	context, _ := gin.CreateTestContext(recorder)
	context.Request = request

	// Call the handler, passing the context and the request
	handler.RootHandler(context)

	// Check the status code
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", recorder.Code)
	}

	// Check the response body
	expectedResponse := `{"message":"Halo dari Gin!"}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("expected body %v; got %v", expectedResponse, recorder.Body.String())
	}
}

func TestRootHandlerwGin(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup the router and route
	router := gin.Default()
	router.GET("/", handler.RootHandler)

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/", nil)

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code) // Check if the status code is 200

	expectedBody := `{"message":"Halo dari Gin!"}`
	assert.JSONEq(t, expectedBody, w.Body.String()) // Check if the response body matches the expected JSON
}

type JsonRequest struct {
	Message string `json:"message"`
}

func TestPostHandlerGin_PositifNegatifCase(t *testing.T) {
	// Setup router
	r := gin.Default()
	r.POST("/", handler.PostHandler)

	t.Run("Positive Case", func(t *testing.T) {
		// Persiapan data JSON
		requestBody := JsonRequest{Message: "Hello from test!"}
		requestBodyBytes, _ := json.Marshal(requestBody)

		// Buat permintaan HTTP POST
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Buat ResponseRecorder untuk merekam respons
		w := httptest.NewRecorder()

		// Lakukan permintaan
		r.ServeHTTP(w, req)

		// Periksa status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Periksa body respons
		expectedBody := `{"message":"Hello from test!"}`
		assert.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("Negative Case - EOF Error", func(t *testing.T) {
		// Persiapan data JSON yang salah
		requestBody := ""
		requestBodyBytes := []byte(requestBody)
		// Buat permintaan HTTP POST
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")
		// Buat ResponseRecorder untuk merekam respons
		w := httptest.NewRecorder()
		// Lakukan permintaan
		r.ServeHTTP(w, req)
		// Periksa status code
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// Periksa body respons
		assert.Contains(t, w.Body.String(), "{\"error\":\"EOF\"}")
	})
}

func TestPostHandlerwGin_PositifCase(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup the router and route
	router := gin.Default()
	router.POST("/", handler.PostHandler)

	// Create a new HTTP request
	bodyJson := JsonRequest{Message: "Hello, World!"}
	bodyBit, _ := json.Marshal(bodyJson)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(bodyBit))

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code) // Check if the status code is 400

	// Periksa body respons
	expectedBody := `{"message":"Hello, World!"}`
	assert.JSONEq(t, expectedBody, w.Body.String()) // Check if the response body matches the expected JSON
}
