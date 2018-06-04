FROM ntrrg/hugo:0.41 as build
COPY . /site
RUN hugo -d /public --baseUrl /

FROM ntrrg/nginx:http
COPY --from=build /public /usr/share/nginx/html
EXPOSE 80

