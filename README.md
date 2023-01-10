# Paste My Taste 🎶

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/dewey/paste-my-taste/LICENSE)

Is a replacement for the now defunct Paste My Taste site. I build this replacement to play around with VueJS. Additionally
to providing the same features as the old site this one also includes the links to the artists on Last.FM in a Reddit-compatible post format. Feature requests are welcome! 👨🏻‍💻

**It's currently available here: [pastemytaste.site](https://pastemytaste.site)**

Up until December 2022 it was available on pastemytaste[.]com but unfortunately due to an accident the domain expired.

If you need a demo username try mine: [tehwey](https://www.last.fm/user/tehwey)


## Configuration and Operation

### Environment

The following environment variables are available, they all have sensible defaults and don't need to be set explicity, except the API token.

- `API_KEY`: The Last.FM API key
- `ENVIRONMENT`: Environment can be `prod` or `develop`. `develop` sets the loglevel to `info` (Default: `develop`)
- `PORT`: Port that PMT is running on (Default: `8080`)


There are two available storage backends right now. An in-memory and a disk backed implementation. Depending on which one you choose
there are additional options you can set.

- `STORAGE_BACKEND`: Set to `memory` to keep everything in-memory or `persistent` to persist the cache to disk. (Default: `memory`)

**In Memory**

- `CACHE_EXPIRATION`: The expiration time of the cache in minutes (Default: 30)
- `CACHE_EXPIRED_PURGE`: The interval at which the expired cache elements will be purged in minutes (Default: 60)

**Persistent**

- `STORAGE_PATH`: Set the storage location of the cache on disk. (Default: `./pmt-data`)



### Running in development

In the root directory of the project:

```
npm --prefix web/ run build && PORT=9999 API_KEY=changeme go run pmt.go
```


### Run with Docker

You can change all these options in the included docker-compose files and use `docker-compose -f docker-compose.yml up -d` to run the project.