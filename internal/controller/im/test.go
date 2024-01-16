package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ImHandler struct {
	AppKey    string
	AppSecret string
}

func (h *ImHandler) GetImHead() (AppKey string, Nonce string, CheckSum string, CurTime string) {
	AppKey = h.AppKey

	//UTC时间戳获取
	CurTime = strconv.FormatInt(time.Now().Unix(), 10)

	//随机数获取[0-100000000)
	rand.NewSource(time.Now().UnixNano())
	Nonce = strconv.Itoa(rand.Intn(100000000))

	//CheckSum计算方式AppSecret + Nonce + CurTime后去sh1哈希计算
	CheckSum =
		func() string {
			o := sha1.New()
			o.Write([]byte(h.AppSecret + Nonce + CurTime))
			return hex.EncodeToString(o.Sum(nil))
		}()

	println("AppKey =", AppKey)
	println("Nonce =", Nonce)
	println("CurTime =", CurTime)
	println("CheckSum =", CheckSum)
	return
}

func (h *ImHandler) httpDo(reader string) {
	AppKey := h.AppKey

	//UTC时间戳获取
	CurTime := strconv.FormatInt(time.Now().Unix(), 10)

	//随机数获取[0-100000000)
	rand.NewSource(time.Now().UnixNano())
	Nonce := strconv.Itoa(rand.Intn(100000000))

	//CheckSum计算方式AppSecret + Nonce + CurTime后去sh1哈希计算
	CheckSum :=
		func() string {
			o := sha1.New()
			o.Write([]byte(h.AppSecret + Nonce + CurTime))
			return hex.EncodeToString(o.Sum(nil))
		}()

	println("AppKey =", AppKey)
	println("Nonce =", Nonce)
	println("CurTime =", CurTime)
	println("CheckSum =", CheckSum)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.netease.im/nimserver/user/create.action", strings.NewReader(reader))

	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("AppKey", AppKey)
	req.Header.Add("Nonce", Nonce)
	req.Header.Add("CurTime", CurTime)
	req.Header.Add("CheckSum", CheckSum)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {

		// handle error

	}

	fmt.Println(string(body))

	return
}

func main() {
	im := ImHandler{AppKey: "05b4c8c3883a6426f3590e5f8fba3762", AppSecret: "e3961aba4f93"}
	im.httpDo("accid=test1")

	//AppKey, Nonce, CheckSum, CurTime := im.GetImHead()

	//targetUrl := "https://api.netease.im/nimserver/user/create.action"
	//
	//payload := strings.NewReader("{\"accid\":\"test1\"")
	//req, _ := http.NewRequest("POST", targetUrl, payload)
	//
	//req.Header.Add("AppKey", AppKey)
	//req.Header.Add("Nonce", Nonce)
	//req.Header.Add("CurTime", CurTime)
	//req.Header.Add("CheckSum", CheckSum)
	//
	//response, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	println("Error")
	//}
	//println(response)
}
