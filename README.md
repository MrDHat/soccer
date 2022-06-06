# soccer-manager

The purpose of this project is to show best practices for building a graphQL API with Golang & Postgres

## Requirements

- Users must be able to create an account and log in using the API.
-  Each user can have only one team (user is identified by an email).
-  When the user is signed up, they should get a team of 20 players (the system should generate players):
   - 3 goalkeepers
   - 6 defenders
   - 6 midfielders
   - 5 attackers
- Each player has an initial value of $1,000,000
- Each team has an additional $5,000,000 to buy other players.
- When logged in, a user can see their team and player information
- The team has the following information:
   - Team name and team country (can be edited).
   - Team value (sum of player values).
- The Player has the following information
   - First name, last name, country (can be edited by a team owner).
   - Age (random number from 18 to 40) and market value.
- A team owner can set the player on a transfer list
- When a user places a player on a transfer list, they must set the asking price/value for
this player. This value should be listed on a market list. When another user/team buys
this player, they must be bought for this price.
- Each user should be able to see all players on a transfer list.
- With each transfer, team budgets are updated.
- When a player is transferred to another team, their value should be increased between 10 and 100 percent.



### Note about Migrations

Migrations should be run using [migrate]([https://github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)) by running bash inside the service container. The default folder for migrations is `db/migrations`
