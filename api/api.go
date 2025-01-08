package api

// Balance Params
type BankBalanceParams struct {
	Username string
}

// Coin Balance Response
type BankBalanceResponse struct {
	// Success Code
	Code int

	//Account Balance
	Balance int64
}

type Error struct {
	// Error code
	Code int

	// Error message
	Message string
}
