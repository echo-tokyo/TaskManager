FROM nginx:1.25.3-alpine

RUN chown nginx:nginx /var/cache/nginx/
RUN apk add npm

# --------
# frontend
# --------

# install vue requirements
WORKDIR /app/frontend/vue
COPY ./frontend/package.json .
RUN npm install

# install react requirements
WORKDIR /app/frontend/react
COPY ./frontend_react/package.json .
RUN npm install

# build vue
WORKDIR /app/frontend/vue
COPY ./frontend .
RUN npm run build

# build react
WORKDIR /app/frontend/react
COPY ./frontend_react .
RUN npm run build

# -----
# nginx
# -----

# main nginx conf
COPY ./web_server/nginx.conf /etc/nginx/nginx.conf
# included nginx confs
COPY ./web_server/http_site.conf /etc/nginx/conf.d/http_site.conf
