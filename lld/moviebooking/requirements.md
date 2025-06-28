Requirements 

1. User should be able to book a movie.
2. from the choice of theater.
3. and should be able to select a seat.
4. seat can have type like normal deluxe.
5. user can search by movie.
6. user can search by theater.


Entities:
Movie
    Id
    Name

Theater 
    Id 
    Name 
    Location 
    list of shows

Show   
    id 
    start time 
    end time
    movie Movie
    theater 
    screen

screen
    []seat


    