// Copyright 2013-2017 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lua

import (
	"github.com/aerospike/aerospike-client-go/internal/lua/resources"
	"github.com/aerospike/aerospike-client-go/logger"
	"github.com/aerospike/aerospike-client-go/types"
	"github.com/yuin/gopher-lua"
)

// SetPath sets the interpreter's current Lua Path
func SetPath(lpath string) {
	lua.LuaPath = lpath
}

// LuaPath returns the interpreter's current Lua Path
func LuaPath() string {
	return lua.LuaPath
}

type LuaInstance struct {
	instance *lua.LState
}

// Global LState pool
var LuaPool = types.NewPool(64)

func newInstance(params ...interface{}) interface{} {
	L := lua.NewState()

	registerLuaAerospikeType(L)
	registerLuaStreamType(L)
	registerLuaListType(L)
	registerLuaMapType(L)

	if err := L.DoString(luaLib.LibStreamOps); err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}

	if err := L.DoString(luaLib.LibAerospike); err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}

	return L
}

func finalizeInstance(instance interface{}) {
	if instance != nil {
		instance.(*lua.LState).Close()
	}
}

func init() {
	LuaPool.New = newInstance
	LuaPool.Finalize = finalizeInstance
}
