FROM node:14.16.0-alpine3.10

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY . ./

RUN npm install

CMD ["npm", "start"]
