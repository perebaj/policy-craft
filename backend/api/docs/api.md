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
        "value": 10
        }'
```
