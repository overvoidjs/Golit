package main

import(
  "bytes"
  "testing"
  "strings"
  "bufio"
)

var testOk = `1
2
3
3
4
5`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T){

  in := bufio.NewReader(strings.NewReader(testOk))
  out := new(bytes.Buffer)

  err := uniq(in, out)
  if err != nil {
    t.Errorf("test for Ok Failed")
  }

  result := out.String()
  if result != testOkResult {
    t.Errorf("test result not match \n %v %v", result , testOkResult)
  }
}
