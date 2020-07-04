import { Application, Router } from "https://deno.land/x/oak/mod.ts";
import { getQuery } from "https://deno.land/x/oak/helpers.ts";
import { oakCors } from "https://deno.land/x/cors/mod.ts";

import { makeContentSearchResponse } from "./iiif.js";

const solr = "http://solr:8983/solr";
const core = "ocr";

const router = new Router();
router.get("/search/:document", async (context) => {
  const document = context.params.document;
  const query = getQuery(context);

  var url = new URL(`${solr}/${core}/select`),
    params = {
      q: `ocr_text:"${query.q}"~10`,
      fq: `id:${document}`,
      hl: true,
      "hl.ocr.fl": "ocr_text",
      "hl.snippets": 5,
    };
  Object.keys(params).forEach((key) =>
    url.searchParams.append(key, params[key])
  );
  const res = await fetch(url).then((res) => res.json());

  const snippets = res.ocrHighlighting[document].ocr_text.snippets;
  const total = res.ocrHighlighting[document].ocr_text.numTotal;
  const annotations = makeContentSearchResponse(
    snippets,
    total,
    query.q,
    document
  );
  context.response.body = JSON.stringify(annotations, null, 4);
});

const app = new Application();

app.addEventListener("error", (evt) => {
  console.log(evt.error);
});

app.use(oakCors());
app.use(router.routes());
app.use(router.allowedMethods());

console.log("http://localhost:8000/search");
await app.listen({ port: 8000 });
