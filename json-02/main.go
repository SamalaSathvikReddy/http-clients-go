package main

import (
	"encoding/json"
)


const issueList = `
[
    {
        "id" : 0,
        "name" : "Fix the thing",
        "estimate" : 0.5,
        "completed" : false
    },
    {
        "id" : 1, 
        "name" : "Unstick the widget", 
        "estimate" : 30, 
        "completed" : false
    }
]
`

const userObject = `
{
    "name" : "Sathvik", 
    "role" : "Developer",
    "remote" : true
}
`

func marshalAll [T any] (items []T) ([][] byte, error) {
    var res [][]byte
    for v := range items {
        item := items[v]
        data, err := json.Marshal(item)
        if err != nil {
            return nil, err
        }
        res = append(res, data)
    }
    return res, nil
}
