FROM denoland/deno:latest

WORKDIR /app

USER deno

COPY . .

# The deno cache command has been merged into the deno install command under the --entrypoint option.
RUN deno install --entrypoint main.ts

CMD ["run", "--allow-net", "main.ts"]
