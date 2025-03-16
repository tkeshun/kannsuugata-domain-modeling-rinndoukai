package shareinfra

import (
	"fmt"

	shareinterfaces "example-ch7_8/internal/share/domain/interfaces"
)

func PublishEvent(event shareinterfaces.Event) {
	// TODO(char5742): イベントを発行する処理を実装
	// 例: メッセージキューにイベントを送信するなど
	// ここでは単純に標準出力に出力する例を示します
	fmt.Printf("Publishing event: %+v\n", event)
}
