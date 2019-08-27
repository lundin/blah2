package main

import "testing"


func TestSayMyName(t *testing.T) {

actual:=sayMyName("")

if (actual!="hello unknown") {
t.Errorf("Failed")
} else {

t.Logf("Success")
}

}
