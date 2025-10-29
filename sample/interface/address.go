package interface_sample

type PostalAPI interface {
	GetAddress(zip string) (string, error)
}

type RealPostalAPI struct{}

func (RealPostalAPI) GetAddress(zip string) (string, error) {
	// 実際はネットワーク通信してAPIを呼ぶ
	return "東京都千代田区千代田", nil
}

type MockPostalAPI struct{}

func (MockPostalAPI) GetAddress(zip string) (string, error) {
	return "テスト用の住所", nil
}
