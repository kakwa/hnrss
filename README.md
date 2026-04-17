hnrss — Hacker News RSS
========================

hnrss provides custom, realtime RSS feeds for [Hacker News](https://news.ycombinator.com/).

The [project page](https://hnrss.kakwalab.ovh/) explains all available RSS feeds and options, and is served directly by the running instance.

## Building

To build for the local platform without stripping:

```bash
go build -o hnrss .
```

## Running

```bash
./hnrss [-bind HOST:PORT]
```

The default bind address is `127.0.0.1:9000`. Override it with `-bind`:

```bash
./hnrss -bind 0.0.0.0:8080
```

The server handles `SIGINT` (Ctrl-C) with a 5-second graceful shutdown.

## Deploying

### SystemD Daemon + Nginx

Copy the binary:

```bash
install -m 755 hnrss /usr/local/bin
```
create SystemD unit `/etc/systemd/system/hnrss.service`:

```ini
[Unit]
Description=Hacker News RSS
After=network.target

[Service]
Type=simple
Restart=always
ExecStart=/usr/local/bin/hnrss -bind 127.0.0.1:9001
User=daemon
Group=daemon
StandardOutput=journal
StandardError=journal
SyslogIdentifier=hnrss

[Install]
WantedBy=multi-user.target
```

```bash
systemctl daemon-reload
sudo systemctl enable --now hnrss
```

And put a reverse proxy (nginx, Caddy, etc.) in front if it, for example:

```apache
server {
	server_name hnrss.kakwalab.ovh;

   location / {
     proxy_pass http://127.0.0.1:9001;
   }

    # TODO: check if it works without certificate
    listen 443 ssl;
    #ssl_certificate /etc/letsencrypt/live/hnrss.kakwalab.ovh/fullchain.pem; # managed by Certbot
    #ssl_certificate_key /etc/letsencrypt/live/hnrss.kakwalab.ovh/privkey.pem; # managed by Certbot
    #include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    #ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
}
  
server {
	server_name hnrss.kakwalab.ovh;
  listen 80;

  location / {
    proxy_pass http://127.0.0.1:9001;
  }
}
```

finally, run certbot to get a proper certificate.

## Dependencies

Runtime dependencies are fetched from the Algolia Hacker News Search API and the Hacker News website directly — no local database or cache is required. The service is stateless.
