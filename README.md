# interview_restapi

### Running api

baseUrl = https://belle-saucisson-47408.herokuapp.com/

Get all movies

    baseUrl/movies?page=&record=10

Get single movie by id

    baseUrl/movies/{id}

Get Movies by date range

    baseUrl/movies/date?start_date=2022-07-11&end_date=2022-07-12

Get movies per page selected, pass a parameter ==> page to get a particular page ==> pass a parameter record for number of records per page

    baseUrl/movies?page=2&record=10


### Build a Docker Image to run app locally

In the root folder of the app run:

    docker build -t interview_app .

Then run:

    docker run -it -p 3000:3000  --rm --name my-running-app interview_app
