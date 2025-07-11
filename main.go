package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var temperatureGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "city_district_temperature_celsius",
		Help: "Temperature in Celsius for different districts in various cities",
	},
	[]string{"city", "district"},
)

var cityDistricts = map[string][]string{
	"广州": {"天河", "白云", "越秀", "番禺"},
	"深圳": {"宝安", "南山", "福田", "龙岗"},
	"北京": {"海淀", "朝阳", "西城", "丰台"},
	"上海": {"浦东", "徐汇", "长宁", "静安"},
}

func simulateTemperature() {
	for {
		for city, districts := range cityDistricts {
			for _, district := range districts {
				temp := 15 + rand.Float64()*25 // 模拟 15~40℃
				temperatureGauge.WithLabelValues(city, district).Set(temp)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	prometheus.MustRegister(temperatureGauge)
	go simulateTemperature()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
