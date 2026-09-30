package main

import (
	"fmt"
	"net/http"

	httpTransport "ai-worker-pool/internal/transport/http"
)

func main() {
	fmt.Println("=== АСИНХРОННЫЙ ИИ-ШЛЮЗ ЗАПУЩЕН НА ПОРТУ 8080 ===")

	transpHandler := httpTransport.NewHandler()

	http.HandleFunc("/api/v1/agent", transpHandler.HandlePrompt)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ошибка запуска сервера:", err.Error())
	}
}
