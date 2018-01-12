package db

var engMap map[string]Engine

func init() {
	engMap = make(map[string]Engine)
}

func getEngine(uri string) (Engine, error) {
	// TODO: db.getEngine
	return nil, nil
}

func Register(name string, eng Engine) {
	// TODO: db.Register
	log.D("register engine %s", name)
	_, registered := engMap[name]
	if registered {
		log.E("register duplicate engine %s", name)
		return
	}
	engMap[name] = eng
}
