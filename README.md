# iOffice CLI Reservation Tool

Lists and creates reservations of spaces

## Usage

1. `cp ioffice.example.yaml ioffice.yaml`
2. Set your username and password from the iOffice app
3. Set your instance's hostname (`example` in https://example.ioffice.com)
4. Run the application to get a list of your future bookings (`go run .` or `./ioffice`)
5. From your list of reservations, take the room ID of the desk you prefer
6. Now you can also reserve this desk for a full day (`go run . 2022-03-13` or `./ioffice 2022-03-13`)
