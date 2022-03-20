# Golang Company Management APIs
    REST API microservice to handle Companies

## Setup Postgres Database

    $ cd migrations
    $ goose postgres "<POSTGRES_DB_URL>" up
    $ cd ..

## Run Development Environment

    Set the necessary environment variables in development.env
    
    $ source config/development.env
    $ go run main.go

## To Run Tests

    $ go test ./...

## To run the complete project

    The complete project can be run with a single command with the tests using the below command

    $ chmod +x ./startup.sh // Give run permission for the script
    $ ./startup.sh

## API Reference

    1. GET /company/{company_ids}?filterKey=filterValue 
        
        To get the details of companies using list of company ids
        * filterKey - company attributes
        * filterValue - value of attributes
    
    2. POST /company

        To create a company
        Sample Post Data:
        ```
            {
                "name":"XM",  // Mandatory
                "code":"1234",  // Mandatory
                "country":"USA",  // Mandatory
                "website":"XM.com"  // Not Mandatory
                "phone":"123456"  // Mandatory
            }
        ```

    3. PATCH /company/{company_id}

        To update company details
        Supported json fields include all the company attributes

    4. DELETE /company/{company_id}

        To delete a company from database with the given company id
        
    