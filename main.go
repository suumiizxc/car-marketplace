package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/gin-bookstore/config"
	"github.com/suumiizxc/gin-bookstore/controllers"
	client "github.com/suumiizxc/gin-bookstore/controllers/client"
	customer "github.com/suumiizxc/gin-bookstore/controllers/core/customer"
	furniture "github.com/suumiizxc/gin-bookstore/controllers/furniture"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
	"github.com/suumiizxc/gin-bookstore/helper/redis"
)

func main() {
	r := gin.Default()

	// Connect to database

	config.ConnectDatabase()
	helper_core.CH.Init()
	redis.RedisConfig()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/furnitures", furniture.FindFurnitures)
	r.POST("/furnitures", furniture.CreateFurniture)

	r.GET("/clients", client.FindClients)
	r.POST("/client", client.CreateClient)
	r.POST("/client/login-phone", client.LoginPhone)
	r.POST("/client/login-email", client.LoginEmail)
	r.GET("/client", client.ProfileClient)

	r.POST("/core/customer/create-test", customer.CreateTest)
	r.POST("/core/customer/create", customer.CreateCustomer)
	r.POST("/core/customer/countryCodes/:limit/:page", customer.GetCountryCodes)

	r.GET("/core/customer/education-degrees", customer.EducationDegreeList)
	r.GET("/core/customer/education-degree/:id", customer.EducationDegreeGet)
	r.POST("/core/customer/education-degree/create", customer.EducationDegreeCreate)
	r.DELETE("/core/customer/education-degree/delete/:id", customer.EducationDegreeDelete)

	r.GET("/core/customer/education-levels", customer.EducationLevelList)
	r.GET("/core/customer/education-level/:id", customer.EducationLevelGet)
	r.POST("/core/customer/education-level/create", customer.EducationLevelCreate)
	r.DELETE("/core/customer/education-level/delete/:id", customer.EducationLevelDelete)

	r.GET("/core/customer/employments", customer.EmploymentList)
	r.GET("/core/customer/employment/:id", customer.EmploymentGet)
	r.POST("/core/customer/employment/create", customer.EmploymentCreate)
	r.DELETE("/core/customer/employment/delete/:id", customer.EmploymentDelete)

	r.GET("/core/customer/type-of-organizations", customer.TypeOfOrganizationList)
	r.GET("/core/customer/type-of-organization/:id", customer.TypeOfOrganizationGet)
	r.POST("/core/customer/type-of-organization/create", customer.TypeOfOrganizationCreate)
	r.DELETE("/core/customer/type-of-organization/delete/:id", customer.TypeOfOrganizationDelete)

	r.GET("/core/customer/route-of-activities/:page/:limit/:value", customer.RouteOfActivityList)
	r.GET("/core/customer/route-of-activity/:id", customer.RouteOfActivityGet)
	r.POST("/core/customer/route-of-activity/create", customer.RouteOfActivityCreate)
	r.DELETE("/core/customer/route-of-activity/delete/:id", customer.RouteOfActivityDelete)

	r.GET("/core/customer/type-of-contacts", customer.TypeOfContactList)

	r.GET("/core/customer/employees", customer.EmployeeList)
	r.GET("/core/customer/employee/:id", customer.EmployeeGet)
	r.POST("/core/customer/employee/create", customer.EmployeeCreate)
	r.DELETE("/core/customer/employee/delete/:id", customer.EmployeeDelete)

	r.GET("/core/customer/relation-customer-companys", customer.RelationCustomerCompanyList)
	r.GET("/core/customer/relation-customer-company/:id", customer.RelationCustomerCompanyGet)
	r.POST("/core/customer/relation-customer-company/create", customer.RelationCustomerCompanyCreate)
	r.DELETE("/core/customer/relation-customer-company/delete/:id", customer.RelationCustomerCompanyDelete)

	// Run the server
	r.Run()
}
