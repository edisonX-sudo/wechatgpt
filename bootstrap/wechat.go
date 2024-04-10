package bootstrap

import (
	"os"
	"wechatbot/handler/wechat"

	"github.com/eatmoreapple/openwechat"
	log "github.com/sirupsen/logrus"
)

func StartWebChat() {
	log.Info("Start WebChat Bot")
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.MessageHandler = wechat.Handler
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	tokenPath := "data/token.json"
	reloadStorage := openwechat.NewJsonFileHotReloadStorage(tokenPath)
	err := bot.HotLogin(reloadStorage, true)
	if err != nil {
		file, _ := os.ReadFile(tokenPath)
		log.Printf("1st HotLogin err:%s\ntoken:%s", err.Error(), string(file))
		//err = os.Remove(tokenPath)
		if err != nil {
			log.Printf("os remove err:%s", err.Error())
			return
		}

		reloadStorage = openwechat.NewJsonFileHotReloadStorage(tokenPath)
		err = bot.HotLogin(reloadStorage)
		if err != nil {
			log.Printf("2rd HotLogin err:%s", err.Error())
			return
		}
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Fatal(err)
		return
	}

	friends, err := self.Friends()
	for i, friend := range friends {
		log.Println(i, friend)
	}

	groups, err := self.Groups()
	for i, group := range groups {
		log.Println(i, group)
	}

	err = bot.Block()
	if err != nil {
		log.Fatal(err)
		return
	}
}
