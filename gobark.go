package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "bytes"
  "strings"
  "os"
  "gopkg.in/xmlpath.v2"
)

type BarkClient struct {
  Protocol string
  User string
  Password string
  Subdomain string
  Host string
  Port string
}

type Monit struct {

}

func CreateClient(protocol, user, password, subdomain, host, port string) BarkClient {
  client := &BarkClient{
    Protocol: protocol,
    User: user,
    Password: password,
    Subdomain: subdomain,
    Host: host,
    Port: port,
  }

  return *client
}

func toUtf8(iso8859_1_buf []byte) string {
    buf := make([]rune, len(iso8859_1_buf))
    for i, b := range iso8859_1_buf {
        buf[i] = rune(b)
    }
    return string(buf)
}

func main() {
  var buffer bytes.Buffer

  client := CreateClient("http", "admin", "monit", "", "localhost", "2812",)

  array := [10]string{client.Protocol, "://", client.User, ":", client.Password, "@", client.Host, ":", client.Port, "/_status?format=xml"}

  for _, elem := range array {
    buffer.WriteString(elem)
  }

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

    utf_string_xml := toUtf8(contents)
    node, err := xmlpath.Parse(strings.NewReader(utf_string_xml))
    if err != nil {
      fmt.Println(err)
    }

    fmt.Println(node)
  }
}
