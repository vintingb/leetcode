// Code generated by https://github.com/j178/leetgo.

package main

import (
    "testing"

    . "github.com/j178/leetgo/testutils/go"
)

var testcases = `
input:
[1,2,3,4,5]
2
output:
[1,2,3,5]

input:
[1]
1
output:
[]

input:
[1,2]
1
output:
[1]
`

func Test_removeNthFromEnd(t *testing.T) {
    targetCaseNum := 0
    // targetCaseNum := -1
    if err := RunTestsWithString(t, removeNthFromEnd, testcases, targetCaseNum); err != nil {
        t.Fatal(err)
    }
}
