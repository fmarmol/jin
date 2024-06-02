package jin_test

import (
	"net/http/httptest"
	"testing"

	"github.com/fmarmol/jin"
	"github.com/gin-gonic/gin"
)

func BenchmarkEngine(b *testing.B) {
	jin.SetRealeaseMode()
	e := jin.New()
	e.GET("/hello", func(c jin.Context) (any, error) {
		return "world", nil
	})

	s := httptest.NewServer(e)
	defer s.Close()
	client := s.Client()

	url := s.URL + "/hello"

	for i := 0; i < b.N; i++ {
		resp, err := client.Get(url)
		if err != nil {
			b.Fatal(err)
		}
		if resp.StatusCode != 200 {
			b.Fatal(resp.Status)
		}
	}

}
func BenchmarkEngineGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.GET("/hello", func(c *gin.Context) {
		c.JSON(200, "world")
		return
	})

	s := httptest.NewServer(e)
	defer s.Close()
	client := s.Client()

	url := s.URL + "/hello"
	for i := 0; i < b.N; i++ {
		resp, err := client.Get(url)
		if err != nil {
			b.Fatal(err)
		}
		if resp.StatusCode != 200 {
			b.Fatal(resp.Status)
		}
	}
}
