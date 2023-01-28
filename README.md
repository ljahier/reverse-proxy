# Local Reverse Proxy

Local Reverse Proxy allows you to use domain names rather than ports in your browser's url bar.

## Installation
```bash
git clone git@github.com:ljahier/reverse-proxy.git

cd reverse-proxy

go build -o rp cmd/main.go

mv ../reverse-proxy /opt/reverse-proxy

ln -s /opt/reverse-proxy/rp /bin/rp
```

## Configuration
```bash
cp /opt/reverse-proxy/config.json.sample ~/.reverse-proxy.json
```

You can copy the configuration sample to fill with your real configuration