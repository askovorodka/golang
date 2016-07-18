package main

import "fmt"

type employee interface  {
	get_experience() int
	get_education() string
}

type medical struct {
	education string
	experience int
	medical_practic int
	medical_institute int
}

type it struct  {
	education string
	experience int
	it_cources int
}

func (m medical) get_experience() int {
	return m.experience + m.medical_practic + m.medical_institute
}

func (i it) get_experience() int {
	return i.experience + i.it_cources
}

func (m medical) get_education() string {
	return m.education
}

func (i it) get_education() string {
	return i.education
}

func get_employee(e employee)  {
	fmt.Println("Object:", e)
	fmt.Println("Education: ", e.get_education())
	fmt.Println("Experience:", e.get_experience())
}

func main()  {
	m := medical{education:"Medical Inst", experience:6, medical_practic:1, medical_institute:6}
	i := it{education:"It inst", experience:5, it_cources:1}

	get_employee(m)
	get_employee(i)

}