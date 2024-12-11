package mockdata

import "altech-omega-api/domain"

var Books = []domain.Book{
	{
		ID:          "1156e695-8043-4c51-a9ae-d05e0eba9b2c",
		AuthorID:    "2dd2784b-31d5-4bb7-9429-644ba619476a",
		Title:       "The 5AM Club",
		Description: "Bangun pagi supaya sukses",
		PublishDate: "2005-08-03",
		Genre:		 "Self Development",
		Pages:		 100,
	},
	{
		ID:          "35a16b5d-8eca-4055-a4d3-a62a355aeab8",
		AuthorID:    "b409d8d6-701a-4a31-b0b0-8db88f09f218",
		Title:       "Hyouka",
		Description: "Oreki Houtarou dan Chitanda",
		PublishDate: "2007-03-05",
		Genre:		 "Light Novel",
		Pages:		 200,
	},
}
