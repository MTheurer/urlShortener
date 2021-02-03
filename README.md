**ShortyResty Challenge**

**General Overview**

This is a REST API that will redirect us from an assigned short ID to a full URL. 



**Usage**

 * ShortyResty can be given a url in the form of Json, and assign an 8 digit key that can be used as a shortened version of its original value. 

 * Your shortened key can be used through localhost port 8080, through the endpoint /(key)



**Endpoints and their Uses**

* The endpoint  will act as a Homepage, producing a small text banner.
* The endpoint ""/(key)" acts as a redirection middleman, redirecting you to the value (url) of the key.
* The endpoint "/shorten" creates a POST request, randomly selecting 8 digits for the key to your new value (url).



 **Execution**

To execute the the Url Shortener file:

1. Run "go build -o outputname" in command line to compile the open file

2. "./outputname" runs the compiled file



**Testing**

Using an Existing Key



1. Go to "http://127.0.0.1:8080"

2. add endpoint /(key) after confirming port 8080 is active

You should be redirected to your cooresponding URL, or be given a 404 error if the URL does not exist.



Creating a Key for New URL (Using Postman)

1. After opening Postman, go to "http://127.0.0.1:8080/shorten"

2. You will be entering your new URL within the "body", through a POST method

3. Enter your URL in Json format

4. If your URL is invalid, you will recieve an error 406.

5. If the URL cannot be decoded, you will recieve an error 400.

6. If your POST request is successful, you will receive the new key in Json format, example:

{"short_url":"http://127.0.0.1:8080/(key)}

