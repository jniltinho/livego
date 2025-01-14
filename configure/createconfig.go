package configure

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func CreateConfigFile() {

	yaml := `# # Logger level
# level: info

# # FLV Options
# flv_archive: false
# flv_dir: "./tmp"
# httpflv_addr: ":7001"

# # RTMP Options
# rtmp_noauth: false
# rtmp_addr: ":1935"
# enable_rtmps: true
# rtmps_cert: server.crt
# rtmps_key: server.key
# read_timeout: 10
# write_timeout: 10

# # HLS Options
# hls_addr: ":7002"
#use_hls_https: true

# # API Options
# api_addr: ":8090"
# http_addr: ":3001"
# live_url: "http://localhost:7001/live/movie.flv"
server:
- appname: live
  live: true
  hls: true
  api: true
  flv: true
`

	f, err := os.Create(Config.GetString("config_file"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(yaml)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(Config.GetString("config_file"), " created")

}
