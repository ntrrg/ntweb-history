FROM ntrrg/hugo:0.59.0 as build
COPY . .
RUN hugo --baseUrl /

FROM ntrrg/nginx:http
COPY --from=build /site/public /usr/share/nginx/html

