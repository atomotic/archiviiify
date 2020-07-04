FROM hayd/alpine-deno

EXPOSE 8000
WORKDIR /app
USER deno
ADD . .
RUN deno cache main.js
CMD ["run", "--allow-net", "main.js"]