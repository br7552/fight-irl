# router
Simple HTTP Go router based on [this gist](https://gist.github.com/benhoyt/a8ecc3f03b86d2d25543750cdcc77732) by Ben Hoyt.

    router := router.New()
  
    router.HandlerFunc(http.MethodPost, "/v1/users", registerUserHandler)
    router.HandlerFunc(http.MethodGet, "/v1/users/:id", showUserHandler)
  
    router.NotFound = http.HandlerFunc(customNotFoundHandler)
    router.MethodNotAllowed = http.HandlerFunc(customMethodNotAllowedHandler)
