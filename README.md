# To Delete Bot
Скрипт на Golang, который позволяет удалять свои сообщения в ВК используя триггер-слово.

Функционал взят [отсюда](https://github.com/P2LOVE/VK-UserSide-Bot).
Используется библиотека [VK SDK](https://github.com/SevereCloud/vksdk) от [SevereCloud](https://github.com/SevereCloud).

# Установка

Для начала необходимо установить компилятор Golang, а также скачать библиотеку [VK SDK](https://github.com/SevereCloud/vksdk).

# Настройка для одного аккаунта

main.go
```go
func main() {
	start("Токен", "триггер-слово")
}
```

# Несколько аккаунтов

main.go
```go
func main() {
	go start("Токен", "триггер-слово")
	start("Токен", "триггер-слово") // Очень важно, чтобы последний запуск был без слова "go"
}
```

# Запуск

После установки, настройки, скрипт готов к работе.

```shell
go run main.go
```

Также вы можете скомпилировать заранее и легко его запускать (можно будет запускать там где не будет компилятора).
В интернете можно найти гайды о том, как компилировать скрипты для другой платформы.

```shell
go build main.go
```
