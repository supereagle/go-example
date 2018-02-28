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

**Note:**

You can refer to the usage of reverse proxy in the real project [Harbor](https://github.com/vmware/harbor/blob/release-1.4.0/src/ui/proxy/proxy.go).
It intercepts the requests to Docker Registry, and handles them through the chains.
This is a very good example to illustrate the usage of reverse proxy.
