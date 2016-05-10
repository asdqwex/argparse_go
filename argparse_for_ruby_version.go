package argparse

import (
    "fmt"
    "strings"
)

func Parse(s []string) (m map[string]string)  {
  m = make(map[string]string)
  pass := 0
  s = append(s[:0], s[1:]...)
  for i := range s {
    if pass == 1 {
      pass = 0
      continue
    }
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
      pass = 1
    } else {
      m[s[i]] = "action"
    }
  }
  action_counter := 0
  for _ , v := range m {
    if v == "action" {
      action_counter++
    }
  }
  if action_counter > 1 {
    fmt.Println()
    panic(fmt.Sprintf("%v", "there were more than on actions called"))
  }
  return m
}
