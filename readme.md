gRPC Project Description

Follow these setup to run the Project:

1. Get the code from Github:
          $ git clone https://github.com/Nagarjunhabbu/TCorp.git
          
2. Install docker :
        - $ sudo apt install docker.io
        - $ sudo snap install docker
        - $ sudo docker images
        - $ sudo docker ps -a
        - $ sudo docker ps  (To check for containers in a running state, use the following command)


3. Run Server:
    $ docker-compose up
  this will run both http & grpc services

4. http endpoints for testing: 

1. ### GetUserByID
```bash
curl --request GET \
  --url http://localhost:8000/v1/user/1
```
         
Response:
```json
{
	"user": {
		"id": 1,
		"name": "Steve",
		"city": "LA",
		"phone": "1234567890",
		"height": 5.8,
		"married": true
	}
}
```

2. ### GetUserByIds

```bash
curl --request POST \
  --url http://localhost:8000/v1/user \
  --header 'Content-Type: application/json' \
  --data '{
	"ids":[
		1,2
	]
}'
```

Response:
```json
{
	"user": [
		{
			"id": 1,
			"name": "Steve",
			"city": "LA",
			"phone": "1234567890",
			"height": 5.8,
			"married": true
		},
		{
			"id": 2,
			"name": "Virat",
			"city": "CA",
			"phone": "4567890233",
			"height": 6.2,
			"married": true
		}
	]
}
```

3. ### SearchUserByFilter

```bash
curl --request GET \
  --url 'http://localhost:8000/v1/user?name=GANESH'
  ```

Response:
```json
{

"user":[
  {
		ID:      5,
		Fname:   "Ganesh",
		City:    "SA",
		Phone:   "4587890233",
		Height:  6.2,
		Married: true,
  },
],
	}
  ```

          

        
