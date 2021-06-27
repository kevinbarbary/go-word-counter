# go-word-counter

A web server that provides word counting service

It accepts arbitrary text via POST request, and counts number of different word occurrences, storing it internally to be able to report it using a separate GET request.

For example:

Given request:
```
POST /text
Content-Type: text/plain

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud
exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute
irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla 
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui 
officia deserunt mollit anim id est laborum.
```
It counts a number of times each word occurs in the text and stores the numbers internally (in some memory structure).

Then, given the request like:
```
GET /counts
Accepts: application/json
```
It returns a response like:
```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "adipiscing": 1,
  ...
  "dolor": 2,
  ...
  "elit": 1,
  ...
  "ut": 3
}
```