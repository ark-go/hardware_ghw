package main

import (
	"fmt"
	"github.com/jaypipes/ghw"
)

func init() {
}

func main() {
	block, err := ghw.Block()
	if err != nil {
		fmt.Printf("Ошибка при получении информации о блочном хранилище : %v", err)
	}

	fmt.Printf("%v\n", block)

	for _, disk := range block.Disks {
		fmt.Printf(" %v\n", disk.SerialNumber)
		if disk.SerialNumber != "" {
			break
		}
	}
}
