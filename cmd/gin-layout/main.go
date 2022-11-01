package main

func main() {
	r, cleanup, _ := wireApp()
	defer cleanup()
	_ = r.Run(":8080")
}
