package model

import "time"

type Booking struct {
  ID             int       `json:"id"`
  UserID         int       `json:"user_id"`
  EventID        int       `json:"event_id"`
  TicketQuantity int       `json:"ticket_quantity"`
  TotalPrice     int       `json:"total_price"`
  Status         string    `json:"status"`
  CreatedAt      time.Time `json:"created_at"`
  UpdatedAt      time.Time `json:"updated_at"`
}
