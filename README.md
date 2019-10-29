<div align="center">
  <img src="https://github.com/dappface/www.dappface.com/raw/master/static/icon-128x128.png" alt="DAPPFACE Logo" />

  <h1>DAPPFACE News loader</h1>

  <strong>
    <p>Collect new articles and tweets</p>
  </strong>

  <p>
    <a href="https://github.com/dappface/news-loader/actions?workflow=Deploy">
      <img src="https://github.com/dappface/news-loader/workflows/Deploy/badge.svg" />
    </a>
  </p>
</div>

## Start Locally

```sh
PROJECT_ID=dappface-dev \
  TWITTER_ACCESS_TOKEN=<access-token> \
  TWITTER_ACCESS_TOKEN_SECRET=<access-toekn-secret> \
  TWITTER_API_KEY=<api-key> \
  TWITTER_API_SECRET=<api-secret> \
  go run . <rss | twitter>
```

## Start Docker Container

```sh
docker build -t news-loader .
docker run --rm \
  -p 8080:8080 \
  -e HOME=$HOME \
  -v $HOME:$HOME \
  -e PROJECT_ID=dappface-dev  \
  -e TWITTER_ACCESS_TOKEN=<access-token> \
  -e TWITTER_ACCESS_TOKEN_SECRET=<access-toekn-secret> \
  -e TWITTER_API_KEY=<api-key> \
  -e TWITTER_API_SECRET=<api-secret> \
  app <rss | twitter>
```
