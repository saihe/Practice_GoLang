package employee

import (
	"testing"
)

func TestEmployee(t *testing.T) {
	_, err := Get(1)
	if err != nil {
		t.Error("従業員（employee）からレコード取得失敗")
	}
	t.Log("employeeのテスト終了")
}
