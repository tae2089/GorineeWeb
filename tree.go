package GorineeWeb

type node struct {
	//메소드 관련 맵
	methods map[METHOD]handlersChain
	//경로
	path       string
	childNodes []*node
}

type tree struct {
	rootNode *node
}
