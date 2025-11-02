# Car Rental Service

This is a car rental service application with a web-based client and a Go backend.

## Project Structure

- `Client/`: Contains the frontend application (HTML, CSS, JavaScript) and an Nginx configuration.
- `server/`: Contains the backend Go application.
- `docker-compose.yml`: Defines the services, networks, and volumes for the application.
- `.github/workflows/ci.yml`: Contains the CI pipeline configuration for GitHub Actions.

## Environment Configuration

Before running the application, you need to create a `.env` file in the root directory. You can use the `.env.example` file as a template:

```bash
cp .env.example .env
```

Then, you can modify the `.env` file with your database credentials.

## Running the Application

To run the application, you need to have Docker and Docker Compose installed.

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   ```

2. **Navigate to the project directory:**

   ```bash
   cd car_rental_service
   ```

3. **Create the `.env` file:**

    ```bash
    cp .env.example .env
    ```
    
    Update the `.env` file with your database credentials.

4. **Run the application using Docker Compose:**

   ```bash
   docker-compose up -d
   ```

   This will build the Docker images for the client and server and start the containers in detached mode.

5. **Access the application:**

   - The client is available at [http://localhost](http://localhost).
   - The server API is available at [http://localhost:8080](http://localhost:8080).
   - The admin panel is available at [http://localhost:8080/admin](http://localhost:8080/admin).

## Running Tests

### Server

To run the server-side tests, navigate to the `server` directory and run the following command:

```bash
go test ./...
```

### Client

To run the client-side tests, navigate to the `Client` directory and run the following command:

```bash
npm install
npm test
```

## Services

The application consists of the following services:

- `client`: An Nginx container that serves the frontend application.
- `server`: A Go container that provides the backend API.
- `db`: A MySQL container for the database.

## CI/CD

A CI pipeline is set up using GitHub Actions. The pipeline is triggered on every push to the `main` branch and performs the following steps:

- Checks out the code.
- Sets up Go and Docker.
- Builds the Docker images for the client and server.
- Runs the server-side and client-side tests.

## Database

The application uses a MySQL database to store data. The database is automatically created when you run `docker-compose up`. The database credentials are configured in the `docker-compose.yml` file and the `.env` file.