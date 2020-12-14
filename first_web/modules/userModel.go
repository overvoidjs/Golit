package modules

import ("fmt")

type User struct {
  Name string
  Age uint16
  Money int16
  Avg_grades, Happiness float64
  Hobbies []string
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is : %s. He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string){
  u.Name = newName
}
