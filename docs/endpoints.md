# Endpoints

## Whoami

```bash
$ http localhost:3000/whoami
[
    "127.0.0.1/8",
    "::1/128",
    "172.18.0.4/16",
    "fe80::42:acff:fe12:4/64"
]
```

## Metrics

Exposes [Prometheus](https://prometheus.io/) Metrics.

## Read notes

```bash
$ http http://localhost:3000/read/note
[
    "Eat",
    "Sleep",
    "Code",
    "Repeat"
]

```

## Insert note

```bash
$ http http://localhost:3000/insert/note/Hello

[
  "Eat",
  "Sleep",
  "Code",
  "Repeat",
  "Hello"
]
```

## Delete note

```bash
$ http http://localhost:3000/delete/note/Hello

[
  "Eat",
  "Sleep",
  "Code",
  "Repeat",
]
```

## Health endpoint

```bash
$ http http://localhost:3000/health

{
    "redis-master-0": "ok",
    "redis-slave-0": "ok",
    "self": "ok"
}
```
