FROM node:14.17.2

ENV PATH=/node_modules/.bin:$PATH

COPY package*.json ./
RUN npm install -g

WORKDIR "/app"

EXPOSE 3000

ENTRYPOINT ["npm", "start"]
