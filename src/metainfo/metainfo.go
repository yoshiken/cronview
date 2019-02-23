package metaInfo

import (
  "strings"
  "regexp"
)
func GetOpTime(str string)  (string){
  if isRegexp("# @optime ", str) {
    return strings.Split(str, "# @optime ")[1]
  }
  return ""
}

func isRegexp(reg, str string) bool {
  return regexp.MustCompile(reg).Match([]byte(str))
}
