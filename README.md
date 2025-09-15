# SOCKS5 Proxy Server

simple SOCKS5 proxy server implemented in Go with authentication only support (default credentials).

## Default Credentials
<h4>username: admin </br> password: password</h4>

## Features

- SOCKS5 protocol implementation
- Username/password authentication
- TCP tunneling
- Connection logging
- Configurable listening port

## How to Run

1. **Clone the repository**:
   ```bash
   git clone <repo URL>
   cd socks5-proxy
   go build -o socks5_proxy
   ./socks5_proxy
   ```


```
curl --socks5 admin:passwod@localhost:1080 https://ipinfo.io/ip

```