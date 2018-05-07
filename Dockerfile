FROM ntrrg/hugo as build
COPY . /site
RUN hugo

FROM ntrrg/nginx:http
COPY --from=build /site/public /usr/share/nginx/html

