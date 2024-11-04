# Buffet POS Backend

## Prerequisites

- **[GoFiber](https://gofiber.io/)** - An Express-inspired web framework for building fast and scalable applications in Go.
- **[Gorm](https://gorm.io/)** - An ORM library for Go, making it easy to work with databases through object relational mapping.
- **[PostgreSQL](https://www.postgresql.org/)** - A powerful, open source object-relational database system.
- **[Docker](https://www.docker.com/)** - A platform for developing, shipping, and running applications in containers.

## Features

### **Store Features**

- **Login**

  - Allows store personnel to securely log into the system to manage store functions and prevent unauthorized access.

- **Add Menu**

  - Enables the store to add or remove menu items, categorize them, and upload images for each menu item, making menu information clear and appealing.

- **Table Reserve**

  - Allows the store to reserve tables for customers, specifying the number of guests and table number for convenient seat management.

- **Print QR Code**

  - Enables the store to print receipts with QR codes so customers can scan to view order details and place orders directly through the system.

- **Check Order**

  - Allows the store to view incoming customer orders with details of items ordered and quantities needed.

- **Payment**

  - Enables the store to check the payment status for each table and cancel reservations if customers cancel their table booking.

- **Order History**

  - Allows the store to check the history of served orders for future reference.

- **Setting**
  - Enables the store to manage and add tables within the premises, assign table numbers, and set per-person service charges.

### **Customer Features**

- **Invoice**

  - Customers can place orders by scanning a QR code or through other order methods without needing to create an account directly with the store.

- **Menu**

  - Customers can view the restaurant's menu, including details like the name, image, and category of each item, allowing for easy selection.

- **Add Order**

  - Customers can select desired menu items and specify quantities, adding items directly to their order from the menu.

- **Order List**

  - Displays the list of items the customer has chosen, along with specified quantities, allowing the store to prepare items according to the customer’s order.

- **Order History**
  - Displays a history of all orders placed by the customer, enabling them to review past orders.

## Contributors

- นายณกรณ์ บุญประสงค์ 6510405458
- นายทัตพง์ วงศ์ไชยทา 6510405521
- นายภีรวิช ภักดีภิญโญ 6510405741
- นายภูมิระพี เสริญวณิชกุล 6510405750
- นายปิยะ กองศรี 6510450666
- นางสาวปุญญิศา ธัญญพงษ์ 6510450674

## Setup Instructions

### Clone the repository

```bash
git clone https://github.com/cs471-buffetpos/buffet-pos-backend
```

### Navigate to the project directory

```bash
cd buffet-pos-backend
```

### Install dependencies

```bash
go mod tidy
```

### Setup environment variables using .env

```bash
cp .env.example .env
```

### Run database using Docker Compose

```bash
docker compose up -d
```

## Running the server

```bash
go run main.go
```

## Database Migration

```bash
make db-migrate
```
