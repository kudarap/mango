package user

import "github.com/javinc/playgo/goryo/resources/user"

func Create(m *user.Model) user.Model {
    m.Name += " appended on service"

    row, _ := user.Create(m)

    return row
}
