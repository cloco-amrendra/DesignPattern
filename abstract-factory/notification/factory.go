package notification

// NotificationProviderFactory defines methods to create email and SMS notifications
type NotificationProviderFactory interface {
	CreateEmailNotification() EmailNotification
	CreateSMSNotification() SMSNotification
}
