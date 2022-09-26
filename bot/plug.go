package bot

import (
	"os"
	"plugin"

	"github.com/mosqu1t0/Amigo-bot/utils/logcat"
)

const pluginsPath = "./plugins/"

type BotPlugin interface {
	Action(bot *Bot, v interface{})
}

type PluginManager struct {
	plugins []BotPlugin
}

func InitPlugin(bot *Bot) {
	bot.plugManager = new(PluginManager)
	bot.plugManager.LoadPlugins()
}

func (pl *PluginManager) LoadPlugins() {
	files, err := os.ReadDir(pluginsPath)
	if err != nil {
		logcat.Error("无法读取文件夹，插件加载失败!")
	}

	for _, file := range files {
		p, errIo := plugin.Open(pluginsPath + file.Name())
		plugSymbot, errLo := p.Lookup("Plug")
		if errIo != nil || errLo != nil {
			logcat.Error("读取插件 ", file.Name(), " 发生错误...")
			continue
		}
		plugin, ok := plugSymbot.(BotPlugin)
		if ok {
			pl.plugins = append(pl.plugins, plugin)
		}
	}
	logcat.Good("读取插件完成，共加载了", len(pl.plugins), "个插件")
}
