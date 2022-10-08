package bot

import "github.com/mosqu1t0/Amigo-bot/utils/logcat"

type BotPlugin interface {
	GetType() string
	Init()
	Action(b *Bot, v interface{})
	Destroy()
}

type pluginManager struct {
	plugins map[string][]BotPlugin
}

var PluginMgr *pluginManager

func init() {
	PluginMgr = new(pluginManager)
	PluginMgr.plugins = make(map[string][]BotPlugin)
}

func (pm *pluginManager) AddPlugin(plugin BotPlugin) {
	postType := plugin.GetType()
	pm.plugins[postType] = append(pm.plugins[postType], plugin)
}

func (pm *pluginManager) startInit() {
	count := 0
	for _, vp := range pm.plugins {
		count += len(vp)
		for _, p := range vp {
			p.Init()
		}
	}
	logcat.Good("> 插件加载完毕，已加载 ", count, " 个插件 <")
}

func (pm *pluginManager) startPluck() {
	for _, vp := range pm.plugins {
		for _, p := range vp {
			p.Destroy()
		}
	}
}
