<div align="center">
<img src="https://s3api.wanz.site/image/reverse1999.jpg"
     alt="St.Pavlov Data Department - Logo"
     width="256" height="256" />

# backend

The Backend of St.Pavlov Data Department. Built in Golang.

</div>

## Getting started

- prerequisite game resources repository  
	clone the game resource to wherever you like
    ```bash
    git clone git@github.com:yuanyan3060/1999GameResource.git
    ```
  and change the `game_resource_path` in stpavlov_backend config:
	
	```toml
 	# pavlov.toml
 
 	# ...
 	game_resource_path = "/path/to/github.com/1999GameResource"
 	# ...
	```
 

- start stpavlov_backend with
    ```bash
    ./stpavlov_backend
    ```
