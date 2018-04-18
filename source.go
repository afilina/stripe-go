package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/form"
)

const (
	SourceCodeVerificationFlowStatusFailed    string = "failed"
	SourceCodeVerificationFlowStatusPending   string = "pending"
	SourceCodeVerificationFlowStatusSucceeded string = "succeeded"

	SourceFlowCodeVerification string = "code_verification"
	SourceFlowNone             string = "none"
	SourceFlowReceiver         string = "receiver"
	SourceFlowRedirect         string = "redirect"

	SourceRedirectFlowFailureReasonDeclined        string = "declined"
	SourceRedirectFlowFailureReasonProcessingError string = "processing_error"
	SourceRedirectFlowFailureReasonUserAbort       string = "user_abort"

	SourceRedirectFlowStatusFailed      string = "failed"
	SourceRedirectFlowStatusNotRequired string = "not_required"
	SourceRedirectFlowStatusPending     string = "pending"
	SourceRedirectFlowStatusSucceeded   string = "succeeded"

	SourceRefundAttributesMethodEmail  string = "email"
	SourceRefundAttributesMethodManual string = "manual"

	SourceRefundAttributesStatusAvailable string = "available"
	SourceRefundAttributesStatusMissing   string = "missing"
	SourceRefundAttributesStatusRequested string = "requested"

	SourceStatusCanceled   string = "canceled"
	SourceStatusChargeable string = "chargeable"
	SourceStatusConsumed   string = "consumed"
	SourceStatusFailed     string = "failed"
	SourceStatusPending    string = "pending"

	SourceUsageReusable  string = "reusable"
	SourceUsageSingleUse string = "single_use"
)

type SourceOwnerParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

type RedirectParams struct {
	ReturnURL *string `form:"return_url"`
}

type SourceObjectParams struct {
	Params              `form:"*"`
	Amount              *int64             `form:"amount"`
	Currency            *string            `form:"currency"`
	Customer            *string            `form:"customer"`
	Flow                *string            `form:"flow"`
	OriginalSource      *string            `form:"original_source"`
	Owner               *SourceOwnerParams `form:"owner"`
	Redirect            *RedirectParams    `form:"redirect"`
	StatementDescriptor *string            `form:"statement_descriptor"`
	Token               *string            `form:"token"`
	Type                *string            `form:"type"`
	TypeData            map[string]string  `form:"-"`
	Usage               *string            `form:"usage"`
}

// SourceObjectDetachParams is the set of parameters that can be used when detaching
// a source from a customer.
type SourceObjectDetachParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"`
}

type SourceOwner struct {
	Address         *Address `json:"address,omitempty"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	Phone           string   `json:"phone"`
	VerifiedAddress *Address `json:"verified_address,omitempty"`
	VerifiedEmail   string   `json:"verified_email"`
	VerifiedName    string   `json:"verified_name"`
	VerifiedPhone   string   `json:"verified_phone"`
}

// ReceiverFlow informs of the state of a redirect authentication flow.
type RedirectFlow struct {
	FailureReason string `json:"failure_reason"`
	ReturnURL     string `json:"return_url"`
	Status        string `json:"status"`
	URL           string `json:"url"`
}

// ReceiverFlow informs of the state of a receiver authentication flow.
type ReceiverFlow struct {
	Address                string `json:"address"`
	AmountCharged          int64  `json:"amount_charged"`
	AmountReceived         int64  `json:"amount_received"`
	AmountReturned         int64  `json:"amount_returned"`
	RefundAttributesMethod string `json:"refund_attributes_method"`
	RefundAttributesStatus string `json:"refund_attributes_status"`
}

// CodeVerificationFlow informs of the state of a verification authentication flow.
type CodeVerificationFlow struct {
	AttemptsRemaining int64  `json:"attempts_remaining"`
	Status            string `json:"status"`
}

type SourceMandateAcceptance struct {
	Date      string `json:"date"`
	IP        string `json:"ip"`
	Status    string `json:"status"`
	UserAgent string `json:"user_agent"`
}

type SourceMandate struct {
	Acceptance         SourceMandateAcceptance `json:"acceptance"`
	NotificationMethod string                  `json:"notification_method"`
	Reference          string                  `json:"reference"`
	URL                string                  `json:"url"`
}

type Source struct {
	Amount              int64                 `json:"amount"`
	ClientSecret        string                `json:"client_secret"`
	CodeVerification    *CodeVerificationFlow `json:"code_verification,omitempty"`
	Created             int64                 `json:"created"`
	Currency            Currency              `json:"currency"`
	Flow                string                `json:"flow"`
	ID                  string                `json:"id"`
	Livemode            bool                  `json:"livemode"`
	Mandate             SourceMandate         `json:"mandate"`
	Metadata            map[string]string     `json:"metadata"`
	Owner               SourceOwner           `json:"owner"`
	Receiver            *ReceiverFlow         `json:"receiver,omitempty"`
	Redirect            *RedirectFlow         `json:"redirect,omitempty"`
	StatementDescriptor string                `json:"statement_descriptor"`
	Status              string                `json:"status"`
	Type                string                `json:"type"`
	TypeData            map[string]interface{}
	Usage               string `json:"usage"`
}

// AppendTo implements custom encoding logic for SourceObjectParams so that the special
// "TypeData" value for is sent as the correct parameter based on the Source type
func (p *SourceObjectParams) AppendTo(body *form.Values, keyParts []string) {
	if len(p.TypeData) > 0 && p.Type == nil {
		panic("You can not fill TypeData if you don't explicitly set Type")
	}

	for k, vs := range p.TypeData {
		body.Add(form.FormatKey(append(keyParts, StringValue(p.Type), k)), vs)
	}
}

// UnmarshalJSON handles deserialization of an Source. This custom unmarshaling
// is needed to extract the type specific data (accessible under `TypeData`)
// but stored in JSON under a hash named after the `type` of the source.
func (s *Source) UnmarshalJSON(data []byte) error {
	type source Source
	var ss source
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}
	*s = Source(ss)

	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	if d, ok := raw[s.Type]; ok {
		if m, ok := d.(map[string]interface{}); ok {
			s.TypeData = m
		}
	}

	return nil
}
