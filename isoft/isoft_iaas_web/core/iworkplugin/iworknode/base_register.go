package iworknode

import (
	"reflect"
	"strings"
)

var typeMap = make(map[string]reflect.Type, 0)

func init() {
	vs := []interface{}{
		WorkStartNode{},
		WorkEndNode{},
		WorkSubNode{},
		SQLExecuteNode{},
		SQLQueryNode{},
		SQLQueryPageNode{},
		JsonRenderNode{},
		JsonParserNode{},
		HttpRequestNode{},
		MapperNode{},
		FileReadNode{},
		FileWriteNode{},
		FileSyncNode{},
		FileDeleteNode{},
		HrefParserNode{},
		EntityParserNode{},
		DBParserNode{},
		MemoryMapCacheNode{},
		CalHashNode{},
		SetEnvNode{},
		GetEnvNode{},
		RunCmdNode{},
		SftpUploadNode{},
		SSHShellNode{},
		TarGzUnCompressNode{},
		TarGzCompressNode{},
		IniReadNode{},
		IniWriteNode{},
		IFNode{},
		EmptyNode{},
		Base64EncodeNode{},
		Base64DecodeNode{},
		DefineVarNode{},
		AssignVarNode{},
	}
	for _, v := range vs {
		t := reflect.ValueOf(v).Type()
		typeMap[strings.ToUpper(t.Name())] = t
	}
}
