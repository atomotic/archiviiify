FROM hayd/alpine-deno

EXPOSE 8000
WORKDIR /app
USER deno
ADD . .
# RUN deno cache index.js
CMD ["run", "--allow-net", "index.js"]