package sdkclient

var clientMethods = map[string]map[string]ClientMethod{}

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
