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
	CreateSubscription  = "subscription"
	SubscriptionStatus  = "subscriptionStatus"
	CancelSubscription  = "cancelSubscription"
)

//Transaction Statuses
const (
	TransactionBlocked                  = "Blocked"
	TransactionCardAdded                = "CardAdded"
	TransactionCardAddedForSubscription = "CardAddedForSubscription"
	TransactionCommitFailed             = "CommitFailed"
	TransactionCommitted                = "Committed"
	TransactionCreated                  = "Created"
	TransactionNothing                  = "Nothing"
	TransactionError                    = "Error"
	TransactionPlatformReceived         = "PlatformReceived"
	TransactionRefunded                 = "Refunded"
	TransactionRefundedPartially        = "RefundedPartially"
	TransactionRejected                 = "Rejected"
	TransactionTimeout                  = "Timeout"
)

const (
	SubscriptionCreated   = "Created"
	SubscriptionStarted   = "Started"
	SubscriptionActive    = "Active"
	SubscriptionCancelled = "Cancelled"
)

const (
	SubscriptionIntervalTypeMonth = "Month"
	SubscriptionIntervalTypeDay   = "Day"
	SubscriptionIntervalTypeWeek  = "Week"
	SubscriptionIntervalTypeYear  = "Year"
)

const (
	ChannelEmail = "EMAIL"
	ChannelSMS   = "SMS"
)

//Currencies
const (
	GEL = "GEL"
	USD = "USD"
)
