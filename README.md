# INFO441Project

## Project Description

With [Project Name], we plan on providing plant-owners with a simple way to keep track of their household or garden plants. Multiple plants have different watering schedules so it's hard to keep track of when each plant needs flowering. Our app aims to help keep track of that by setting up regular watering intervals so that they don't water their plants too little or too much. We, as a team, understand these issues of keeping our vegetation growing at a healthy rate so we wish to provide an app that will answer anyone who shares our problems.

## Technical Description

### User Stories 
| Priority | User | Description |
| --- | --- | --- |
| P0 | As a user | I want to see the status of my plants that shows if it needs to be watered |
| P0 | As a user | I want to see a list of all my plants that I've added |
| P0 | As a user | I want to add new plants to my account with different watering schedules |
| P0 | As a user | I want to see when was the last time and when is the next time I need to water the plant |
| P1 | As a user | I want to track when my plant was last fertilized |
| P1 | As a user | I want to know how much water is needed for each plant |
| P1 | As a user | I want to list other information about the plant such as placement, pot size, etc. |
| P2 | As a user | I want to filter out the plants by certain traits like fruits, flowers, etc. |
| P2 | As a user | I want to search for a single plant from my list |
| P3 | As a user | I want to upload images of my plants to capture their growth |

### Technical Implementations
- To show when the plants need watering, we'll have a timer that goes down based on the water interval that the user provides and when it reaches the end, it will show a status that says "Needs watering"
- We'll have a **MySQL** database to store a table that holds information about each plant connected through a **Redis** network
- We'll connect to the database and add Plant structs to the DB
- We'll have timestamps of when the last watering was and use the water interval to calculate the next time the plant will need watering
- We'll have timestamps for the last fertilization as well
- We plan to experiment a plant API called **Trefle** to acquire any more information about a certain plant
- We'll have fields in both our structs and the DB to include information like placement or pot size that user wants to add
- We'll set up a function to grab only the rows of our DB that meet the criteria of the filter
- We'll set up a search bar to do the same thing
- If there is still time, we may add a microservice to our API to store images on