### Step 1: Use Curl to test the API

Open a terminal and use the following Curl commands to test each route:

1. **Test the `/hello` route:**
   ```bash
   curl -X GET http://localhost:3000/hello
   ```

2. **Test the `/allBooks` route:**
   ```bash
   curl -X GET http://localhost:3000/allBooks
   ```

3. **Test the `/addBook` route:**
   Assuming your `AddBook` route expects a JSON payload with book details, you can use:
   ```bash
   curl -X POST http://localhost:3000/addBook \
   -H "Content-Type: application/json" \
   -d '{"title": "The Catcher in the Rye", "author": "J.D. Salinger"}'
   ```

4. **Test the `/book` route:**
   To get a specific book, modify the payload accordingly:
   ```bash
   curl -X POST http://localhost:3000/book \
   -H "Content-Type: application/json" \
   -d '{"id": 1}'
   ```

5. **Test the `/update` route:**
   If the update requires a book ID and new details:
   ```bash
   curl -X PUT http://localhost:3000/update \
   -H "Content-Type: application/json" \
   -d '{"id": 1, "title": "Updated Title", "author": "Updated Author"}'
   ```

6. **Test the `/delete` route:**
   To delete a book by ID:
   ```bash
   curl -X DELETE http://localhost:3000/delete \
   -H "Content-Type: application/json" \
   -d '{"id": 1}'
   ```
