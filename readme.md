
# capybara-api

capybara-api is a sweet, cute, fast API to get a whole bunch of images or a random image of a capybara

## Documentation

Note: 
- Pages are cached for 7 days
- API Base URL is `https://api.capy.lol`


### Get Capybara
Get random capy and return as an image
```
GET /v1/capybara
```

Get random capy and return as JSON
```
GET /v1/capybara?json=true
```

### Get Capybara by image index/ID
Get capy by image index/ID as an image
```
GET /v1/capybara/:index
```

Get capy by image index/ID as json
```
GET /v1/capybara/:index?json=true
```


### Get Capybaras
Get a bunch of capys
```
GET /v1/capybaras
```

Get a bunch of capys from index
```
GET /v1/capybaras?from=index
```

Get a bunch of capys but to a certain limit
```
GET /v1/capybaras?take=limit
```

Or combine both!
```
GET /v1/capybaras?take=limit&from=index
```