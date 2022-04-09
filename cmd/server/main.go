package main

import (
	// "fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"os"
	"strconv"
	// "bytes"
	// "encoding/json"
	"github.com/uniplaces/carbon"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mm-saito/birthday-line-notify"
)

type User struct {
	Id    int `json:"id"`
	Name  string `json:"name"`
	Month int `json:"month"`
	Day int `json:"day"`
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

// Basic認証
// func checkAuth(r *http.Request) bool {
//     // 認証情報取得
//     clientID, clientSecret, ok := r.BasicAuth()
//     if ok == false {
//         return false
//     }
//     return clientID == os.Getenv("BASIC_AUTH_USER") && clientSecret == os.Getenv("BASIC_AUTH_PASS")
// }

func Index(w http.ResponseWriter, r *http.Request) {
	// 認証
	// if checkAuth(r) == false {
	// 	w.Header().Add("WWW-Authenticate", `Basic realm="SECRET AREA"`)
	// 	w.WriteHeader(http.StatusUnauthorized) // 401
	// 	http.Error(w, "Unauthorized", 401)
	// 	return
	// }
	name := ""
	now := carbon.Now()
	nowMonth, _ := strconv.Atoi(now.Format("01"))
	nowDay, _ := strconv.Atoi(now.Format("02"))
	db := database.OpenDB(os.Getenv("DRIVER"), os.Getenv("DSN"))
	if err := db.Ping(); err != nil {
		log.Fatal("db.Ping failed:", err)
	}
	selected, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("select failed:", err)
	}
	defer database.CloseDB(db)
	data := []User{}
	for selected.Next() {
		user := User{}
		err = selected.Scan(&user.Id, &user.Name, &user.Month, &user.Day)
		if err != nil {
			log.Fatal("loop failed:", err)
		}
		// 対象者名設定
		if nowMonth == user.Month && nowDay == user.Day {
			if name == "" {
				name = user.Name
			} else {
				name = name + "、" + user.Name
			}
		}
		data = append(data, user)
	}
	selected.Close()
	if name != "" {
		// LINE API Request
		LineNotifyApi(name)
	}
	// JSONレスポンス
	// var buf bytes.Buffer
	// enc := json.NewEncoder(&buf)
	// if err := enc.Encode(&data); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(buf.String())
	// _, err = fmt.Fprint(w, buf.String())
	// if err != nil {
	// 	return
	// }
}

func LineNotifyApi(name string) {
	accessToken := os.Getenv("LINE_TOKEN")
	var msg string
	if name == "あけおめ" {
		msg = "\n🎍あけましておめでとう！🎍"
	} else {
		msg =  "\n" + name + "\n\n" + "誕生日おめでとうございます🎂🎉"
	}
	URL := "https://notify-api.line.me/api/notify"

	apiUrl, err := url.ParseRequestURI(URL)
	if err != nil {
		log.Fatal(err)
	}

	c := &http.Client{}
	form := url.Values{}
	form.Add("message", msg)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", apiUrl.String(), body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer " + accessToken)

	_, err = c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}