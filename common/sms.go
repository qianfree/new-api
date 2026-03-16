package common

import (
	"errors"
	"sync"
)

type SMSSender interface {
	Send(phone, code string) error
}

var (
	SMSProvider      string
	smsProviders     = make(map[string]SMSSender)
	smsProviderMutex sync.RWMutex
)

func RegisterSMSProvider(name string, provider SMSSender) {
	smsProviderMutex.Lock()
	defer smsProviderMutex.Unlock()
	smsProviders[name] = provider
}

func SendSMS(phone, code string) error {
	if SMSProvider == "" {
		return errors.New("SMS provider not configured")
	}
	smsProviderMutex.RLock()
	provider, ok := smsProviders[SMSProvider]
	smsProviderMutex.RUnlock()
	if !ok {
		return errors.New("SMS provider not configured: " + SMSProvider)
	}
	return provider.Send(phone, code)
}
