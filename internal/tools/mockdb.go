package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"robert": {
		AuthToken: "123ABC",
		Username:  "robert",
	},
	"lauren": {
		AuthToken: "456DEF",
		Username:  "lauren",
	},
	"kendall": {
		AuthToken: "789GHI",
		Username:  "kendall",
	},
	"matt": {
		AuthToken: "123JKL",
		Username:  "matt",
	},
}

var mockBankDetails = map[string]BankDetails{
	"robert": {
		Balance:  100,
		Username: "robert",
	},
	"lauren": {
		Balance:  200,
		Username: "lauren",
	},
	"kendall": {
		Balance:  300,
		Username: "kendall",
	},
	"matt": {
		Balance:  400,
		Username: "matt",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUserBalance(username string) *BankDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = BankDetails{}
	clientData, ok := mockBankDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}
