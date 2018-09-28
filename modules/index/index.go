package index

import (
	. "github.com/xirtah/gopa-framework/core/config"
	api "github.com/xirtah/gopa-framework/core/http"
	core "github.com/xirtah/gopa-framework/core/index"
	"github.com/xirtah/gopa-framework/core/model"
	"github.com/xirtah/gopa-ui/modules/index/ui"
	common "github.com/xirtah/gopa-ui/modules/index/ui/common"
)

type IndexModule struct {
}

func (this IndexModule) Name() string {
	return "Index"
}

var signalChannel chan bool

var (
	defaultConfig = common.IndexConfig{
		Elasticsearch: &core.ElasticsearchConfig{
			Endpoint:    "http://localhost:9200",
			IndexPrefix: "gopa-",
		},
		UIConfig: &common.UIConfig{
			Enabled:     true,
			SiteName:    "GOPA",
			SiteFavicon: "/static/assets/img/favicon.ico",
			SiteLogo:    "/static/assets/img/logo.svg",
		},
	}
)

func (module IndexModule) Start(cfg *Config) {

	indexConfig := defaultConfig
	cfg.Unpack(&indexConfig)

	signalChannel = make(chan bool, 1)
	//client := core.ElasticsearchClient{Config: indexConfig.Elasticsearch}

	//register UI
	if indexConfig.UIConfig.Enabled {
		ui := ui.UserUI{}
		ui.Config = indexConfig.UIConfig
		ui.SearchClient = &core.ElasticsearchClient{Config: indexConfig.Elasticsearch}
		api.HandleUIMethod(api.GET, "/", ui.IndexPageAction)
		api.HandleUIMethod(api.GET, "/m/", ui.MobileIndexPageAction)
		api.HandleUIMethod(api.GET, "/ajax_more_item/", ui.AJAXMoreItemAction)
		api.HandleUIMethod(api.GET, "/snapshot/:id", api.NeedPermission(model.PERMISSION_SNAPSHOT_VIEW, ui.GetSnapshotPayloadAction))
		api.HandleUIMethod(api.GET, "/suggest/", ui.SuggestAction)
	}

	// go func() {
	// 	defer func() {

	// 		if !global.Env().IsDebug {
	// 			if r := recover(); r != nil {

	// 				if r == nil {
	// 					return
	// 				}
	// 				var v string
	// 				switch r.(type) {
	// 				case error:
	// 					v = r.(error).Error()
	// 				case runtime.Error:
	// 					v = r.(runtime.Error).Error()
	// 				case string:
	// 					v = r.(string)
	// 				}
	// 				log.Error("error in indexer,", v)
	// 			}
	// 		}
	// 	}()

	// 	for {
	// 		select {
	// 		case <-signalChannel:
	// 			log.Trace("indexer exited")
	// 			return
	// 		default:
	// 			log.Trace("waiting index signal")
	// 			er, v := queue.Pop(config.IndexChannel)
	// 			log.Trace("got index signal, ", string(v))
	// 			if er != nil {
	// 				log.Error(er)
	// 				continue
	// 			}
	// 			//indexing to es or blevesearch
	// 			doc := model.IndexDocument{}
	// 			err := json.Unmarshal(v, &doc)
	// 			if err != nil {
	// 				panic(err)
	// 			}

	// 			client.Index(doc.Index, doc.ID, doc.Source)
	// 		}

	// 	}
	// }()
}

func (module IndexModule) Stop() error {
	signalChannel <- true
	return nil
}
