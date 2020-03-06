package main

import (
	"deleter/bot"
)

func main() {
	/*
	ВАЖНО!

	Если вы хотите одновременно запустить несколько аккаунтов, то необходимо сделать примерно так:
	...
	go bot.Start()
	bot.Start()

	Иначе, если поставить наоборот, то запустится только один аккаунт. В for {} смысла нет из-за его бесполезной нагрузки.
	*/
	// Токен и слово-триггер
	// go bot.Start("", "dd")
	bot.Start("", "хм")
}
