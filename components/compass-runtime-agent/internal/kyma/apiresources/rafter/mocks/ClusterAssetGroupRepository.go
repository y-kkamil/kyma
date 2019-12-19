// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	apperrors "kyma-project.io/compass-runtime-agent/internal/apperrors"
	clusterassetgroup "kyma-project.io/compass-runtime-agent/internal/kyma/apiresources/rafter/clusterassetgroup"
)

// ClusterAssetGroupRepository is an autogenerated mock type for the ClusterAssetGroupRepository type
type ClusterAssetGroupRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: assetGroup
func (_m *ClusterAssetGroupRepository) Create(assetGroup clusterassetgroup.Entry) apperrors.AppError {
	ret := _m.Called(assetGroup)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(clusterassetgroup.Entry) apperrors.AppError); ok {
		r0 = rf(assetGroup)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *ClusterAssetGroupRepository) Delete(id string) apperrors.AppError {
	ret := _m.Called(id)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *ClusterAssetGroupRepository) Get(id string) (clusterassetgroup.Entry, apperrors.AppError) {
	ret := _m.Called(id)

	var r0 clusterassetgroup.Entry
	if rf, ok := ret.Get(0).(func(string) clusterassetgroup.Entry); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(clusterassetgroup.Entry)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string) apperrors.AppError); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: assetGroup
func (_m *ClusterAssetGroupRepository) Update(assetGroup clusterassetgroup.Entry) apperrors.AppError {
	ret := _m.Called(assetGroup)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(clusterassetgroup.Entry) apperrors.AppError); ok {
		r0 = rf(assetGroup)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}
