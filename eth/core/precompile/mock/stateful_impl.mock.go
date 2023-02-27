// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"pkg.berachain.dev/stargazer/eth/core/precompile"
	"sync"
)

// StatefulImplMock is a mock implementation of precompile.StatefulImpl.
//
//	func TestSomethingThatUsesStatefulImpl(t *testing.T) {
//
//		// make and configure a mocked precompile.StatefulImpl
//		mockedStatefulImpl := &StatefulImplMock{
//			ABIEventsFunc: func() map[string]abi.Event {
//				panic("mock out the ABIEvents method")
//			},
//			ABIMethodsFunc: func() map[string]abi.Method {
//				panic("mock out the ABIMethods method")
//			},
//			CustomValueDecodersFunc: func() precompile.ValueDecoders {
//				panic("mock out the CustomValueDecoders method")
//			},
//			PrecompileMethodsFunc: func() precompile.Methods {
//				panic("mock out the PrecompileMethods method")
//			},
//			RegistryKeyFunc: func() common.Address {
//				panic("mock out the RegistryKey method")
//			},
//		}
//
//		// use mockedStatefulImpl in code that requires precompile.StatefulImpl
//		// and then make assertions.
//
//	}
type StatefulImplMock struct {
	// ABIEventsFunc mocks the ABIEvents method.
	ABIEventsFunc func() map[string]abi.Event

	// ABIMethodsFunc mocks the ABIMethods method.
	ABIMethodsFunc func() map[string]abi.Method

	// CustomValueDecodersFunc mocks the CustomValueDecoders method.
	CustomValueDecodersFunc func() precompile.ValueDecoders

	// PrecompileMethodsFunc mocks the PrecompileMethods method.
	PrecompileMethodsFunc func() precompile.Methods

	// RegistryKeyFunc mocks the RegistryKey method.
	RegistryKeyFunc func() common.Address

	// calls tracks calls to the methods.
	calls struct {
		// ABIEvents holds details about calls to the ABIEvents method.
		ABIEvents []struct {
		}
		// ABIMethods holds details about calls to the ABIMethods method.
		ABIMethods []struct {
		}
		// CustomValueDecoders holds details about calls to the CustomValueDecoders method.
		CustomValueDecoders []struct {
		}
		// PrecompileMethods holds details about calls to the PrecompileMethods method.
		PrecompileMethods []struct {
		}
		// RegistryKey holds details about calls to the RegistryKey method.
		RegistryKey []struct {
		}
	}
	lockABIEvents           sync.RWMutex
	lockABIMethods          sync.RWMutex
	lockCustomValueDecoders sync.RWMutex
	lockPrecompileMethods   sync.RWMutex
	lockRegistryKey         sync.RWMutex
}

// ABIEvents calls ABIEventsFunc.
func (mock *StatefulImplMock) ABIEvents() map[string]abi.Event {
	if mock.ABIEventsFunc == nil {
		panic("StatefulImplMock.ABIEventsFunc: method is nil but StatefulImpl.ABIEvents was just called")
	}
	callInfo := struct {
	}{}
	mock.lockABIEvents.Lock()
	mock.calls.ABIEvents = append(mock.calls.ABIEvents, callInfo)
	mock.lockABIEvents.Unlock()
	return mock.ABIEventsFunc()
}

// ABIEventsCalls gets all the calls that were made to ABIEvents.
// Check the length with:
//
//	len(mockedStatefulImpl.ABIEventsCalls())
func (mock *StatefulImplMock) ABIEventsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockABIEvents.RLock()
	calls = mock.calls.ABIEvents
	mock.lockABIEvents.RUnlock()
	return calls
}

// ABIMethods calls ABIMethodsFunc.
func (mock *StatefulImplMock) ABIMethods() map[string]abi.Method {
	if mock.ABIMethodsFunc == nil {
		panic("StatefulImplMock.ABIMethodsFunc: method is nil but StatefulImpl.ABIMethods was just called")
	}
	callInfo := struct {
	}{}
	mock.lockABIMethods.Lock()
	mock.calls.ABIMethods = append(mock.calls.ABIMethods, callInfo)
	mock.lockABIMethods.Unlock()
	return mock.ABIMethodsFunc()
}

// ABIMethodsCalls gets all the calls that were made to ABIMethods.
// Check the length with:
//
//	len(mockedStatefulImpl.ABIMethodsCalls())
func (mock *StatefulImplMock) ABIMethodsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockABIMethods.RLock()
	calls = mock.calls.ABIMethods
	mock.lockABIMethods.RUnlock()
	return calls
}

// CustomValueDecoders calls CustomValueDecodersFunc.
func (mock *StatefulImplMock) CustomValueDecoders() precompile.ValueDecoders {
	if mock.CustomValueDecodersFunc == nil {
		panic("StatefulImplMock.CustomValueDecodersFunc: method is nil but StatefulImpl.CustomValueDecoders was just called")
	}
	callInfo := struct {
	}{}
	mock.lockCustomValueDecoders.Lock()
	mock.calls.CustomValueDecoders = append(mock.calls.CustomValueDecoders, callInfo)
	mock.lockCustomValueDecoders.Unlock()
	return mock.CustomValueDecodersFunc()
}

// CustomValueDecodersCalls gets all the calls that were made to CustomValueDecoders.
// Check the length with:
//
//	len(mockedStatefulImpl.CustomValueDecodersCalls())
func (mock *StatefulImplMock) CustomValueDecodersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockCustomValueDecoders.RLock()
	calls = mock.calls.CustomValueDecoders
	mock.lockCustomValueDecoders.RUnlock()
	return calls
}

// PrecompileMethods calls PrecompileMethodsFunc.
func (mock *StatefulImplMock) PrecompileMethods() precompile.Methods {
	if mock.PrecompileMethodsFunc == nil {
		panic("StatefulImplMock.PrecompileMethodsFunc: method is nil but StatefulImpl.PrecompileMethods was just called")
	}
	callInfo := struct {
	}{}
	mock.lockPrecompileMethods.Lock()
	mock.calls.PrecompileMethods = append(mock.calls.PrecompileMethods, callInfo)
	mock.lockPrecompileMethods.Unlock()
	return mock.PrecompileMethodsFunc()
}

// PrecompileMethodsCalls gets all the calls that were made to PrecompileMethods.
// Check the length with:
//
//	len(mockedStatefulImpl.PrecompileMethodsCalls())
func (mock *StatefulImplMock) PrecompileMethodsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockPrecompileMethods.RLock()
	calls = mock.calls.PrecompileMethods
	mock.lockPrecompileMethods.RUnlock()
	return calls
}

// RegistryKey calls RegistryKeyFunc.
func (mock *StatefulImplMock) RegistryKey() common.Address {
	if mock.RegistryKeyFunc == nil {
		panic("StatefulImplMock.RegistryKeyFunc: method is nil but StatefulImpl.RegistryKey was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRegistryKey.Lock()
	mock.calls.RegistryKey = append(mock.calls.RegistryKey, callInfo)
	mock.lockRegistryKey.Unlock()
	return mock.RegistryKeyFunc()
}

// RegistryKeyCalls gets all the calls that were made to RegistryKey.
// Check the length with:
//
//	len(mockedStatefulImpl.RegistryKeyCalls())
func (mock *StatefulImplMock) RegistryKeyCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRegistryKey.RLock()
	calls = mock.calls.RegistryKey
	mock.lockRegistryKey.RUnlock()
	return calls
}
