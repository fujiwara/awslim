package sdkclient

import "runtime/debug"

var (
	MarshalJSON = marshalJSON
)

func SetClientMethod(service, method string, fn ClientMethod) {
	sv := clientMethods[service]
	if sv == nil {
		sv = make(map[string]ClientMethod)
		clientMethods[service] = sv
	}
	sv[method] = fn
}

func ClientMethods() map[string]map[string]ClientMethod {
	return clientMethods
}

type ClientMethodParam = clientMethodParam

func SetBuildInfo(info *debug.BuildInfo) {
	readBuildInfo = func() (*debug.BuildInfo, bool) {
		return info, info != nil
	}
}
