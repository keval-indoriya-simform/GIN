package main

func main() {
	err := Router().Run("localhost:8080")
	if err != nil {
		return
	}
}
