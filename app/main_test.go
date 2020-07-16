package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	env "github.com/caarlos0/env/v6"
)

func TestHello(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(Hello))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode expected %v, got %v", http.StatusOK, resp.StatusCode)
	}

	if string(body) != "Hello World" {
		t.Errorf("Body expected %q, got %q", "Hello World", string(body))
	}
}

func TestCreateDBConnection(t *testing.T) {
	dbConf := DBConfig{}
	if err := env.Parse(&dbConf); err != nil {
		t.Error(err)
	}

	err := CreateDBConnection(dbConf)
	defer SQLConn.Close()

	if err != nil {
		t.Error(err)
	}
}
