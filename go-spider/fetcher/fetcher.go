package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//请求速率 10毫秒
var rateLimter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimter
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")
	c1 := http.Cookie{
		Name:  "sid",
		Value: "a98add8a-1644-4335-bd06-6884b10df620",
	}
	req.AddCookie(&c1)
	resp, err := client.Do(req)
	//resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("http status code  error %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(bodyReader, determineEncoding(bodyReader).NewEncoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		println(err)
	}
	return all, nil
}

/**
自动发现网页编码UTF
*/
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	byte, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher determineEncoding err")
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(byte, "")
	return e
}
