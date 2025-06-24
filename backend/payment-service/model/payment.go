package model

import "time"

type Payment struct {
    ID         int       `json:"id"`
    BookingID  int       `json:"booking_id"`
    UserID     int       `json:"user_id"`
    Amount     int       `json:"amount"`
    Status     string    `json:"status"`
    PaymentID  string    `json:"payment_id"`
    PaymentURL string    `json:"payment_url"`
    CreatedAt  time.Time `json:"created_at"`
}
