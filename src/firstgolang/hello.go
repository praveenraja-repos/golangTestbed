package main
import ("fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "encoding/json")
type output struct {
    T1 string
	T2 int}
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    data := string(reqBody)
	for _, v := range strings.Split(r.URL.Path[1:len(r.URL.Path)], ",") {
		if v == "uppercase" {data = uppercase(data)}
		if v == "exclaimer" {data = exclaimer(data)}
	}
    fmt.Fprintln(w,  string(data))}
func main() {
    http.HandleFunc("/", ServeHTTP)
    http.ListenAndServe(":8080", nil)
}
func uppercase(s string) string {return strings.ToUpper(s)}
func exclaimer(s string) string {
	str, _ := json.Marshal(output{T1: "wow " + s + " !!!!.", T2: len("wow " + s + " !!!!.")})
	return string(str)}

