type MongoManage struct {
    {{range .}}{{.Type}}   model.{{.Type}}Model
    {{end}}
}

func NewMongoManage(monCfg model.MonConfig) *MongoManage {
	return &MongoManage{
        {{range .}}{{.Type}}:  model.New{{.Type}}Model(monCfg),
        {{end}}
	}
}