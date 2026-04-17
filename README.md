hnrss — Hacker News RSS
========================

hnrss provides custom, realtime RSS feeds for [Hacker News](https://news.ycombinator.com/).

The [project page](https://hnrss.github.io/) explains all available RSS feeds and options.

## Building

**Requirements:** Go 1.25.0+

Build a stripped Linux amd64 binary (embeds the current git tag as version string):

```bash
make hnrss_linux_amd64
```

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

There is no containerization or process supervisor configuration included. A minimal deployment looks like:

1. Build the binary on a machine with Go installed (or cross-compile with `make hnrss_linux_amd64`).
2. Copy the binary to the target host.
3. Run it under a process supervisor (systemd, runit, etc.).

**Example systemd unit** (`/etc/systemd/system/hnrss.service`):

```ini
[Unit]
Description=hnrss Hacker News RSS service
After=network.target

[Service]
ExecStart=/usr/local/bin/hnrss -bind 127.0.0.1:9000
Restart=on-failure
DynamicUser=yes

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl enable --now hnrss
```

Put a reverse proxy (nginx, Caddy, etc.) in front if you need TLS or a public-facing port.

## Dependencies

Runtime dependencies are fetched from the Algolia Hacker News Search API and the Hacker News website directly — no local database or cache is required. The service is stateless.
