<img src="https://docuver.se/assets/archiviiify.png" width=50% height=50%>

**Archiviiify** is a streamlined command-line interface tool designed to empower you to effortlessly acquire digitized books from the Internet Archive in the JPEG2000 format and access them via IIIF on your local machine, even when disconnected from the internet. While not intended to serve as a public IIIF service for the Internet Archive, it proves exceptionally beneficial for personal research, academic pursuits, and as a resource for honing your IIIF application development skills.

In its earlier iteration, Archiviiify was a more complex amalgamation of scripts and components, as documented [here](https://literarymachin.es/archiviiify/), with the code residing in the [v1](https://github.com/atomotic/archiviiify/tree/v1) branch. The updated version features a unified command-line interface (CLI) and a single Docker image, making the installation and usage process smoother.

Please be aware that the full-text search service is not currently available in this release.

# Quickstart

Download a [release](https://github.com/atomotic/archiviiify/releases) for your platform, or build the cli

```
https://github.com/atomotic/archiviiify/
cd archiviiify
go build
```

Set up the project structure and make any necessary adjustments to the `.env` file.

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

Start the docker container. The official [iipsrv image](https://hub.docker.com/r/iipsrv/iipsrv) will be downloaded from Docker Hub

```
~ docker-compose up -d

~ docker-compose ps
NAME            IMAGE           COMMAND            SERVICE   CREATED          STATUS              PORTS
test-iipsrv-1   iipsrv/iipsrv   "/bin/sh -c run"   iipsrv    54 minutes ago   Up About a minute   9000/tcp, 0.0.0.0:9000->80/tcp, :::9000->80/tcp
```

Download an [item](https://internetarchive.readthedocs.io/en/stable/items.html) from Internet Archive. The IIIF manifest will be generated, providing access to a local Mirador viewer for browsing.

```
~ ./archiviiify get -i wholeearthcatalo00unse_1

archiviiify

[1/3] Downloading Whole Earth Catalog   Spring 1970:
 Source: https://ia903200.us.archive.org/23/items/wholeearthcatalo00unse_1/wholeearthcatalo00unse_1_jp2.zip
 236.29 MiB / 236.29 MiB [===================================================================================================================================] 100.00% 1m56s
[2/3] Generating IIIF manifest
[3/3] View the IIIF manifest at:
 http://localhost:9000/?manifest=wholeearthcatalo00unse_1
```

# Limitations

* The command-line interface (CLI) design is currently a work in progress and needs improvement in terms of aesthetics.
* The metadata embedding in IIIF manifest has not been fully tested, and may require further optimization.
* Currently, the IIIF manifest generation does not support items containing multiple subitems. This issue can be resolved by using an IIIF collection manifest.
