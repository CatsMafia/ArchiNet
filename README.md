# LolScroll

Social network LolScroll

# Documentation
## Download :
  Firstly need to install golang:  
  sudo apt-get install golang  
  Secondly set GOPATH:  
  exprot GOPATH=$HOME/go  
  Next need to download repository:  
  go get github.com/CatsMafia/LolScroll  
## Run :
  To run the server on Linux you need to write in terminal:  
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
  To add new you lol need to send POST request with values:   
  userId: user id  
  text: text of lol  
### Get lols:  
  /api/getlols  
  For getting lol you need to send Get request with values:  
  id: Lol id (it's need to search for one lol). Optionaly.  
  start: Starting id, from which the countup begins. Optionaly, default 0.  
  count: Num of lols.Optionaly, default 10.    
  hashtag: Search by hashtag.  
  Example for search by hashtag #разбудименяв420, need to be send hashtag="разбудименяв420". Optionaly.  
  linkpeople: search by linkpeople.  
  Example for searching by @Google need to send linkpeople="Google"
### Put kek(like)
  /api/putkek  
  For put kek(like) need to send post request with values:   
  id: lols id  
  userid: user, who puts kek  
  kek: the number by which the number of keks changes

### All Examples:
```html
<form action="/api/newlol" method="post">
	Text:
	<input type="text" name="text">
	<input type="hidden" name="userId" value="1"> 
	<button type="submit">Hello WOrld</button>
</form>
<br>
<form action="/api/getlol">
	<input type="text" name="hashtag">
	<input type="submit">
</form>

<form action="/api/putkek" method="post">
	<input type="hidden" name="id" value="0">
	<input type="hidden" name="lol" value="1">
	<input type="hidden" name="user" value="1">
	<input type="submit">
</form>

<form action="/api/putkek" method="post">
	<input type="hidden" name="id" value="0">
	<input type="hidden" name="lol" value="-1">
	<input type="hidden" name="user" value="1">
	<input type="submit">
</form>
```
