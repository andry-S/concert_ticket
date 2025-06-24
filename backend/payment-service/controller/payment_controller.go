package controller

import (
    "payment-service/config"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "fmt"
)

func CreatePayment(c *gin.Context) {
    var req struct {
        BookingID int `json:"booking_id"`
        UserID    int `json:"user_id"`
        Amount    int `json:"amount"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if req.BookingID <= 0 || req.UserID <= 0 || req.Amount <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "BookingID, UserID, and Amount must be greater than 0"})
        return
    }

  
    var exists bool
    err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM payments WHERE booking_id = ?)", req.BookingID).Scan(&exists)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing payment"})
        return
    }
    if exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Payment already exists for this booking"})
        return
    }

    now := time.Now()
    paymentID := fmt.Sprintf("PAY-%d", now.UnixNano())
    paymentURL := fmt.Sprintf("http://dummy-gateway/pay/%s", paymentID)

    _, err = config.DB.Exec(`INSERT INTO payments
        (booking_id, user_id, amount, status, payment_id, payment_url, created_at)
        VALUES (?, ?, ?, 'pending', ?, ?, ?)`,
        req.BookingID, req.UserID, req.Amount, paymentID, paymentURL, now)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":     "Payment created",
        "payment_id":  paymentID,
        "payment_url": paymentURL,
        "status":      "pending",
    })
}
