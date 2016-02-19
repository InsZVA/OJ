package main
import "fmt"
import "net/http"
import "io/ioutil"
import "net/url"

func main() {
    resp, err := http.PostForm("http://127.0.0.1:1996/judge",
		url.Values{
            "api":      {"http://127.0.0.1:1234/problem/"},
            "problemId":{"2"},
            "submission":{"1234"},
            "compiler": {"gcc"},
        })
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}