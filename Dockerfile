FROM ntrrg/hugo:0.58.3 as build
COPY . .
RUN hugo --baseUrl /

FROM ntrrg/nginx:http
COPY --from=build /tmp/site/public /usr/share/nginx/html
EXPOSE 80

