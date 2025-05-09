FROM node:20

WORKDIR /boredomboard


COPY package.json package-lock.json angular.json tsconfig.json ./
RUN npm install

COPY projects/app-client/ ./projects/app-client/

RUN npm install -g @angular/cli

CMD [ "ng", "serve", "--host", "0.0.0.0" ]
