# Chat Application

A real-time chat application built with Go, featuring WebSocket communication, user authentication, chat room management, direct messaging, and message storage in a file. The app is containerized using Docker and can be deployed easily.

---

## **Features**

1. **Real-Time Communication**:
   - Supports WebSocket communication between clients and the server for real-time messaging.

2. **User Authentication**:
   - Register and authenticate users via username and password.

3. **Chat Rooms**:
   - Users can create, join, and leave chat rooms.
   - Broadcast messages to all users in a room.

4. **Direct Messaging (DM)**:
   - Send private messages to specific users.

5. **Message Storage**:
   - Stores chat messages in a file (`messages.log`) for persistence.

6. **Dockerized Deployment**:
   - Deploy the application easily using Docker and Docker-Compose.

---

## **Prerequisites**

- [Docker](https://www.docker.com/) installed on your machine.
- Optional: [Postman](https://www.postman.com/) or `wscat` for testing WebSocket communication.

---

## **How to Run**

### **Step 1: Clone the Repository**
Clone this repository to your local machine:
```bash
git clone <repository-url>
cd <repository-folder>
```

### **Step 2: Build and Run with Docker**

1. **Build the Docker Image**:
   ```bash
   docker-compose build
   ```
2. **Run the Application:**:
   ```bash
   docker-compose up
   ```
   The application will be accessible at http://localhost:8080.
3. **Stop the Application:**:
   ```bash
   docker-compose down
   ```

---

## **Endpoints**

### **1. WebSocket Endpoint**
- **URL**: `ws://localhost:8080/ws`
- **Purpose**: Real-time communication.

#### **WebSocket Commands**
- **Authenticate**:
  ```json
  {"username": "<username>", "password": "<password>"}
  ```
- **Join a Room:**:
  ```json
  {"from": "<username>", "room": "<room-name>", "action": "join"}
  ```
- **Send a Message to a Room:**:
  ```json
  {"from": "<username>", "room": "<room-name>", "content": "<message>", "action": "message"}
  ```
- **Send a Direct Message:**:
  ```json
  {"from": "<username>", "to": "<recipient>", "content": "<message>", "action": "message"}
  ```
- **Leave a Room:**:
  ```json
  {"from": "<username>", "room": "<room-name>", "action": "leave"}
  ```

### **2. User Registration Endpoint**
- **URL**: `http://localhost:8080/register`
- **Method**: `POST`
- **Request Body**:
  ```json
  {"username": "<username>", "password": "<password>"}
  ```
- **Response**: `201 Created` if successful.

### **3. Chat History Endpoint**
- **URL**: `http://localhost:8080/history?room=<room-name>`
- **Method**: `GET`
- **Purpose**: Retrieve all messages from a specific room.
- **Response**:
  ```json
  [
    {"from": "user1", "room": "general", "content": "Hello!", "action": "message"},
    {"from": "user2", "room": "general", "content": "Hi!", "action": "message"}
  ]
  ```

---

## **Testing the Application**

### **Using WebSocket Client (`wscat`)**
1. Install `wscat`:
   ```bash
   npm install -g wscat
   ```
2. Connect to the WebSocket endpoint:
   ```bash
   wscat -c ws://localhost:8080/ws
   ```
3. Test WebSocket commands (e.g., authenticate, send messages).

### **Using Postman**
- Use the Postman WebSocket client for real-time communication.
- Use Postman for HTTP endpoints (`/register`, `/history`).

---

## **Project Structure**

    ```bash
    chat_app/
    ├── main.go        # Entry point of the application (server setup and handlers)
    ├── db.go          # File-based storage (saveMessageToFile, getMessagesFromFile)
    ├── client.go      # Client struct and related operations
    ├── message.go     # Message processing logic (handleMessages)
    ├── websocket.go   # WebSocket connection handling logic (handleConnections)
    ├── auth.go        # User registration and authentication
    ├── messages.log   # File to store chat messages (auto-created at runtime)
    ├── Dockerfile     # Dockerfile for containerizing the application
    ├── docker-compose.yml # Docker Compose configuration
    ```
