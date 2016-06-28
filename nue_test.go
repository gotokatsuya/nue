package nue

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNue(t *testing.T) {
	nue := New()
	nue.AddHandler("/foo", "/bar/baz", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world"))
	})
	nue.ShowNodes()

	req, err := http.NewRequest("GET", "/foo/bar/baz", nil)
	if err != nil {
		t.Fatalf("NewReuqest err:%v", err)
	}
	w := httptest.NewRecorder()
	nue.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected %d got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != "hello world" {
		t.Fatalf("expected %s got %s", "hello world", w.Body.String())
	}
}
