package stripe

import "encoding/json"

const (
	ErrorTypeAPI            string = "api_error"
	ErrorTypeAPIConnection  string = "api_connection_error"
	ErrorTypeAuthentication string = "authentication_error"
	ErrorTypeCard           string = "card_error"
	ErrorTypeInvalidRequest string = "invalid_request_error"
	ErrorTypePermission     string = "more_permissions_required"
	ErrorTypeRateLimit      string = "rate_limit_error"

	ErrorCodeCardDeclined       string = "card_declined"
	ErrorCodeExpiredCard        string = "expired_card"
	ErrorCodeIncorrectCVC       string = "incorrect_cvc"
	ErrorCodeIncorrectZip       string = "incorrect_zip"
	ErrorCodeIncorrectNumber    string = "incorrect_number"
	ErrorCodeInvalidCVC         string = "invalid_cvc"
	ErrorCodeInvalidExpiryMonth string = "invalid_expiry_month"
	ErrorCodeInvalidExpiryYear  string = "invalid_expiry_year"
	ErrorCodeInvalidNumber      string = "invalid_number"
	ErrorCodeInvalidSwipeData   string = "invalid_swipe_data"
	ErrorCodeMissing            string = "missing"
	ErrorCodeProcessingError    string = "processing_error"
	ErrorCodeRateLimit          string = "rate_limit"
)

// Error is the response returned when a call is unsuccessful.
// For more details see  https://stripe.com/docs/api#errors.
type Error struct {
	ChargeID string `json:"charge,omitempty"`
	Code     string `json:"code,omitempty"`

	// Err contains an internal error with an additional level of granularity
	// that can be used in some cases to get more detailed information about
	// what went wrong. For example, Err may hold a ChargeError that indicates
	// exactly what went wrong during a charge.
	Err error `json:"-"`

	HTTPStatusCode int    `json:"status,omitempty"`
	Msg            string `json:"message"`
	Param          string `json:"param,omitempty"`
	RequestID      string `json:"request_id,omitempty"`
	Type           string `json:"type"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// APIConnectionError is a failure to connect to the Stripe API.
type APIConnectionError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIConnectionError) Error() string {
	return e.stripeErr.Error()
}

// APIError is a catch all for any errors not covered by other types (and
// should be extremely uncommon).
type APIError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIError) Error() string {
	return e.stripeErr.Error()
}

// AuthenticationError is a failure to properly authenticate during a request.
type AuthenticationError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *AuthenticationError) Error() string {
	return e.stripeErr.Error()
}

// PermissionError results when you attempt to make an API request
// for which your API key doesn't have the right permissions.
type PermissionError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *PermissionError) Error() string {
	return e.stripeErr.Error()
}

// CardError are the most common type of error you should expect to handle.
// They result when the user enters a card that can't be charged for some
// reason.
type CardError struct {
	stripeErr   *Error
	DeclineCode string `json:"decline_code,omitempty"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *CardError) Error() string {
	return e.stripeErr.Error()
}

// InvalidRequestError is an error that occurs when a request contains invalid
// parameters.
type InvalidRequestError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *InvalidRequestError) Error() string {
	return e.stripeErr.Error()
}

// RateLimitError occurs when the Stripe API is hit to with too many requests
// too quickly and indicates that the current request has been rate limited.
type RateLimitError struct {
	stripeErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *RateLimitError) Error() string {
	return e.stripeErr.Error()
}
