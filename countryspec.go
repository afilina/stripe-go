package stripe

// VerificationFieldsList lists the fields needed for an account verification.
// For more details see https://stripe.com/docs/api#country_spec_object-verification_fields.
type VerificationFieldsList struct {
	AdditionalFields []string `json:"additional"`
	Minimum          []string `json:"minimum"`
}

// CountrySpec is the resource representing the rules required for a Stripe account.
// For more details see https://stripe.com/docs/api/#country_specs.
type CountrySpec struct {
	DefaultCurrency                Currency                          `json:"default_currency"`
	ID                             string                            `json:"id"`
	SupportedBankAccountCurrencies map[Currency][]string             `json:"supported_bank_account_currencies"`
	SupportedPaymentCurrencies     []Currency                        `json:"supported_payment_currencies"`
	SupportedPaymentMethods        []string                          `json:"supported_payment_methods"`
	VerificationFields             map[string]VerificationFieldsList `json:"verification_fields"`
}

// CountrySpecList is a list of country specs as retrieved from a list endpoint.
type CountrySpecList struct {
	ListMeta
	Data []*CountrySpec `json:"data"`
}

// CountrySpecListParams are the parameters allowed during CountrySpec listing.
type CountrySpecListParams struct {
	ListParams `form:"*"`
}
