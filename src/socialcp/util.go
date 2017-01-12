package main

import (
	"encoding/json"
	"io/ioutil"
  "bytes"
)

type Configuration struct {
  Host string
  User string
  Slots map[int]string
}

var filePath = "configuration.json"
var configuration *Configuration;

func LoadConfiguration() {
  data, error := ioutil.ReadFile(filePath)

  if error != nil {
    InitializeConfiguration()
    return
  }

  error = json.Unmarshal(data, &configuration)

  if error != nil {
    InitializeConfiguration()
    return
  }
}

func SaveConfiguration() {
  if configuration == nil {
    InitializeConfiguration()
  }

  serializedConfiguration, error := json.Marshal(configuration)

  if error != nil {
    return
  }

  var indentedJSON bytes.Buffer
  json.Indent(&indentedJSON, serializedConfiguration, "", "  ")
  ioutil.WriteFile(filePath, indentedJSON.Bytes(), 0644)
}

func InitializeConfiguration() {
  configuration = new(Configuration)
  configuration.Host = "http://socialcp.cloud.dev.globoi.com/api/"
  configuration.User = ""
  configuration.Slots = make(map[int]string)
}

func SetHost(host string) {
  if configuration == nil {
    LoadConfiguration()
  }

  configuration.Host = host
}

func GetHost() (string) {
  if configuration == nil {
    LoadConfiguration()
  }

  return configuration.Host
}

func SetUser(user string) {
  if configuration == nil {
    LoadConfiguration()
  }

  configuration.User = user
}

func GetUser() (string) {
  if configuration == nil {
    LoadConfiguration()
  }

  return configuration.User
}

func SetRecipient(slot int, email string) {
  if configuration == nil {
    LoadConfiguration()
  }

  configuration.Slots[slot] = email
}

func GetRecipient(slot int) (string) {
  if configuration == nil {
    LoadConfiguration()
  }

  return configuration.Slots[slot]
}
