package nue

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/naoina/denco"
)

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

func nueHandle(_ http.ResponseWriter, _ *http.Request) {}

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := new(mockResponseWriter)
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(w, r)
	}
}

func BenchmarkNue(b *testing.B) {
	router := New()
	router.Add("/users", "/match", nueHandle)
	req, _ := http.NewRequest("GET", "/users/match", nil)
	benchRequest(b, router, req)
}

// HttpRouter
func httpRouterHandle(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}

func BenchmarkHttpRouter(b *testing.B) {
	router := httprouter.New()
	router.GET("/users/match", httpRouterHandle)

	req, _ := http.NewRequest("GET", "/users/match", nil)
	benchRequest(b, router, req)
}

// gorilla/mux
func gorillaHandle(_ http.ResponseWriter, _ *http.Request) {
}

func BenchmarkGorilla(b *testing.B) {
	router := mux.NewRouter()
	router.HandleFunc("/users/match", gorillaHandle).Methods("GET")

	req, _ := http.NewRequest("GET", "/users/match", nil)
	benchRequest(b, router, req)
}

// Denco
func dencoHandler(w http.ResponseWriter, r *http.Request, params denco.Params) {}

func BenchmarkDenco(b *testing.B) {
	mux := denco.NewMux()
	router, _ := mux.Build([]denco.Handler{mux.Handler("GET", "/users/match", dencoHandler)})

	req, _ := http.NewRequest("GET", "/users/match", nil)
	benchRequest(b, router, req)
}
