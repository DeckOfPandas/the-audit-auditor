package models

type Audit struct {
    Id          int     `form:"-"`
    Title       string  `form:"title,text,Title:"`
    Summary     string  `form:"summary,text,Summary:"`
    Author      string  `form:"author,text,Author:"`
    Email       string  `form:"email,text,Email:"`
    Location    string  `form:"location,text,Location:"`
    Days        int     `form:"days,number,Days:"`
}
