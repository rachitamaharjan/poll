# Polling Application

This repository contains the code for a real-time polling application that allows users to create polls, cast votes, and view poll results. The app is built using **Golang (backend)** and **React (frontend)**, with **PostgreSQL** as the database. The application implements a range of advanced features, including concurrency, Real-time updates are powered by **Server-Sent Events (SSE)**. The backend also ensures thread safety using **mutex** for locking.

## Features

### 1. **Create Polls**
- Users can dynamically create new polls through the app's interface.
- Polls can be customized with single or multiple-choice questions.
- Admins can configure whether multiple votes per user are allowed using a checkbox.

### 2. **Vote on Polls**
- Polls can be voted on in real time, allowing multiple users to cast their votes simultaneously.
- Each user is restricted to one vote per request.

### 3. **View Poll Results**
- Poll results are displayed using **Chart.js** in a visually appealing pie chart format.
- The poll results page updates dynamically with real-time votes using **Server-Sent Events (SSE)**.
- Total votes and vote distributions are clearly shown.

### 4. **Real-time Updates**
- Real-time updates for poll results using **Server-Sent Events (SSE)**.

### 5. **Security**
- SQL injection prevention and Cross-Site Request Forgery (CSRF) protection are implemented to secure user data.
- Passwords and other sensitive information are encrypted using industry-standard cryptography techniques.

### 6. **Internationalization**
- The app supports multiple languages, with English and Nepali available out of the box.
- A toggle button allows users to switch between languages seamlessly.

### 7. **Containerization**
- Docker is used to containerize the application, ensuring smooth deployment and scalability.

### 8. **Concurrency and Parallelism**
- The backend makes use of Go’s concurrency primitives, including Goroutines and channels, to manage multiple vote submissions concurrently.
- Mutexes are used to handle race conditions and ensure data consistency during vote casting.


## Technologies Used
### Backend
- **Golang** with **Gin** framework for routing.
- **PostgreSQL** as the database.
- **GORM** for ORM support.
- **SSE (Server-Sent Events)** for real-time communication.
- **Mutex** for safe concurrent access.
- **logrus** for logging.

### Frontend
- **React** for the user interface.
- **Chart.js** for visualizing poll results.
- **React Context API** for state management.
- **React Hooks** for managing side effects and custom hooks.
- **React Router** for client-side routing.

### Other Libraries/Tools
- **WebSockets** for real-time communication.
- **logrus** for structured logging.
- **Docker** for containerization.
- **i18next** for internationalization and localization support.
- **UUID** for unique poll and user identification.

## Setup Instructions

### Prerequisites
- **Go** version 1.19 or higher
- **Node.js** and **npm** for frontend development
- **PostgreSQL** database
- **Docker** (optional, for containerized setup)

### Steps to Run Locally

1. **Clone the Repository**
   ```
   git clone https://github.com/rachitamaharjan/poll.git
   cd poll

### Backend Setup
1. Navigate to the backend directory:
   ```cd app```
2. Set up the database and update environment variables:
   ```export DB_HOST=localhost
   export DB_USER=your_db_user
   export DB_PASSWORD=your_db_password
   export DB_NAME=poll_db
3. Run database migrations:
   ```go run main.go migrate```
4. Start the backend server:
   ```go run main.go```

### Frontend Setup
1. Navigate to the frontend directory:
   ```cd client```
2. Install dependencies:
   ```npm install```  
3. Start the React development server:
   ```npm start```

### Database Configuration
Ensure your PostgreSQL is configured with the appropriate database connection string. You can set this in the .env file:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=poll
```
## Backend Details

The backend is built using **Gin**, and the application handles the following:
- Poll creation and management.
- User voting and real-time update of results using **SSE**.
- Mutex locks for concurrency safety during voting operations.
- Database management using **PostgreSQL** and **GORM** for ORM.

### Backend Structure
- `routes/`: Contains the API routes for the polling application.
- `controllers/`: Handles business logic for poll creation, voting, and retrieving poll results.
- `models/`: Contains the database models for Poll, Vote, etc.
- `services/`: Contains service functions for the polling application.

## Frontend Details

The frontend is built with **React**, using the following key features:
- **React Context API**: Manages global state for polls and votes.
- **Chart.js**: Displays poll results in real-time with interactive pie charts.
- **React Router**: Handles routing for pages such as poll creation, voting, and results.

### Frontend Structure
- `src/components/`: Contains the React components for the application (PollList, PollDetail, CreatePoll).
- `src/hooks/`: Custom hooks for managing poll data (`usePollData`).
- `src/context/`: Global context for managing application state.
- `src/locales/`: Global context for managing application state.

## Database Schema

- `polls`: Stores poll information (id, question, options, etc.).
- `votes`: Stores user votes, associated with poll options.
- `users`: Stores information about users and vote tracking.

## API Endpoints

### Polls
- `GET /polls`: Get all active polls.
- `POST /polls`: Create a new poll.

### Votes
- `POST /polls/:id/vote`: Cast a vote on a poll.

### SSE
- `GET /polls/:id/stream`: SSE endpoint for real-time poll updates.




