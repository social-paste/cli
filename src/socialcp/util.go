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

var FilePath = "configuration.json"
var configuration *Configuration;

func loadConfiguration() {
  data, error := ioutil.ReadFile(FilePath)

  if error != nil {
    initializeConfiguration()
    return
  }

  error = json.Unmarshal(data, &configuration)

  if error != nil {
    initializeConfiguration()
    return
  }
}

func saveConfiguration() {
  if configuration == nil {
    initializeConfiguration()
  }

  serializedConfiguration, error := json.Marshal(configuration)

  if error != nil {
    return
  }

  var indentedJSON bytes.Buffer
  json.Indent(&indentedJSON, serializedConfiguration, "", "  ")
  ioutil.WriteFile(FilePath, indentedJSON.Bytes(), 0644)
}

func initializeConfiguration() {
  configuration = new(Configuration)
  configuration.Host = "http://socialcp.cloud.dev.globoi.com/api/"
  configuration.User = ""
  configuration.Slots = make(map[int]string)
}

func SetHost(host string) {
  if configuration == nil {
    loadConfiguration()
  }

  configuration.Host = host
  saveConfiguration()
}

func GetHost() (string) {
  if configuration == nil {
    loadConfiguration()
  }

  return configuration.Host
}

func SetUser(user string) {
  if configuration == nil {
    loadConfiguration()
  }

  configuration.User = user
  saveConfiguration()
}

func GetUser() (string) {
  if configuration == nil {
    loadConfiguration()
  }

  return configuration.User
}

func SetRecipient(slot int, email string) {
  if configuration == nil {
    loadConfiguration()
  }

  configuration.Slots[slot] = email
  saveConfiguration()
}

func GetRecipient(slot int) (string) {
  if configuration == nil {
    loadConfiguration()
  }

  return configuration.Slots[slot]
}
