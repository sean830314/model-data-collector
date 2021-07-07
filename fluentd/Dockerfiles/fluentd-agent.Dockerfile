FROM fluent/fluentd:v1.13.0-1.0
USER root
# add mongo plugin
RUN apk add --no-cache bash make gcc libc-dev ruby-dev \
    && apk del make gcc libc-dev ruby-dev \
    && rm -rf /var/cache/apk/* \
    && rm -rf /tmp/* /var/tmp/* /usr/lib/ruby/gems/*/cache/*.gem

RUN mkdir /var/log/test