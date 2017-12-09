# Reverse Proxy

This is a demo for reverse proxy server to illustrate how to proxy the requests.

What the demo does during the proxy:
* Change the request path
* Get the body from quest body
* Update the body from quest body

There are two handlers:
* `/hello`: The real backend handler.
  * For `GET` request, directly write the message
  * For other requests, print the data of the request body
* `/proxy`: The proxy handler, will redirect the request to `/hello` handler.
  * For `GET` request, print the request body, and redirect the request
  * For `PATCH` request, print the request body, then update the request body before redirect
  * For other requests, just redirect the request
