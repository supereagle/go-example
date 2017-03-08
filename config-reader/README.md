# Config Reader

Config is an important part of an application. Most applications with configuration needs to first read config 
to initialize the instance before really start work. 

This example illustrates how to read config for an application from following 4 ways:
- **Environment Variables**: The environment variables must be set before start the application. 
If the application runs in another place, these environment variables must be set again. 
Some container orchestration tools such [Kubernetes](https://github.com/kubernetes/kubernetes) and 
[Compose](https://github.com/docker/compose) can help to set the environment variables when schedule the application.
- **[ETCD](https://github.com/coreos/etcd)**: ETCD is a distributed reliable key-value store, 
which can be used as a config center. It can keep high consistent and available when works as cluster.
- JSON file
- YAML file