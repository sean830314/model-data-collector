FROM fluent/fluentd:v1.13.0-1.0
USER root
# add mongo plugin
RUN apk add --no-cache bash make gcc libc-dev ruby-dev \
    && gem install fluent-plugin-mongo \
    && gem install  fluent-plugin-redis-store \
    && gem install  fluent-plugin-elasticsearch \
    && gem install fluent-plugin-out-http \
    && apk del make gcc libc-dev ruby-dev \
    && rm -rf /var/cache/apk/* \
    && rm -rf /tmp/* /var/tmp/* /usr/lib/ruby/gems/*/cache/*.gem

RUN mkdir /var/log/test