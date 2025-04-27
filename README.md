# oolio-assignment-backend

## Getting Started

Follow the steps below to set up and run the project:

### Step 1: Clone the Repository
```bash
git clone <repository-url>
cd oolio-assignment-backend
```

### Step 2: Add Data Files
Place the following files into the `data` directory:
[couponbase1.gz](https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase1.gz)
[couponbase2.gz](https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase2.gz)
[couponbase3.gz](https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase3.gz)

### Step 3: Install Dependencies
Run the following command to download and tidy up Go modules:
```bash
go mod tidy
```

### Step 4: Explore the API
- Use the provided **Postman collection** to interact with the API endpoints.
- The **OpenAPI specification** is updated and now includes security handling.

#### New Endpoints:
- **`/register`**: Register a new user.
- **`/token`**: Obtain a bearer token to authorize and access protected endpoints.
