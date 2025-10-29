package interface_sample

import "testing"

func TestMockPostalAPI_GetAddress(t *testing.T) {
	api := &MockPostalAPI{}
	zip := "123-4567"
	expected := "テスト用の住所"

	addr, err := api.GetAddress(zip)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if addr != expected {
		t.Errorf("expected address %q, got %q", expected, addr)
	}
}

func TestRealPostalAPI_GetAddress(t *testing.T) {
	api := &RealPostalAPI{}
	zip := "100-0001"
	expected := "東京都千代田区千代田"

	addr, err := api.GetAddress(zip)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if addr != expected {
		t.Errorf("expected address %q, got %q", expected, addr)
	}
}
