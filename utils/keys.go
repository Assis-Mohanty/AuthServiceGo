package utils

type jwtContextKey struct{}

var JwtContextKey =jwtContextKey{}

type  requireRoles struct{}

var RequireRolesKey =requireRoles{}

type loginkeyStruct struct{}
var LoginKeyStruct =loginkeyStruct{}

type createRequestkeyStruct struct{}
var  CreateRequestkeyStruct=createRequestkeyStruct{}