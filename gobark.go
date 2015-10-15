package main

import "fmt"
import "net/http"
import "io/ioutil"
import "bytes"
import "os"

type BarkClient struct {
  Protocol string
  User string
  Password string
  Subdomain string
  Host string
  Port string
}

func CreateClient(host, port, user, password string) BarkClient {
  client := &BarkClient{
    Host: host,
    Port: port,
    User: user,
    Password: password,
  }

  return *client
}

func main() {
  var buffer bytes.Buffer
  client := CreateClient("http", "admin", "monit", "", "localhost", "2812",)

  ["http://")]

  buffer.WriteString(
  buffer.WriteString(client.User)
  buffer.WriteString(":")
  buffer.WriteString(client.Password)
  buffer.WriteString("@")
  buffer.WriteString(client.Host)
  buffer.WriteString(":")
  buffer.WriteString(client.Port)
  buffer.WriteString("/_status")

  url := buffer.String()

  response, err := http.Get(url)
  if err != nil {
      fmt.Printf("%s", err)
      os.Exit(1)
  } else {
      defer response.Body.Close()
      contents, err := ioutil.ReadAll(response.Body)
      if err != nil {
          fmt.Printf("%s", err)
          os.Exit(1)
      }
      fmt.Printf("%s\n", string(contents))
  }
}
