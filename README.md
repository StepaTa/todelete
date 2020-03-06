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

```go
func main() {
  go start("Токен", "триггер-слово")
	start("Токен", "триггер-слово") // Очень важно, чтобы последний запуск был без слова "go"
}
```
