# Read API Project

## Description

The Read API Project provides a RESTful API for reading data from a database. It allows users to retrieve information from various endpoints based on their requirements.

## Table of Contents

1. [Design Considerations](#design-considerations)
2. [Assumptions](#assumptions)
3. [Trade-offs](#trade-offs)
4. [Usage](#usage)
5. [Installation](#installation)
6. [Configuration](#configuration)
7. [Endpoints](#endpoints)
8. [Examples](#examples)

## Design Considerations

The API was designed with a focus on simplicity, ease of use, and performance. It follows the REST architectural style, using HTTP methods for communication and JSON for data interchange. The database schema was optimized for read-heavy operations to ensure efficient retrieval of data.

## Assumptions

- Users have prior knowledge of RESTful APIs and can understand request/response formats.
- The database is properly configured and contains the necessary data for retrieval.
- Authentication and authorization mechanisms are handled by an external service and are not part of this API.

## Trade-offs

- **Simplicity vs. Flexibility**: The API prioritizes simplicity over flexibility, limiting the number of configuration options to reduce complexity.
- **Performance**: Certain design choices were made to optimize performance

## Usage

To use the Read API, send HTTP requests to the provided endpoints with the appropriate parameters. 

## Installation

To install and run the API locally, follow these steps:

1. Clone this repository to your local machine.
2. Install Docker 
3. Run "make run-service"

## Configuration

The API can be configured using environment variables. Some of the key variables include database connection settings. Refer to the `config/config.yaml` file for a list of required variables.

## Endpoints

The API exposes the following endpoints:

- `GET /services`: Retrieves a list of resources.
- `GET /versions`: Retrieves a specific resource by ID.

For detailed information about each endpoint, refer to the API documentation.

## Examples

### Retrieve Resources

**Request:**

```
GET /services
```
curl --location --request GET 'localhost:8081/services?page=1&limit=2&sort_by=name&sort_type=D'

**Response:**

```json
[
    {
        "service_id": 3,
        "name": "Service C",
        "description": "Description for Service C",
        "version_count": 3
    },
    {
        "service_id": 2,
        "name": "Service B",
        "description": "Description for Service B",
        "version_count": 1
    }
]

```
curl --location --request GET 'localhost:8081/services?page=1&limit=2&name=Service A'

```json
[
    {
        "service_id": 1,
        "name": "Service A",
        "description": "Description for Service A",
        "version_count": 2
    }
]
```


### Retrieve Resource by ID

**Request:**

GET /versions

```
curl --location --request GET 'localhost:8081/versions?service_id=1'

```

**Response:**

```json
[
    {
        "name": "v1",
        "description": "Description for Service A Version 1"
    },
    {
        "name": "v2",
        "description": "Description for Service A Version 2"
    }
]
```



