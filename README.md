# Todo APIs + Routing frameworks

Here are 5 key factors I would consider when evaluating a Golang router framework:

1. Performance - How many requests can the framework handle and how fast it responds, is it suitable for high load scenarios?
2. Features - Does the framework provide complete and powerful features like routing, middleware, parameter validation etc.
3. Ease of use - Is the framework API and syntax simple and easy to learn and use, is the documentation comprehensive?
4. Community support - Is the framework actively maintained, are issues responded timely, is it easy to get support?
5. Compatibility - Does the framework work well with other technologies used in the project, is it easy to integrate?

Some other aspects I would look at:

- Whether there are template scaffolds for quickly getting started.
- Are there mature middleware ecosystems.

By comprehensively considering these factors, I can have a good understanding of the pros and cons of the framework, and choose the one that best fits the project requirements. Need to evaluate both performance and usability, as well as ecosystem support.

| Package    | demo | routing                                              | controller               | context        | json binding                             | params/path | middleware |
| ---------- | ---- | ---------------------------------------------------- | ------------------------ | -------------- | ---------------------------------------- | ----------- | ---------- |
| gin        | ✅   | engine.GET() engine.Run()                            | c.JSON()                 | ctx            | c.ShouldBindJSON(&newTodo)               |             |            |
| fiber      | ✅   | app.Get() app.Listen()                               | c.Status().JSON(newTodp) | ctx            | c.BodyParser(newTodo)                    |             |            |
| iris       | ✅   | app.Get() app.Listen()                               | c.JSON(newTodo)          | ctx            | c.ReadJSON(&newTodo)                     |             |            |
| echo       | ✅   | e.GET() E.Start()                                    | c.JSON(code, newTodo)    | ctx            | c.Bind(newTodo)                          | todo/:id    |            |
| chi        | ✅   | chi.NewRouter() router.Get() http.ListenAndServe()   | w.Write()                | reader, writer | json.NewDecoder(r.Body).Decode(&newTodo) | todo/{id}   |            |
| mux        | ✅   | mux.NewRouter() Methods("GET") http.ListenAndServe() | w.Write()                | reader, writer | json.NewDecoder(r.Body).Decode(&newTodo) |             |            |
| HttpRouter | ✅   | httprouter.New() router.GET() http.ListenAndServe()  | w.Write()                | reader, writer | json.NewDecoder(r.Body).Decode(&newTodo) |             |            |
| FastHTTP   | ✅   | router.New() route.POST() fastHttp.ListenAndServe()  | c.SetBody                | ctx            | json.Maschal() json.unmarshal()          | todo/:id    |            |

## Performance

1 fastHttp - 800k reqs/sec
2 fiber - 600k reqs/sec
3 echo - 400k reqs/sec
4 iris - 300k reqs/sec
5 gin - 200k reqs/sec
6 chi - 150k reqs/sec
7 gorilla mux - 100k reqs/sec
8 httpRouter - 100k reqs/sec

## API User friendly

1 Gin
2 Echo
3 Iris
4 Chi
5 Gorilla Mux
6 httpRouter
7 fiber
8 fastHttp

## Middleware

| Middleware                      | Usage                            | Benefits                                  |
| ------------------------------- | -------------------------------- | ----------------------------------------- |
| Logging middleware              | Record request logs              | Debugging and monitoring                  |
| Exception handling middleware   | Handle exceptions                | Avoid application crashes                 |
| Parameter validation middleware | Validate request parameters      | Improve data quality                      |
| Authentication middleware       | Authentication and authorization | Ensure security                           |
| Rate limiting middleware        | Rate limit APIs                  | Prevent overload                          |
| Load balancing middleware       | Traffic distribution             | Share server load                         |
| Caching middleware              | Cache requests                   | Reduce database queries                   |
| CORS middleware                 | Support cross-origin requests    | Enable separation of frontend and backend |
| Formatting middleware           | Format request/response          | Simplify processing                       |
| Retry middleware                | Retry failed requests            | Improve fault tolerance                   |
| Metrics middleware              | Collect metrics                  | Analyze and monitor                       |
| Authentication middleware       | User authentication              | Ensure security                           |

## Tips

### config sql.db for better performance

- https://www.alexedwards.net/blog/configuring-sqldb

> limit on the number of database connections: PostgreSQL(default: 100), MySQL/MariaDB(default: 150 + 1), MongoDB(default: 100)

### DB connections

For small-to-medium web applications, use the following settings as a starting point, and then optimize them later:

- db.SetMaxOpenConns(25)
- db.SetMaxIdleConns(25)
- db.SetConnMaxLifetime(5\*time.Minute)
