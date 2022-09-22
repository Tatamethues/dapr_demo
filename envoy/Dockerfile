FROM envoyproxy/envoy-dev:latest

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get -qq update \
    && apt-get -qq install --no-install-recommends -y curl \
    && apt-get -qq autoremove -y \
    && apt-get clean \
    && rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/*
COPY envoy/gw-envoy.yaml /etc/gw-envoy.yaml
RUN chmod go+r /etc/gw-envoy.yaml
CMD ["/usr/local/bin/envoy", "-c", "/etc/gw-envoy.yaml"]