package tests

import (
	"net"
)

type Mock struct {
	ReturnValues map[string]interface{}	
	CalledMethods map[string]interface{}
}

func (mock *Mock) GetMethodParams(method string) (interface{}) {
	return mock.CalledMethods[method]
}

func (mock *Mock) WasMethodCalled(method string) (bool) {
	_, called := mock.CalledMethods[method]
	return called
}

func (mock *Mock) SetReturnValue(method string, val interface{}) {
	mock.ReturnValues[method] = val
}

type MockConn struct {
	net.Conn
	Mock
}

func NewMockConn() (*MockConn) {
	conn := new(MockConn)
	conn.CalledMethods = make(map[string]interface{})
	return conn
}

func (conn *MockConn) Close() (error) {
	conn.CalledMethods["Close"] = true
	return nil
}

func (conn *MockConn) Read(b []byte) (int, error) {
	conn.CalledMethods["Read"] = true
	return 0, nil
}

func (conn *MockConn) Write(b []byte) (int, error) {
	conn.CalledMethods["Write"] = b
	return 0, nil
}