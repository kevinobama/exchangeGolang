package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/kevingates/exchange/models"
	"github.com/pkg/errors"
)

// Following naming logic is implemented in Buffalo:
// Model: Singular (Robot)
// DB Table: Plural (robots)
// Resource: Plural (Robots)
// Path: Plural (/robots)
// View Template Folder: Plural (/templates/robots/)

// RobotsResource is the resource for the Robot model
type RobotsResource struct {
	buffalo.Resource
}

func UsersRegisterGet(c buffalo.Context) error {
	//return c.Render(200, r.JSON([]))
	return c.Render(200, r.String(c.Param("name")))
}

func (v RobotsResource) robotListAll(context buffalo.Context) error {

	db, ok := context.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	robots := &models.Robots{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	query := db.PaginateFromParams(context.Params())

	// Retrieve all Robots from the DB
	if err := query.All(robots); err != nil {
		return errors.WithStack(err)
	}

	robot := &models.Robot{}

	// To find the Robot the parameter robot_id is used.
	db.Find(robot, "668a37f1-c47d-48a3-b8c8-d456ec98aee7")

	//fmt.Printf("%+v\n", robots)
	fmt.Printf("robots = \n")
	fmt.Printf("%+v\n", robots)
	fmt.Printf("%+v\n", "robot=")
	fmt.Printf("%+v\n", robot)

	//return c.Render(200, r.Auto(c, robots))
	return context.Render(200, r.JSON(robots))
}

// List gets all Robots. This function is mapped to the path
// GET /robots
func (v RobotsResource) List(c buffalo.Context) error {

	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	robots := &models.Robots{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	query := db.PaginateFromParams(c.Params())

	// Retrieve all Robots from the DB
	if err := query.All(robots); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", query.Paginator)

	//fmt.Printf("%+v\n", robots)
	fmt.Printf("robot = \n")
	fmt.Printf("%+v\n", robots)
	fmt.Printf("%+v\n", "data")

	return c.Render(200, r.Auto(c, robots))
}

// Show gets the data for one Robot. This function is mapped to
// the path GET /robots/{robot_id}
func (v RobotsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Robot
	robot := &models.Robot{}

	// To find the Robot the parameter robot_id is used.
	if err := tx.Find(robot, c.Param("robot_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, robot))
}

// New renders the form for creating a new Robot.
// This function is mapped to the path GET /robots/new
func (v RobotsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Robot{}))
}

// Create adds a Robot to the DB. This function is mapped to the
// path POST /robots
func (v RobotsResource) Create(c buffalo.Context) error {
	// Allocate an empty Robot
	robot := &models.Robot{}

	// Bind robot to the html form elements
	if err := c.Bind(robot); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(robot)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, robot))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Robot was created successfully")

	// and redirect to the robots index page
	return c.Render(201, r.Auto(c, robot))
}

// Edit renders a edit form for a Robot. This function is
// mapped to the path GET /robots/{robot_id}/edit
func (v RobotsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Robot
	robot := &models.Robot{}

	if err := tx.Find(robot, c.Param("robot_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, robot))
}

// Update changes a Robot in the DB. This function is mapped to
// the path PUT /robots/{robot_id}
func (v RobotsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Robot
	robot := &models.Robot{}

	if err := tx.Find(robot, c.Param("robot_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Robot to the html form elements
	if err := c.Bind(robot); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(robot)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, robot))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Robot was updated successfully")

	// and redirect to the robots index page
	return c.Render(200, r.Auto(c, robot))
}

// Destroy deletes a Robot from the DB. This function is mapped
// to the path DELETE /robots/{robot_id}
func (v RobotsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Robot
	robot := &models.Robot{}

	// To find the Robot the parameter robot_id is used.
	if err := tx.Find(robot, c.Param("robot_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(robot); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Robot was destroyed successfully")

	// Redirect to the robots index page
	return c.Render(200, r.Auto(c, robot))
}

// RobotsGetListAll default implementation.
func RobotsGetListAll(c buffalo.Context) error {
	return c.Render(200, r.HTML("robots/get_list_all.html"))
}

// RobotsGetListAllOne default implementation.
func RobotsGetListAllOne(c buffalo.Context) error {
	return c.Render(200, r.HTML("robots/get_list_all_one.html"))
}

func RobotsGetListAllTwo(c buffalo.Context) error {
	return c.Render(200, r.HTML("robots/get_list_all_one.html"))
}