FROM ntrrg/hugo:0.59.0
COPY . /srv
CMD ["server", "-D", "-E", "-F", "--noHTTPCache", "--bind", "0.0.0.0", "--baseUrl", "/", "--appendPort=false"]

