package prometheus

import "github.com/prometheus/client_golang/prometheus"

var (
	RegisterApiPingCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "register_api_ping_counter",
		Help: "Number of pings made to the Register API endpoint",
	})
	LoginApiPingCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "login_api_ping_counter",
		Help: "Number of pings made to the Login API endpoint",
	})
	AdCreateCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "create_ad_counter",
		Help: "Number of pings made to the Create Ad API endpoint",
	})
	AdGetCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "get_ad_counter",
		Help: "Number of pings made to the Get Ad API endpoint",
	})
	AdGetAllCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "get_all_ads_counter",
		Help: "Number of pings made to the Get All Ads API endpoint",
	})
)

func Init() {
	prometheus.MustRegister(RegisterApiPingCounter)
	prometheus.MustRegister(LoginApiPingCounter)
	prometheus.MustRegister(AdGetCounter)
	prometheus.MustRegister(AdCreateCounter)
	prometheus.MustRegister(AdGetAllCounter)
}
