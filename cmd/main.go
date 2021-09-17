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
		args := strings.Split(scanner.Text(), " ")
		var command Command

		switch args[0] {
		case exitCmd:
			command = NewExitCommand()
			break
		case checkoutCmd:
			command = NewCheckoutCommand(args[1:], checkoutService)
			break
		}

		if command == nil {
			fmt.Println("unknown command")
		}

		command.Execute()
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

type exitCommand struct {
}

func NewExitCommand() Command {
	return exitCommand{}
}

func (e exitCommand) Execute() {
	os.Exit(0)
}

type CheckoutCommand struct {
	items           []string
	checkoutService checkout.Server
}

func NewCheckoutCommand(args []string, checkoutService checkout.Server) Command {
	items := strings.Split(strings.Join(args, " "), ",")
	return CheckoutCommand{
		checkoutService: checkoutService,
		items:           items,
	}
}

func (c CheckoutCommand) Execute() {
	result := c.checkoutService.Run(c.items)
	fmt.Println(result)
}

type Command interface {
	Execute()
}
