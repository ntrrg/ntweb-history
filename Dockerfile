FROM ntrrg/hugo:latest as build
USER 0
COPY . /site
RUN hugo -d /public --baseUrl /

FROM ntrrg/nginx:http
COPY --from=build /public /usr/share/nginx/html
EXPOSE 80

