FROM ntrrg/hugo:0.40.3 as build
COPY . /site
RUN hugo -d /public

FROM ntrrg/nginx:http
COPY --from=build /public /usr/share/nginx/html
EXPOSE 80

