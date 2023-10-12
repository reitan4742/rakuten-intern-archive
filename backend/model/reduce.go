package model

type UserReduce struct {
	Username    string `json:"username" gorm:"primary_key"`
	Exp         int    `json:"exp"`
	ReduceScore int    `json:"reduce_score"`
}

type UserReduceRequest struct {
	Username string `json:"username"`
}

type UserReduceResponse struct {
	ReduceScore int `json:"reduce_score"`
}

type AllReduce struct {
	AllReduceId int `json:"allreduceid" gorm:"primary_key"`
	AllReduce   int `json:"allreduce"`
}

type AllReduceResponse struct {
	AllReduce int `json:"allreduce"`
}
