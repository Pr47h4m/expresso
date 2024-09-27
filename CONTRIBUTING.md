# Contributing to Expresso

Thank you for considering contributing to Expresso! We welcome all kinds of contributions including bug reports, feature suggestions, and code improvements. Please follow the guidelines below to ensure a smooth contribution process.

## How to Contribute

### 1. Reporting Bugs
If you find a bug or an issue in the framework, please help us by submitting an issue to our [GitHub Issue Tracker](https://github.com/Pr47h4m/expresso/issues). Be sure to include:
- A clear and descriptive title.
- Steps to reproduce the problem.
- Expected and actual results.
- Any error logs or screenshots.

### 2. Suggesting Features or Enhancements
We also welcome ideas for new features and improvements. If you have suggestions, please open an issue labeled as **Feature Request** and include:
- A detailed description of the feature.
- Why this feature is important or how it improves the framework.
- Any relevant examples or references.

### 3. Contributing Code
To contribute code, please follow these steps:

#### 3.1 Fork the Repository
- Fork [Expresso](https://github.com/Pr47h4m/expresso) to your GitHub account.
- Clone the repository locally using:
  ```bash
  git clone https://github.com/<your-username>/expresso.git
  cd expresso
  ```

#### 3.2 Set Up Your Environment
- Ensure you have Go installed on your system.
- Install dependencies using:
  ```bash
  go mod download
  ```

#### 3.3 Create a Feature Branch
- Create a new branch for your changes:
  ```bash
  git checkout -b feature/your-feature-name
  ```

#### 3.4 Write and Test Your Code
- Write clear, concise, and well-documented code.
- Ensure that you write unit and/or integration tests for your changes.
- Run the test suite locally to confirm everything works:
  ```bash
  go test ./...
  ```

#### 3.5 Commit Your Changes
- Make sure your commit messages are descriptive.
- Format your Go code using:
  ```bash
  go fmt ./...
  ```

#### 3.6 Push and Create a Pull Request
- Push your changes to your fork:
  ```bash
  git push origin feature/your-feature-name
  ```
- Go to the main repository and submit a pull request. Ensure your pull request:
  - Links to any relevant issues.
  - Provides a clear summary of the changes.
  - Describes any tests you've added or modified.

### 4. Code Review Process
Once you submit your pull request:
- A maintainer will review your code and may request changes or improvements.
- After approval, your code will be merged into the main branch.

### 5. Setting Up GitHub Actions
- Ensure that any code you contribute triggers relevant GitHub Actions workflows, such as automated tests or linters.
- Follow the guidelines for setting up CI workflows.

### 6. Writing Tests
We aim for high test coverage. Please include appropriate unit and integration tests in your contributions. You can run the test suite locally as described earlier, and it will also run automatically via GitHub Actions.

### 7. Style Guide
- Follow Go's standard formatting (`gofmt`).
- Use clear variable and function names.
- Ensure that the code is well-documented and comments are meaningful.
