package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func HttpGetJson(url string, obj interface{}) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if obj != nil {
		err = json.Unmarshal(body, obj)
	}
	return
}

func GetSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := r.Intn(max-min) + min
	fmt.Println("int code -----> ", code)
	return code
}

func CurrentTime() int64 {
	var curTime = time.Now().Unix()
	return curTime
}

func CurrentTimeString() string {
	var curTimeInt = time.Now().Unix()
	var curTime = strconv.FormatInt(curTimeInt, 10)
	return curTime
}

func EntityTime(timeStr string) (int64, bool) {
	the_time, err := time.Parse("2006-01-02T15:04:05-0700", timeStr)
	if err != nil {
		the_time, err = time.Parse("2006-01-02 15:04:05", timeStr)
	}

	if err != nil {
		the_time, err = time.Parse("2006-01-02 15:04", timeStr)
	}

	if err != nil {
		fmt.Println(err)
		return 0, false
	} else {
		unix_time := the_time.Unix()
		return unix_time, true
	}
}
