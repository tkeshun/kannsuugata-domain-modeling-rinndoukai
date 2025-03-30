package order

var PlaceOrder PlaceOrderWorkflow = func(order PlaceOrderCommand) ([]Event, error) {
	event, err := ValidateOrder(order.UnvalidatedOrder)
	if err != nil {
		return nil, err
	}
	// 検証が成功した場合、次のステップに進む
	events := []Event{event, NewPlaceOrderStartedEvent()}

	return events, nil
}

func NewPlaceOrderStartedEvent() *PlaceOrderStartedEvent {
	return &PlaceOrderStartedEvent{
		// ランダム生成
		ID: WorkflowID("12345"),
	}
}
func (PlaceOrderStartedEvent) EventType() EventType {
	return "PlaceOrderStarted"
}
func (p PlaceOrderStartedEvent) WorkflowID() WorkflowID {
	return p.ID
}

func (CallValidateAddressEvent) EventType() EventType {
	return "CallValidateAddress"
}
func (c CallValidateAddressEvent) WorkflowID() WorkflowID {
	return c.ID
}

func NewAddressValidatedEvent(id WorkflowID, address string) *AddressValidatedEvent {
	return &AddressValidatedEvent{
		ID:      id,
		Address: address,
	}
}

func (AddressValidatedEvent) EventType() EventType {
	return "AddressValidated"
}
func (v AddressValidatedEvent) WorkflowID() WorkflowID {
	return v.ID
}

var ValidateOrder ValidateOrderStep = func(order UnvalidatedOrder) (*CallValidateAddressEvent, error) {
	// 検証ロジックを実装

	// ドメインで閉じている検証可能なフィールドはここで検証する
	// 文字数, 形式, など

	// 検証が成功した場合、リモートサービスを呼び出すためのイベントを返す
	return &CallValidateAddressEvent{}, nil
}
