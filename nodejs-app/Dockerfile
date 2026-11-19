FROM node:20-alpine

WORKDIR /app

COPY --chown=node:node package.json index.js ./

USER node

EXPOSE 8080

CMD ["node", "index.js"]
