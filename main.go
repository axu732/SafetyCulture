package main

import (
	"github.com/georgechieng-sc/interns-2022/folders"
)

func main() {
	res := folders.GenerateData()

	folders.PrettyPrint(res)

	folders.WriteSampleData(res)
}