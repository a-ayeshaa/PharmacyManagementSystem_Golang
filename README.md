# PharmacyManagementSystem_Golang

	This API based system has two user roles:
	1)Admin
	2)Customer
	
	Note:Postgresql has been used in this service as the database.
	After the user logs in with the credentials, the user is allowed access to the system according to his role registered.
	1)Admin features:
		-Add medicine 
		-Delete medicine
		-Update medicine
		-Check detail of each medicine
		-Check all medicine from inventory
	2)Customer features:
		-Add medicine to cart
		-Delete medicine from cart
		-Check Cart
		-Confirm Order
	
		     
	To run this project open terminal and run ./run to run the project and ./worker.sh to start the worker.

	e.g, If Admin wants to add medicine to the inventory list:
	Request URL: http://localhost:3000/medicine/add
	Method: POST
	Body: {
		"name":"<name>"
		"price":<price>
	}

	e.g, If Admin wants to delete a medicine from the list:
	Request URL: http://localhost:3000/medicine/delete/<id>
	Method: DELETE
	
