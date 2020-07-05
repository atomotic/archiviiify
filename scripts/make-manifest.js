import { expandGlob } from "https://deno.land/std/fs/mod.ts";

const manifest = {
  "@context": "http://iiif.io/api/presentation/2/context.json",
  "@type": "sc:Manifest",
  "@id": "",
  label: "",
  service: {
    "@context": "http://iiif.io/api/search/0/context.json",
    "@id": "",
    profile: "http://iiif.io/api/search/0/search",
  },
  sequences: [
    {
      "@id": "",
      "@type": "sc:Sequence",
      label: "default",
      startCanvas: "",
      canvases: [],
    },
  ],
  thumbnail: {
    "@id": "",
    format: "image/jpeg",
  },
};

function Canvas() {
  return {
    "@id": "",
    "@type": "sc:Canvas",
    label: "",
    metadata: [],
    images: [
      {
        "@type": "oa:Annotation",
        motivation: "sc:painting",
        resource: {
          "@id": "",
          "@type": "dctypes:Image",
          service: {
            "@context": "http://iiif.io/api/image/2/context.json",
            "@id": "",
            profile: "http://iiif.io/api/image/2/level0.json",
          },
          format: "image/jpeg",
          width: 0,
          height: 0,
        },
        on: "",
      },
    ],
    width: 0,
    height: 0,
  };
}

const item = Deno.args[0];
if (item == null) {
  console.log("Usage: ./make-manifest.js ITEM");
  Deno.exit();
}
const base = "http://localhost:8094/iiif";

manifest["@id"] = `${base}/${item}/manifest`;
manifest.label = "";
manifest.service["@id"] = `http://localhost:8094/search/${item}`;
manifest.sequences[0]["@id"] = `${base}/${item}/seq/1`;

for await (const page of expandGlob(`./data/${item}/*.jp2`)) {
  const iiifapi = await fetch(`${base}/${item}/${page.name}/info.json`);
  const image = await iiifapi.json();

  const canvas = Canvas();
  canvas["@id"] = `${base}/${item}/${page.name}`;
  canvas.images[0].on = `${base}/${item}/${page.name}`;
  canvas.images[0].resource["@id"] = `${base}/${item}/${page.name}`;
  canvas.images[0].resource.service["@id"] = `${base}/${item}/${page.name}`;
  canvas.images[0].resource.height = image.height;
  canvas.images[0].resource.width = image.width;
  canvas.height = image.height;
  canvas.width = image.width;

  manifest.sequences[0].canvases.push(canvas);
}

console.log(JSON.stringify(manifest));
