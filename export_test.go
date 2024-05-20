package sdkclient

func SetClientMethod(key string, fn ClientMethod) {
	clientMethods[key] = fn
}

func ClientMethods() map[string]ClientMethod {
	return clientMethods
}

type ClientMethodParam = clientMethodParam
