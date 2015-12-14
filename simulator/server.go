package main
import "gossip/members"
import "sync"
import "net/http"
import "encoding/json"
import "html"
import "log"
import "fmt"
import "io/ioutil"
func ServeSummary(lock * sync.Mutex, members * map[string]members.GossipMember) {

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      b, err := ioutil.ReadFile("page.html")
    if (err == nil) {
      w.Write(b)
    } else {
      fmt.Fprintf(w, "%q", html.EscapeString(err.Error()))
    }
  })

  http.HandleFunc("/jquery.js", func(w http.ResponseWriter, r *http.Request) {
      fmt.Println("jq")
      b, err := ioutil.ReadFile("jquery-2.1.4.min.js")
    if (err == nil) {
      w.Write(b)
    } else {
      fmt.Fprintf(w, "%q", html.EscapeString(err.Error()))
    }
  })


  http.HandleFunc("/data.json", func(w http.ResponseWriter, r *http.Request) {
    lock.Lock()
    b, err := json.Marshal(members)
    lock.Unlock()
    if (err == nil) {
      w.Write(b)
    } else {
      fmt.Fprintf(w, "%q", html.EscapeString(err.Error()))
    }
  })

  log.Fatal(http.ListenAndServe(":7070", nil))
}
