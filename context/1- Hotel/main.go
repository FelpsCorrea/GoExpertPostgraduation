package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// Setando timeout de 3 segundos para esse contexto
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// O select serve como um case, mas aguardando resultados de forma assincrona
	select {
	// Caso o timeout de 3 segundos passe, cancela a reserva
	case <-ctx.Done():
		fmt.Println("Hotel booking canceller. Timeout reached.")

	// Caso passe 1 segundo, reserva
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel Booked.")
	}

}
