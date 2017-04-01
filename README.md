# LolScroll

Social network LolScroll

# Documentation
## Download :
  Firstly need to install golang:  
  sudo apt-get install golang  
  Secondly set GOPATH:  
  exprot GOPATH=$HOME/go  
  Next need download repository:  
  go get github.com/CatsMafia/LolScroll  
## Run :
  For run server on Linux need in terminal write:  
    ./$GOPATH/bin/LolScroll
## Registration and log in:
#### Registration:
  Post request on /registartion with values:  
  Username, Password  
  ``` html
  <form action="/registration" method="post">
	<label>Username:</label>
	<input type="text" name="Username"><br>
	<label>Password</label>
	<input type="password" name="Password"><br>
	<button type="submit"></button>
</form>
  ```
#### Log in:
  Post request on /login with values:  
  Username, Password  
``` html
  <form action="/login" method="post">
	<label>Username:</label>
	<input type="text" name="Username"><br>
	<label>Password</label>
	<input type="password" name="Password"><br>
	<button type="submit"></button>
</form>
```
## API :
  All api are available on URL begin with /api/
### Add new lol:
  /api/newlol  
  For add new lol need send POST request with values:   
  userId: user id  
  text: text of lol  
### Get lols:  
  /api/getlols  
  For get lol need send Get request with values:  
  id: Lol id (need for search for one lol). Optionaly.  
  start: Starting id, from which the countdown. begins(). Optionaly, default 0.  
  count: num of lols.Optionaly, default 10.    
  hashtag: search by hashtag.  
  Example for searh by hashtag #разбудименяв420, need send hashtag="разбудименяв420". Optionaly.  
  linkpeople: searsh by linkpeaople.  
  Example for search by @Google need send linkpeople="Google"
### Put kek(like)
  /api/putkek
  For put kek(like) need to send post request with values:   
  id: lols id  
  userid: user, who puts kek  
  deltaKek: the number by which the number of keks changes  
