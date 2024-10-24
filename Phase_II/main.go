package main

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Niyonkuru", Quantity: 2},
	{ID: "1", Title: "The Great Gatsby", Author: "NiyonkuruII", Quantity: 5},
	{ID: "1", Title: "War and Peace", Author: "NiyonkuruIII", Quantity: 6},
}
