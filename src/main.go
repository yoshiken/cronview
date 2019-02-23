package main

import (
  "./metainfo"

  "bufio"
  "fmt"
  "os"
  "flag"
  "regexp"
)

func main() {
    var (
      readFilePath = flag.String("f", "", "path to crontabfile")
    )
    flag.Parse()
    file, err := os.Open(*readFilePath)
    if err != nil {
        fmt.Printf("can't open file")
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            fmt.Println("crontab format error")
        }
        line_str := sc.Text()
        fmt.Println()
        is_context, is_meta := checkCrontabContext(line_str)
        if is_context {
          if is_meta {
            optime := metaInfo.GetOpTime(line_str)
            if optime != "" {
              fmt.Println("is meta")
              fmt.Println(optime)
            }
          } else {

          }
        }
    }
}

func checkCrontabContext(str string) (is_context bool, is_meta bool) {
  ismeta := false
  iscontext := false
  if str == "" {
    return iscontext, ismeta
  }
  iscontext = true
  if isRegexp("#", string(getRuneAt(str, 0))) {
    ismeta = true
  }
  return iscontext, ismeta
}

func isRegexp(reg, str string) bool {
  return regexp.MustCompile(reg).Match([]byte(str))
}

func getRuneAt(s string, i int) rune {
    rs := []rune(s)
    return rs[i]
}
