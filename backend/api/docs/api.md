# API examples

# Policies
Endpoints for managing policies.

## POST policies/

The `criteria` field can be one of the following values: `>`, `<`, `>=`, `<=`, `==`.

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

## GET policies/

Returns a list of all policies.

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

## POST execution-engine/

```bash
curl -i -X POST http://localhost:8080/execution-engine \
     -H "Content-Type: application/json" \
     -d '{
        "age": 20,
        "income": 1000
     }'
```
