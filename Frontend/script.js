document.getElementById('textForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const textInput = document.getElementById('textInput').value;

    fetch('/submit', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({
            textInput: textInput
        })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('output').textContent = data.message;
    })
    .catch(error => {
        console.error('Error:', error);
    });
});
