GoTFTP application:
========================
Requirements: go1.3.3 darwin/amd64

##Go totorial:
	* Run TFTP Server: $:go run main.go

##Application commands:
### server:
	* "ls" : show all the filenames and the size received 
### client:
	since the server bind tftp port randomly, start the server first then read the port P binded
	* $: tftp
	  $: connect 127.0.0.1 P
	  $: put targetFile
	  $: get targetFile <dst filename>