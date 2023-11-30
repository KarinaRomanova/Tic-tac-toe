package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const indexHTML = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Крестики-нолики</title>
    <style>
        body {
            background-color: black;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }

        .game-board {
            display: grid;
            grid-template-columns: repeat(3, 150px);
            grid-template-rows: repeat(3, 150px);
            gap: 0;
            width: 450px;
            height: 450px;
        }

        .game-board div {
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 60px;
            font-family: 'Arial Black', Gadget, sans-serif; /* изменение шрифта на мультяшный */
            cursor: pointer;
            border: 2px solid white;
            color: white;
        }

        .game-board div:nth-child(3n + 1) {
            border-left: none;
        }

        .game-board div:nth-child(3n) {
            border-right: none;
        }

        .game-board div:nth-child(-n + 3) {
            border-top: none;
        }

        .game-board div:nth-last-child(-n + 3) {
            border-bottom: none;
        }
    </style>
</head>
<body>
    <div class="game-board">
        <div onclick="makeMove(this, 0, 0)"></div>
        <div onclick="makeMove(this, 0, 1)"></div>
        <div onclick="makeMove(this, 0, 2)"></div>
        <div onclick="makeMove(this, 1, 0)"></div>
        <div onclick="makeMove(this, 1, 1)"></div>
        <div onclick="makeMove(this, 1, 2)"></div>
        <div onclick="makeMove(this, 2, 0)"></div>
        <div onclick="makeMove(this, 2, 1)"></div>
        <div onclick="makeMove(this, 2, 2)"></div>
    </div>

    <script>
        let currentPlayer = 'X';
        let movesCount = 0;
        let gameBoard = [
            ['', '', ''],
            ['', '', ''],
            ['', '', '']
        ];

        function makeMove(cell, row, col) {
            if (cell.innerHTML === '' && !checkWin()) {
                cell.innerHTML = currentPlayer;
                gameBoard[row][col] = currentPlayer;
                movesCount++;

                if (checkWin()) {
                    setTimeout(function() {
                        alert('Игрок ' + currentPlayer + ' победил!');
                    }, 100); // Пауза перед алертом победы
                } else if (movesCount === 9) {
                    alert('Ничья!');
                } else {
                    currentPlayer = currentPlayer === 'X' ? 'O' : 'X';
                }
            }
        }

        function checkWin() {
            for (let i = 0; i < 3; i++) {
                if (gameBoard[i][0] !== '' && gameBoard[i][0] === gameBoard[i][1] && gameBoard[i][1] === gameBoard[i][2]) {
                    return true;
                }
                if (gameBoard[0][i] !== '' && gameBoard[0][i] === gameBoard[1][i] && gameBoard[1][i] === gameBoard[2][i]) {
                    return true;
                }
            }
            if (gameBoard[0][0] !== '' && gameBoard[0][0] === gameBoard[1][1] && gameBoard[1][1] === gameBoard[2][2]) {
                return true;
            }
            if (gameBoard[0][2] !== '' && gameBoard[0][2] === gameBoard[1][1] && gameBoard[1][1] === gameBoard[2][0]) {
                return true;
            }
            return false;
        }
    </script>
</body>
</html>
`

func main() {
	http.HandleFunc("/", serveIndex)
	fmt.Println("Сервер запущен. Перейдите по адресу http://localhost:8080 в вашем браузере.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl := template.Must(template.New("index").Parse(indexHTML))
	tmpl.Execute(w, nil)
}
