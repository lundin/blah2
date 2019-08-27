package main

import "testing"


func TestSayMyName(t *testing.T) {

actual:=sayMyName("")

if (actual!="Hello unknown") {
t.Errorf("Failed")
} else {

t.Logf("Success")
}

}
