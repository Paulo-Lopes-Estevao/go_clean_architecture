package repository_mock

import (
	"reflect"

	gomock "github.com/golang/mock/gomock"
)

type MockParkRepository struct {
	controller *gomock.Controller
	recorder   *MockParkRepositoryMockRecorder
}

type MockParkRepositoryMockRecorder struct {
	mock *MockParkRepository
}

func NewMockParkRepository(controller *gomock.Controller) *MockParkRepository {
	mock := &MockParkRepository{controller: controller}
	mock.recorder = &MockParkRepositoryMockRecorder{mock}
	return mock
}

func (m *MockParkRepository) EXPECT() *MockParkRepositoryMockRecorder {
	return m.recorder
}

func (m *MockParkRepository) Registre(name string, limit, vague int32, status bool) error {
	m.controller.T.Helper()
	ref := m.controller.Call(m, "Registre", name, limit, vague, status)
	refl, _ := ref[0].(error)
	return refl
}

func (mr *MockParkRepositoryMockRecorder) Registre(name, limit, vague, status interface{}) *gomock.Call {
	mr.mock.controller.T.Helper()
	return mr.mock.controller.RecordCallWithMethodType(mr.mock, "Registre", reflect.TypeOf((*MockParkRepository)(nil).Registre), name, limit, vague, status)
}
