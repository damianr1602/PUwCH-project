// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple API",
    "title": "Simple movies API",
    "contact": {
      "email": "damianr1602@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "virtserver.swaggerhub.com",
  "basePath": "/damianr1602/chmury/1.0.0",
  "paths": {
    "/movies": {
      "get": {
        "description": "By passing in the appropriate options, you can search for\navailable movies in the system\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiUsers"
        ],
        "summary": "Show movies",
        "operationId": "AllMovies",
        "parameters": [
          {
            "type": "string",
            "description": "pass an optional search string for looking up movie",
            "name": "searchString",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "number of records to skip for pagination",
            "name": "skip",
            "in": "query"
          },
          {
            "maximum": 15,
            "type": "integer",
            "format": "int32",
            "description": "maximum number of records to return",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "search results matching criteria",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/movie"
              }
            }
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      },
      "put": {
        "description": "Update existing movie",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiUsers"
        ],
        "summary": "Update existing movie",
        "operationId": "UpdateMovie",
        "parameters": [
          {
            "description": "Movie to update",
            "name": "Movie",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/movie"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "movie updated successfully",
            "schema": {
              "$ref": "#/definitions/result"
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      },
      "post": {
        "description": "Adds a movie to the system",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiUsers"
        ],
        "summary": "adds a movie",
        "operationId": "CreateMovie",
        "parameters": [
          {
            "description": "Movie to add",
            "name": "Movie",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/movie"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "item created",
            "schema": {
              "$ref": "#/definitions/movie"
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "409": {
            "description": "an existing item already exists"
          }
        }
      },
      "delete": {
        "description": "Delete existing movie",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiUsers"
        ],
        "summary": "Delete existing movie",
        "operationId": "DeleteMovie",
        "parameters": [
          {
            "description": "Movie to delete",
            "name": "Movie",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/movie"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "movie deleted successfully",
            "schema": {
              "$ref": "#/definitions/result"
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      }
    },
    "/movies/{id}": {
      "get": {
        "description": "By passing in the appropriate options, you can search for\navailable movie in the system\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiUsers"
        ],
        "summary": "Search best result matching criteria",
        "operationId": "FindMovie",
        "parameters": [
          {
            "type": "string",
            "description": "pass an optional search string for looking up movie",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "search result matching criteria",
            "schema": {
              "$ref": "#/definitions/movie"
            }
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      }
    }
  },
  "definitions": {
    "movie": {
      "type": "object",
      "properties": {
        "_id": {
          "type": "object",
          "properties": {
            "$oid": {
              "type": "string"
            }
          },
          "example": [
            "5ed2882a10ab722e54e07f7f"
          ]
        },
        "actors": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              [
                "Tom Hanks",
                "Jan Kowalski"
              ]
            ]
          }
        },
        "awards": {
          "type": "object",
          "properties": {
            "nominations": {
              "type": "integer"
            },
            "text": {
              "type": "string"
            },
            "wins": {
              "type": "integer"
            }
          }
        },
        "countries": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "USA"
            ]
          }
        },
        "director": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            "John Lasseter"
          ]
        },
        "genres": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            [
              "Adventure",
              "Drama"
            ]
          ]
        },
        "imdb": {
          "type": "object",
          "properties": {
            "id": {
              "type": "string"
            },
            "rating": {
              "type": "integer"
            },
            "votes": {
              "type": "integer"
            }
          }
        },
        "metacritic": {
          "type": "integer"
        },
        "plot": {
          "type": "string"
        },
        "poster": {
          "type": "string"
        },
        "rated": {
          "type": "string"
        },
        "released": {
          "type": "object",
          "properties": {
            "$date": {
              "type": "string"
            }
          }
        },
        "reviews": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "comment": {
                "type": "string"
              },
              "date": {
                "type": "object",
                "properties": {
                  "$date": {
                    "type": "string"
                  }
                }
              },
              "name": {
                "type": "string"
              },
              "rating": {
                "type": "integer"
              }
            }
          }
        },
        "runtime": {
          "type": "integer"
        },
        "title": {
          "type": "string"
        },
        "tomato": {
          "type": "object",
          "properties": {
            "consensus": {
              "type": "string"
            },
            "fresh": {
              "type": "integer"
            },
            "image": {
              "type": "string"
            },
            "meter": {
              "type": "integer"
            },
            "rating": {
              "type": "integer"
            },
            "reviews": {
              "type": "integer"
            },
            "userMeter": {
              "type": "integer"
            },
            "userRating": {
              "type": "integer"
            },
            "userReviews": {
              "type": "integer"
            }
          }
        },
        "type": {
          "type": "string"
        },
        "writers": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            [
              "John Lasseter"
            ]
          ]
        },
        "year": {
          "type": "integer"
        }
      }
    },
    "result": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/result"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/result"
      }
    }
  },
  "tags": [
    {
      "description": "Operations available to regular apiUsers",
      "name": "apiUsers"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple API",
    "title": "Simple movies API",
    "contact": {
      "email": "damianr1602@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "virtserver.swaggerhub.com",
  "basePath": "/damianr1602/chmury/1.0.0",
  "paths": {
    "/movies": {
      "get": {
        "description": "By passing in the appropriate options, you can search for\navailable movies in the system\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiUsers"
        ],
        "summary": "Show movies",
        "operationId": "AllMovies",
        "parameters": [
          {
            "type": "string",
            "description": "pass an optional search string for looking up movie",
            "name": "searchString",
            "in": "query"
          },
          {
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "description": "number of records to skip for pagination",
            "name": "skip",
            "in": "query"
          },
          {
            "maximum": 15,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "description": "maximum number of records to return",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "search results matching criteria",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/movie"
              }
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/result"
            }
          }
        }
      },
      "put": {
        "description": "Update existing movie",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiUsers"
        ],
        "summary": "Update existing movie",
        "operationId": "UpdateMovie",
        "parameters": [
          {
            "description": "Movie to update",
            "name": "Movie",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/movie"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "movie updated successfully",
            "schema": {
              "$ref": "#/definitions/result"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/result"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/result"
            }
          }
        }
      },
      "post": {
        "description": "Adds a movie to the system",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiUsers"
        ],
        "summary": "adds a movie",
        "operationId": "CreateMovie",
        "parameters": [
          {
            "description": "Movie to add",
            "name": "Movie",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/movie"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "item created",
            "schema": {
              "$ref": "#/definitions/movie"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/result"
            }
          },
          "409": {
            "description": "an existing item already exists"
          }
        }
      },
      "delete": {
        "description": "Delete existing movie",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiUsers"
        ],
        "summary": "Delete existing movie",
        "operationId": "DeleteMovie",
        "parameters": [
          {
            "description": "Movie to delete",
            "name": "Movie",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/movie"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "movie deleted successfully",
            "schema": {
              "$ref": "#/definitions/result"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/result"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/result"
            }
          }
        }
      }
    },
    "/movies/{id}": {
      "get": {
        "description": "By passing in the appropriate options, you can search for\navailable movie in the system\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiUsers"
        ],
        "summary": "Search best result matching criteria",
        "operationId": "FindMovie",
        "parameters": [
          {
            "type": "string",
            "description": "pass an optional search string for looking up movie",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "search result matching criteria",
            "schema": {
              "$ref": "#/definitions/movie"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/result"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "MovieAwards": {
      "type": "object",
      "properties": {
        "nominations": {
          "type": "integer"
        },
        "text": {
          "type": "string"
        },
        "wins": {
          "type": "integer"
        }
      }
    },
    "MovieID": {
      "type": "object",
      "properties": {
        "$oid": {
          "type": "string"
        }
      },
      "example": [
        "5ed2882a10ab722e54e07f7f"
      ]
    },
    "MovieImdb": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "rating": {
          "type": "integer"
        },
        "votes": {
          "type": "integer"
        }
      }
    },
    "MovieReleased": {
      "type": "object",
      "properties": {
        "$date": {
          "type": "string"
        }
      }
    },
    "MovieReviewsItems0": {
      "type": "object",
      "properties": {
        "comment": {
          "type": "string"
        },
        "date": {
          "type": "object",
          "properties": {
            "$date": {
              "type": "string"
            }
          }
        },
        "name": {
          "type": "string"
        },
        "rating": {
          "type": "integer"
        }
      }
    },
    "MovieReviewsItems0Date": {
      "type": "object",
      "properties": {
        "$date": {
          "type": "string"
        }
      }
    },
    "MovieTomato": {
      "type": "object",
      "properties": {
        "consensus": {
          "type": "string"
        },
        "fresh": {
          "type": "integer"
        },
        "image": {
          "type": "string"
        },
        "meter": {
          "type": "integer"
        },
        "rating": {
          "type": "integer"
        },
        "reviews": {
          "type": "integer"
        },
        "userMeter": {
          "type": "integer"
        },
        "userRating": {
          "type": "integer"
        },
        "userReviews": {
          "type": "integer"
        }
      }
    },
    "movie": {
      "type": "object",
      "properties": {
        "_id": {
          "type": "object",
          "properties": {
            "$oid": {
              "type": "string"
            }
          },
          "example": [
            "5ed2882a10ab722e54e07f7f"
          ]
        },
        "actors": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              [
                "Tom Hanks",
                "Jan Kowalski"
              ]
            ]
          }
        },
        "awards": {
          "type": "object",
          "properties": {
            "nominations": {
              "type": "integer"
            },
            "text": {
              "type": "string"
            },
            "wins": {
              "type": "integer"
            }
          }
        },
        "countries": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "USA"
            ]
          }
        },
        "director": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            "John Lasseter"
          ]
        },
        "genres": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            [
              "Adventure",
              "Drama"
            ]
          ]
        },
        "imdb": {
          "type": "object",
          "properties": {
            "id": {
              "type": "string"
            },
            "rating": {
              "type": "integer"
            },
            "votes": {
              "type": "integer"
            }
          }
        },
        "metacritic": {
          "type": "integer"
        },
        "plot": {
          "type": "string"
        },
        "poster": {
          "type": "string"
        },
        "rated": {
          "type": "string"
        },
        "released": {
          "type": "object",
          "properties": {
            "$date": {
              "type": "string"
            }
          }
        },
        "reviews": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/MovieReviewsItems0"
          }
        },
        "runtime": {
          "type": "integer"
        },
        "title": {
          "type": "string"
        },
        "tomato": {
          "type": "object",
          "properties": {
            "consensus": {
              "type": "string"
            },
            "fresh": {
              "type": "integer"
            },
            "image": {
              "type": "string"
            },
            "meter": {
              "type": "integer"
            },
            "rating": {
              "type": "integer"
            },
            "reviews": {
              "type": "integer"
            },
            "userMeter": {
              "type": "integer"
            },
            "userRating": {
              "type": "integer"
            },
            "userReviews": {
              "type": "integer"
            }
          }
        },
        "type": {
          "type": "string"
        },
        "writers": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            [
              "John Lasseter"
            ]
          ]
        },
        "year": {
          "type": "integer"
        }
      }
    },
    "result": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/result"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/result"
      }
    }
  },
  "tags": [
    {
      "description": "Operations available to regular apiUsers",
      "name": "apiUsers"
    }
  ]
}`))
}