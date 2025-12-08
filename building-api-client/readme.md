## Details

**We will build a simple client with an elementary API server here. In this case, the API server will be an API that `returns a random number between 1 and 100`, `allows the seed to be set`, and `has a login process`.**

For that, we are going to need three endpoints. The first one is /random, which will receive a GET
request and return JSON containing the random value like this: {"value": 42}. The second one
will be /seed, which will receive a POST request with a JSON body containing the seed like this:
{"seed": 123} . And the last one will be /login, which will receive a POST request with a JSON body
containing the user and the password like this: {"user": "user", "password": "password"},
and it will return a token.