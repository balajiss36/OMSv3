package models

type RazorpayWebhook struct {
	Entity    string   `json:"entity,omitempty"`
	AccountID string   `json:"account_id,omitempty"`
	Event     string   `json:"event,omitempty"`
	Contains  []string `json:"contains,omitempty"`
	Payload   struct {
		Payment struct {
			Entity struct {
				ID                string `json:"id,omitempty"`
				Entity            string `json:"entity,omitempty"`
				Amount            int    `json:"amount,omitempty"`
				Currency          string `json:"currency,omitempty"`
				BaseAmount        int    `json:"base_amount,omitempty"`
				Status            string `json:"status,omitempty"`
				OrderID           string `json:"order_id,omitempty"`
				InvoiceID         any    `json:"invoice_id,omitempty"`
				International     bool   `json:"international,omitempty"`
				Method            string `json:"method,omitempty"`
				AmountRefunded    int    `json:"amount_refunded,omitempty"`
				AmountTransferred int    `json:"amount_transferred,omitempty"`
				RefundStatus      any    `json:"refund_status,omitempty"`
				Captured          bool   `json:"captured,omitempty"`
				Description       any    `json:"description,omitempty"`
				CardID            any    `json:"card_id,omitempty"`
				Bank              any    `json:"bank,omitempty"`
				Wallet            any    `json:"wallet,omitempty"`
				Vpa               string `json:"vpa,omitempty"`
				Email             string `json:"email,omitempty"`
				Contact           string `json:"contact,omitempty"`
				Notes             []any  `json:"notes,omitempty"`
				Fee               int    `json:"fee,omitempty"`
				Tax               int    `json:"tax,omitempty"`
				ErrorCode         any    `json:"error_code,omitempty"`
				ErrorDescription  any    `json:"error_description,omitempty"`
				ErrorSource       any    `json:"error_source,omitempty"`
				ErrorStep         any    `json:"error_step,omitempty"`
				ErrorReason       any    `json:"error_reason,omitempty"`
				AcquirerData      struct {
					Rrn string `json:"rrn,omitempty"`
				} `json:"acquirer_data,omitempty"`
				CreatedAt int `json:"created_at,omitempty"`
				Upi       struct {
					PayerAccountType string `json:"payer_account_type,omitempty"`
					Vpa              string `json:"vpa,omitempty"`
					Flow             string `json:"flow,omitempty"`
				} `json:"upi,omitempty"`
			} `json:"entity,omitempty"`
		} `json:"payment,omitempty"`
	} `json:"payload,omitempty"`
	CreatedAt int `json:"created_at,omitempty"`
}
