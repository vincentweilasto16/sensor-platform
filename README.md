# 🛰️ Sensor Platform Backend


This project is a backend system for processing data streams from multiple sensors and human inputs.  
The system is built using **Go (Echo framework)**, **MySQL**, and follows a **microservices architecture**.  
It is designed to generate, receive, process, and store sensor data, while exposing REST APIs for analytics and management.

---

## 📋 Table of Contents


1. [Project Scopes](#project-scopes)  
2. [Assumptions](#assumptions)  
3. [Process Flow](#process-flow)  
4. [Tech Stack](#tech-stack)  
5. [Prerequisites](#prerequisites)
6. [Software Architecture](#software-architecture)
7. [Database Entity Relationship Diagram (ERD)](#database-entity-relationship-diagram)  
8. [Setup & Run](#setup--run)  
9. [API Contract](#api-contract)
10. [Postman Collection](#postman-collection)
11. [Testing Application](#testing-application)  
12. [Limitations & Future Improvements](#limitations--future-improvements)

---

## 📍 Project Scopes 


The project will deliver the following:  

1. A microservice (**Microservice A**) that simulates sensors by generating data records with the following fields:  
   - Sensor value (float)  
   - Sensor type (string)  
   - Device code (renamed from ID1, alphabets in capital letters)  
   - Device number (renamed from ID2, integer)  
   - Timestamp  

2. The frequency of data generation can be configured through a REST API endpoint.  

3. Multiple instances of **Microservice A** can run at the same time, with each instance producing data for only one fixed sensor type (e.g., temperature, humidity, pressure).  

4. Another microservice (**Microservice B**) will be created to receive and process the data stream from Microservice A using gRPC, MQTT, or an equivalent protocol.  

5. **Microservice B** will compile and store the received data in a MySQL database in a structured format, making it easy to manage and retrieve.  

6. **Microservice B** will provide REST API endpoints to:  
   - Retrieve data based on device code and device number.  
   - Retrieve data based on a time range.  
   - Retrieve data based on a combination of device identifiers and timestamps.  
   - Edit and delete data using the same filtering options.  
   - Support pagination when retrieving large sets of data.  

7. Authentication and authorization will be implemented to ensure secure access to the API endpoints.  

8. **Microservice B** will be designed to handle input from multiple Microservice A instances simultaneously and support scalability as the number of devices grows.  

9. The system will be tested by running many instances of **Microservice A** to push the scalability limits of **Microservice B**.  

---

## ⚡ Assumptions  


- Sensors are simulated and not connected to real hardware.  
- Device code will be randomly generated from A–Z.  
- Device number from sensor data will be randomly generated in the range of 1–1000 and cannot be zero.  
- Sensor data generation starts automatically when the service runs and continues until shutdown, with APIs available to manually start or stop the process.  
- Sensor value will be randomly generated based on the sensor type and within the possible range for that sensor type.  
- Sensor data is produced at regular intervals (default: 5 seconds) and generated from multiple instances.  
- The regular interval for producing sensor data cannot be less than 5 seconds.  
- A login process is required before accessing features such as viewing, editing, or deleting sensor data.  
- The system defines two roles: **admin** and **user**. Admins can view, edit, and delete data, while users are limited to viewing and editing only. Since deletion carries higher risk, it is restricted to admins.  
- A registration feature is available to create new users by providing a username, password, and role for simplicity the input. Meantime, registration must be completed before login.  
- Sensor data can be retrieved through a single API that supports filtering (by device code, device number, start time, and end time) and pagination (page and limit).  
- Sensor data can be deleted or updated using filtering, with restrictions:  
  - Deletion or update cannot occur if the start time is later than the end time.  
  - Deletion or update cannot be performed without a filter.  
- For the update feature, only specific fields can be modified: **sensor type**, **sensor value**, and **timestamp**. 


---


## 🧑‍💻 Process Flow

### 📊 Sensor Data Collection Process

![Sensor Data Collection Process Flow](docs/flowchart/sensor_data_collection_process_flow.svg)

This diagram shows how sensor system automatically collects and stores data from multiple types of sensors. Think of it like a smart monitoring system that never sleeps and continuously tracks environmental conditions.

**How It Works**

🌡️ Smart Sensors Generate Data
- Multiple sensor devices work independently (temperature, humidity, pressure sensors)
- Each sensor takes readings automatically every few seconds
- The timing can be adjusted - you can make sensors read faster or slower as needed
- Each sensor creates a data record with the measurement, device information, and timestamp

📬 Message Delivery System
- All sensor data goes through a message queue system (like a smart postal service)
- This system ensures no data gets lost, even when many sensors send data at the same time
- It organizes and delivers data efficiently to the processing center

🏭 Data Processing Center
- Multiple data processors work together to handle incoming sensor information
- They validate the data to make sure it's correct and complete
- The system can handle many sensors sending data simultaneously
- Processes data in real-time as it arrives

🗄️ Secure Data Storage
- All validated sensor data is stored safely in a database
- Data is organized so it can be easily searched and retrieved later
- The storage system keeps track of when each reading was taken
- Data remains accessible for analysis and reporting


### 🔐 User Access and Data Management Process

![User Access And Data Management Process Flow](docs/flowchart/user_access_and_data_management_process_flow.svg)

This diagram explains how people interact with the sensor data - from creating accounts to viewing and managing the collected information. It's designed to be secure while remaining easy to use.

**User Journey**

👤 Account Creation
- Users start by creating their account with a username and password
- They choose their access level (Administrator or Regular User)
- The system automatically encrypts passwords for security
- Simple signup process focused on getting access quickly

🔐 Secure Sign-In
- Users log in with their username and password
- The system verifies their identity against secure stored information
- After successful login, users receive a secure access key (like a digital badge)
- This access key allows them to use the system features safely

🛡️ Permission-Based Access
- The system checks what each user is allowed to do based on their role
- Administrators have full access to all features
- Regular Users have limited access for data protection
- Every action is checked against user permissions automatically


**Available Features**

📊 View Sensor Data (Available to: Admin ✓ | User ✓)
- Browse through all collected sensor readings
- Search Options: Find data by specific device code, device number, and date ranges
- Page Navigation: View large amounts of data in manageable chunks
- Smart Filtering: Combine multiple search criteria to find exactly what you need

✏️ Edit Sensor Data (Available to: Admin ✓ | User ✓)
- Modify existing sensor readings when corrections are needed
- Find and Update: Locate specific records and change their values
- Bulk Changes: Update multiple records that match your search criteria
- Safe Editing: Changes are tracked and can be audited

🗑️ Remove Sensor Data (Available to: Admin ✓ | User ✗)
- Delete sensor readings that are no longer needed
- Admin-Only Feature: Only administrators can remove data for security
- Smart Deletion: Remove data based on the same filtering system used for viewing
- Safe Removal: Data is marked as deleted but can potentially be recovered


**Security Features**

🔒 Password Protection
- All passwords are encrypted using industry-standard security methods
- Users receive secure access tokens that expire automatically
- Every request to view or change data requires proper authorization

👥 Role-Based Access
- Administrator Role: Complete control over all sensor data and system features
- User Role: Can view and edit data but cannot delete anything
- Automatic Enforcement: The system automatically prevents unauthorized actions
- Clear Permissions: Users always know what they can and cannot do


---


## 🛠 Tech Stack


### Core
- [Go](https://golang.org/) `v1.23.0` – Main programming language
- [Echo](https://echo.labstack.com/) – Web framework for building REST APIs

### Database & Migration
- [MySQL](https://www.mysql.com/) `v8.0.43` – Relational database
- [golang-migrate](https://github.com/golang-migrate/migrate) – Database migrations
- [sqlc](https://sqlc.dev/) – Generate type-safe Go code from SQL

### Messaging
- [Apache Kafka](https://kafka.apache.org/) – Message broker for sensor data
- [Zookeeper](https://zookeeper.apache.org/) – Coordination service for Kafka

### Authentication
- [JWT](https://jwt.io/) – Token-based authentication

### Containerization & DevOps
- [Docker](https://www.docker.com/) & Docker Compose – Containerization and orchestration

### Developer Tools
- [Postman](https://www.postman.com/) – API testing
- [DBeaver](https://dbeaver.io/) – Database management
- [Visual Studio Code](https://code.visualstudio.com/) – Code editor

---


## 📦 Prerequisites  


Before running the application, make sure you have the following installed:

- [Go](https://go.dev/) >= 1.23.0  
- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)  
  (used to run MySQL, Kafka, and Zookeeper — no local installation required)  
- [golang-migrate CLI](https://github.com/golang-migrate/migrate) (for database migrations)  
- [sqlc CLI](https://sqlc.dev/) (for generating type-safe Go code from SQL)  
- [Git](https://git-scm.com/) (for version control)  
- [Postman](https://www.postman.com/) (optional, for API testing)  
- [DBeaver](https://dbeaver.io/) (optional, for database management)  

---


## 🏗️ Software Architecture

![Software Architecture Schematic Diagram](docs/software_architecture/software_architecture_schematic_diagram.svg)


### Architecture Diagram Overview

This schematic illustrates a microservice-based IoT sensor data processing system designed for scalability and high performance. The architecture follows a clean separation of concerns with distinct layers for data generation, message streaming, processing, and storage.


**Key Data Flow**
- Data Generation → Multiple Microservice A instances generate sensor data (temperature, humidity, pressure) with configurable frequencies through REST APIs.

- Message Streaming → Apache Kafka handles high-throughput data streaming between producers and consumers, ensuring reliable message delivery.

- Data Processing → Scalable Microservice B cluster processes incoming Kafka streams, implements JWT authentication, and provides comprehensive REST APIs for CRUD operations.

- Storage → MySQL database stores sensor data with advanced filtering capabilities by ID combinations, time ranges, and pagination support.


**Current Implementation**
The system is containerized with Docker for local development and uses JWT authentication for secure API access. Components marked as "Not implemented yet" (API Gateway, Redis Cache, Kubernetes, Monitoring) represent planned enhancements for future development.


---


## 🗄 Database Entity Relationship Diagram (ERD)


The following Entity-Relationship Diagram (ERD) illustrates the database structure for sensor platform spesifically on **Microservice B**:

![Sensor ERD](docs/erd/sensor_erd.svg)

### Table Structure Overview

**sensor_data Table (Blue)**
- *Purpose*: Stores IoT sensor readings and measurements
- *Primary Key*: id (BIGINT) - uniquely identifies each sensor reading
- *Core Data*: Contains sensor type, measured values, and device identification
- *Key Fields*:
  - sensor_type: What kind of sensor (temperature, humidity, etc.)
  - sensor_value: The actual measurement reading
  - device_code & device_number: Identify which physical device sent the data
  - timestamp: When the reading was taken


| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | BIGINT | PRIMARY KEY, AUTO_INCREMENT | Unique identifier for each sensor reading |
| `sensor_type` | VARCHAR(50) | NOT NULL | Type of sensor (e.g., temperature, humidity, pressure) |
| `sensor_value` | DOUBLE | NOT NULL | The actual measurement value from the sensor |
| `device_code` | VARCHAR(20) | NOT NULL | Alphanumeric code identifying the device |
| `device_number` | INT | NOT NULL | Numeric identifier for the device |
| `timestamp` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | When the sensor reading was recorded |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP ON UPDATE | Last modification timestamp |
| `deleted_at` | TIMESTAMP | NULL | Soft delete timestamp (NULL = active record) |


**users Table (Green)**
- *Purpose*: Manages user authentication and access control
- *Primary Key*: id (INT) - uniquely identifies each user
- *Core Data*: User credentials and role-based permissions
- *Key Fields*:
  - username: Unique login identifier
  - password: Encrypted authentication credential
  - role: Access level (admin or regular user)

| Column | Data Type | Constraints | Description |
|--------|-----------|-------------|-------------|
| `id` | INT | PRIMARY KEY, AUTO_INCREMENT | Unique identifier for each user |
| `username` | VARCHAR(50) | NOT NULL, UNIQUE | User login identifier |
| `password` | VARCHAR(255) | NOT NULL | Encrypted password hash |
| `role` | ENUM('admin', 'user') | NOT NULL, DEFAULT 'user' | User access level |
| `created_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Account creation timestamp |
| `updated_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE | Last profile update timestamp |
| `deleted_at` | TIMESTAMP | NULL | Soft delete timestamp (NULL = active account) |


### Key Features

**🔐 Security & Access Control**
- **Role-based permissions**: Users can be either 'admin' or 'user'
- **Unique usernames**: Prevents duplicate accounts
- **Password encryption**: Stored as hashed values for security

**📈 Data Integrity**
- **Primary keys**: Each table has a unique identifier
- **NOT NULL constraints**: Critical fields cannot be empty
- **Auto-increment**: Automatic ID generation for new records

**🕒 Audit Trail**
Both tables implement comprehensive audit logging:
- **created_at**: Tracks when records are first added
- **updated_at**: Automatically updates when records are modified
- **deleted_at**: Implements soft delete (records marked as deleted but not physically removed)

**📡 IoT Data Management**
The sensor_data table is optimized for high-volume IoT data:
- **BIGINT primary key**: Supports billions of sensor readings
- **Device identification**: Dual identifiers (code + number) for flexible device management
- **Flexible sensor types**: VARCHAR field accommodates various sensor categories
- **Precise timestamps**: Ensures accurate temporal data tracking

**This ERD represents the current database structure. Future iterations may include additional tables and relationships to support enhanced functionality.**
---

## 🚀 Setup & Run

1. **Clone the repository**
   Clone the project and open it in your preferred IDE or code editor
   
   ```bash
   git clone https://github.com/vincentweilasto16/sensor-platform.git
   ```

2. **Install dependencies**
   Make sure Go is installed, then install the required dependencies:
   
   ```bash
   go mod tidy
   go mod vendor
   ```
   
3. **Set up environment variables**
   Before building or running the project, create a `.env` file from the example:
   
   ```bash
   cp .env.example .env
   ```
   
4. Start Docker Daemon
   Ensure Docker is running on your machine. The project includes a docker-compose.yml file that defines all required services:
   - Zookeeper
   - Kafka
   - MySQL
   - Microservice B (Data Receiver & API)
   - Microservice A instances (Temperature, Humidity, Pressure)

5. Prepare the database
   Before running all services, make sure MySQL is ready and the database exists:

   1. **Build and start MySQL**

      Start only the MySQL service using Docker Compose:
   
      ```bash
      docker compose up mysql
      ```

      > This will build and start the MySQL container using the credentials defined in your `.env` file (`DB_USER`, `DB_PASSWORD`, `DB_PORT`).

   2. **Connect to MySQL**
      
      Once the container is running, connect to MySQL using your preferred client (e.g., MySQL CLI, TablePlus, DBeaver) with the following details:
      
      - **Host:** `localhost`  
      - **Port:** `${DB_PORT}` (from `.env`)  
      - **Username:** `${DB_USER}`  
      - **Password:** `${DB_PASSWORD}`  

      > Ensure the connection is successful before proceeding.

   3. **Create the database**
      
      After a successful connection, create the database defined in your `.env` file (`DB_NAME`) if it doesn’t already exist:

      ```sql
      CREATE DATABASE IF NOT EXISTS your_db_name;
      ```

6. Start All Services
   Build and start all services:
   
   ```bash
   docker compose up --build
   ```
   
   - Database migration runs automatically after Microservice B starts.
   - Kafka topic (sensor-topic) is automatically created during the Docker build.
   - Sensor data generation starts automatically once all services are running.


---

## 📜 API Contract

### Service B

1. **Sensor Service (Main Service)**
   
   1. **Get Sensors**
      Endpoint:
      ```bash
      GET api/v1/sensors
      ```

      Authorization Header:
      ```bash
      Bearer <token>
      ```
   
      Query Params:
      
      | Key | Value | Data Type | Description |
      |--------|-----------|-------------|-------------|
      | `device_code` | A | string | ID1, optional |
      | `device_number` | 1 | int | ID2, optional |
      | `start_time` | 2025-08-28T15:51:00Z | time.Time | optional |
      | `end_time` | 2025-08-28T15:51:59Z | time.Time | optional |
      | `limit` | 10 | int | required, min=1 |
      | `page` | 1 | int | required, min=1 |
   
   
      Response (Success - 200 OK):
      ```bash
      {
       "status": "success",
       "message": "ok",
       "data": [
           {
               "id": 4,
               "sensor_type": "humidity",
               "sensor_value": 55.8,
               "device_code": "B",
               "device_number": 1,
               "timestamp": "2025-08-24T01:22:48Z",
               "created_at": "2025-08-24T01:22:48Z",
               "updated_at": "2025-08-24T02:40:31Z",
               "deleted_at": null
           },
           ....
       ],
       "paginator": {
           "total_items": 3,
           "total_pages": 1,
           "item_from": 1,
           "item_to": 3,
           "current_page": 1,
           "limit": 10
       }
      }
      ```

   2. **Update Sensors**
      Endpoint:
      ```bash
      PUT api/v1/sensors
      ```
  
      Authorization Header:
      ```bash
      Bearer <token>
      ```
   
      Request Body:
      ```bash
      {
          "criteria": {
              "device_code": "I",
              "device_number": 1,
              "start_time": "2025-08-28T15:51:00Z",
              "end_time": "2025-08-28T15:51:59Z"
          },
          "changes": {
              "sensor_value": 33,
              "sensor_type": "humidity",
              "timestamp": "2009-05-22T10:27:00Z"
          }
      } 
      ```
   
      Response (Success - 200 OK):
      ```bash
      {
          "status": "success",
          "message": "ok"
      }
      ```

      Response (Failure - No Filter Provided - 400 Bad Request):
      ```bash
      {
          "status": "error",
          "message": "at least one filter is required for update"
      }
      ```

      Response (Failure - No Changes Data Provided - 400 Bad Request):
      ```bash
      {
          "status": "error",
          "message": "at least one field to update must be provided"
      }
      ```
      
   3. **Delete Sensors**
      Endpoint:
      ```bash
      DELETE api/v1/sensors
      ```
  
      Authorization Header:
      ```bash
      Bearer <token>
      ```
   
      Request Body:
      ```bash
      {
        "device_code": "I",
        "device_number": 1,
        "start_time": "2025-08-28T15:51:00Z",
        "end_time": "2025-08-28T15:51:59Z"
      } 
      ```
   
      Response (Success - 200 OK):
      ```bash
      {
          "status": "success",
          "message": "ok"
      }
      ```

      Response (Failure - No Filter Provided - 400 Bad Request):
      ```bash
      {
          "status": "error",
          "message": "at least one filter is required for deletion"
      }
      ```

      Response (Failure - Start Time Greater Than End Time - 400 Bad Request):
      ```bash
      {
          "status": "error",
          "message": "start time cannot be greater than end time"
      }
      ```

2. **Auth Service (Simplfy Only For Authentication & Authorization)**

   1. **Register**
      Endpoint:
      ```bash
      POST api/v1/auth/register
      ```
   
      Request Body:
      ```bash
      {
        "username": "admin",
        "password": "admin",
        "role": "admin"        // can be admin or user
      }
      ```
   
      Response (Success - 200 OK):
      ```bash
      {
          "status": "success",
          "message": "ok"
      }
      ```

   2. **Login**
      Endpoint:
      ```bash
      POST api/v1/auth/login
      ```
   
      Request Body:
      ```bash
      {
        "username": "user",
        "password": "user"
      }
      ```
   
      Response (Success - 200 OK):
      ```bash
      {
          "status": "success",
          "message": "ok",
          "data": {
              "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTY0NDU2NjksInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6MX0.HY-moLeuY7r704tLL-LwEhrGbGKvBU7KpcdTX-_jezk",
              "expires_at": 1756445669
          }
      }
      ```

### Service A

1. **Sensor Service**

   1. **Start Sensor Generator**
      Endpoint:
      ```bash
      POST api/v1/sensors/generate/start
      ```
   
      Response (Success - 200 OK):
      ```bash
      {
          "status": "success",
          "message": "sensor generator started"
      }
      ```

   2. **Stop Sensor Generator**
      Endpoint:
      ```bash
      POST api/v1/sensors/generate/stop
      ```
   
      Response (Success - 200 OK):
      ```bash
      {
          "status": "success",
          "message": "sensor generator stopped"
      }
      ```

   3. **Update Generate Sensor Frequency**
      Endpoint:
      ```bash
      PUT api/v1/sensors/generate/frequency
      ```
  
      Request Body:
      ```bash
      {
        "frequency": 25   // in seconds
      }
      ```
   
      Response (Success - 200 OK):
      ```bash
      {
          "status": "success",
          "message": "ok"
      }
      ```


---


## 📡 Postman Collection


You can test all API endpoints by importing the Postman collection and environment into Postman:

### 1. Import Collection
- File: [sensor_api_postman_collection.json](./docs/postman/sensor_api_postman_collection.json)
- This collection contains all your API requests including examples.

### 2. Import Environment
- File: [sensor_api_local postman_environment.json](./docs/postman/sensor_api_local_postman_environment.json)
- The environment contains variables used in the requests, such as:
  - `{{token}}` – JWT token for authorization

### 3. Usage
1. Import both files into Postman.  
2. Select the imported environment from the top-right dropdown.    
3. Set `{{token}}` by logging in via the `/auth/login` endpoint or manually adding a valid JWT.  
4. Run requests — all endpoints will use the variables automatically.


---


## 🧪 Testing Application


// Coming Soon


---

## 🔮 Limitations & Future Improvements

While the current implementation meets the core project requirements, several limitations remain, and there are opportunities for enhancement:

1. **Swagger API Documentation**
   - Limitation: Swagger documentation for all endpoints has not been fully implemented.
   - Future Improvement: Generate complete Swagger docs for all REST API endpoints in Microservice B to improve developer experience and ease of testing.

2. **Load Balancing & Auto-Scaling**
   - Limitation: No API Gateway, load balancing, or auto-scaling implemented.
   - Future Improvement:
     - Introduce API Gateway for centralized request entry.
     - Implement load balancing across Microservice B instances.
     - Enable auto-scaling using Kubernetes or other orchestration tools to handle high loads.

3. **Unit Testing & CI/CD**
   - Limitation: Unit tests and integration tests are not implemented.
   - Future Improvement:
     - Add unit tests for controller and service layers.
     - Test messaging systems and automatic sensor data generation.
     - Set up CI/CD pipelines (e.g., GitHub Actions) for automated build and test.

4. **Payload Validation & Rate Limiting**
   - Limitation: Minimal payload validation and no rate limiting.
   - Future Improvement:
     - Enforce stricter validation for sensor values, IDs, timestamps, and device codes.
     - Restrict sensor frequency updates below a minimum threshold.
     - Implement rate limiting to prevent API abuse.

5. **JWT Authentication & Authorization**
   - Limitation: Only access tokens implemented; refresh token mechanism is missing.
   - Future Improvement:
     - Implement refresh tokens with single-use and soft deletion.
     - Use cron jobs to hard delete expired tokens.
     - Expand authorization using roles and permissions.
     - Consider kid in JWT for key management.

6. **Kafka Data Stream & Reliability**
   - Limitation: Single partition with replication factor = 1; no persistent storage.
   - Future Improvement:
     - Increase partitions and replication factor for fault tolerance.
     - Store Kafka data on disk for persistence.
     - Utilize message keys to maintain order and optimize throughput.
     - Monitor broker health and performance.

7. **Database Design & ERD**
   - Limitation: Minimal tables; limited validation.
   - Future Improvement:
     - Separate tables for users, roles, permissions, and sensor metadata (sensors, devices, sensor_types).
     - Validate device codes, sensor value ranges, and IDs.
     - Apply constraints to ensure data integrity.

8. **Environments**
   - Limitation: No clear separation of configs for development, staging, production.
   - Future Improvement:
     - Create separate configuration files or environment variables per environment.
     - Ensure proper security, logging, and database connections per environment.

9. **Monitoring & Observability**
   - Limitation: Limited health checks, logging, and metrics collection.
   - Future Improvement:
     - Add centralized logging, monitoring, and alerting.
     - Track API usage, database performance, and sensor data throughput.
     - Integrate Prometheus/Grafana for observability.

10. **Error Handling**
    - Limitation: Error responses are minimal
    - Future Improvement:
      - Ensure consistency across all endpoints and services.

11. **Deployment**
    - Limitation: Deployment is manual; no orchestration.
    - Future Improvement:
      - Use Kubernetes manifests for deploying microservices.
      - Implement rolling updates, service discovery, and scaling policies.

12. **Performance Optimization**
    - Limitation: No caching or query optimization implemented.
    - Future Improvement:
      - Introduce Redis caching for frequently accessed data.
      - Optimize database queries and indexes for large datasets.
      - Apply batching or streaming techniques for high-frequency sensor data.
  

---

## 📦 Note

// Coming Soon

---
