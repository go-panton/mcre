# How do I use repository_test.go

## First time User:
Uncomment the TestNewDB function, this function will :-
 1. Create a new Database named "foo" if it does not exist
 2. Create the table required for the test
 3. Close the db COnnection after it's done

 Note: sql.Open("driverName","username:password@/")

4. Now if you need some dummy data you can uncomment the TestInsert function which will insert some dummy data into node, nodelink, fmedia, fverinfo, fmpolicy

 ## Old User:
If you don't like the database to be named foo, change it on the constant global fakeDBName variable.

Be reminded that if you already have dummy data in the table do not run the TestInsert function again because duplicate data will throw an error

The exec function take it the test, db object, statement to run, arguments  to run which will fail the test immediately if there is an error

If you need to drop the database entirely, make a new test, and instantiate a db object and include the dropDatabase function, this will drop the database entirely together with it's data

 Nothing left to say, just create the test function you want and run the test.
