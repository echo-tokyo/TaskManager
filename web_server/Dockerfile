FROM nginx:1.25.3-alpine

# -----
# nginx
# -----

RUN chown nginx:nginx /var/cache/nginx/

# main nginx conf
COPY ./nginx.conf /etc/nginx/nginx.conf
# included nginx confs
COPY ./http_site.conf /etc/nginx/conf.d/http_site.conf

RUN apk add npm


# ---
# vue
# ---

WORKDIR /app/frontend/vue

COPY ./frontend/package.json .
RUN npm install

COPY ./frontend .

RUN npm run build


# -----
# react
# -----

WORKDIR /app/frontend/react

COPY ./frontend_react/package.json .
RUN npm install

COPY ./frontend/react .

RUN npm run build
