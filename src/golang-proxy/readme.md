
# Golang Proxy

This project provides a proxy for port 80. Requests are
fed through alternate ports based on request hostname.

e.g.
- serve google.com from port 8001
- serve images.google.com from port 8002
- serve videos.google.com from port 8003
