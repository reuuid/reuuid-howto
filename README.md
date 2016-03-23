# ReUUID How-To
Find out how to use ReUUID from your langauge of choice.

---

If you can't find it here, this is how our REST API works:

####To Get One
URL:

    http://reuuid.org/get/

Response:

    029e05cf-fbf8-4fcb-819b-247b4f26426f
    
---

#### To Get Five
URL:

    http://reuuid.org/get/5

Response:

    3b54969c-d9fa-4ac9-aa38-4c69590ebaa5
    1237168c-35c4-437f-94fe-f48fe972eafa
    ac1c7d22-4903-4231-ae9a-c042c3a6211d
    b63f76ac-3d7c-43fd-b966-38ce938a126e
    03e65fe3-21d2-4ef4-bbf1-f14bb42f06e3

If there is an error, the Null UUID will be returned. If there are not enough UUIDs, less than the requested amount will be returned; please conserve them. Note that you cannont get more than 100 UUIDs at a time.

---

### To Give Some
URL:

    http://reuuid.org/give/

Request:
  
    19936217-7c8c-4c9e-9d5b-8fa4fecb3bf3
    a1bf277c-6cb8-4676-936a-4a32e1c908bb
    ...

Send us a POST message to the URL with as many as you want to give us.
