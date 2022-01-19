package gorineeweb

//
//type router struct {
//	trees    map[string]*node
//	cache    map[string]*matchResult
//	cacheLen int
//	mutex    sync.RWMutex
//	notFound handlersChain
//	settings *Settings
//	pool     sync.Pool
//	handlers map[string]map[string]fasthttp.RequestHandler
//}

type matchResult struct {
	handlers handlersChain
	params   map[string]string
}

type Handler func(*gorineeWebContext)

//func (r *router) HandleFunc(method, pattern string, h fasthttp.RequestHandler) {
//	// http 메서드로 등록된 맵이 있는지 확인
//	m, ok := r.handlers[method]
//	if !ok {
//		// 등록된 맵이 없으면 새 맵을 생성
//		m = make(map[string]fasthttp.RequestHandler)
//		r.handlers[method] = m
//	}
//	// http 메서드로 등록된 맵에 URL 패턴과 핸들러 함수 등록
//	m[pattern] = h
//}
//
//func (r *router) ServeHTTP(ctx *fasthttp.RequestCtx) {
//	if m, ok := r.handlers[string(ctx.Method())]; ok {
//		if h, ok := m[string(ctx.Path())]; ok {
//			// 요청 URL에 해당하는 핸들러 수행
//			h(ctx)
//			return
//		}
//	}
//}
