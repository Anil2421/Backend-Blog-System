# Steps
1. clone the repo
1. open app.env file to change the basic settings like port/db_name
1. make sure same db_name is used as provided in docker-compose file line 21
1. use this cmd, where the docker is installed:
`docker-compose -f docker-compose.yml up`
1. It will bring up 2 services, postgres and simpleblog service
1. Sample output can be seen in file `Output.txt`
1. Further more we can check the logs of the server inside a container in `./blogs.log` file


## Testing
1. Try running the postman for following requests
### POST:
	Request: localhost:8083/articles
	Body: {
      "id": 1,
      "title": "DsdasJggeru",
      "content": "Dhgis is the content of the second blog",
      "author": "DGegrry Mulligan"
    }
    Response: 
    {
    "data": {
        "id": 1,
        "title": "DsdasJggeru",
        "content": "Dhgis is the content of the second blog",
        "author": "DGegrry Mulligan"
        },
    "message": "Success",
    "status": 200
    }

### Get:(single record)
	Request: localhost:8083/article/3
	Response:
	{
    "data": 
        {
        "id": 3,
        "title": "DsdasJggeru",
        "content": "Dhgis is the content of the second blog",
        "author": "DGegrry Mulligan"
        },
    "message": "success",
    "status": 200
    }
	Request: localhost:8083/article/90
	Response:
	{
    "data": null,
    "message": "Fail",
    "status": 404
    }

### GET: (all)
	localhost:8083/articles
	Response:
	{
    "data": [
        {
        "id": 33,
        "title": "DsdasJggeru",
        "content": "Dhgis is the content of the second blog",
        "author": "DGegrry Mulligan"
        },
        {
        "id": 3,
        "title": "DsdasJggeru",
        "content": "Dhgis is the content of the second blog",
        "author": "DGegrry Mulligan"
        },
        {
        "id": 30,
        "title": "DsdasJggeru",
        "content": "Dhgis is the content of the second blog",
        "author": "DGegrry Mulligan"
        },
        {
        "id": 1,
        "title": "DsdasJggeru",
        "content": "Dhgis is the content of the second blog",
        "author": "DGegrry Mulligan"
        }
    ],
    "message": "success",
    "status": 200
    }
