FROM node:20-alpine
WORKDIR /app
COPY app.js .
EXPOSE 3000
ENTRYPOINT ["node","app.js"]
