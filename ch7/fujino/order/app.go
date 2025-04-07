package order

import "net/http"

func main() {
	processManager := NewProcessManager()

	// イベントを購読する
	SubscribeEvent(
		processManager,
		CallValidateAddressEventHandler,
	)

	go Listen(processManager)
}

// 注文確定コマンドを処理する
func PlaceOrderCommandHandler(w http.ResponseWriter, r *http.Request) ([]Event, error) {
	cmd := PlaceOrderCommand{}
	events, err := PlaceOrder(cmd)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// アドレス検証イベントを処理する
func CallValidateAddressEventHandler(event CallValidateAddressEvent) ([]Event, error) {
	// アドレス検証の処理を実装
	println("Validating address:", event.Address)

	validated := NewAddressValidatedEvent(event.ID, event.Address)
	return []Event{validated}, nil
}
