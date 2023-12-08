FROM node:18.12-alpine


WORKDIR /app/
COPY ./frontend/package.json ./

RUN npm install
