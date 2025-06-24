package controller

import (
    "booking-service/config"
    "booking-service/model"
    "encoding/json"
    "net/http"
    "time"
    "strconv"
    "bytes"

    "github.com/gin-gonic/gin"
)

const pricePerTicket = 250000

func CreateBooking(c *gin.Context) {
    var req struct {
        UserID         int `json:"user_id"`
        EventID        int `json:"event_id"`
        TicketQuantity int `json:"ticket_quantity"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    total := req.TicketQuantity * pricePerTicket

    res, err := config.DB.Exec(`INSERT INTO bookings (user_id, event_id, ticket_quantity, total_price, status)
        VALUES (?, ?, ?, ?, 'pending')`, req.UserID, req.EventID, req.TicketQuantity, total)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
        return
    }

    insertID, _ := res.LastInsertId()
    booking := model.Booking{
        ID:             int(insertID),
        UserID:         req.UserID,
        EventID:        req.EventID,
        TicketQuantity: req.TicketQuantity,
        TotalPrice:     total,
        Status:         "pending",
        CreatedAt:      time.Now(),
        UpdatedAt:      time.Now(),
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Booking created",
        "booking": booking,
    })
}

func GetAllBookings(c *gin.Context) {
    rows, err := config.DB.Query("SELECT id, user_id, event_id, ticket_quantity, total_price, status, created_at, updated_at FROM bookings")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
        return
    }
    defer rows.Close()

    var bookings []model.Booking
    for rows.Next() {
        var b model.Booking
        if err := rows.Scan(&b.ID, &b.UserID, &b.EventID, &b.TicketQuantity, &b.TotalPrice, &b.Status, &b.CreatedAt, &b.UpdatedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan failed"})
            return
        }
        bookings = append(bookings, b)
    }

    c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func GetBookingByID(c *gin.Context) {
    id := c.Param("id")
    row := config.DB.QueryRow("SELECT id, user_id, event_id, ticket_quantity, total_price, status, created_at, updated_at FROM bookings WHERE id = ?", id)

    var b model.Booking
    if err := row.Scan(&b.ID, &b.UserID, &b.EventID, &b.TicketQuantity, &b.TotalPrice, &b.Status, &b.CreatedAt, &b.UpdatedAt); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
        return
    }

    c.JSON(http.StatusOK, b)
}

func GetBookingsByUserID(c *gin.Context) {
    userID := c.Param("id")

    rows, err := config.DB.Query(`
        SELECT id, user_id, event_id, ticket_quantity, total_price, status, created_at, updated_at
        FROM bookings
        WHERE user_id = ?`, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user bookings"})
        return
    }
    defer rows.Close()

    var bookings []model.Booking
    for rows.Next() {
        var b model.Booking
        if err := rows.Scan(&b.ID, &b.UserID, &b.EventID, &b.TicketQuantity, &b.TotalPrice, &b.Status, &b.CreatedAt, &b.UpdatedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan failed"})
            return
        }
        bookings = append(bookings, b)
    }

    c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func DeleteBooking(c *gin.Context) {
    id := c.Param("id")
    _, err := config.DB.Exec("DELETE FROM bookings WHERE id = ?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Booking deleted"})
}

func PayBooking(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
        return
    }

    var b model.Booking
    row := config.DB.QueryRow("SELECT id, user_id, event_id, ticket_quantity, total_price, status FROM bookings WHERE id = ?", id)
    if err := row.Scan(&b.ID, &b.UserID, &b.EventID, &b.TicketQuantity, &b.TotalPrice, &b.Status); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
        return
    }

    if b.Status != "pending" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Booking already paid or invalid status"})
        return
    }

    payload := map[string]interface{}{
        "booking_id": b.ID,
        "user_id":    b.UserID,
        "amount":     b.TotalPrice,
    }

    payloadBytes, err := json.Marshal(payload)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal payment payload"})
        return
    }

    resp, err := http.Post("http://localhost:8082/payments", "application/json", bytes.NewBuffer(payloadBytes))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to payment service"})
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        c.JSON(http.StatusBadGateway, gin.H{"error": "Payment service returned error", "code": resp.StatusCode})
        return
    }

    var paymentResp struct {
        PaymentID  string `json:"payment_id"`
        PaymentURL string `json:"payment_url"`
        Status     string `json:"status"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&paymentResp); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response from payment service"})
        return
    }

    _, err = config.DB.Exec("UPDATE bookings SET payment_id = ?, status = 'waiting_payment', updated_at = ? WHERE id = ?", paymentResp.PaymentID, time.Now(), b.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":      "Payment initiated",
        "payment_id":   paymentResp.PaymentID,
        "payment_url":  paymentResp.PaymentURL,
        "booking_id":   b.ID,
        "total_amount": b.TotalPrice,
    })
  }
