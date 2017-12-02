# Refer to https://gist.github.com/htp/fbce19069187ec1cc486b594104f01d0

# The Origin header should be provided, otherwise there will be error: 'Origin' header value not allowed.

curl --include \
     --no-buffer \
     --header "Connection: Upgrade" \
     --header "Upgrade: websocket" \
     --header "Origin: localhost" \
     --header "Host: localhost:8080" \
     --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
     --header "Sec-WebSocket-Version: 13" \
     --header "X-Tenant: caicloud" \
     http://localhost:7088/api/v1/workspaces/ws1/pipelines/pipeline1/records/5a1be4ff2785251bd63c0350/logstream