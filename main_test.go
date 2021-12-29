package main

import "testing"

func TestUpdateAliDns(t *testing.T) {
	err := UpdateAliDns(GetInternetIp())

	if err != nil {
		t.Fatal(err)
	}
}
