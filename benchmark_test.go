package nue

import (
	"fmt"
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
	for i := 0; i < 50; i++ {
		router.Add("/user", fmt.Sprintf("/show%d", i), nueHandle)
	}
	for i := 0; i < 50; i++ {
		router.Add(fmt.Sprintf("/user%d", i), fmt.Sprintf("/match%d", i), nueHandle)
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("/user/show%d", 5), nil)
	benchRequest(b, router, req)
}

// HttpRouter
func httpRouterHandle(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}

func BenchmarkHttpRouter(b *testing.B) {
	router := httprouter.New()
	for i := 0; i < 50; i++ {
		router.GET(fmt.Sprintf("/user/show%d", i), httpRouterHandle)
	}
	for i := 0; i < 50; i++ {
		router.GET(fmt.Sprintf("/user%d/show%d", i, i), httpRouterHandle)
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("/user/show%d", 5), nil)
	benchRequest(b, router, req)
}

// gorilla/mux
func gorillaHandle(_ http.ResponseWriter, _ *http.Request) {
}

func BenchmarkGorilla(b *testing.B) {
	router := mux.NewRouter()
	for i := 0; i < 50; i++ {
		router.HandleFunc(fmt.Sprintf("/user/show%d", i), gorillaHandle).Methods("GET")
	}
	for i := 0; i < 50; i++ {
		router.HandleFunc(fmt.Sprintf("/user%d/show%d", i, i), gorillaHandle).Methods("GET")
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("/user/show%d", 5), nil)
	benchRequest(b, router, req)
}

// Denco
func dencoHandler(w http.ResponseWriter, r *http.Request, params denco.Params) {}

func BenchmarkDenco(b *testing.B) {
	mux := denco.NewMux()
	var handlers []denco.Handler
	for i := 0; i < 50; i++ {
		handlers = append(handlers, mux.Handler("GET", fmt.Sprintf("/user/show%d", i), dencoHandler))
	}
	for i := 0; i < 50; i++ {
		handlers = append(handlers, mux.Handler("GET", fmt.Sprintf("/user%d/show%d", i, i), dencoHandler))
	}
	router, _ := mux.Build(handlers)
	req, _ := http.NewRequest("GET", fmt.Sprintf("/user/show%d", 5), nil)
	benchRequest(b, router, req)
}
