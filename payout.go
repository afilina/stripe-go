package stripe

import "encoding/json"

const (
	PayoutDestinationTypeBankAccount string = "bank_account"
	PayoutDestinationTypeCard        string = "card"

	PayoutTypeBank string = "bank_account"
	PayoutTypeCard string = "card"

	PayoutFailureCodeAccountClosed         string = "account_closed"
	PayoutFailureCodeAccountFrozen         string = "account_frozen"
	PayoutFailureCodeBankAccountRestricted string = "bank_account_restricted"
	PayoutFailureCodeBankOwnershipChanged  string = "bank_ownership_changed"
	PayoutFailureCodeCouldNotProcess       string = "could_not_process"
	PayoutFailureCodeDebitNotAuthorized    string = "debit_not_authorized"
	PayoutFailureCodeInsufficientFunds     string = "insufficient_funds"
	PayoutFailureCodeInvalidAccountNumber  string = "invalid_account_number"
	PayoutFailureCodeInvalidCurrency       string = "invalid_currency"
	PayoutFailureCodeNoAccount             string = "no_account"

	PayoutMethodInstant  string = "instant"
	PayoutMethodStandard string = "standard"

	PayoutSourceTypeAlipayAccount   string = "alipay_account"
	PayoutSourceTypeBankAccount     string = "bank_account"
	PayoutSourceTypeBitcoinReceiver string = "bitcoin_receiver"
	PayoutSourceTypeCard            string = "card"

	PayoutStatusCanceled  string = "canceled"
	PayoutStatusFailed    string = "failed"
	PayoutStatusInTransit string = "in_transit"
	PayoutStatusPaid      string = "paid"
	PayoutStatusPending   string = "pending"
)

// PayoutDestination describes the destination of a Payout.
// The Type should indicate which object is fleshed out
// For more details see https://stripe.com/docs/api/go#payout_object
type PayoutDestination struct {
	BankAccount *BankAccount `json:"-"`
	Card        *Card        `json:"-"`
	ID          string       `json:"id"`
	Type        string       `json:"object"`
}

// PayoutParams is the set of parameters that can be used when creating or updating a payout.
// For more details see https://stripe.com/docs/api#create_payout and https://stripe.com/docs/api#update_payout.
type PayoutParams struct {
	Params              `form:"*"`
	Amount              *int64  `form:"amount"`
	Currency            *string `form:"currency"`
	Destination         *string `form:"destination"`
	Method              *string `form:"method"`
	SourceType          *string `form:"source_type"`
	StatementDescriptor *string `form:"statement_descriptor"`
}

// PayoutListParams is the set of parameters that can be used when listing payouts.
// For more details see https://stripe.com/docs/api#list_payouts.
type PayoutListParams struct {
	ListParams       `form:"*"`
	ArrivalDate      *int64            `form:"arrival_date"`
	ArrivalDateRange *RangeQueryParams `form:"arrival_date"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	Destination      *string           `form:"destination"`
	Status           *string           `form:"status"`
}

// Payout is the resource representing a Stripe payout.
// For more details see https://stripe.com/docs/api#payouts.
type Payout struct {
	Amount                    int64               `json:"amount"`
	ArrivalDate               int64               `json:"arrival_date"`
	Automatic                 bool                `json:"automatic"`
	BalanceTransaction        *BalanceTransaction `json:"balance_transaction"`
	BankAccount               *BankAccount        `json:"bank_account"`
	Card                      *Card               `json:"card"`
	Created                   int64               `json:"created"`
	Currency                  Currency            `json:"currency"`
	Destination               string              `json:"destination"`
	FailureBalanceTransaction *BalanceTransaction `json:"failure_balance_transaction"`
	FailureCode               string              `json:"failure_code"`
	FailureMessage            string              `json:"failure_message"`
	ID                        string              `json:"id"`
	Livemode                  bool                `json:"livemode"`
	Metadata                  map[string]string   `json:"metadata"`
	Method                    string              `json:"method"`
	SourceType                string              `json:"source_type"`
	StatementDescriptor       string              `json:"statement_descriptor"`
	Status                    string              `json:"status"`
	Type                      string              `json:"type"`
}

// PayoutList is a list of payouts as retrieved from a list endpoint.
type PayoutList struct {
	ListMeta
	Data []*Payout `json:"data"`
}

// UnmarshalJSON handles deserialization of a Payout.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *Payout) UnmarshalJSON(data []byte) error {
	type payout Payout
	var tb payout
	err := json.Unmarshal(data, &tb)
	if err == nil {
		*t = Payout(tb)
	} else {
		// the id is surrounded by "\" characters, so strip them
		t.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// UnmarshalJSON handles deserialization of a PayoutDestination.
// This custom unmarshaling is needed because the specific
// type of destination it refers to is specified in the JSON
func (d *PayoutDestination) UnmarshalJSON(data []byte) error {
	type dest PayoutDestination
	var dd dest
	err := json.Unmarshal(data, &dd)
	if err == nil {
		*d = PayoutDestination(dd)

		switch d.Type {
		case PayoutDestinationTypeBankAccount:
			err = json.Unmarshal(data, &d.BankAccount)
		case PayoutDestinationTypeCard:
			err = json.Unmarshal(data, &d.Card)
		}

		if err != nil {
			return err
		}
	} else {
		// the id is surrounded by "\" characters, so strip them
		d.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// MarshalJSON handles serialization of a PayoutDestination.
// This custom marshaling is needed because we can only send a string
// ID as a destination, even though it can be expanded to a full
// object when retrieving
func (d *PayoutDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ID)
}
