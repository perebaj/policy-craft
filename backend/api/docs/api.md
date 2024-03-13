# API examples

Here it's possible to find some useful examples of how to use the API.

# Policies
Endpoints for managing policies.

## POST policies/

This endpoint creates a new policy.

- The `criteria` field can be one of the following values: `>`, `<`, `>=`, `<=`, `==`.
- The id field must be a UUID.
- The `value` and `priority` fields must be integers.

curl request example:


```bash
curl -i -X POST http://localhost:8080/policies \
     -H "Content-Type: application/json" \
     -d '{
        "id": "a43cafc3-87ad-4e13-9e42-fbd7113b7e82",
        "name": "Sample Policy",
        "criteria": ">",
        "value": 10,
        "success_case": true,
        "priority": 1
        }'
```

Response:

```bash
HTTP/1.1 200 OK
HTTP/1.1 400 Bad Request
HTTP/1.1 500 Internal Server Error
```

## GET policies/

Returns a list of all policies that have been created.

curl request example:

```bash
curl -i -X GET http://localhost:8080/policies
```

Response example:

```json
[
    {
        "id": "a43cafc3-87ad-4e13-9e42-fbd7113b7e82",
        "name": "Sample Policy",
        "criteria": ">",
        "value": 10,
        "success_case": true,
        "priority": 1
    }
]
```

# Execution Engine

Endpoints for managing the execution engine.

## POST /execution-engine

The `POST /execution-engine` will return errors if the value is not an integer, or if the key doesn't have a respective created policy.

```bash
curl -i -X POST http://localhost:8080/execution-engine \
     -H "Content-Type: application/json" \
     -d '{
        "age": 20,
        "income": 1000
     }'
```

Response example:

```json
{
    "result": true
}

HTTP 1.1 200 OK
```
