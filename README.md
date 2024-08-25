# go-yaml

Clone the repository.

```sh
git clone https://github.com/chrisbradleydev/go-yaml.git .
```

Build Docker image.

```sh
docker build -t go-yaml .
```

Run Docker container.

```sh
docker run -v $(pwd):/app -v /app/tmp go-yaml
```
