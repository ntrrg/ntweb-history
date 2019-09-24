FROM ntrrg/hugo:0.58.3 as build
USER 0
WORKDIR /srv
COPY . .
RUN hugo --baseUrl /

FROM ntrrg/nginx:http
COPY --from=build /srv/public /usr/share/nginx/html
EXPOSE 80

