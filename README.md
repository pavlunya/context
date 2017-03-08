# fluc 

[![Build Status](https://img.shields.io/travis/pavlunya/fluc.svg?colorB=1b98e0&style=flat-square)](https://travis-ci.org/pavlunya/fluc)
[![Coverage Status](https://img.shields.io/coveralls/pavlunya/fluc.svg?colorB=1b98e0&style=flat-square)](https://coveralls.io/github/pavlunya/fluc)
[![Go Report Card](https://img.shields.io/badge/go_report-A+-1b98e0.svg?style=flat-square)](https://goreportcard.com/report/github.com/pavlunya/fluc)
[![GoDoc](https://img.shields.io/badge/go-documentation-1b98e0.svg?style=flat-square)](https://godoc.org/github.com/pavlunya/fluc)
[![MIT License](https://img.shields.io/badge/license-MIT-1b98e0.svg?style=flat-square)](https://raw.githubusercontent.com/pavlunya/fluc/master/LICENSE)

**fluc** provides fluent interface wrapper to work with `context.Context()`. 
The main idea is to simplify injection. 
Enough said, let's write some code.

```go
...

ctx := fluc.Context().
    With("redisClient", rc).
    With("mongoSess", ms).
    With("rmqConn", rmqc).
    Get()
    
...

func ContextInjector(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

...

http.ListenAndServe(":8080", ContextInjector(ctx, http.HandlerFunc(Handler)))

...

func Handler(rw http.ResponseWriter, r *http.Request) {
	ms, ok := r.Context().Value("mongoSess").(*mgo.Session)
	if !ok {
		// In this case panic will never be caused 
		// because mongoSess is present in the context.
		panic("Panic! Everything is broken!")
	}

	...
	
	ctx, cancel = fluc.Context(r.Context()).
		With("user", user).
		With("articles", articles).
		WithTimeout(10 * time.Second)
	defer cancel()
		
	someFunctionThatWorksHard(ctx)
	
	...
}

...
```

You can find more about contexts in [official documentation](https://golang.org/pkg/context).
