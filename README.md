# Guess-It 1

A competitive number guessing AI that predicts ranges for sequence continuation.

## 🚀 Student Solution

The solution in `student/solution.go` implements an **adaptive statistical approach** designed to handle diverse data distributions (from small ranges like [100,200] to massive ranges like [-1B, +1B]).

### Algorithm Strategy

1.  **Sliding Window**: Uses a window of the last **50** values to adapt quickly to data trends and forget old outliers.
2.  **Trimmed Percentiles**: Removes the top and bottom 10% of data in the window to filter out extreme outliers.
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

## 🛠️ How to Run

### 1. Start the Tester

Open a terminal in the project folder and run:

```bash
docker compose up --build
```

### 2. Open the Interface

Go to [http://localhost:3000](http://localhost:3000) in your browser.

### 3. Choose an AI Opponent

You must add a `?guesser=NAME` parameter to the URL to select which AI opponent to compete against.

**Click the links below to test directly:**

| AI Opponent | Difficulty | Link |
| :--- | :--- | :--- |
| **Big Range** | Standard | [Test vs Big Range](http://localhost:3000/?guesser=big-range) |
| **Average** | Basic | [Test vs Average](http://localhost:3000/?guesser=average) |
| **Median** | Solid | [Test vs Median](http://localhost:3000/?guesser=median) |
| **Nic** | Bonus | [Test vs Nic](http://localhost:3000/?guesser=nic) |
| **Linear Regr** | Advanced | [Test vs Linear Regression](http://localhost:3000/?guesser=linear-regr) |
| **MSE** | Metric | [Test vs MSE](http://localhost:3000/?guesser=mse) |

### 4. Run the Test

1.  Click **"Test Data 1"** (or 2, 3...)
2.  Click **"Quick"** to skip the animation and see the final result.
3.  Click **"Clean"** before running another test.

---

## 📂 Project Structure

```console
guess-it/
├── ai/                 # AI opponents (binaries)
├── data_sets/          # Test data (1-9)
├── student/
│   ├── solution.go     # The logic (THIS IS THE SOLUTION)
│   └── script.sh       # Script to run the solution
├── index.html          # Web interface
├── server.js           # Backend tester
├── Dockerfile          # Environment setup
└── README.md           # This file
```
