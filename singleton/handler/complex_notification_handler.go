package handler

import (
	"github.com/gin-gonic/gin"
	"go-design-patterns/abstract-factory/notification"
)

// NotificationHandler handles sending notifications using the NotificationService
type ComplexNotificationHandler struct {
	EmailService notification.EmailNotification
	SMSService   notification.SMSNotification
}

// NewNotificationHandler creates a new NotificationHandler with the appropriate services
func NewComplexNotificationHandler(factory notification.NotificationProviderFactory) *ComplexNotificationHandler {
	emailService := factory.CreateEmailNotification()
	smsService := factory.CreateSMSNotification()
	return &ComplexNotificationHandler{
		EmailService: emailService,
		SMSService:   smsService,
	}
}

// SendEmail handles the request to send an email notification
func (h *ComplexNotificationHandler) SendEmail(c *gin.Context) {
	message := c.DefaultQuery("message", "Default email message")
	result := h.EmailService.SendEmail(message)
	c.JSON(200, gin.H{"email_result": result})
}

// SendSMS handles the request to send an SMS notification
func (h *ComplexNotificationHandler) SendSMS(c *gin.Context) {
	message := c.DefaultQuery("message", "Default SMS message")
	result := h.SMSService.SendSMS(message)
	c.JSON(200, gin.H{"sms_result": result})
}
