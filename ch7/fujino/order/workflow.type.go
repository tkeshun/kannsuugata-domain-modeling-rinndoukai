package order

// 注文確定のコマンド
type PlaceOrderCommand struct {
	UnvalidatedOrder
}

// 注文確定ワークフロー
type PlaceOrderWorkflow func(order PlaceOrderCommand) ([]Event, error)
type PlaceOrderStartedEvent struct {
	ID WorkflowID
	UnvalidatedOrder
}
type PlaceOrderCompletedEvent struct {
	ID WorkflowID
}

// 注文の検証
type ValidateOrderStep func(order UnvalidatedOrder) (*CallValidateAddressEvent, error)

// 検証されていない注文
type UnvalidatedOrder struct{}

// 検証された注文
type ValidatedOrder struct{}

// アドレス検証のためのコールイベント
type CallValidateAddressEvent struct {
	ID      WorkflowID
	Address string
}

type AddressValidatedEvent struct {
	ID      WorkflowID
	Address string
}

// 注文の価格計算イベント
type PriceOrderStep func(order ValidatedOrder) (*CallGetProductPriceEvent, error)

// 商品価格取得のためのコールイベント
type CallGetProductPriceEvent struct{}
