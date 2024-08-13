package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func callOllamaAPI(input string) (string, error) {
	// Replace with the actual IP address of your WSL environment and the port number your Flask app is running on
	//url := "http:///172.18.69.141:5000/generate"
	url := "http://172.18.69.141:5000/generate"  // Correct the URL here
	// Create JSON data to send in the request body
	jsonData := map[string]string{"text": input}
	jsonValue, _ := json.Marshal(jsonData)

	// Make the POST request to the Flask API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the JSON response
	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result["response"], nil
}

func main() {
	// Example input text
	inputText := "Hello, Llama!"

	// Call the Ollama API and get the response
	response, err := callOllamaAPI(inputText)
	if err != nil {
		log.Fatalf("Error calling Ollama API: %v", err)
	}

	// Print the response from the Llama model
	fmt.Println("Response from Llama:", response)
}





// working code
/*package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
)

var tmpl = template.Must(template.ParseFiles("../Frontend/index.html"))

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/submit", submitHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Frontend"))))

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		textInput := r.FormValue("textInput")
		response := Response{Message: fmt.Sprintf("You entered: %s", textInput)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
*/


/*

To integrate Ollama (a hypothetical or specific AI service) into your Go web application, you'll need to:

1. Capture the user input from the form.
2. Send this input to the Ollama API.
3. Process the response from Ollama and return it to the user.

Here's a step-by-step approach to achieve this:

### 1. Capture User Input
You've already captured user input in the `submitHandler`. We will use this input to send a request to Ollama.

### 2. Send Input to Ollama API
Assuming Ollama provides a REST API, you can use Go's `net/http` package to send the input and receive a response.

### 3. Process the Response
Once you get the response from Ollama, process it and send it back to the client.

Here’s how you can modify your `submitHandler` to integrate with Ollama:

### Updated `submitHandler` to Call Ollama API
```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"path/filepath"
)

var tmpl = template.Must(template.ParseFiles(filepath.Join("Frontend", "index.html")))

type Response struct {
	Message string `json:"message"`
}

type OllamaRequest struct {
	Text string `json:"text"`
}

type OllamaResponse struct {
	ProcessedText string `json:"processed_text"`
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/submit", submitHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Frontend"))))

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		textInput := r.FormValue("textInput")

		// Create request to Ollama API
		ollamaRequest := OllamaRequest{Text: textInput}
		requestBody, err := json.Marshal(ollamaRequest)
		if err != nil {
			http.Error(w, "Unable to create request", http.StatusInternalServerError)
			return
		}

		// Send request to Ollama API
		ollamaResponse, err := http.Post("https://api.ollama.ai/process", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			http.Error(w, "Failed to send request to Ollama", http.StatusInternalServerError)
			return
		}
		defer ollamaResponse.Body.Close()

		// Read response from Ollama
		responseBody, err := ioutil.ReadAll(ollamaResponse.Body)
		if err != nil {
			http.Error(w, "Failed to read response from Ollama", http.StatusInternalServerError)
			return
		}

		var response OllamaResponse
		err = json.Unmarshal(responseBody, &response)
		if err != nil {
			http.Error(w, "Failed to parse response from Ollama", http.StatusInternalServerError)
			return
		}

		// Send processed text back to client
		responseMessage := Response{Message: response.ProcessedText}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseMessage)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
```

### JavaScript to Handle the Response
Make sure your `script.js` can handle the JSON response and update the HTML accordingly.

### `script.js`
```javascript
document.getElementById('textForm').addEventListener('submit', async function (event) {
    event.preventDefault();

    const textInput = document.getElementById('textInput').value;
    const response = await fetch('/submit', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({
            'textInput': textInput,
        }),
    });

    if (response.ok) {
        const jsonResponse = await response.json();
        document.getElementById('output').innerText = jsonResponse.message;
    } else {
        document.getElementById('output').innerText = 'Error: Unable to process your request';
    }
});
```

### Explanation
1. **Submit Handler**:
   - Captures the input from the form.
   - Creates a JSON payload and sends a POST request to Ollama's API.
   - Reads and processes the response from Ollama.
   - Sends the processed text back to the client as a JSON response.

2. **JavaScript**:
   - Prevents the default form submission.
   - Sends an AJAX request to the `/submit` endpoint.
   - Updates the page with the response from the server.

Ensure that the URL for Ollama's API (`https://api.ollama.ai/process` in the example) is correct and that the API request and response formats match those expected by Ollama.

With these updates, your web application should now be able to capture user input, send it to Ollama for processing, and display the result to the user.

*/