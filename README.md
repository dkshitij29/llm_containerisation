# LLM Containerisation

This project demonstrates how to containerize an application that leverages a Large Language Model (LLM), specifically LLaMA, to process user inputs from a frontend and serve responses via a backend API.

## Project Structure

- **Backend**: Built with Go, the backend handles API requests from the frontend, processes input, and interacts with the LLaMA model to generate responses.
  - `server.go`: Main Go server code that listens for requests and forwards them to the model.
  - `go.mod`: Go module dependencies.
  
- **Frontend**: A simple web interface where users can input text and receive responses from the LLaMA model.
  - `index.html`: The HTML structure for the web interface.
  - `script.js`: JavaScript that interacts with the backend, sending user input and displaying the responses.
  
- **Model**: Contains the Python code responsible for interacting with the LLaMA model.
  - `app.py`: The main application file that serves the model and handles requests from the backend.
  - `model.py`: Code that loads and processes the LLaMA model for generating text responses.

## Installation

### Backend

1. Navigate to the `Backend/` directory:
   ```bash
   cd Backend
   ```
2. Install the necessary Go dependencies:
   ```bash
   go mod tidy
   ```
3. Run the Go server:
   ```bash
   go run server.go
   ```

### Frontend

1. Navigate to the `Frontend/` directory:
   ```bash
   cd Frontend
   ```
2. Open `index.html` in a browser to start interacting with the application.

### Model

1. Navigate to the `Model/` directory:
   ```bash
   cd Model
   ```
2. Install the necessary Python dependencies (ensure you have `virtualenv` installed):
   ```bash
   virtualenv venv
   source venv/bin/activate
   pip install -r requirements.txt
   ```
3. Run the model server:
   ```bash
   python app.py
   ```

## Usage

1. Start the Go backend and Python model server as explained in the Installation section.
2. Open the frontend (`index.html`) in a web browser.
3. Enter text in the input field. The backend will send the input to the model, and the model's response will be displayed on the frontend.
