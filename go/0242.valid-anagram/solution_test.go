// Code generated by https://github.com/j178/leetgo.

package main

import (
    "testing"

    . "github.com/j178/leetgo/testutils/go"
)

var testcases = `
input:
"anagram"
"nagaram"
output:
true

input:
"rat"
"car"
output:
false
`

func Test_isAnagram(t *testing.T) {
    targetCaseNum := 0
    // targetCaseNum := -1
    if err := RunTestsWithString(t, isAnagram, testcases, targetCaseNum); err != nil {
        t.Fatal(err)
    }
}
