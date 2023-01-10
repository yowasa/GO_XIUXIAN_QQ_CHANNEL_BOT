package item

import "GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"

var ItemNameMap map[string]cfg.Item

var ItemIdMap map[int]cfg.Item

func init() {
	ItemNameMap = make(map[string]cfg.Item)
	ItemIdMap = make(map[int]cfg.Item)
	//初始化item映射
	for _, i := range cfg.ItemList {
		ItemNameMap[i.Name] = i
		ItemIdMap[i.Id] = i
	}

}
