package algo_interview

import (
	"testing"
)

func TestNewRichestAccounts(t *testing.T) {
	r := NewRichestAccounts()
	if r == nil {
		t.Errorf("NewRichestAccounts() returned nil")
	}
}

func TestUpdateRichest(t *testing.T) {
	r := NewRichestAccounts()
	err := r.UpdateRichest("ID1", Type1, 10, 0)
	if err != nil {
		t.Errorf("UpdateRichest() returned an error: %v", err)
	}
	err = r.UpdateRichest("ID1", "InvalidType", 10, 0)
	if err == nil {
		t.Errorf("UpdateRichest() should return an error for invalid type")
	}
}

func TestGetRichestByTime(t *testing.T) {
	r := NewRichestAccounts()
	r.UpdateRichest("ID1", Type1, 10, 0)
	r.UpdateRichest("ID2", Type1, 5, 100)
	r.UpdateRichest("ID3", Type1, 11, 200)

	accounts, err := r.GetRichestByTime(110)
	if err != nil {
		t.Errorf("GetRichestByTime() returned an error: %v", err)
		return
	}
	if len(accounts) == 0 {
		t.Errorf("GetRichestByTime() returned an empty array")
		return
	}
	if accounts[0].Token != 10 {
		t.Errorf("GetRichestByTime() returned incorrect Token: got %v, want %v", accounts[0].Token, 10)
	}
}

func TestGetAccountTypeTimeMap(t *testing.T) {
	r := NewRichestAccounts()
	_, err := r.GetAccountTypeTimeMap(Type1)
	if err != nil {
		t.Errorf("GetAccountTypeTimeMap() returned an error: %v", err)
	}
	_, err = r.GetAccountTypeTimeMap("InvalidType")
	if err == nil {
		t.Errorf("GetAccountTypeTimeMap() should return an error for invalid type")
	}
}

func TestCheckTime(t *testing.T) {
	r := NewRichestAccounts()
	r.UpdateRichest("ID1", Type1, 10, 0)
	r.UpdateRichest("ID2", Type1, 5, 100)

	timeMap, _ := r.GetAccountTypeTimeMap(Type1)

	account, err := checkTime(timeMap, 50)
	if err != nil {
		t.Errorf("checkTime() returned an error: %v", err)
	}
	if account.Token != 10 {
		t.Errorf("checkTime() returned incorrect Token: got %v, want %v", account.Token, 10)
	}

	_, err = checkTime(timeMap, 1000)
	if err == nil {
		t.Errorf("checkTime() should return an error for non-matching time")
	}
}

func TestGetMaxTokensForTime(t *testing.T) {
	r := NewRichestAccounts()
	r.UpdateRichest("ID1", Type1, 10, 0)
	r.UpdateRichest("ID2", Type1, 5, 100)

	timeMap, _ := r.GetAccountTypeTimeMap(Type1)
	tokens, exists := getMaxTokensForTime(timeMap, 0)

	if !exists || tokens != 10 {
		t.Errorf("getMaxTokensForTime() returned incorrect Token: got %v, want %v", tokens, 10)
	}
}

func TestUpdateRichestWithLesserValue(t *testing.T) {
	r := NewRichestAccounts()
	err := r.UpdateRichest("ID1", Type1, 10, 0)
	if err != nil {
		t.Fatalf("UpdateRichest() returned an error: %v", err)
	}

	// Trying to update with a lesser token value
	err = r.UpdateRichest("ID2", Type1, 5, 0)
	if err != nil {
		t.Fatalf("UpdateRichest() returned an error: %v", err)
	}

	account, err := checkTime(r.Type1, 0)
	if err != nil {
		t.Fatalf("checkTime() returned an error: %v", err)
	}
	if account.Token != 10 {
		t.Errorf("Token should not be updated: got %v, want %v", account.Token, 10)
	}
}
