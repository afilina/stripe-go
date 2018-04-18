package stripe

import "encoding/json"

const (
	ReviewReasonApproved        string = "approved"
	ReviewReasonDisputed        string = "disputed"
	ReviewReasonManual          string = "manual"
	ReviewReasonRefunded        string = "refunded"
	ReviewReasonRefundedAsFraud string = "refunded_as_fraud"
	ReviewReasonRule            string = "rule"
)

type Review struct {
	Charge   *Charge `json:"charge"`
	Created  int64   `json:"created"`
	ID       string  `json:"id"`
	Livemode bool    `json:"livemode"`
	Open     bool    `json:"open"`
	Reason   string  `json:"reason"`
}

func (r *Review) UnmarshalJSON(data []byte) error {
	type review Review
	var rr review

	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Review(rr)
	} else {
		// Otherwise...we have to strip the escaping
		r.ID = string(data[1 : len(data)-1])
	}

	return nil
}
