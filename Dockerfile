######################################################################################

FROM node:22-alpine3.19 as frontend-builder

ADD frontend /frontend
RUN cd /frontend && rm -rf node_modules && npm install && npm run build


######################################################################################

FROM golang:1.22.4-bookworm as backend-builder

ADD backend /backend
RUN cd /backend && go mod tidy && go build main.go


######################################################################################

FROM debian

COPY --from=frontend-builder /frontend/dist /html
COPY --from=backend-builder /backend/main /main

WORKDIR /

CMD ["/main"]
