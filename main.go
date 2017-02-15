package main

import (
  "crypto/rand"
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "path"
  "os"
)

var items []interface{}

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func UUID() string {
  b := make([]byte, 16)
  _, err := rand.Read(b)
  check(err)

  b[6] = (b[6] & 0x0f) | 0x40
  b[8] = (b[8] & 0x3f) | 0x80

  return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func main() {
  var f string
  var o string

  flag.StringVar(&f, "f", "", "the path to the json file");
  flag.StringVar(&o, "o", "", "the output folder");
  flag.Parse();

  if len(f) > 0 && len(o) > 0 {
    file, err := ioutil.ReadFile(f)
    check(err)

    os.MkdirAll(o, 0755)
    json.Unmarshal(file, &items)

    for _, item := range items {
      b, err := json.MarshalIndent(item, "", "  ")
      check(err)

      filepath := path.Join(o, fmt.Sprintf("%s.md", UUID()))
      newfile, err := os.Create(filepath)
      check(err)

      newfile.Write(b)
      fmt.Printf("%s ... saved\n", filepath)
    }
  } else {
    flag.Usage()
  }
}
