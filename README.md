# My telegram bot for training
Bot is under construction so don't judge too harshly :)

Bot repeats the text sent to it. Also, there are some commands available: 

* `/help` — help info with the list of the commands  
* `/list` — list of products:  
  [ID 0] one  
  [ID 1] two  
  [ID 2] three  
  [ID 3] four  
  [ID 4] five  
  At the moment this list of products is hardcoded.  

* `/get [ID]` — get product info by its ID. Example of usage:  
  `/get 2`  
  result: three

* `/delete [ID]` — delete an existing entity by ID. Example of usage:  
  `/delete 5`  
  result: Product with ID = 5 was successfully deleted

* `/new [product_name]` — add new product. New product name should be written. Example of usage:  
  `/new abc`  
  result: Product with ID = 5 was successfully added  

* `/edit` — edit an existing product by ID. ID of existing product and its new name should be written. Example of usage:  
  `/edit 5 qwerty`  
  result: Product was successfully updated

For launching the bot .env file with your bot token should be created in the project's root directory. Format of .env file:  
TOKEN=[your token]  
