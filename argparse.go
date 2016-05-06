package main

import (
    "fmt"
    "strings"
    "os"
)

func main() {
  m := make(map[string]string)
  s := os.Args
  s = append(s[:0], s[1:]...)
  for i := range s {
    if strings.HasPrefix(s[i], "-") {
      s[i] = s[i][1:]
      if strings.Contains(s[i], "=") {
        p := strings.Split(s[i], "=")
        m[p[0]] = p[1]
        continue
      }
      i2 := i + 1
      next := s[i2]
      m[s[i]] = next
    } else {
      m[s[i]] = "action"
    }
  }
  for key, value := range m {
    fmt.Println("Key:", key, "Value:", value)
  }
}
