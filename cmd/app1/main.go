package main

import "gin-layout/internal/app1/configs"

func main() {

	r, cleanup, _ := wireApp(&configs.Bootstrap{})
	defer cleanup()
	_ = r.Run(":8080")
}
