FROM ntrrg/hugo:0.55.6
COPY . /srv
CMD ["server", "-D", "-E", "-F", "--noHTTPCache", "--bind", "0.0.0.0", "--baseUrl", "/", "--appendPort=false"]

