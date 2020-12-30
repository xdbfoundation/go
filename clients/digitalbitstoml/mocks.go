package digitalbitstoml

import "github.com/stretchr/testify/mock"

// MockClient is a mockable digitalbitstoml client.
type MockClient struct {
	mock.Mock
}

// GetDigitalBitsToml is a mocking a method
func (m *MockClient) GetDigitalBitsToml(domain string) (*Response, error) {
	a := m.Called(domain)
	return a.Get(0).(*Response), a.Error(1)
}

// GetDigitalBitsTomlByAddress is a mocking a method
func (m *MockClient) GetDigitalBitsTomlByAddress(address string) (*Response, error) {
	a := m.Called(address)
	return a.Get(0).(*Response), a.Error(1)
}
