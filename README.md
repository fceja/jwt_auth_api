# Desciption
Go API that provides endponts to: create, retrieve, and log-in users. </br>
Implemented with JWT authentication.

# Installation
1. Download and install:
   - Go (Golang)
     - https://go.dev/dl/
   - Docker
     - Install Docker: https://docs.docker.com/get-docker/
     - Pull Postgres Docker image: ```docker pull postgres```
2. Clone repo
3. Navigate to project root, install all packages
   - ```go get```
4. Initialize Postgres Docker container
   - ```- docker run --name postgres-db -e POSTGRES_PASSWORD=ecretPassword -p 5432:5432 -d postgres```
5. Start server, at project root run:
   - ```make run```

# Technologies & Tools
<a href="https://go.dev/" target="_blank" rel="noreferrer">
  <img
    src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg"
    width="40"
    height="40"
    alt="golang"
  />
</a>
<a href="https://www.postgresql.org/" target="_blank" rel="noreferrer">
  <img
    src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original-wordmark.svg"
    width="40"
    height="40"
    alt="posgres"
  />
</a>
<a href="https://www.docker.com/" target="_blank" rel="noreferrer">
  <img
    src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-plain-wordmark.svg"
    alt="docker"
    width="40"
    height="40"
  />
</a>
<a href="https://jwt.io/" target="_blank" rel="noreferrer">
  <img
    src="https://jwt.io/img/favicon/apple-icon-72x72.png"
    width="40"
    height="40"
    alt="jsonwebtoken"
  />
</a>
