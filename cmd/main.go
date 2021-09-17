package main

import (
	"bufio"
	"fmt"
	"myshop/repository"
	"myshop/service/checkout"
	"os"
	"strings"
)

const (
	exitCmd     = "exit"
	checkoutCmd = "checkout"
)

func main() {
	checkoutService := setupServer()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == exitCmd {
			break
		}
		args := strings.Split(scanner.Text(), " ")
		if len(args) > 1 && args[0] == checkoutCmd {
			fmt.Println(checkoutService.Run(args[1:]))
		} else {
			fmt.Println("no item added")
		}
	}
}

func setupServer() checkout.Server {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	inventory, err := repository.NewInventoryRepository(fmt.Sprintf("%s/%s", dir, "repository/data/sample_inventory.json"))
	if err != nil {
		panic(err)
	}

	promotion, err := repository.NewPromotionRepository(fmt.Sprintf("%s/%s", dir, "repository/data/sample_promotion.json"))
	if err != nil {
		panic(err)
	}

	return checkout.NewCheckoutService(inventory, promotion)
}
