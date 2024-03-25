package structs

type SignDanneyen struct{
	Id string `bson:"_id"`
	Name string
	Email string
	Login string
	Password string
	Permission string
}
type AddServiceData struct{
	Id string `bson:"_id"`
	Name string
	CourseCount int
}
type AddMapData struct{
	Id string `bson:"_id"`
	Year int
	Description string
}
type AddMemberData struct{
	Id string `bson:"_id"`
	Name string
	Position string
}


