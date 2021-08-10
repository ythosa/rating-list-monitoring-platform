FROM node:latest as build
RUN mkdir /app

ADD ./frontend/src /app/src/
ADD ./frontend/public /app/public/
COPY ./frontend/package*.json ./frontend/tsconfig.json /app/

WORKDIR /app

RUN npm install
RUN npm run build

FROM nginx:stable

EXPOSE 80

COPY --from=build /app/build /var/www/
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf

CMD ["nginx","-g","daemon off;"]
