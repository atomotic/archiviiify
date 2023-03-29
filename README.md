<img src="https://docuver.se/assets/archiviiify.png" width=50% height=50%>

**Archiviiify** is a minimalistic command-line interface tool that allows you to download digitized books from the Internet Archive in JPEG2000 format and view them with IIIF on your local computer, even when offline. Although it is not intended to be a public IIIF service for Internet Archive, it is highly useful for personal research and study, as well as for training or developing IIIF applications.

While an initial version was released several years ago, it was created using a combination of scripts and various components. This updated version is highly simplified, consisting of a single CLI and a ready-to-use Docker image. Please note that the full-text search service is not currently available, but will be included in future updates.

# Quickstart

Download a [release](https://github.com/atomotic/archiviiify/releases) for your platform, or build the cli

```
https://github.com/atomotic/archiviiify/
cd archiviiify
go build
```

Initialise the project structure (make adjustments to `.env` if needed)

```
./archiviiify init

# init — the following directories have been created
├── archiviiify
├── data
│   ├── .env
│   ├── images
│   ├── manifests
│   └── www
│       └── index.html
└── docker-compose.yml
```

Start the docker container. A [ready image](https://hub.docker.com/r/atomotic/iipsrv) will be download from Docker Hub. Repository here:  [iipsrv.docker](https://github.com/atomotic/iipsrv.docker/)

```
docker-compose up -d

docker-compose ps
NAME                IMAGE                    COMMAND             SERVICE             CREATED             STATUS              PORTS
test-iipsrv-1       atomotic/iipsrv:latest   "/init"             iipsrv              2 hours ago         Up 2 hours          0.0.0.0:9000->80/tcp

```

Download an item from Internet Archive. An IIIF manifest will be generated, and a local Mirador viewer is available to browse it.

```
./archiviiify run -i codici-immaginari-1

archiviiify
· downloading  Codici Immaginari 1
· from         https://ia903100.us.archive.org/35/items/codici-immaginari-1/codici-immaginari-1_jp2.zip
 11.86 MiB / 11.86 MiB [===================================================================] 100.00% 5s
· generating IIIF manifest
view http://localhost:9000/?manifest=codici-immaginari-1
```

