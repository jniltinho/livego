<p align='center'>
    <img src='./logo.png' width='200px' height='80px'/>
</p>


[![Test](https://github.com/jniltinho/livego/workflows/Test/badge.svg)](https://github.com/jniltinho/livego/actions?query=workflow%3ATest)
[![Release](https://github.com/jniltinho/livego/workflows/Release/badge.svg)](https://github.com/jniltinho/livego/actions?query=workflow%3ARelease)

Simple and efficient live broadcast server:
- Very simple to install and use;
- Pure Golang, high performance, and cross-platform;
- Supports commonly used transmission protocols, file formats, and encoding formats;

#### Supported transport protocols
- RTMP
- AMF
- HLS
- HTTP-FLV

#### Supported container formats
- FLV
- TS

#### Supported encoding formats
- H264
- AAC
- MP3

## Installation
After directly downloading the compiled [binary file](https://github.com/jniltinho/livego/releases), execute it on the command line.

#### Boot from Docker
Run `docker run -p 1935:1935 -p 7001:7001 -p 7002:7002 -p 8090:8090 -p 3001:3001 -d jniltinho/livego` to start

#### Compile from source
1. Download the source code `git clone https://github.com/jniltinho/livego.git`
2. Go to the livego directory and execute `go build -ldflags="-s -w"` or `make build`

## Use
1. Start the service: execute the livego binary file or `make run` to start the livego service;
2. Get a channelkey(used for push the video stream) from `http://localhost:8090/control/get?room=movie` or `curl http://localhost:8090/control/get?room=movie` and copy data like your channelkey.
3. Upstream push: Push the video stream to `rtmp://localhost:1935/{appname}/{channelkey}` through the` RTMP` protocol(default appname is `live`), for example, use `ffmpeg -re -i demo.flv -c copy -f flv rtmp://localhost:1935/{appname}/{channelkey}` push([download demo flv](https://s3plus.meituan.net/v1/mss_7e425c4d9dcb4bb4918bbfa2779e6de1/mpack/default/demo.flv));
4. Downstream playback: The following three playback protocols are supported, and the playback address is as follows:
    - `RTMP`:`rtmp://localhost:1935/{appname}/movie`
    - `FLV`:`http://127.0.0.1:7001/{appname}/movie.flv`
    - `HLS`:`http://127.0.0.1:7002/{appname}/movie.m3u8`
    - `HTTP`:`http://127.0.0.1:3001/live/{room}`
5. Use hls via https: generate ssl certificate(server.key, server.crt files), place them in directory with executable file, change "use_hls_https" option in livego.yaml to true (false by default)

All options: 
```bash
./livego  -h
Usage of ./livego:
      --api_addr string       HTTP manage interface server listen address (default ":8090")
      --config_file string    configure filename (default "livego.yaml")
      --create_config         Create yaml config file
      --enable_rtmps          enable server session RTMPS
      --enable_tls_verify     Use system root CA to verify RTMPS connection, set this flag to false on Windows (default true)
      --flv_dir string        output flv file at flvDir/APP/KEY_TIME.flv (default "tmp")
      --gop_num int           gop num (default 1)
      --hls_addr string       HLS server listen address (default ":7002")
      --hls_keep_after_end    Maintains the HLS after the stream ends
      --http_addr string      HTTP server listen address (default ":3001")
      --httpflv_addr string   HTTP-FLV server listen address (default ":7001")
      --level string          Log level (default "info")
      --live_url string       Live URL (default "http://localhost:7001/live/movie.flv")
      --read_timeout int      read time out (default 10)
      --rtmp_addr string      RTMP server listen address (default ":1935")
      --rtmps_cert string     cert file path required for RTMPS (default "server.crt")
      --rtmps_key string      key file path required for RTMPS (default "server.key")
      --write_timeout int     write time out (default 10)

```

### [Use with flv.js](https://github.com/xqq/mpegts.js/tree/master/demo)
