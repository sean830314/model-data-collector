# Fluentd-agent and aggregator

## Build image

```
docker build -t fluentd:agent . -f Dockerfiles/fluentd-agent.Dockerfile
docker build -t fluentd:aggregator . -f Dockerfiles/fluentd-aggregator.Dockerfile
```

## Local run image

run agent
```
docker run -v <path>/fluentd/config/server_fluentd-agent.conf:/fluentd/etc/fluent.conf -v <log-folder>:/var/log/app --net=host -d fluentd:agent
```
run aggregator
```
docker run -v <path>/fluentd/config/server_fluentd-aggregator.conf:/fluentd/etc/fluent.conf -v <log-folder>:/var/log/app -p 24225:24224  -d fluentd:aggregator
```
run test-aggregator
```
docker run -v <path>/fluentd/deprecated/test-aggregator.conf:/fluentd/etc/fluent.conf -v <log-folder>:/var/log/test-service --net=host -d fluentd:aggregator
```