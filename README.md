# Nue
Fast router.


[![wercker status](https://app.wercker.com/status/6135ebbc86ffbe8fc6b370f18241bbea/m "wercker status")](https://app.wercker.com/project/bykey/6135ebbc86ffbe8fc6b370f18241bbea)

## usage

```go
package main

import (
	"net/http"
	
	"github.com/gotokatsuya/nue"
)

func main() {
	handler := nue.New()
	handler.Add("/user", "/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", handler)
}
```
