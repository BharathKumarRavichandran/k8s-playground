FROM node:8.12.0

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install --silent

COPY . .

ENV NODE_ENV=production

EXPOSE 3000

CMD ["npm", "run", "docker"]
