# archiviiify

Download digitized books from Internet Archive and rehost on your own infrastructure using IIIF with full-text search 

Required: Docker, [Deno](https://deno.land), [jq](https://github.com/stedolan/jq), [pup](https://github.com/ericchiang/pup), [sd](https://github.com/chmln/sd), [ia](https://github.com/jjjake/internetarchive), [parallel](https://www.gnu.org/software/parallel/), [tesseract](https://github.com/tesseract-ocr/tesseract)

Start the containers:

```
git clone https://github.com/atomotic/archiviiify
cd archiviiify
docker-compose up
```

Given an Internet Archive books at https://archive.org/details/ITEM run:

```
./scripts/get ITEM
./scripts/iiif ITEM
./scripts/ocr ITEM
./scripts/ocr-fix ITEM
./scripts/index ITEM
```

Read the book with Mirador3 at `http://localhost:8094/index.html?manifest=ITEM`