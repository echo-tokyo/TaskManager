FROM nginx:1.25.3-alpine

# -----
# nginx
# -----

RUN chown nginx:nginx /var/cache/nginx/

# main nginx conf
COPY ./web_server/nginx.conf /etc/nginx/nginx.conf
# included nginx confs
COPY ./web_server/http_site.conf /etc/nginx/conf.d/http_site.conf

RUN apk add npm


# --------
# frontend
# --------

WORKDIR /app/frontend


# install vue requirements
COPY ./frontend/package.json ./vue
RUN npm install

# install react requirements
COPY ./frontend_react/package.json ./react
RUN npm install


# build vue
COPY ./frontend ./vue
RUN npm run build

# build react
COPY ./frontend_react ./react
RUN npm run build
