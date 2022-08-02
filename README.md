# basic_crud

Find below architectural sketch of the implementation.
This will help give you a mental picture before you go through 

Tools used.
1. MySQL Database
2. GORM
3. JSON Marshalling/Unmarshalling
4. Gorilla Mux

Project Structure
1. cmd
    main
      main.go
2. pkg
    config
      app.go
    controllers
      book_controller.go
    models
      book.go
    routes
      book_routes.go
    utils
      utils.go

<img width="614" alt="image" src="https://user-images.githubusercontent.com/27703937/182417370-f49414bd-a1c9-4cf8-bde5-855c86d828ce.png">
