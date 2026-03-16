package main


import "fmt"


func main() {
	name := "api-server"
	port := 8080
	isHealthy := true


	var timeout int = 30
	var version string = "v1.10"


	const maxRetries = 3

	fmt.Println(name, port, isHealthy, timeout, version, maxRetries)


}
