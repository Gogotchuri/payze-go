package payze_go

//Connection based constants
const (
	PayZeApiV1URL     = "https://payze.io/api/v1"
	DefaultHTTPMethod = "POST"
)

//Action Method names
const (
	JustPay             = "justPay"
	JustPayAndSplit     = JustPay
	AddCard             = "addCard"
	PayWithCard         = "payWithCard"
	PayWithCardAndSplit = PayWithCard
	TransactionInfo     = "getTransactionInfo"
	Refund              = "refund"
	Balance             = "getBalance"
	CommitTransaction   = "commit"
)

//Currencies
const (
	GEL = "GEL"
	USD = "USD"
)
