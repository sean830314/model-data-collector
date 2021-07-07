FROM fluent/fluentd:v1.4-2
USER root
# add mongo plugin
RUN apk add --no-cache bash make gcc libc-dev ruby-dev \
    && gem install fluent-plugin-mongo \
    && gem install  fluent-plugin-redis-store \
    && gem install  fluent-plugin-elasticsearch \
    && apk del make gcc libc-dev ruby-dev \
    && rm -rf /var/cache/apk/* \
    && rm -rf /tmp/* /var/tmp/* /usr/lib/ruby/gems/*/cache/*.gem

RUN mkdir /var/log/test