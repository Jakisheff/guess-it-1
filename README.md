# Guess-It-1

A competitive number guessing AI that predicts ranges for sequence continuation.

## Student Solution

The solution in `student/solution.go` implements an **adaptive statistical approach** designed to handle diverse data distributions (from small ranges like [100,200] to massive ranges like [-1B, +1B]).

### Algorithm Stratgey

1.  **Sliding Window**: Uses a window of the last **50** values to adapt quickly to changing data trends and forget old outliers.
2.  **Trimmed Percentiles**: Removes the top and bottom 10% of data in the window to filter out extreme outliers before calculating spread.
3.  **Adaptive Range**: Calculates the range using **Median ± 1.5 × IQR** (Interquartile Range) of the trimmed data.
4.  **Safety Floor**: Ensures the predicted range is never too narrow (minimum 5% of observed spread).

### Performance

The solution achieves consistent high scores across all test datasets:

| Dataset | Score | Correct % |
| :--- | :--- | :--- |
| **Data 1** (Range 100-200) | **~83,000** | 96% |
| **Data 2** (Range -1B to +1B) | **~83,000** | 96% |
| **Data 3** (Range 100-1M) | **~83,000** | 96% |

---

## Usage

You will first need copy the `student/` folder (provided by the student) in
which you will see the student guessing program along with a file called
`script.sh`. This file should be an executable shell script that runs the
student program if you are in the root folder `guess-it/`. The filesystem
should look somethings like this:

```console
─ guess-it/
├── ai/
│   ├── big-range
│   └── ...
├── index.html
├── index.js
└── ...
└── student/
    ├── ...
    └── script.sh

```

To test the student program, these commands should be ran to have the
dependencies needed and to start the webpage on the port 3000:

```console
docker compose up --build
```

After opening your browser of preference in the port
[3000](http://localhost:3000/), if you try clicking on any of the `Test Data`
buttons, you will notice that in the Dev Tool/ Console there is a message which
tells you that you need another guesser besides the student.

Adding a guesser is simple. You need to add in the URL a guesser, in other
words, the name of one of the files present in the `ai/` folder:

```console
?guesser=<name_of_guesser>
```

For example:

```console
?guesser=big-range
```

After that, choose which of the random data set to test. After that you can
wait for the program to test all of the values (boooooring), or you can click
`Quick` to skip the waiting and be presented with the results.

Since the website uses big data sets, we advise you to clear the displays
clicking on the `Clean` button after each test.
