# Expenses Tracker

This is a simple full-stack application I built following [this Vue.js tutorial](https://www.youtube.com/watch?v=hNPwdOZ3qFU), to which I added a backend server developed in Golang. The app serves as a simple expense tracker, allowing users to record both their spending and income entries.

## Technologies Used

- **Vue.js** Used to build the frontend and create a smooth, interactive user experience.

- **CSS** Handles the app’s look and layout, keeping it simple, clean, and responsive.

- **PostgreSQL** Stores all the data safely and efficiently, making it easy to manage expenses and income records.

- **Gin** A fast and lightweight Go framework that powers the backend and handles API requests.

- **Go SQL package** Used to connect the backend with the PostgreSQL database and manage queries.


## Features

- Create and delete transactions  
- Calculate total income and expenses  
- Calculate the overall balance  

## Project Structure

- **/client** Vue.sj frontend.  
- **/server** Golang backend that also contains the DB schema.  

## Project Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/expense-tracker.git

# Navigate to the project folder
cd expense-tracker

# --- Frontend setup ---
cd client

# Install dependencies
npm install

# Run the frontend (default: http://localhost:5173)
npm run dev

# --- Backend setup ---
cd ../server

# Download and clean Go dependencies
go mod download
go mod tidy

# Run the backend (default: http://localhost:8080)
go run .

```
At this point, the frontend should be running on port 5173, and the backend on port 8080.
You can interact with the API using Postman or cURL, or simply use the web client to test all features.

## Future improvements

- Add user authentication
- Add endpoints to compute the totals in the server. 
