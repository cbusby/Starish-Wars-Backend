**Game status and messages**

**Possible Game Statuses:**
```
	AWAITING_SHIPS
	PLAYER_1_ACTIVE
	PLAYER_2_ACTIVE
	GAME_OVER
```

**Game Play**

The first player requests a game by sending a `POST` message to the server. The server sends a response with HTTP status code `201` (created), the gameId in a Location header, and the current status of the game and empty player grids.

[Create Game Response](./create_game_response.json)

At any time either player can get the current state of the game with a `GET` request and the gameId.

Until two players join and place ships, the `GET` request returns the `AWAITING_SHIPS` state.

A player can send a `PUT` message with his ship placements. The server will not overwrite an existing ship placement, so any changes to existing ship placements are ignored.

[Place Ships Request](./place_ships_request.json)

If player 2 submits her ship placement before refreshing her state with a `GET` request, her ship placement will have no information for player 1's ships. The server will NOT overwrite an existing ship placement, so it will simply take player 2's ship placement, add it to the current game state, and change the status to `PLAYER_1_ACTIVE`.

Once both players have placed ships, the server will return state `PLAYER_1_ACTIVE`. player 1 goes first, and the server returns the active state to `GET` requests to that gameId:

[Player 1 Shot Request](./player_1_shot_request.json)

Only the active player can send `PUT` requests to the server with her shots:

[Player 1 Shot Response](./player_1_shot_response.json)

Turns alternate until all of one player's ships are sunk. Then the `GAME_OVER` state is returned.

[Game Over Response](./game_over_response.json)

**Server validation:**

- The grid is 10x10, A1-J10. All ships and shots must be fully on this grid.
- All ships must be placed ("carrier": 5, "battleship": 4, cruister: 3, "submarine": 3, "destroyer": 2)
- All ship spaces must be adjacent and in a straight line (vertical or horizontal).
- Ships cannot overlap. A grid space can only be occupied by one ship square.
- Only the active player can PUT a shot.
- The server maintains a record of the current status and state of the game.
- All ship placements must match from turn to turn.
- All earlier shots must be the same.
- The game is over when all of one player's ship spaces have been hit by a shot.