// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import pager "github.com/kyma-project/kyma/components/console-backend-service/internal/pager"

import v1alpha1 "github.com/kyma-project/helm-broker/pkg/apis/addons/v1alpha1"

// addonsCfgLister is an autogenerated mock type for the addonsCfgLister type
type addonsCfgLister struct {
	mock.Mock
}

// List provides a mock function with given fields: namespace, pagingParams
func (_m *addonsCfgLister) List(namespace string, pagingParams pager.PagingParams) ([]*v1alpha1.AddonsConfiguration, error) {
	ret := _m.Called(namespace, pagingParams)

	var r0 []*v1alpha1.AddonsConfiguration
	if rf, ok := ret.Get(0).(func(string, pager.PagingParams) []*v1alpha1.AddonsConfiguration); ok {
		r0 = rf(namespace, pagingParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.AddonsConfiguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, pager.PagingParams) error); ok {
		r1 = rf(namespace, pagingParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
