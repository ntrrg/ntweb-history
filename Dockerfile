FROM ntrrg/hugo:0.40.3 as build
COPY . /site
RUN hugo

FROM ntrrg/nginx:http
COPY --from=build /site/public /usr/share/nginx/html
EXPOSE 80

