# Guessing Game in Go

This is a simple guessing game implemented in the Go programming language. The game generates a random number, and the player's task is to guess that number. After each guess, the game provides feedback on whether the guess is too high, too low, or correct.

## How to Play

1. Clone or download this repository to your local machine.

   ```shell
   git clone https://github.com/your-username/guessing-game-go.git
   ```

2. Navigate to the project directory.

   ```shell
   cd guessing-game-go
   ```

3. Build the game using the `go build` command.

   ```shell
   go build
   ```

4. Run the game.

   ```shell
   ./guessing-game-go
   ```

5. The game will generate a random number between a predefined range and ask you to guess it. You will be prompted to enter your guess.

6. Enter a number and press Enter.

7. The game will provide feedback, telling you if your guess is too high, too low, or correct.

8. Continue guessing until you guess the correct number.

9. After you guess the correct number, the game will display the number of attempts it took you to guess correctly and ask if you want to play again.

## Customizing the Game

You can customize the game by modifying the following parameters in the `main.go` file:

- You can  change the message prompts and feedback messages to make the game more personalized.

```go
fmt.Printf("Guess the number between %d and %d: ", min, max)
fmt.Println("Congratulations! You guessed the number in", attempts, "attempts.")
```

Feel free to modify the game logic and user interface to suit your preferences.


Enjoy the game and happy guessing!