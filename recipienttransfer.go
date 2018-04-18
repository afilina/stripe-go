package stripe

import "encoding/json"

const (
	RecipientTransferDestinationBankAccount string = "bank_account"
	RecipientTransferDestinationCard        string = "card"

	RecipientTransferFailureCodeAccountClosed         string = "account_closed"
	RecipientTransferFailureCodeAccountFrozen         string = "account_frozen"
	RecipientTransferFailureCodeBankAccountRestricted string = "bank_account_restricted"
	RecipientTransferFailureCodeBankOwnershipChanged  string = "bank_ownership_changed"
	RecipientTransferFailureCodeDebitNotAuthorized    string = "debit_not_authorized"
	RecipientTransferFailureCodeCouldNotProcess       string = "could_not_process"
	RecipientTransferFailureCodeInsufficientFunds     string = "insufficient_funds"
	RecipientTransferFailureCodeInvalidAccountNumber  string = "invalid_account_number"
	RecipientTransferFailureCodeInvalidCurrency       string = "invalid_currency"
	RecipientTransferFailureCodeNoAccount             string = "no_account"

	RecipientTransferMethodInstant  string = "instant"
	RecipientTransferMethodStandard string = "standard"

	RecipientTransferSourceTypeAlipayAccount   string = "alipay_account"
	RecipientTransferSourceTypeBankAccount     string = "bank_account"
	RecipientTransferSourceTypeBitcoinReceiver string = "bitcoin_receiver"
	RecipientTransferSourceTypeCard            string = "card"

	RecipientTransferStatusFailed    string = "failed"
	RecipientTransferStatusInTransit string = "in_transit"
	RecipientTransferStatusPaid      string = "paid"
	RecipientTransferStatusPending   string = "pending"

	RecipientTransferTypeBank string = "bank_account"
	RecipientTransferTypeCard string = "card"
)

// RecipientTransferDestination describes the destination of a RecipientTransfer.
// The Type should indicate which object is fleshed out
// For more details see https://stripe.com/docs/api/go#recipient_transfer_object
type RecipientTransferDestination struct {
	BankAccount *BankAccount `json:"-"`
	Card        *Card        `json:"-"`
	ID          string       `json:"id"`
	Type        string       `json:"object"`
}

// RecipientTransfer is the resource representing a Stripe recipient_transfer.
// For more details see https://stripe.com/docs/api#recipient_transfers.
type RecipientTransfer struct {
	Amount              int64                     `json:"amount"`
	AmountReversed      int64                     `json:"amount_reversed"`
	BalanceTransaction  *BalanceTransaction       `json:"balance_transaction"`
	BankAccount         *BankAccount              `json:"bank_account"`
	Card                *Card                     `json:"card"`
	Created             int64                     `json:"created"`
	Currency            Currency                  `json:"currency"`
	Date                int64                     `json:"date"`
	Description         string                    `json:"description"`
	Destination         string                    `json:"destination"`
	FailureCode         string                    `json:"failure_code"`
	FailureMessage      string                    `json:"failure_message"`
	ID                  string                    `json:"id"`
	Livemode            bool                      `json:"livemode"`
	Metadata            map[string]string         `json:"metadata"`
	Method              string                    `json:"method"`
	Recipient           *Recipient                `json:"recipient"`
	Reversals           *ReversalList             `json:"reversals"`
	Reversed            bool                      `json:"reversed"`
	SourceTransaction   *BalanceTransactionSource `json:"source_transaction"`
	SourceType          string                    `json:"source_type"`
	StatementDescriptor string                    `json:"statement_descriptor"`
	Status              string                    `json:"status"`
	Type                string                    `json:"type"`
}

// UnmarshalJSON handles deserialization of a RecipientTransfer.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *RecipientTransfer) UnmarshalJSON(data []byte) error {
	type transfer RecipientTransfer
	var tb transfer
	err := json.Unmarshal(data, &tb)
	if err == nil {
		*t = RecipientTransfer(tb)
	} else {
		// the id is surrounded by "\" characters, so strip them
		t.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// UnmarshalJSON handles deserialization of a RecipientTransferDestination.
// This custom unmarshaling is needed because the specific
// type of destination it refers to is specified in the JSON
func (d *RecipientTransferDestination) UnmarshalJSON(data []byte) error {
	type dest RecipientTransferDestination
	var dd dest
	err := json.Unmarshal(data, &dd)
	if err == nil {
		*d = RecipientTransferDestination(dd)

		switch d.Type {
		case RecipientTransferDestinationBankAccount:
			err = json.Unmarshal(data, &d.BankAccount)
		case RecipientTransferDestinationCard:
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

// MarshalJSON handles serialization of a RecipientTransferDestination.
// This custom marshaling is needed because we can only send a string
// ID as a destination, even though it can be expanded to a full
// object when retrieving
func (d *RecipientTransferDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ID)
}
