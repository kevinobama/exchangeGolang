
## Database Setup

It looks like you chose to set up your application using a mysql database! Fantastic!

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start mysql for you.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started mysql, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a
## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.


buffalo g resource orders order_id name email bio:nulls.Text
####################################################
Run the dev server on a custom port
Sometimes you will already have an app working on the 3000 port. You can configure the dev server port by providing the PORT environment variable:

 PORT=3001 buffalo dev

####################################################
Querying
Find By ID
user := User{}
err := db.Find(&user, id)
Find All
users := []User{}
err := db.All(&users)
err = db.Where("id in (?)", 1, 2, 3).All(&users)
Find Where
users := []models.User{}
query := db.Where("id = 1").Where("name = 'Mark'")
err := query.All(&users)

err = tx.Where("id in (?)", 1, 2, 3).All(&users)

####################################################
fmt.Printf("%+v\n", robots)
To print the name of the fields in a struct:

fmt.Printf("%+v\n", yourProject)
From the fmt package:

when printing structs, the plus flag (%+v) adds field names

That supposes you have an instance of Project (in 'yourProject')

The article JSON and Go will give more details on how to retrieve the values from a JSON struct.

This Go by example page provides another technique:

type Response2 struct {
  Page   int      `json:"page"`
  Fruits []string `json:"fruits"`
}

res2D := &Response2{
    Page:   1,
    Fruits: []string{"apple", "peach", "pear"}}
res2B, _ := json.Marshal(res2D)
fmt.Println(string(res2B))
That would print:

{"Page":1,"Fruits":["apple","peach","pear"]}
If you don't have any instance, then you need to use reflection to display the name of the field of a given struct, as in this example.

type T struct {
    A int
    B string
}

t := T{23, "skidoo"}
s := reflect.ValueOf(&t).Elem()
typeOfT := s.Type()

for i := 0; i < s.NumField(); i++ {
    f := s.Field(i)
    fmt.Printf("%d: %s %s = %v\n", i,
        typeOfT.Field(i).Name, f.Type(), f.Interface())
}