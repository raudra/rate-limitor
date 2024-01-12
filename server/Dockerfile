FROM fabiocicerchia/nginx-lua

RUN apk add git unzip luarocks gcc make

RUN luarocks install lua-resty-redis&& \
    luarocks install json.lua

COPY default /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/
COPY rule.lua /etc/nginx/

EXPOSE 80/tcp

CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
