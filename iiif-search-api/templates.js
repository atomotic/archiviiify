export function Annotations() {
  return {
    "@context": [
      "http://iiif.io/api/presentation/2/context.json",
      "http://iiif.io/api/search/0/context.json",
    ],
    "@id": "",
    "@type": "sc:AnnotationList",
    within: {
      "@type": "sc:Layer",
      total: 1,
      ignored: [],
    },
    resources: [],
    hits: [],
  };
}

export function Annotation() {
  return {
    "@id": "",
    "@type": "oa:Annotation",
    motivation: "sc:painting",
    resource: {
      "@type": "cnt:ContentAsText",
      chars: "",
    },
    on: "",
  };
}

export function Hit() {
  return {
    "@type": "search:Hit",
    annotations: [],
    match: "",
    before: "",
    after: "",
  };
}
