# Details
This app uses a map as a database for simplicity. You can test the app with following commands:
```
    // to get all of the books
	curl localhost:8080/books
    
    // to get a specific book
	curl localhost:8080/book/2

    // to create a new book instance
	curl -X POST -d '{"Id":5, "Name":"Reza"}' -L localhost:8080/book/5
	
    // to delete a specific book instance
        curl -X DELETE -L localhost:8080/book/5
    
    // to search for a specific book name
	curl localhost:8080/book?name="Elevation"

```

# Refrences
- middleware stuff:
    * Philipp Tanlak   [Middleware (Basic)](https://gowebexamples.com/basic-middleware/)
    * Philipp Tanlak   [Middleware (Advanced)](https://gowebexamples.com/advanced-middleware/) 
- Context:
    * Mat Ryer         [Context keys in Go](https://medium.com/@matryer/context-keys-in-go-5312346a868d)

- Overall structure:
    * Sandeep Kalra    [git repo: mv-backend](https://github.com/sandeepkalra/mv-backend)
- Overall Go Porgramming:
    * William Kennedy  [Ultimate Go Programming](https://www.oreilly.com/library/view/ultimate-go-programming/9780135261651/)
    * Todd McLeod      [Learn How To Code](https://www.udemy.com/course/learn-how-to-code/)
    * Todd McLeod      [Web Development with Google’s Go](https://www.udemy.com/go-programming-language/)