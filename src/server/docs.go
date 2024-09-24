package server

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Load swagger configuration from yaml an export it as string.
func LoadDoc(urlApi string) string {
	dir, _ := os.Getwd()
	path := "/src/docs/swagger.yaml"
	docsPath := filepath.Join(dir, path)
	swaggerFile, err := os.ReadFile(docsPath)
	if err != nil {
		log.Print("Remember, your must create a folder docs inside of src folder and create swagger.yaml file")
		log.Panicf("Error on load swagger doc: %v", err)
	}
	return string(setProperties(&swaggerFile, urlApi))
}

// Set dynamic configuration that must be expose on the documentation endpoint.
// For example: Set the url of the API.
func setProperties(swaggerFile *[]byte, urlApi string) []byte {
	var config map[string]interface{}
	err := yaml.Unmarshal(*swaggerFile, &config)
	if err != nil {
		log.Fatalf("Error on parsing yaml file: %v", err)
	}
	servers := config["servers"].([]interface{})
	if len(servers) > 0 {
		server := servers[0].(map[interface{}]interface{})
		server["url"] = urlApi
	}
	updatedYAML, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Error on update yaml: %v", err)
	}
	return updatedYAML
}
