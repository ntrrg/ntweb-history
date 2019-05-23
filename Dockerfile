FROM ntrrg/hugo:0.55.6 as build
USER 0
COPY . /srv
RUN hugo --baseUrl /

FROM ntrrg/nginx:http
COPY --from=build /srv/public /usr/share/nginx/html
EXPOSE 80

