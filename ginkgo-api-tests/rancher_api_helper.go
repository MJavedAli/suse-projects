package main_test

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
	"crypto/tls"
)

type TestConfig struct {
    RancherURL  string
    Username    string
    Password    string
}

func loginAndGetToken(config TestConfig) (string, error) {
    loginData := map[string]string{
        "username": config.Username,
        "password": config.Password,
    }

    loginDataJSON, err := json.Marshal(loginData)
    if err != nil {
        return "", err
    }

	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

	client := &http.Client{Transport: tr}

	resp, err := client.Post(config.RancherURL+"/v3-public/localProviders/local?action=login", "application/json", bytes.NewBuffer(loginDataJSON))

    // resp, err := http.Post(config.RancherURL+"/v3-public/localProviders/local?action=login", "application/json", bytes.NewBuffer(loginDataJSON))
    
	if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var loginResponse map[string]interface{}
    if err := json.Unmarshal(body, &loginResponse); err != nil {
        return "", err
    }

    token, ok := loginResponse["token"].(string)
    if !ok {
        return "", fmt.Errorf("failed to get token from the login response")
    }

    return token, nil
}
