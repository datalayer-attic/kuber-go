package rest

import (
	"github.com/datalayer-attic/kuber-go/cloud"
	"github.com/datalayer-attic/kuber-go/config"
	"github.com/datalayer-attic/kuber-go/google"

	//	"github.com/datalayer-attic/kuber-go/helm"
	"github.com/datalayer-attic/kuber-go/about"
	k "github.com/datalayer-attic/kuber-go/k8s"
	"github.com/datalayer-attic/kuber-go/microsoft"
	"github.com/datalayer-attic/kuber-go/spl"
	"github.com/datalayer-attic/kuber-go/twitter"
	"github.com/datalayer-attic/kuber-go/user"
	wso "github.com/datalayer-attic/kuber-go/ws"
	restful "github.com/emicklei/go-restful"
)

func SetupGoRestful(wsContainer *restful.Container) {

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  AllowedHeaders(),
		AllowedHeaders: AllowedHeaders(),
		AllowedMethods: AllowedMethods(),
		AllowedDomains: AllowedOrigins(),
		CookiesAllowed: true,
		Container:      wsContainer,
	}
	wsContainer.Filter(cors.Filter)
	// Add container filter to respond to OPTIONS.
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	// Web Socket Resources.
	wsoc := wso.WsResource{wso.WsMessage{}}
	wsContainer.Add(wsoc.WebService())

	// Config Resources.
	conf := config.ConfigResource{config.Config{}}
	wsContainer.Add(conf.WebService())

	// Cloud Resources.
	clo := cloud.CloudResource{}
	wsContainer.Add(clo.WebService())

	// K8S Resources.
	k := k.ClusterResource{}
	wsContainer.Add(k.WebService())

	// Helm Resources.
	//	h := helm.HelmResource{}
	//	wsContainer.Add(h.WebService())

	// Google Resources.
	g := google.GoogleResource{}
	wsContainer.Add(g.WebService())

	// Microsoft Resources.
	m := microsoft.MicrosoftResource{}
	wsContainer.Add(m.WebService())

	// Twitter Resources.
	tw := twitter.TwitterResource{}
	wsContainer.Add(tw.WebService())

	// User Resources.
	u := user.UserResource{map[string]user.User{}}
	wsContainer.Add(u.WebService())

	// Sample Resources.
	s := spl.SplResource{map[int]spl.Spl{}}
	wsContainer.Add(s.WebService())

	// About Resources.
	a := about.AboutResource{}
	wsContainer.Add(a.WebService())

}

/*
func SetupGoRestful2(ws *restful.WebService, wsContainer *restful.Container) {

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  AllowedHeaders(),
		AllowedHeaders: AllowedHeaders(),
		AllowedMethods: AllowedMethods(),
		AllowedDomains: AllowedOrigins(),
		CookiesAllowed: true,
		Container:      wsContainer,
	}
	wsContainer.Filter(cors.Filter)
	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	// spl
	s := spl.SplResource{map[int]spl.Spl{}}
	ws.Route(
		ws.GET("spl").
			To(s.FindAllSpls).
			Writes(spl.SplResource{}))

	u := user.UserResource{map[string]user.User{}}
	ws.Route(
		ws.GET("user").
			To(u.FindAllUsers).
			Writes(user.UserResource{}))

	c := config.ConfigResource{config.Config{}}
	ws.Route(
		ws.GET("conf").
			To(c.GetConf).
			Writes(config.Config{}))

	wsss := wso.WsResource{wso.WsMessage{}}
	ws.Route(ws.GET("/").To(wsss.ToWs))

//		wsock := wso.WsResource{wso.WsMessage{}}
//		ws.Route(
//			ws.GET("ws").
//				To(wsock.ToWs))

}
*/
