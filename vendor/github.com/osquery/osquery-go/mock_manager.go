// Automatically generated by mockimpl. DO NOT EDIT!

package osquery

import "github.com/osquery/osquery-go/gen/osquery"

var _ ExtensionManager = (*MockExtensionManager)(nil)

type CloseFunc func()

type PingFunc func() (*osquery.ExtensionStatus, error)

type CallFunc func(registry string, item string, req osquery.ExtensionPluginRequest) (*osquery.ExtensionResponse, error)

type ExtensionsFunc func() (osquery.InternalExtensionList, error)

type RegisterExtensionFunc func(info *osquery.InternalExtensionInfo, registry osquery.ExtensionRegistry) (*osquery.ExtensionStatus, error)

type DeregisterExtensionFunc func(uuid osquery.ExtensionRouteUUID) (*osquery.ExtensionStatus, error)

type OptionsFunc func() (osquery.InternalOptionList, error)

type QueryFunc func(sql string) (*osquery.ExtensionResponse, error)

type GetQueryColumnsFunc func(sql string) (*osquery.ExtensionResponse, error)

type MockExtensionManager struct {
	CloseFunc        CloseFunc
	CloseFuncInvoked bool

	PingFunc        PingFunc
	PingFuncInvoked bool

	CallFunc        CallFunc
	CallFuncInvoked bool

	ExtensionsFunc        ExtensionsFunc
	ExtensionsFuncInvoked bool

	RegisterExtensionFunc        RegisterExtensionFunc
	RegisterExtensionFuncInvoked bool

	DeRegisterExtensionFunc        DeregisterExtensionFunc
	DeRegisterExtensionFuncInvoked bool

	OptionsFunc        OptionsFunc
	OptionsFuncInvoked bool

	QueryFunc        QueryFunc
	QueryFuncInvoked bool

	GetQueryColumnsFunc        GetQueryColumnsFunc
	GetQueryColumnsFuncInvoked bool
}

func (m *MockExtensionManager) Close() {
	m.CloseFuncInvoked = true
	m.CloseFunc()
}

func (m *MockExtensionManager) Ping() (*osquery.ExtensionStatus, error) {
	m.PingFuncInvoked = true
	return m.PingFunc()
}

func (m *MockExtensionManager) Call(registry string, item string, req osquery.ExtensionPluginRequest) (*osquery.ExtensionResponse, error) {
	m.CallFuncInvoked = true
	return m.CallFunc(registry, item, req)
}

func (m *MockExtensionManager) Extensions() (osquery.InternalExtensionList, error) {
	m.ExtensionsFuncInvoked = true
	return m.ExtensionsFunc()
}

func (m *MockExtensionManager) RegisterExtension(info *osquery.InternalExtensionInfo, registry osquery.ExtensionRegistry) (*osquery.ExtensionStatus, error) {
	m.RegisterExtensionFuncInvoked = true
	return m.RegisterExtensionFunc(info, registry)
}

func (m *MockExtensionManager) DeregisterExtension(uuid osquery.ExtensionRouteUUID) (*osquery.ExtensionStatus, error) {
	m.DeRegisterExtensionFuncInvoked = true
	return m.DeRegisterExtensionFunc(uuid)
}

func (m *MockExtensionManager) Options() (osquery.InternalOptionList, error) {
	m.OptionsFuncInvoked = true
	return m.OptionsFunc()
}

func (m *MockExtensionManager) Query(sql string) (*osquery.ExtensionResponse, error) {
	m.QueryFuncInvoked = true
	return m.QueryFunc(sql)
}

func (m *MockExtensionManager) GetQueryColumns(sql string) (*osquery.ExtensionResponse, error) {
	m.GetQueryColumnsFuncInvoked = true
	return m.GetQueryColumnsFunc(sql)
}
