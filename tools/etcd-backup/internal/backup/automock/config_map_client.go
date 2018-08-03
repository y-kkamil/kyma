// Code generated by mockery v1.0.0
package automock

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/core/v1"

// configMapClient is an autogenerated mock type for the configMapClient type
type configMapClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *configMapClient) Create(_a0 *v1.ConfigMap) (*v1.ConfigMap, error) {
	ret := _m.Called(_a0)

	var r0 *v1.ConfigMap
	if rf, ok := ret.Get(0).(func(*v1.ConfigMap) *v1.ConfigMap); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ConfigMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.ConfigMap) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: name, options
func (_m *configMapClient) Get(name string, options metav1.GetOptions) (*v1.ConfigMap, error) {
	ret := _m.Called(name, options)

	var r0 *v1.ConfigMap
	if rf, ok := ret.Get(0).(func(string, metav1.GetOptions) *v1.ConfigMap); ok {
		r0 = rf(name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ConfigMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, metav1.GetOptions) error); ok {
		r1 = rf(name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *configMapClient) Update(_a0 *v1.ConfigMap) (*v1.ConfigMap, error) {
	ret := _m.Called(_a0)

	var r0 *v1.ConfigMap
	if rf, ok := ret.Get(0).(func(*v1.ConfigMap) *v1.ConfigMap); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ConfigMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.ConfigMap) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
