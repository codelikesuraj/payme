/*
Copyright © 2024 Abdulbaki Suraj <surajabdulbaki19@gmail.com>
*/
package main

import (
	"fmt"
	"os"
	"payme/cmd"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Execute()
}
