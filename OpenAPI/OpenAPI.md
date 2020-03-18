# OpenAPI

https://github.com/OAI/OpenAPI-Specification



Some examples of possible media type definitions:

```
  text/plain; charset=utf-8
  application/json
  application/vnd.github+json
  application/vnd.github.v3+json
  application/vnd.github.v3.raw+json
  application/vnd.github.v3.text+json
  application/vnd.github.v3.html+json
  application/vnd.github.v3.full+json
  application/vnd.github.v3.diff
  application/vnd.github.v3.patch
```



| [`type`](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#dataTypes) | [`format`](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#dataTypeFormat) | Comments                                                     |
| ------------------------------------------------------------ | :----------------------------------------------------------- | ------------------------------------------------------------ |
| `integer`                                                    | `int32`                                                      | signed 32 bits                                               |
| `integer`                                                    | `int64                                                       | signed 64 bits (a.k.a long)                                  |
| `number`                                                     | `float`                                                      |                                                              |
| `number`                                                     | `double`                                                     |                                                              |
| `string`                                                     |                                                              |                                                              |
| `string`                                                     | `byte`                                                       | base64 encoded characters                                    |
| `string`                                                     | `binary`                                                     | any sequence of octets                                       |
| `boolean`                                                    |                                                              |                                                              |
| `string`                                                     | `date`                                                       | As defined by `full-date` - [RFC3339](https://xml2rfc.ietf.org/public/rfc/html/rfc3339.html#anchor14) |
| `string`                                                     | `date-time`                                                  | As defined by `date-time` - [RFC3339](https://xml2rfc.ietf.org/public/rfc/html/rfc3339.html#anchor14) |
| `string`                                                     | `password`                                                   | A hint to UIs to obscure input.                              |

#### JSON vs YAML

JSON:

```json
{
    "swagger": "2.0",
    "info": {
        "version": "1.0.0",
        "title": "Simple API",
        "description": "A simple API to learn how to write OpenAPI Specification"
    },
    "schemes": [
        "https"
    ],
    "host": "simple.api",
    "basePath": "/openapi101",
    "paths": {
        "/persons": {
            "get": {
                "summary": "Gets some persons",
                "description": "Returns a list containing all persons.",
                "responses": {
                    "200": {
                        "description": "A list of Person",
                        "schema": {
                            "type": "array",
                            "items": {
                                "properties": {
                                    "firstName": {
                                        "type": "string"
                                    },
                                    "lastName": {
                                        "type": "string"
                                    },
                                    "username": {
                                        "type": "string"
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
```

YAML:

```yaml
swagger: "2.0"

info:
  version: 1.0.0
  title: Simple API
  description: A simple API to learn how to write OpenAPI Specification

schemes:
  - https
host: simple.api
basePath: /openapi101

paths:
  /persons:
    get:
      summary: Gets some persons
      description: Returns a list containing all persons.
      responses:
        200:
          description: A list of Person
          schema:
            type: array
            items:
              required:
                - username
              properties:
                firstName:
                  type: string
                lastName:
                  type: string
                username:
                  type: string
```



#### Swagger Editor

http://editor.swagger.io/



---





```YAML
swagger: "2.0"

info:
  version: 1.0.0
  title: Simple API
  description: A simple API to learn how to write OpenAPI Specification

schemes:
  - https
host: simple.api
basePath: /openapi101

paths: {}
```

##### OpenAPI 規範版本

```yaml
swagger: "2.0"
```

可以填入 ```openapi: "3.0.0"```



```
info:
  version: 1.0.0
  title: Simple API
  description: A simple API to learn how to write OpenAPI Specification
```



```
schemes:
  - https
host: simple.api
basePath: /openapi101
```









