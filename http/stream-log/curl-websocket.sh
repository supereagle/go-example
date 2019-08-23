# Refer to https://gist.github.com/htp/fbce19069187ec1cc486b594104f01d0

# The Origin header should be provided, otherwise there will be error: 'Origin' header value not allowed.

curl --include \
     --no-buffer \
     --header "Connection: Upgrade" \
     --header "Upgrade: websocket" \
     --header "Host: localhost:8080" \
     --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
     --header "Sec-WebSocket-Version: 13" \
     --header "X-Tenant: caicloud" \
     http://localhost:8080/apis/v1/proxyws?aaa=111\&bbb=222
