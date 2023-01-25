# PharmacyManagementSystem_Golang

The system has two user roles:
	1)Admin
	2)Customer
	
	Note: Database was not used in this project instead it is done with a file system
	After the user logs in with the credentials, the user is allowed access to the system according to his role registered.
	1)Admin features:
		-Add medicine 
		-Delete medicine
		-Update medicine
		-Check detail of each medicine
	2)Customer features:
		-Add medicine to cart
		-Delete medicine from cart
		-Check Cart
		-Confirm Order
		
	meds.go : CRUD of entity "medicine"
	cart.go : CRUD of entity "cart"
	orders.go : CRUD of entity "order"
	
	Description: A medicine list will be seeded initially from which the admin can add, remove, update and get indivitual medicines. 
		     If the user is a customer then the customer can buy, remove from cart, check the cart and also confirm the order.
		     
To run this project open terminal in "task" and write "go run *"

e.g, If Admin wants to add medicine to the list:
	- After admin logs in, he is given 5 options (with the actions mentioned).
	- Choose option 1 by entering "1" (as this option indicates Adding medicine)
	- User is asked to give the name of the medicine and the price that the user wants to set
	- Medicine is added
e.g, If Admin wants to delete a medicine from the list:
	- Choose option 2 by entering "2" (as this option indicates deleting medicine)
	- User is asked to enter the "ID" of the medicine he wants to delete.
	- Medicine is deleted if the ID matches with the existing IDs
If User wants to exit, selecting option "5 : Exit" or Ctrl+C from keyboard will halt the program. 
