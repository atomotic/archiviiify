import { Annotation, Annotations, Hit } from "./templates.js";
import ShortUniqueId from "https://cdn.jsdelivr.net/npm/short-unique-id@latest/short_uuid/mod.ts";

export function makeContentSearchResponse(snippets, total, query, document) {
  const uid = new ShortUniqueId();

  const annotations = Annotations();
  annotations["@id"] = "http://archiviiify.loc";
  annotations.within.total = total;

  for (const [index, snippet] of snippets.entries()) {
    const text = snippet.text;
    for (const [index, hlspan] of snippet.highlights.entries()) {
      for (const hlbox of hlspan) {
        const region = snippet.regions[hlbox.parentRegionIdx];
        const page = snippet.pages[region.pageIdx].id;
        const x = region.ulx + hlbox.ulx;
        const y = region.uly + hlbox.uly;
        const w = hlbox.lrx - hlbox.ulx;
        const h = hlbox.lry - hlbox.uly;

        const annotation = Annotation();
        annotation["@id"] = `https://archiviiify.loc/iiif/anno/${uid()}`;
        annotation.resource.chars = query;
        annotation.on = `http://localhost:8094/iiif/${document}/${page}#xywh=${x},${y},${w},${h}`;
        annotations.resources.push(annotation);

        const hit = Hit();
        const beforeafter = region.text.match(/(.*)\<em\>(.*)\<\/em\>(.*)/);
        hit.match = hlbox.text;
        hit.before = beforeafter[1];
        hit.after = beforeafter[3];
        hit.annotations.push(annotation["@id"]);
        annotations.hits.push(hit);
      }
    }
  }

  return annotations;
}
