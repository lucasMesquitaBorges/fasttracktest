<h1>Fast Track Code Test | Quiz</h1>
Simple multiple choice quiz game with 5 questions and single answers developed with Golang as a back-end API and Cobra as a CLI that talks with the API.
<h1>Usage</h1>

1. Clone the repository into a local folder.<br><br>
2. Inside quiz-api open a terminal and run main.go to start the API<br>
<code>go run main.go</code><br><br>
3. After starting the api head to <b>cli</b> folder and start the test with the following command<br>
<code>./cli quiz start</code>

<h1>EndPoints</h1>

1. GET '/quiz':<br>
Retrieves the quiz questions and choices

2. POST '/quiz':<br>
Posts the user answers creating a record for every user.
