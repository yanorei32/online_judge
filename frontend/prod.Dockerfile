FROM node:14.17.2 as build-env
WORKDIR /work
COPY . .
RUN npm install --frozen-lockfile
RUN npm run build

FROM node:14.17.2
WORKDIR /app
COPY --from=build-env /work/build /app/build
RUN npm install serve

EXPOSE 3000
ENTRYPOINT ["npx", "serve", "-s", "build"]
