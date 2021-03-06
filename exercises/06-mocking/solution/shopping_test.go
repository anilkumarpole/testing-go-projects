package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSuccessfulCheckoutShoppingCart(t *testing.T) {
	givenNotificationMsg := "Successfully purchased 10 books"
	notifier := new(notifierMock)
	notifier.On("SendMessage", givenNotificationMsg).Return(nil)

	purchaseService := PurchaseService{notifier: notifier}
	err := purchaseService.CheckoutShoppingCart(99, givenNotificationMsg)

	notifier.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestFailedCheckoutShoppingCart(t *testing.T) {
	givenNotificationMsg := "Successfully purchased 10 books"
	expectedErrMsg := "failed to send message"
	notifier := new(notifierMock)
	notifier.On("SendMessage", givenNotificationMsg).Return(errors.New(expectedErrMsg))

	purchaseService := PurchaseService{notifier: notifier}
	err := purchaseService.CheckoutShoppingCart(99, givenNotificationMsg)

	notifier.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErrMsg, err.Error())
}

type notifierMock struct {
	mock.Mock
}

func (m *notifierMock) SendMessage(message string) error {
	args := m.Called(message)
	return args.Error(0)
}
