# Rest API Pod

## Overview

The `rest-api-pod` project provides a set of APIs for user authentication and event management. The server-side code includes endpoints for user signup, login, and event operations. The client folder contains HTTP files that can be easily executed using the REST API extension in Visual Studio Code.

## Available APIs

### User Authentication

- **Signup a User**
  - **Endpoint:** `POST /signup`
  - **Description:** Create a new user account.

- **Login a User**
  - **Endpoint:** `POST /login`
  - **Description:** Authenticate a user and obtain a token.

### Event Management

- **Post Event**
  - **Endpoint:** `POST /events`
  - **Description:** Create a new event. **(Authentication required)**

- **Update Event**
  - **Endpoint:** `PUT /events/<id>`
  - **Description:** Update an existing event. **(Authentication and Authorization required)**

- **Get Events**
  - **Endpoint:** `GET /events`
  - **Description:** Retrieve a list of all events.

- **Get Single Event**
  - **Endpoint:** `GET /events/<id>`
  - **Description:** Retrieve details of a specific event.

- **Delete Event**
  - **Endpoint:** `DELETE /events/<id>`
  - **Description:** Delete a specific event. **(Authentication and Authorization required)**

- **Register for an Event**
  - **Endpoint:** `POST /events/<id>/register`
  - **Description:** Register a user for a specific event. **(Authentication required)**

- **Cancel a Registration**
  - **Endpoint:** `DELETE /events/<id>`
  - **Description:** Cancel a user's registration for an event. **(Authentication and Authorization required)**

## Instructions to Run the Server

1. Clone the repository:
   ```bash
   git clone https://github.com/ssahoo11/rest-api-pod.git
   ```

2. Navigate to the project directory:
   ```bash
   cd rest-api-pod
   ```

3. Start the server:
   ```bash
   go run .
   ```

4. The server will start on `localhost` at port `8080`.

## Instructions to Run REST API Server on Kubernetes Pod

This guide explains how to deploy the REST API server from this repository to a Kubernetes pod.

### Prerequisites

- **Docker**: To build the Docker image.
- **Kubernetes**: A Kubernetes cluster (you can use Minikube for local development).
- **kubectl**: Command-line tool for interacting with Kubernetes.

### Steps to Deploy on Kubernetes Pod

Build the Docker image for the REST API server:

```bash
docker build -t rest-api-pod:latest .
```
Create Kubernetes Deployment
```bash
kubectl apply -f deployment.yaml
```
Expose the Pod as a Service
```bash
kubectl apply -f service.yaml
```
Access the REST API
```bash
minikube service rest-api-pod --url
```
Or
```bash
kubectl get svc rest-api-service
```


## Client Folder

In the client folder, you will find all the HTTP files required to interact with the APIs. You can run these files using the REST API extension in Visual Studio Code.

Or using curl tool from bash
```bash
curl -get https://<url>:8080/events
```

## Contributing

Feel free to open issues or submit pull requests for any improvements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

Feel free to modify any sections to better fit your project or add additional information as needed!
