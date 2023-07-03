FROM nginx:stable-alpine3.17-slim

RUN rm /etc/nginx/conf.d/default.conf

RUN apk add --no-cache unzip

RUN adduser -D -g 'nginx' nginx

COPY api_gateway.conf /etc/nginx/conf.d/api_gateway.conf
COPY cert /etc/nginx/cert
COPY opentelemetry_module.conf /etc/nginx/conf.d/

WORKDIR /opt
RUN wget -O opentelemetry-webserver-sdk-x64-linux.tgz.zip https://github.com/open-telemetry/opentelemetry-cpp-contrib/releases/download/webserver%2Fv1.0.0/opentelemetry-webserver-sdk-x64-linux.tgz.zip \
    && unzip opentelemetry-webserver-sdk-x64-linux.tgz.zip \
    && tar xvfz opentelemetry-webserver-sdk-x64-linux.tgz \
    && rm opentelemetry-webserver-sdk-x64-linux.tgz opentelemetry-webserver-sdk-x64-linux.tgz.zip \
    && cd opentelemetry-webserver-sdk \
    && ./install.sh \
    && cd ..

ENV LD_LIBRARY_PATH=/opt/opentelemetry-webserver-sdk/sdk_lib/lib:$LD_LIBRARY_PATH

RUN echo "load_module /opt/opentelemetry-webserver-sdk/WebServerModule/Nginx/ngx_http_opentelemetry_module.so;" >> /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
