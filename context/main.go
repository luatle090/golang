package context

import (
	"context"
	"net/http"
	"strconv"
)

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// get query string
	gioHang := r.URL.Query().Get("type")
	if gioHang == "" {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("query string ko tồn tại"))
	}
	result, _ := doSomething(ctx, gioHang)
	// convert int to string and then to byte array
	w.Write([]byte(strconv.Itoa(result)))

}

func doSomething(ctx context.Context, data string) (int, error) {
	// xử lý nghiệp vụ với data
	return 100, nil
}

func main() {

}
