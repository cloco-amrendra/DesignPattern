package handler

import (
	"github.com/gin-gonic/gin"
	"go-design-patterns/factory/notification"
)

// NotificationHandler handles sending notifications using the NotificationService
type NotificationHandler struct {
	Service notification.NotificationService
}

// NewNotificationHandler creates a new NotificationHandler with the appropriate service
func NewNotificationHandler(factory notification.NotificationServiceFactory, notificationType string) *NotificationHandler {
	service := factory.CreateNotificationService(notificationType)
	return &NotificationHandler{Service: service}
}

// SendNotification handles the request to send a notification
func (h *NotificationHandler) SendNotification(c *gin.Context) {
	message := c.DefaultQuery("message", "Default notification message") // Get message from query param or use default
	notificationResult := h.Service.SendNotification(message)
	c.JSON(200, gin.H{"message": notificationResult})
}
