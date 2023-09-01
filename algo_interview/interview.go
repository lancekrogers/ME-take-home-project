// Package algo_interview implements the following instructions
// Can you modify your solution, to give me the ability to query the highest AccountType tokens at any time-point Example:

// 0ms - ID1, TYPE1, token: 10
// 100ms -> ID2, TYPE1, token: 5
// 200ms -> ID3, TYPE1, token 11

// query(time = 1ms, type = TYPE1)  = 10
// query(time = 110ms, type = TYPE1) = 10
// query(time = 200ms, type = TYPE1) = 11
package algo_interview

import (
	"fmt"
	"sync"
)

const (
	Type1     = "escrow"
	Type2     = "auctionData"
	Type3     = "auction"
	Type4     = "masterEdition"
	Type5     = "metadata"
	Type6     = "mint"
	Type7     = "account"
	TypeCount = 7
)

type Account struct {
	Id      string
	ActType string
	Token   int
	Time    int64
}

// RichestAccounts struct   Contains a map of richest accounts for each valid type
type RichestAccounts struct {
	mu sync.Mutex
	// type: time[*Account]
	Type1 map[int64]*Account
	Type2 map[int64]*Account
	Type3 map[int64]*Account
	Type4 map[int64]*Account
	Type5 map[int64]*Account
	Type6 map[int64]*Account
	Type7 map[int64]*Account
}

// NewRichestAccounts function    Initializes NewRichestAccounts()
func NewRichestAccounts() *RichestAccounts {
	return &RichestAccounts{
		Type1: make(map[int64]*Account),
		Type2: make(map[int64]*Account),
		Type3: make(map[int64]*Account),
		Type4: make(map[int64]*Account),
		Type5: make(map[int64]*Account),
		Type6: make(map[int64]*Account),
		Type7: make(map[int64]*Account),
	}
}

// UpdateRichest method    Updates the RichestAccounts struct
func (r *RichestAccounts) UpdateRichest(id string, typ string, token int, time int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	account := &Account{
		Id:      id,
		ActType: typ,
		Token:   token,
		Time:    time,
	}
	timeMap, err := r.GetAccountTypeTimeMap(typ)
	if err != nil {
		return err
	}
	if val, ok := getMaxTokensForTime(timeMap, time); ok {
		if token > val {
			timeMap[time] = account
		}
	} else {
		timeMap[time] = account
	}
	return nil
}

// GetRichestByTime method    Gets the richest account for each type at a particular time
func (r *RichestAccounts) GetRichestByTime(time int64) ([]*Account, error) {
	accountArray := make([]*Account, TypeCount)
	types := []string{Type1, Type2, Type3, Type4, Type5, Type6, Type7}
	for i, name := range types {
		timeMap, err := r.GetAccountTypeTimeMap(name)
		if err != nil {
			return nil, err
		}
		account, err := checkTime(timeMap, time)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		accountArray[i] = account
	}
	return accountArray, nil
}

// GetAccountTypeTimeMap method    Get the timemap for an account type
func (r *RichestAccounts) GetAccountTypeTimeMap(typ string) (map[int64]*Account, error) {
	var timeMap map[int64]*Account
	switch typ {
	case Type1:
		timeMap = r.Type1
	case Type2:
		timeMap = r.Type2
	case Type3:
		timeMap = r.Type3
	case Type4:
		timeMap = r.Type4
	case Type5:
		timeMap = r.Type5
	case Type6:
		timeMap = r.Type6
	case Type7:
		timeMap = r.Type7
	default:
		return nil, fmt.Errorf("Type does not exists.")
	}
	return timeMap, nil
}

func getMaxTokensForTime(timeMap map[int64]*Account, time int64) (int, bool) {
	if val, ok := timeMap[time]; ok {
		return val.Token, true
	}
	return 0, false
}

func checkTime(timeMap map[int64]*Account, time int64) (*Account, error) {
	// See if time exists in timeMap, if it doesn't find the max time that does exists and return it
	// If no time exists in that map return an error
	if val, ok := timeMap[time]; ok {
		return val, nil
	}
	var closestTime int64 = -1
	for t := range timeMap {
		if t <= time {
			if closestTime == -1 || t > closestTime {
				closestTime = t
			}
		}
	}

	if closestTime != -1 {
		return timeMap[closestTime], nil
	}
	return nil, fmt.Errorf("No matching time found")
}
